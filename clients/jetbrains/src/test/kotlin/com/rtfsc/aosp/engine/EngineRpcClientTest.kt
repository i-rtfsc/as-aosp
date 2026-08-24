package com.rtfsc.aosp.engine

import java.io.ByteArrayInputStream
import java.io.ByteArrayOutputStream
import org.junit.Assert.assertTrue
import org.junit.Test

class EngineRpcClientTest {
    @Test
    fun initializeUsesVersionedJsonRpcRequest() {
        val output = ByteArrayOutputStream()
        val response = JsonRpcFraming.encode(
            "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":null}".toByteArray(),
        )
        val client = EngineRpcClient(ByteArrayInputStream(response), output)

        client.initialize()

        val request = output.toByteArray().decodeToString()
        assertTrue(request.contains("\"method\":\"aosp/initialize\""))
        assertTrue(request.contains("\"protocolVersion\":\"1.0\""))
    }

    @Test
    fun rejectsMismatchedResponseId() {
        val response = JsonRpcFraming.encode(
            "{\"jsonrpc\":\"2.0\",\"id\":2,\"result\":null}".toByteArray(),
        )
        val client = EngineRpcClient(
            ByteArrayInputStream(response),
            ByteArrayOutputStream(),
        )

        try {
            client.initialize()
            throw AssertionError("mismatched response id was accepted")
        } catch (_: IllegalArgumentException) {
        }
    }
}
