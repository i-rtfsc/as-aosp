package com.rtfsc.aosp.aidl

import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.actionSystem.CommonDataKeys
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import com.intellij.openapi.vfs.LocalFileSystem
import com.rtfsc.aosp.engine.ProjectEngineService
import com.rtfsc.aosp.profile.EngineProfileSettings
import java.nio.file.Path

class AidlReferencesAction : AnAction("Find AIDL Usages") {
    private val mapper = ObjectMapper()

    override fun update(event: AnActionEvent) {
        event.presentation.isEnabledAndVisible = event.getData(CommonDataKeys.VIRTUAL_FILE)?.fileType == AidlFileType
    }

    override fun actionPerformed(event: AnActionEvent) {
        val project = event.project ?: return
        val file = event.getData(CommonDataKeys.VIRTUAL_FILE) ?: return
        val root = runCatching { Path.of(project.getService(EngineProfileSettings::class.java).current.workspaceViewRoot).toAbsolutePath().normalize() }.getOrNull() ?: return
        val current = runCatching { root.relativize(file.toNioPath().toAbsolutePath().normalize()).toString().replace('\\', '/') }.getOrNull() ?: return
        val settings = project.getService(EngineProfileSettings::class.java).current
        val request = mapper.createObjectNode().apply { put("workspaceViewRoot", root.toString()); put("currentPath", current); put("name", file.nameWithoutExtension) }
        object : Task.Backgroundable(project, "Find AIDL Usages", true) {
            private var count = 0
            override fun run(indicator: ProgressIndicator) {
                val engine = project.getService(ProjectEngineService::class.java)
                val response = runCatching { engine.aidlReferences(request.toString()) }.getOrNull() ?: return
                count = mapper.readTree(response).path("result").path("references").size()
            }
            override fun onSuccess() { if (count == 0) Messages.showInfoMessage(project, "No AIDL usages were found.", "AOSP AIDL Navigation") else Messages.showInfoMessage(project, "$count AIDL usages found.", "AOSP AIDL Navigation") }
        }.queue()
    }
}
