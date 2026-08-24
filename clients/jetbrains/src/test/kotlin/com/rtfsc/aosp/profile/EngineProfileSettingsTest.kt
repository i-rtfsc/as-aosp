package com.rtfsc.aosp.profile

import org.junit.Assert.assertEquals
import org.junit.Test

class EngineProfileSettingsTest {
    @Test
    fun loadsAndUpdatesProjectLocalPathsWithoutNormalization() {
        val settings = EngineProfileSettings()
        settings.loadState(
            EngineProfileState(
                profileId = "framework",
                workspaceViewRoot = "C:\\aosp view",
                engineExecutable = "C:\\tools\\aosp-index.exe",
            ),
        )

        settings.update { it.outRoot = "Z:\\mounted out" }

        assertEquals("C:\\aosp view", settings.current.workspaceViewRoot)
        assertEquals("C:\\tools\\aosp-index.exe", settings.current.engineExecutable)
        assertEquals("Z:\\mounted out", settings.current.outRoot)
    }
}
