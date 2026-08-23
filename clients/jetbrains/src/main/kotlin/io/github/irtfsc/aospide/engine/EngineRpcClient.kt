package io.github.irtfsc.aospide.engine

import java.io.InputStream
import java.io.OutputStream
import java.nio.charset.StandardCharsets
import java.util.concurrent.atomic.AtomicLong

internal class EngineRpcClient(
    private val input: InputStream,
    private val output: OutputStream,
) {
    private val nextId = AtomicLong(1)

    fun initialize(protocolVersion: String = "1.0"): ByteArray {
        return request(
            method = "aosp/initialize",
            params = "{\"protocolVersion\":${quote(protocolVersion)}}",
        )
    }

    fun shutdown(): ByteArray {
        return request(method = "aosp/shutdown", params = "null")
    }

    fun request(method: String, params: String): ByteArray {
        val id = nextId.getAndIncrement()
        val payload = "{" +
            "\"jsonrpc\":\"2.0\"," +
            "\"id\":$id," +
            "\"method\":${quote(method)}," +
            "\"params\":$params" +
            "}"
        output.write(JsonRpcFraming.encode(payload.toByteArray(StandardCharsets.UTF_8)))
        output.flush()
        return JsonRpcFraming.read(input)
    }

    private fun quote(value: String): String {
        return buildString {
            append('"')
            value.forEach { character ->
                when (character) {
                    '\\' -> append("\\\\")
                    '"' -> append("\\\"")
                    '\n' -> append("\\n")
                    '\r' -> append("\\r")
                    '\t' -> append("\\t")
                    else -> append(character)
                }
            }
            append('"')
        }
    }
}
