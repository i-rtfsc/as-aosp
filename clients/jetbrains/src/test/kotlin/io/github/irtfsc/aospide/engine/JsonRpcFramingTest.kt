package io.github.irtfsc.aospide.engine

import java.io.ByteArrayInputStream
import java.nio.charset.StandardCharsets
import org.junit.Assert.assertArrayEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class JsonRpcFramingTest {
    @Test
    fun encodesUtf8ByteLength() {
        val payload = "{\"message\":\"中文\"}".toByteArray(StandardCharsets.UTF_8)
        val encoded = JsonRpcFraming.encode(payload)
        assertArrayEquals(
            payload,
            encoded.copyOfRange(encoded.size - payload.size, encoded.size),
        )
        assertTrue(encoded.toString(StandardCharsets.US_ASCII).startsWith("Content-Length: 20"))
    }

    @Test
    fun readsConsecutiveFrames() {
        val first = JsonRpcFraming.encode("{\"a\":1}".toByteArray())
        val second = JsonRpcFraming.encode("{\"b\":2}".toByteArray())
        val input = ByteArrayInputStream(first + second)
        assertArrayEquals("{\"a\":1}".toByteArray(), JsonRpcFraming.read(input))
        assertArrayEquals("{\"b\":2}".toByteArray(), JsonRpcFraming.read(input))
    }

    @Test
    fun rejectsDuplicateContentLength() {
        val input = "Content-Length: 2\r\nContent-Length: 2\r\n\r\n{}"
            .toByteArray()
        try {
            JsonRpcFraming.read(ByteArrayInputStream(input))
            throw AssertionError("duplicate header was accepted")
        } catch (_: IllegalArgumentException) {
        }
    }
}
