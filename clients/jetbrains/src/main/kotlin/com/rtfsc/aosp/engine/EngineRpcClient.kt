package com.rtfsc.aosp.engine

import com.fasterxml.jackson.databind.ObjectMapper
import java.io.InputStream
import java.io.OutputStream
import java.nio.charset.StandardCharsets
import java.util.concurrent.atomic.AtomicLong

internal class EngineRpcClient(
    private val input: InputStream,
    private val output: OutputStream,
    private val objectMapper: ObjectMapper = ObjectMapper(),
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
        val response = JsonRpcFraming.read(input)
        validateResponse(id, response)
        return response
    }

    private fun validateResponse(id: Long, response: ByteArray) {
        val envelope = objectMapper.readTree(response)
        require(envelope.isObject) { "JSON-RPC response must be an object" }
        require(envelope.path("jsonrpc").asText() == "2.0") {
            "JSON-RPC response version is invalid"
        }
        require(envelope.path("id").asLong(Long.MIN_VALUE) == id) {
            "JSON-RPC response id does not match request $id"
        }
        if (envelope.has("error")) {
            throw EngineRpcException(envelope.path("error").toString())
        }
        require(envelope.has("result")) { "JSON-RPC response has no result" }
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

internal class EngineRpcException(message: String) : IllegalStateException(message)
