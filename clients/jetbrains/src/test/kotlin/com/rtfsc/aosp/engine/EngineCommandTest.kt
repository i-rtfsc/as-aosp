package com.rtfsc.aosp.engine

import org.junit.Assert.assertEquals
import org.junit.Test
import java.nio.file.Files

class EngineCommandTest {
    @Test
    fun usesTheUnifiedServeSubcommand() {
        assertEquals(listOf("serve", "--stdio"), EngineCommand.serveArguments)
    }

    @Test
    fun prefersConfiguredThenSharedEngineEntry() {
        val home = Files.createTempDirectory("aosp-engine-home")
        val shared = home.resolve("engines/aosp-index")
        Files.createDirectories(shared.parent)
        Files.writeString(shared, "engine")
        assertEquals("/custom/aosp-index", EngineCommand.resolveExecutable("/custom/aosp-index", home.toString()))
        assertEquals(shared.toString(), EngineCommand.resolveExecutable("", home.toString()))
    }

    @Test
    fun resolvesBundledWindowsAndMacEnginePaths() {
        val root = Files.createTempDirectory("aosp-plugin")
        assertEquals(
            root.resolve("engine/windows-amd64/aosp-index.exe"),
            EngineCommand.bundledCandidate(root, "Windows 11", "amd64"),
        )
        assertEquals(
            root.resolve("engine/darwin-arm64/aosp-index"),
            EngineCommand.bundledCandidate(root, "Mac OS X", "aarch64"),
        )
        assertEquals(
            "aosp-index.exe",
            EngineCommand.resolveExecutable("", root.resolve("missing-home").toString(), osName = "Windows 11"),
        )
    }
}
