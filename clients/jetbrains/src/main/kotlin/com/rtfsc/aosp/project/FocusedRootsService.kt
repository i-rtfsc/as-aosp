package com.rtfsc.aosp.project

import com.intellij.openapi.Disposable
import com.intellij.openapi.application.WriteAction
import com.intellij.openapi.components.Service
import com.intellij.openapi.module.EmptyModuleType
import com.intellij.openapi.module.Module
import com.intellij.openapi.module.ModuleManager
import com.intellij.openapi.project.Project
import com.intellij.openapi.roots.ModuleRootManager
import com.rtfsc.aosp.profile.EngineProfileSettings

@Service(Service.Level.PROJECT)
internal class FocusedRootsService(
    private val project: Project,
) : Disposable {
    private val settings = project.getService(EngineProfileSettings::class.java)
    private var focusedModule: Module? = null

    fun apply(): WorkspaceRootPlan {
        val plan = WorkspaceRootPlanner.plan(settings.current)
        if (!plan.isUsable) return plan
        val module = WriteAction.compute<Module, RuntimeException> {
            focusedModule ?: ModuleManager.getInstance(project)
                .newNonPersistentModule("AOSP Focused", EmptyModuleType.EMPTY_MODULE)
                .also { focusedModule = it }
        }
        FocusedRootsApplier.apply(module, plan)
        return plan
    }

    override fun dispose() {
        val module = focusedModule ?: return
        WriteAction.run<RuntimeException> {
            ModuleManager.getInstance(project).disposeModule(module)
        }
        focusedModule = null
    }
}

internal object FocusedRootsApplier {
    fun apply(module: Module, plan: WorkspaceRootPlan) {
        WriteAction.compute<Unit, RuntimeException> {
            val rootModel = ModuleRootManager.getInstance(module).modifiableModel
            rootModel.clear()
            plan.sourceRoots.forEach { sourceRoot ->
                val content = rootModel.addContentEntry(sourceRoot.toString())
                content.addSourceFolder(sourceRoot.toString(), false)
                plan.excludedRoots
                    .filter { excluded -> excluded != sourceRoot && excluded.startsWith(sourceRoot) }
                    .forEach { excluded -> content.addExcludeFolder(excluded.toString()) }
            }
            rootModel.commit()
        }
    }
}
