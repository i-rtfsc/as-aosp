package com.rtfsc.aosp.profile

import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class TargetDraftCodecTest {
    @Test
    fun roundTripsArtifactAndProjectionFlags() {
        val input = listOf(
            TargetDraft(
                id = "systemui",
                kind = "android",
                artifactKind = "apk",
                artifact = "SystemUI.apk",
                module = "SystemUI",
                sourceRoots = listOf("frameworks/base/packages/SystemUI"),
                includeDependencies = true,
                includeGeneratedSources = false,
                includeResources = true,
                includeManifest = true,
            ),
        )
        val output = TargetDraftCodec.decode(TargetDraftCodec.encode(input))
        assertEquals(input, output)
        assertTrue(output.single().toTargetSpec().containsKey("selectors"))
    }

    @Test
    fun invalidJsonProducesEmptyDrafts() {
        assertTrue(TargetDraftCodec.decode("not-json").isEmpty())
        assertTrue(TargetDraftCodec.decode("{}").isEmpty())
    }
}
