package com.rtfsc.aosp.engine

import java.io.ByteArrayInputStream
import java.io.ByteArrayOutputStream
import java.nio.file.Path
import java.util.concurrent.TimeUnit
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Test

class EngineProcessTest {
    @Test
    fun startsWithAnArgumentListAndInitializesOnce() {
        val factory = RecordingProcessFactory()
        val process = EngineProcess(factory)

        val response = process.start(
            Path.of("/opt/aosp-index"),
            listOf("--stdio", "path with spaces"),
        )

        assertTrue(response.decodeToString().contains("\"result\":{}"))
        assertEquals(
            listOf("/opt/aosp-index", "--stdio", "path with spaces"),
            factory.command,
        )
        assertTrue(process.isRunning)
        assertEquals(1, factory.process.requestCount("aosp/initialize"))
    }

    @Test
    fun stopSendsShutdownAndClearsSession() {
        val factory = RecordingProcessFactory()
        val process = EngineProcess(factory)
        process.start(Path.of("aosp-index"))

        process.stop()

        assertEquals(1, factory.process.requestCount("aosp/shutdown"))
        assertFalse(process.isRunning)
        assertTrue(factory.process.destroyed)
    }

    @Test
    fun initializeErrorTerminatesTheChild() {
        val factory = RecordingProcessFactory(initializeError = true)
        val process = EngineProcess(factory)

        try {
        process.start(Path.of("aosp-index"))
            throw AssertionError("initialize error was accepted")
        } catch (_: EngineRpcException) {
        }

        assertFalse(process.isRunning)
        assertTrue(factory.process.destroyedForcibly)
    }
}

private class RecordingProcessFactory(
    private val initializeError: Boolean = false,
) : ProcessFactory {
    lateinit var command: List<String>
    lateinit var process: RecordingProcess

    override fun start(command: List<String>): ManagedProcess {
        this.command = command
        process = RecordingProcess(initializeError)
        return process
    }
}

private class RecordingProcess(
    initializeError: Boolean,
) : ManagedProcess {
    private val initialize = if (initializeError) {
        "{\"jsonrpc\":\"2.0\",\"id\":1,\"error\":{\"code\":-32600}}"
    } else {
        "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{}}"
    }
    private val responses =
        JsonRpcFraming.encode(initialize.toByteArray()) +
            JsonRpcFraming.encode(
                "{\"jsonrpc\":\"2.0\",\"id\":2,\"result\":null}".toByteArray(),
            )
    override val inputStream = ByteArrayInputStream(responses)
    val output = ByteArrayOutputStream()
    override val outputStream = output
    override val errorStream = ByteArrayInputStream(ByteArray(0))
    override var isAlive = true
    var destroyed = false
    var destroyedForcibly = false

    override fun waitFor(timeout: Long, unit: TimeUnit): Boolean {
        return !isAlive
    }

    override fun destroy() {
        destroyed = true
        isAlive = false
    }

    override fun destroyForcibly() {
        destroyed = true
        destroyedForcibly = true
        isAlive = false
    }

    fun requestCount(method: String): Int {
        return output.toString(Charsets.UTF_8).windowed(method.length).count { it == method }
    }
}
