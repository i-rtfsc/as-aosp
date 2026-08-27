package com.rtfsc.aosp.engine

import java.nio.file.Path
import java.nio.file.Files

internal object EngineCommand {
    val serveArguments: List<String> = listOf("serve", "--stdio")

    fun start(process: EngineProcess, executable: String): ByteArray {
        return process.start(Path.of(executable), serveArguments)
    }

    fun resolveExecutable(
        configured: String,
        userHome: String,
        bundled: Path? = null,
        osName: String = System.getProperty("os.name"),
    ): String {
        if (configured.isNotBlank()) return configured
        val executableName = if (normalizeOs(osName) == "windows") "aosp-index.exe" else "aosp-index"
        val shared = Path.of(userHome, "engines", executableName)
        if (Files.isRegularFile(shared)) return shared.toString()
        if (bundled != null && Files.isRegularFile(bundled)) return bundled.toString()
        return executableName
    }

    fun bundledCandidate(
        pluginRoot: Path,
        osName: String = System.getProperty("os.name"),
        architecture: String = System.getProperty("os.arch"),
    ): Path {
        val target = "${normalizeOs(osName)}-${normalizeArchitecture(architecture)}"
        val executableName = if (target.startsWith("windows-")) "aosp-index.exe" else "aosp-index"
        return pluginRoot.resolve("engine").resolve(target).resolve(executableName)
    }

    private fun normalizeOs(value: String): String = when {
        value.lowercase().contains("win") -> "windows"
        value.lowercase().contains("mac") || value.lowercase().contains("darwin") -> "darwin"
        else -> "linux"
    }

    private fun normalizeArchitecture(value: String): String = when (value.lowercase()) {
        "x86_64", "x64", "amd64" -> "amd64"
        "aarch64", "arm64" -> "arm64"
        else -> value.lowercase()
    }
}
