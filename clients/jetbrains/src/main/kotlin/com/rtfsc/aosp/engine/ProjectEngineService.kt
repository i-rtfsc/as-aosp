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
        require(executable.isNotBlank()) { "engine executable is not configured" }
        return EngineCommand.start(engineProcess, executable)
    }

    fun stop() {
        engineProcess.stop()
    }

    override fun dispose() {
        stop()
    }
}
