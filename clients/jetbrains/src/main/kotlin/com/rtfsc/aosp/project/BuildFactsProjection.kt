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
        if (packRoot.isEmpty()) {
            return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), 0, listOf("Metadata Pack Root is required."))
        }
        val validation = mapper.readTree(engine.packValidate(mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot)))).path("result")
        if (!validation.path("valid").asBoolean(false)) {
            val diagnostics = validation.path("diagnostics").map { node ->
                "${node.path("code").asText("PACK_INVALID")}: ${node.path("message").asText("invalid pack")}"
            }
            return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), 0, diagnostics)
        }
        val packInfoParams = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot))
        val packInfo = mapper.readTree(engine.packInfo(packInfoParams)).path("result")
        val catalogParams = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot))
        val catalog = mapper.readTree(engine.modulesCatalog(catalogParams)).path("result")
        val diagnostics = mutableListOf<String>()
        collectDiagnostics(catalog.path("diagnostics"), diagnostics)
        val names = catalog.path("modules").map { it.path("name").asText() }.filter(String::isNotBlank)
        val params = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot, "names" to names))
        val closure = mapper.readTree(engine.modulesClosure(params)).path("result")
        collectDiagnostics(closure.path("diagnostics"), diagnostics)
        val sourceRoots = linkedSetOf<Path>()
        val libraries = linkedSetOf<Path>()
        val skipped = intArrayOf(0)
        val workspaceRoot = Paths.get(state.workspaceViewRoot).toAbsolutePath().normalize()
        for (item in closure.path("modules")) {
            collectRepoRoots(item.path("sourceRoots"), workspaceRoot, sourceRoots, skipped)
            collectRepoRoots(item.path("sources"), workspaceRoot, sourceRoots, skipped)
            collectLibraries(item.path("jars"), workspaceRoot, packRoot, packInfo.path("outRoot").asText(), state.outRoot, libraries, skipped)
        }
        val generatedParams = mapper.writeValueAsString(
            mapOf("metadataPackRoot" to packRoot, "cacheRoot" to cacheRoot.toString()),
        )
        val generated = mapper.readTree(engine.generatedMaterialize(generatedParams)).path("result")
        collectDiagnostics(generated.path("diagnostics"), diagnostics)
        val generatedRoots = generated.path("sourceRoots").mapNotNull { node ->
            node.path("root").asText().takeIf { it.isNotBlank() }?.let(Paths::get)
        }.filter { Files.isDirectory(it) }
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
                "generated" -> Paths.get(packRoot).resolve(raw.replace('/', java.io.File.separatorChar))
                "external-absolute" -> Paths.get(raw)
                else -> null
            }?.normalize()
            val relocated = if (outRoot.isNotBlank() && packOutRoot.isNotBlank() &&
                raw.startsWith(packOutRoot.trimEnd('/', '\\') + java.io.File.separator) &&
                !Files.isRegularFile(candidate)) {
                val suffix = raw.removePrefix(packOutRoot.trimEnd('/', '\\') + java.io.File.separator)
                Paths.get(outRoot).resolve(suffix.replace('/', java.io.File.separatorChar)).normalize()
            } else candidate
            if (relocated != null && Files.isRegularFile(relocated) && relocated.toString().endsWith(".jar")) {
                libraries.add(relocated)
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
