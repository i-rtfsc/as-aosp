package com.rtfsc.aosp.engine

import java.io.BufferedReader
import java.io.InputStream
import java.io.InputStreamReader
import java.nio.file.Path
import java.util.concurrent.Callable
import java.util.concurrent.Executors
import java.util.concurrent.ExecutionException
import java.util.concurrent.TimeUnit
import kotlin.concurrent.thread

internal class EngineProcess(
    private val processFactory: ProcessFactory = JvmProcessFactory(),
    private val stderrSink: (String) -> Unit = {},
) {
    private var managedProcess: ManagedProcess? = null
    private var rpcClient: EngineRpcClient? = null

    val isRunning: Boolean
        get() = managedProcess?.isAlive == true

    @Synchronized
    fun start(
        executable: Path,
        arguments: List<String> = emptyList(),
        timeoutMillis: Long = 5_000,
    ): ByteArray {
        check(managedProcess == null) { "engine process is already started" }
        val command = buildList {
            add(executable.toString())
            addAll(arguments)
        }
        val child = processFactory.start(command)
        drainStderr(child.errorStream)
        val client = EngineRpcClient(child.inputStream, child.outputStream)
        try {
            val response = callWithTimeout(timeoutMillis) { client.initialize() }
            managedProcess = child
            rpcClient = client
            return response
        } catch (failure: Throwable) {
            child.destroyForcibly()
            throw failure
        }
    }

    @Synchronized
    fun stop(timeoutMillis: Long = 2_000) {
        val child = managedProcess ?: return
        try {
            rpcClient?.let { client ->
                callWithTimeout(timeoutMillis) { client.shutdown() }
            }
        } catch (_: Throwable) {
            // The process may already have closed its protocol streams.
        }
        if (!child.waitFor(timeoutMillis, TimeUnit.MILLISECONDS) && child.isAlive) {
            child.destroy()
            if (!child.waitFor(timeoutMillis, TimeUnit.MILLISECONDS) && child.isAlive) {
                child.destroyForcibly()
            }
        }
        managedProcess = null
        rpcClient = null
    }

    @Synchronized
    fun request(method: String, params: String, timeoutMillis: Long = 5_000): ByteArray {
        val client = rpcClient ?: error("engine process is not running")
        return callWithTimeout(timeoutMillis) { client.request(method, params) }
    }

    @Synchronized
    fun restart(
        executable: Path,
        arguments: List<String> = emptyList(),
        timeoutMillis: Long = 2_000,
    ): ByteArray {
        stop(timeoutMillis)
        return start(executable, arguments, timeoutMillis)
    }

    private fun drainStderr(stream: InputStream) {
        thread(start = true, isDaemon = true, name = "aosp-indexd-stderr") {
            BufferedReader(InputStreamReader(stream, Charsets.UTF_8)).useLines { lines ->
                lines.forEach(stderrSink)
            }
        }
    }

    private fun <T> callWithTimeout(timeoutMillis: Long, operation: () -> T): T {
        require(timeoutMillis > 0) { "engine timeout must be positive" }
        val executor = Executors.newSingleThreadExecutor { runnable ->
            Thread(runnable, "aosp-indexd-rpc").apply { isDaemon = true }
        }
        return try {
            executor.submit(Callable(operation)).get(timeoutMillis, TimeUnit.MILLISECONDS)
        } catch (failure: ExecutionException) {
            val cause = failure.cause
            if (cause is RuntimeException) throw cause
            if (cause is Error) throw cause
            throw failure
        } finally {
            executor.shutdownNow()
        }
    }
}
