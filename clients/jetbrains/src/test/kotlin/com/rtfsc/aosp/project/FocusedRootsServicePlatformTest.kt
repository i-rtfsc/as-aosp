package com.rtfsc.aosp.project

import com.intellij.openapi.roots.ModuleRootManager
import com.intellij.testFramework.fixtures.BasePlatformTestCase
import com.rtfsc.aosp.profile.EngineProfileSettings
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import java.nio.file.Files
import com.intellij.openapi.application.WriteAction

class FocusedRootsServicePlatformTest : BasePlatformTestCase() {
    fun testAppliesNonPersistentFocusedModule() {
        val viewRoot = Files.createTempDirectory("aosp-focused-view")
        val settings = project.getService(EngineProfileSettings::class.java)
        settings.update {
            it.profileId = "platform-test"
            it.workspaceViewRoot = viewRoot.toString()
        }

        val plan = WorkspaceRootPlanner.plan(settings.current)
        FocusedRootsApplier.apply(getModule(), plan)

        assertTrue(plan.isUsable)
        val contentEntries = ModuleRootManager.getInstance(getModule()).contentEntries
        assertEquals(1, contentEntries.size)
        val viewEntry = contentEntries.single()
        assertTrue(viewEntry.url.endsWith(viewRoot.fileName.toString()))
        assertTrue(viewEntry.sourceFolders.any { it.url.endsWith(viewRoot.fileName.toString()) })
        assertTrue(viewEntry.excludeFolderUrls.any { it.endsWith("/.git") })
        assertTrue(viewEntry.excludeFolderUrls.any { it.endsWith("/.repo") })
        assertTrue(viewEntry.excludeFolderUrls.any { it.endsWith("/out") })
    }
}
