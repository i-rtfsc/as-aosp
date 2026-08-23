package io.github.irtfsc.aospide.engine

import com.intellij.openapi.Disposable
import com.intellij.openapi.components.Service
import com.intellij.openapi.project.Project
import io.github.irtfsc.aospide.profile.EngineProfileSettings
import java.nio.file.Path

@Service(Service.Level.PROJECT)
internal class ProjectEngineService(
    private val project: Project,
) : Disposable {
    private val profileSettings = project.getService(EngineProfileSettings::class.java)
    private val engineProcess = EngineProcess()

    fun start(): ByteArray {
        val executable = profileSettings.current.engineExecutable
        require(executable.isNotBlank()) { "engine executable is not configured" }
        return engineProcess.start(Path.of(executable))
    }

    fun stop() {
        engineProcess.stop()
    }

    override fun dispose() {
        stop()
    }
}
