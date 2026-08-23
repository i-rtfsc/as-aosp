package io.github.irtfsc.aospide.engine

import java.io.ByteArrayOutputStream
import java.io.EOFException
import java.io.InputStream
import java.nio.charset.StandardCharsets

internal object JsonRpcFraming {
    const val maxPayloadSize: Int = 16 * 1024 * 1024

    fun encode(payload: ByteArray): ByteArray {
        require(payload.size <= maxPayloadSize) {
            "JSON-RPC payload exceeds $maxPayloadSize bytes"
        }
        val header = "Content-Length: ${payload.size}\r\n\r\n"
            .toByteArray(StandardCharsets.US_ASCII)
        return header + payload
    }

    fun read(input: InputStream): ByteArray {
        var contentLength: Int? = null
        while (true) {
            val line = readHeaderLine(input)
            if (line.isEmpty()) break
            val separator = line.indexOf(':')
            require(separator > 0) { "malformed JSON-RPC header" }
            val name = line.substring(0, separator).trim().lowercase()
            val value = line.substring(separator + 1).trim()
            if (name != "content-length") continue
            require(contentLength == null) { "duplicate Content-Length header" }
            val length = value.toIntOrNull()
            require(length != null && length >= 0) { "invalid Content-Length" }
            require(length <= maxPayloadSize) {
                "JSON-RPC payload exceeds $maxPayloadSize bytes"
            }
            contentLength = length
        }
        val length = contentLength ?: error("missing Content-Length header")
        val payload = ByteArray(length)
        input.readFully(payload)
        return payload
    }

    private fun readHeaderLine(input: InputStream): String {
        val bytes = ByteArrayOutputStream()
        while (true) {
            val value = input.read()
            if (value < 0) {
                if (bytes.size() == 0) throw EOFException("EOF before JSON-RPC frame")
                throw EOFException("EOF in JSON-RPC header")
            }
            if (value == '\n'.code) {
                val data = bytes.toByteArray()
                val length = if (data.lastOrNull() == '\r'.code.toByte()) {
                    data.size - 1
                } else {
                    data.size
                }
                return data.copyOf(length).toString(StandardCharsets.US_ASCII)
            }
            bytes.write(value)
        }
    }

    private fun InputStream.readFully(target: ByteArray) {
        var offset = 0
        while (offset < target.size) {
            val count = read(target, offset, target.size - offset)
            if (count < 0) throw EOFException("EOF in JSON-RPC payload")
            offset += count
        }
    }
}
