package com.rtfsc.aosp.engine

import com.intellij.openapi.Disposable
import com.intellij.openapi.components.Service
import com.intellij.openapi.project.Project
import com.intellij.ide.plugins.PluginManagerCore
import com.intellij.openapi.extensions.PluginId
import com.rtfsc.aosp.profile.EngineProfileSettings

@Service(Service.Level.PROJECT)
internal class ProjectEngineService(
    private val project: Project,
) : Disposable {
    private val profileSettings = project.getService(EngineProfileSettings::class.java)
    private val engineProcess = EngineProcess()

    fun start(): ByteArray {
        val executable = profileSettings.current.engineExecutable
        val pluginRoot = PluginManagerCore.getPlugin(PluginId.getId("com.rtfsc.aosp"))?.pluginPath
        val bundled = pluginRoot?.let(EngineCommand::bundledCandidate)
        val resolved = EngineCommand.resolveExecutable(executable, System.getProperty("user.home"), bundled)
        return EngineCommand.start(engineProcess, resolved)
    }

    fun stop() {
        engineProcess.stop()
    }

    fun jniPeers(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/jni/peers", params)
    }

    fun aidlResolve(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/aidl/resolve", params)
    }

    fun aidlPeers(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/aidl/peers", params)
    }

    fun aidlReferences(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/aidl/references", params)
    }

    fun projectGraph(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/projects/graph", params)
    }

    fun artifactsCatalog(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/artifacts/catalog", params)
    }

    fun modulesCatalog(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/modules/catalog", params)
    }

    fun sourceModules(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/source/modules", params)
    }

    fun selectionCatalog(): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/selection/catalog", "null")
    }

    fun selectionPlan(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/selection/plan", params)
    }

    fun generatedMaterialize(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/generated/materialize", params)
    }

    fun indexStart(params: String, rebuild: Boolean): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request(if (rebuild) "aosp/index/rebuild" else "aosp/index/start", params)
    }

    fun indexProjection(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/index/projection", params)
    }

    fun compdbOverlay(params: String): ByteArray {
        if (!engineProcess.isRunning) start()
        return engineProcess.request("aosp/compdb/overlay", params)
    }

    override fun dispose() {
        stop()
    }
}
