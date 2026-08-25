package com.rtfsc.aosp.engine

import java.nio.file.Path
import java.nio.file.Files

internal object EngineCommand {
    val serveArguments: List<String> = listOf("serve", "--stdio")

    fun start(process: EngineProcess, executable: String): ByteArray {
        return process.start(Path.of(executable), serveArguments)
    }

    fun resolveExecutable(configured: String, userHome: String, bundled: Path? = null): String {
        if (configured.isNotBlank()) return configured
        val shared = Path.of(userHome, "engines", "aosp-index")
        if (Files.isRegularFile(shared)) return shared.toString()
        if (bundled != null && Files.isRegularFile(bundled)) return bundled.toString()
        return "aosp-index"
    }
}
