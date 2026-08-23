package io.github.irtfsc.aospide.engine

import java.io.InputStream
import java.io.OutputStream
import java.util.concurrent.TimeUnit

internal interface ManagedProcess {
    val inputStream: InputStream
    val outputStream: OutputStream
    val errorStream: InputStream
    val isAlive: Boolean

    fun waitFor(timeout: Long, unit: TimeUnit): Boolean
    fun destroy()
    fun destroyForcibly()
}

internal fun interface ProcessFactory {
    fun start(command: List<String>): ManagedProcess
}

internal class JvmProcessFactory : ProcessFactory {
    override fun start(command: List<String>): ManagedProcess {
        require(command.isNotEmpty()) { "engine command is empty" }
        return JvmManagedProcess(ProcessBuilder(command).start())
    }
}

private class JvmManagedProcess(
    private val process: Process,
) : ManagedProcess {
    override val inputStream: InputStream
        get() = process.inputStream

    override val outputStream: OutputStream
        get() = process.outputStream

    override val errorStream: InputStream
        get() = process.errorStream

    override val isAlive: Boolean
        get() = process.isAlive

    override fun waitFor(timeout: Long, unit: TimeUnit): Boolean {
        return process.waitFor(timeout, unit)
    }

    override fun destroy() {
        process.destroy()
    }

    override fun destroyForcibly() {
        process.destroyForcibly()
    }
}
