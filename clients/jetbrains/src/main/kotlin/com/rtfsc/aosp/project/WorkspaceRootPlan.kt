package com.rtfsc.aosp.project

import com.rtfsc.aosp.profile.EngineProfileState
import java.nio.file.Path
import java.nio.file.Paths

internal data class WorkspaceRootPlan(
    val contentRoot: Path,
    val sourceRoots: List<Path>,
    val excludedRoots: List<Path>,
    val diagnostics: List<String>,
) {
    val isUsable: Boolean
        get() = diagnostics.isEmpty() && sourceRoots.isNotEmpty()
}

internal object WorkspaceRootPlanner {
    private val fixedExcludes = listOf(".git", ".repo", "out", ".aosp-ide")
    private val windowsAbsolute = Regex("^[A-Za-z]:[\\\\/].*")

    fun plan(state: EngineProfileState): WorkspaceRootPlan {
        val diagnostics = mutableListOf<String>()
        val contentRoot = parseAbsolute(state.workspaceViewRoot, "Workspace View Root", diagnostics)
            ?: return WorkspaceRootPlan(
                Paths.get("."), emptyList(), emptyList(), diagnostics,
            )
        val sourceRoots = parseFocusedRoots(contentRoot, state.focusedSourcePaths, diagnostics)
        val excludedRoots = fixedExcludes.map { contentRoot.resolve(it).normalize() }.toMutableList()
        val canonical = state.canonicalRepoRoot.trim()
        if (canonical.isNotEmpty()) {
            val canonicalPath = parseAbsolute(canonical, "Canonical AOSP Root", diagnostics)
            if (canonicalPath != null && canonicalPath == contentRoot) {
                diagnostics += "Canonical AOSP Root must not be the Workspace View Root."
            }
            if (canonicalPath != null && canonicalPath.startsWith(contentRoot)) {
                excludedRoots.add(canonicalPath)
            }
        }
        return WorkspaceRootPlan(contentRoot, sourceRoots, excludedRoots, diagnostics)
    }

    private fun parseFocusedRoots(
        contentRoot: Path,
        raw: String,
        diagnostics: MutableList<String>,
    ): List<Path> {
        val values = raw.split(',').map(String::trim).filter(String::isNotEmpty)
        if (values.isEmpty()) return listOf(contentRoot)
        return values.distinct().mapNotNull { value ->
            if (value.startsWith("/") || windowsAbsolute.matches(value)) {
                diagnostics += "Focused source path must be relative: $value"
                return@mapNotNull null
            }
            val candidate = contentRoot.resolve(value.replace('\\', '/')).normalize()
            if (!candidate.startsWith(contentRoot)) {
                diagnostics += "Focused source path escapes Workspace View Root: $value"
                return@mapNotNull null
            }
            candidate
        }
    }

    private fun parseAbsolute(
        value: String,
        label: String,
        diagnostics: MutableList<String>,
    ): Path? {
        if (value.isBlank()) {
            diagnostics += "$label is required."
            return null
        }
        val path = try {
            Paths.get(value)
        } catch (_: RuntimeException) {
            diagnostics += "$label is invalid: $value"
            return null
        }
        if (!path.isAbsolute) {
            diagnostics += "$label must be absolute: $value"
            return null
        }
        return path.normalize()
    }
}
