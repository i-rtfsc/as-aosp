package com.rtfsc.aosp.engine

import org.junit.Assert.assertEquals
import org.junit.Test

class EngineCommandTest {
    @Test
    fun usesTheUnifiedServeSubcommand() {
        assertEquals(listOf("serve", "--stdio"), EngineCommand.serveArguments)
    }
}
