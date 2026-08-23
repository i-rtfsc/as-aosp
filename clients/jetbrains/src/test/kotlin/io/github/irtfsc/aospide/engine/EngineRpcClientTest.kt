package io.github.irtfsc.aospide.engine

import java.io.ByteArrayInputStream
import java.io.ByteArrayOutputStream
import org.junit.Assert.assertTrue
import org.junit.Test

class EngineRpcClientTest {
    @Test
    fun initializeUsesVersionedJsonRpcRequest() {
        val output = ByteArrayOutputStream()
        val response = JsonRpcFraming.encode("{\"result\":null}".toByteArray())
        val client = EngineRpcClient(ByteArrayInputStream(response), output)

        client.initialize()

        val request = output.toByteArray().decodeToString()
        assertTrue(request.contains("\"method\":\"aosp/initialize\""))
        assertTrue(request.contains("\"protocolVersion\":\"1.0\""))
    }
}
