package com.rtfsc.aosp.project

import com.rtfsc.aosp.profile.EngineProfileState
import java.nio.file.Paths
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class WorkspaceRootPlanTest {
    @Test
    fun defaultsToWorkspaceRootAndFixedExcludes() {
        val plan = WorkspaceRootPlanner.plan(
            EngineProfileState(workspaceViewRoot = "/workspace/aosp"),
        )

        assertEquals(listOf(Paths.get("/workspace/aosp")), plan.sourceRoots)
        assertTrue(plan.excludedRoots.contains(Paths.get("/workspace/aosp/.repo")))
        assertTrue(plan.excludedRoots.contains(Paths.get("/workspace/aosp/out")))
    }

    @Test
    fun projectsExplicitAliasesWithoutFollowingLinks() {
        val plan = WorkspaceRootPlanner.plan(
            EngineProfileState(
                workspaceViewRoot = "/workspace/view",
                focusedSourcePaths = "frameworks/base, ui",
            ),
        )

        assertEquals(
            listOf(Paths.get("/workspace/view/frameworks/base"), Paths.get("/workspace/view/ui")),
            plan.sourceRoots,
        )
        assertTrue(plan.isUsable)
    }

    @Test
    fun rejectsAbsoluteAndEscapingFocusedPaths() {
        val plan = WorkspaceRootPlanner.plan(
            EngineProfileState(
                workspaceViewRoot = "/workspace/view",
                focusedSourcePaths = "/etc,../outside,C:\\aosp",
            ),
        )

        assertEquals(3, plan.diagnostics.size)
        assertTrue(plan.sourceRoots.isEmpty())
    }

    @Test
    fun canonicalRootIsNeverAcceptedAsContentRoot() {
        val plan = WorkspaceRootPlanner.plan(
            EngineProfileState(
                workspaceViewRoot = "/workspace/aosp",
                canonicalRepoRoot = "/workspace/aosp",
            ),
        )

        assertTrue(plan.diagnostics.any { it.contains("Canonical AOSP Root") })
        assertTrue(!plan.isUsable)
    }

    @Test
    fun excludesCanonicalRootWhenItIsNestedInTheView() {
        val plan = WorkspaceRootPlanner.plan(
            EngineProfileState(
                workspaceViewRoot = "/workspace",
                canonicalRepoRoot = "/workspace/full-aosp",
            ),
        )

        assertTrue(plan.excludedRoots.contains(Paths.get("/workspace/full-aosp")))
    }
}
