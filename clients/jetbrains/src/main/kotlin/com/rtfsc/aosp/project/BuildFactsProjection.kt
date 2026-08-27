package com.rtfsc.aosp.project

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.application.WriteAction
import com.intellij.openapi.module.Module
import com.intellij.openapi.roots.ModuleRootManager
import com.intellij.openapi.roots.OrderRootType
import com.intellij.openapi.vfs.LocalFileSystem
import com.rtfsc.aosp.engine.ProjectEngineService
import com.rtfsc.aosp.profile.EngineProfileState
import java.nio.file.Files
import java.nio.file.Path
import java.nio.file.Paths

internal data class BuildFactsProjectionResult(
    val sourceRoots: List<Path>,
    val generatedRoots: List<Path>,
    val libraries: List<Path>,
    val skipped: Int,
    val diagnostics: List<String>,
)

internal object BuildFactsProjection {
    private val mapper = ObjectMapper()

    fun apply(
        module: Module,
        state: EngineProfileState,
        engine: ProjectEngineService,
        cacheRoot: Path,
    ): BuildFactsProjectionResult {
        val packRoot = state.metadataPackRoot.trim()
        val metadataRoot = state.outRoot.trim()
        if (packRoot.isEmpty() && metadataRoot.isEmpty()) {
            return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), 0, listOf("Metadata Pack Root or Out Root is required."))
        }
        if (packRoot.isNotEmpty()) {
            val validation = mapper.readTree(engine.packValidate(mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot)))).path("result")
            if (!validation.path("valid").asBoolean(false)) {
                val diagnostics = validation.path("diagnostics").map { node ->
                    "${node.path("code").asText("PACK_INVALID")}: ${node.path("message").asText("invalid pack")}"
                }
                return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), 0, diagnostics)
            }
        }
        val packInfo = if (packRoot.isNotEmpty()) {
            val params = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot))
            mapper.readTree(engine.packInfo(params)).path("result")
        } else mapper.createObjectNode()
        val diagnostics = mutableListOf<String>()
        val targets = runCatching { mapper.readTree(state.targetsJson) }.getOrNull()
        if (targets == null || !targets.isArray || targets.size() == 0) {
            diagnostics += "PROJECT_GRAPH_TARGETS_MISSING: Targets JSON must contain at least one target."
            return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), 0, diagnostics)
        }
        val graphResponse = mapper.readTree(engine.projectGraph(mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot, "metadataRoot" to metadataRoot, "targets" to targets))))
        collectDiagnostics(graphResponse.path("result").path("diagnostics"), diagnostics)
        val graph = graphResponse.path("result").path("graph")
        val sourceRoots = linkedSetOf<Path>()
        val libraries = linkedSetOf<Path>()
        val skipped = intArrayOf(0)
        val workspaceRoot = Paths.get(state.workspaceViewRoot).toAbsolutePath().normalize()
        val graphItems = graph.path("targets")
        val moduleItems = graphItems
        for (item in moduleItems) {
            collectRepoRoots(item.path("sourceRoots"), workspaceRoot, sourceRoots, skipped)
            collectRepoRoots(item.path("sources"), workspaceRoot, sourceRoots, skipped)
            val jars = item.path("libraries")
            collectLibraries(jars, workspaceRoot, packRoot, packInfo.path("outRoot").asText(), state.outRoot, libraries, skipped)
        }
        val generatedRoots = if (packRoot.isNotEmpty()) {
            val generatedParams = mapper.writeValueAsString(
                mapOf("metadataPackRoot" to packRoot, "cacheRoot" to cacheRoot.toString()),
            )
            val generated = mapper.readTree(engine.generatedMaterialize(generatedParams)).path("result")
            collectDiagnostics(generated.path("diagnostics"), diagnostics)
            generated.path("sourceRoots").mapNotNull { node ->
                node.path("root").asText().takeIf { it.isNotBlank() }?.let(Paths::get)
            }.filter { Files.isDirectory(it) }
        } else emptyList()
        sourceRoots.addAll(generatedRoots)
        applyProjection(module, sourceRoots.toList(), libraries.toList())
        return BuildFactsProjectionResult(sourceRoots.toList(), generatedRoots, libraries.toList(), skipped[0], diagnostics)
    }

    private fun collectRepoRoots(
        refs: JsonNode,
        workspaceRoot: Path,
        roots: MutableSet<Path>,
        skipped: IntArray,
    ) {
        for (ref in refs) {
            if (ref.path("kind").asText() != "repo-relative") {
                skipped[0]++
                continue
            }
            val repoPath = ref.path("repoPath").asText()
            if (repoPath.isBlank()) {
                skipped[0]++
                continue
            }
            val candidate = workspaceRoot.resolve(repoPath.replace('/', java.io.File.separatorChar)).normalize()
            if (candidate.startsWith(workspaceRoot)) roots.add(candidate) else skipped[0]++
        }
    }

    private fun collectLibraries(
        refs: JsonNode,
        workspaceRoot: Path,
        packRoot: String,
        outRoot: String,
        packOutRoot: String,
        libraries: MutableSet<Path>,
        skipped: IntArray,
    ) {
        for (ref in refs) {
            val kind = ref.path("kind").asText()
            val raw = ref.path("raw").asText()
            val candidate = when (kind) {
                "repo-relative" -> workspaceRoot.resolve(ref.path("repoPath").asText().replace('/', java.io.File.separatorChar))
                "generated" -> packRoot.takeIf(String::isNotBlank)
                    ?.let(Paths::get)?.resolve(raw.replace('/', java.io.File.separatorChar))
                "external-absolute" -> Paths.get(raw)
                else -> null
            }?.normalize()
            val relocated = if (outRoot.isNotBlank() && packOutRoot.isNotBlank() &&
                raw.startsWith(packOutRoot.trimEnd('/', '\\') + java.io.File.separator) &&
                !Files.isRegularFile(candidate)) {
                val suffix = raw.removePrefix(packOutRoot.trimEnd('/', '\\') + java.io.File.separator)
                Paths.get(outRoot).resolve(suffix.replace('/', java.io.File.separatorChar)).normalize()
            } else candidate
            val packed = if (packOutRoot.isNotBlank() && raw.startsWith(packOutRoot.trimEnd('/', '\\') + java.io.File.separator)) {
                val suffix = raw.removePrefix(packOutRoot.trimEnd('/', '\\') + java.io.File.separator)
                Paths.get(packRoot).resolve("out").resolve(suffix.replace('/', java.io.File.separatorChar)).normalize()
            } else null
            val localOut = if (outRoot.isNotBlank() && !Paths.get(raw).isAbsolute) {
                val normalized = raw.replace('\\', '/')
                val suffix = normalized.removePrefix("out/")
                Paths.get(outRoot).resolve(suffix.replace('/', java.io.File.separatorChar)).normalize()
            } else null
            val selected = when {
                relocated != null && Files.isRegularFile(relocated) -> relocated
                packed != null && Files.isRegularFile(packed) -> packed
                localOut != null && Files.isRegularFile(localOut) -> localOut
                else -> null
            }
            if (selected != null && selected.toString().endsWith(".jar")) {
                libraries.add(selected)
            } else if (outRoot.isNotBlank() || raw.isNotBlank()) {
                skipped[0]++
            }
        }
    }

    private fun applyProjection(module: Module, roots: List<Path>, libraries: List<Path>) {
        WriteAction.compute<Unit, RuntimeException> {
            val model = ModuleRootManager.getInstance(module).modifiableModel
            model.clear()
            roots.distinct().forEach { root ->
                val content = model.addContentEntry(root.toString())
                content.addSourceFolder(root.toString(), false)
            }
            val libraryTable = model.moduleLibraryTable
            libraries.distinct().forEach { libraryPath ->
                val library = libraryTable.createLibrary("AOSP ${libraryPath.fileName}")
                val libraryModel = library.modifiableModel
                val file = LocalFileSystem.getInstance().refreshAndFindFileByPath(libraryPath.toString())
                if (file != null) {
                    libraryModel.addRoot(file, OrderRootType.CLASSES)
                }
                libraryModel.commit()
            }
            model.commit()
        }
    }

    private fun collectDiagnostics(nodes: JsonNode, destination: MutableList<String>) {
        for (node in nodes) {
            val code = node.path("code").asText("UNKNOWN")
            val message = node.path("message").asText("unknown diagnostic")
            destination += "$code: $message"
        }
    }
}
