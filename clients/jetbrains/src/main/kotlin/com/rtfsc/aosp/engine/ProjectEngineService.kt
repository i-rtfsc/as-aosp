package com.rtfsc.aosp.engine

import com.intellij.openapi.Disposable
import com.intellij.openapi.components.Service
import com.intellij.openapi.project.Project
import com.rtfsc.aosp.profile.EngineProfileSettings

@Service(Service.Level.PROJECT)
internal class ProjectEngineService(
    private val project: Project,
) : Disposable {
    private val profileSettings = project.getService(EngineProfileSettings::class.java)
    private val engineProcess = EngineProcess()

    fun start(): ByteArray {
        val executable = profileSettings.current.engineExecutable
        val resolved = EngineCommand.resolveExecutable(executable, System.getProperty("user.home"))
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

    override fun dispose() {
        stop()
    }
}
