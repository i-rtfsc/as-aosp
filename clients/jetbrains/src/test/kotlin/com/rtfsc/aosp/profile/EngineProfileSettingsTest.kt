package com.rtfsc.aosp.profile

import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.Assert.assertTrue

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

    @Test
    fun notifiesListenersAfterProjectStateChanges() {
        val settings = EngineProfileSettings()
        val roots = mutableListOf<String>()
        val subscription = settings.addListener { roots += it.workspaceViewRoot }

        settings.update { it.workspaceViewRoot = "/workspace/aosp" }
        subscription.close()
        settings.update { it.workspaceViewRoot = "/workspace/other" }

        assertEquals(listOf("", "/workspace/aosp"), roots)
        assertTrue(settings.current.workspaceViewRoot.endsWith("other"))
    }
}
