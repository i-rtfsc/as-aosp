package com.rtfsc.aosp.engine

import java.nio.file.Path

internal object EngineCommand {
    val serveArguments: List<String> = listOf("serve", "--stdio")

    fun start(process: EngineProcess, executable: String): ByteArray {
        return process.start(Path.of(executable), serveArguments)
    }
}
