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
}
