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
    val targetIds: List<String>,
    val skipped: Int,
    val diagnostics: List<String>,
)

internal object BuildFactsProjection {
    private val mapper = ObjectMapper()

    fun apply(
        state: EngineProfileState,
        engine: ProjectEngineService,
        cacheRoot: Path,
        targetModule: (String) -> Module,
        generatedModule: () -> Module,
    ): BuildFactsProjectionResult {
        val outRoot = state.outRoot.trim()
        if (outRoot.isEmpty()) return emptyResult("Out Root is required.")
        val diagnostics = mutableListOf<String>()
        val targets = runCatching { mapper.readTree(state.targetsJson) }.getOrNull()
        if (targets == null || !targets.isArray || targets.size() == 0) {
            diagnostics += "PROJECT_GRAPH_TARGETS_MISSING: Targets JSON must contain at least one target."
            return BuildFactsProjectionResult(emptyList(), emptyList(), emptyList(), emptyList(), 0, diagnostics)
        }
        val graphResponse = mapper.readTree(engine.projectGraph(mapper.writeValueAsString(mapOf("outRoot" to outRoot, "targets" to targets))))
        collectDiagnostics(graphResponse.path("result").path("diagnostics"), diagnostics)
        val graph = graphResponse.path("result").path("graph")
        val sourceRoots = linkedSetOf<Path>()
        val libraries = linkedSetOf<Path>()
        val targetIds = mutableListOf<String>()
        val targetDependencies = mutableMapOf<String, List<String>>()
        val archiveRefs = graph.path("targets").flatMap { target ->
            target.path("libraries").mapNotNull { ref ->
                ref.path("raw").asText().takeIf { it.endsWith(".srcjar", ignoreCase = true) }
            }
        }.distinct()
        val skipped = intArrayOf(0)
        val workspaceRoot = Paths.get(state.workspaceViewRoot).toAbsolutePath().normalize()
        val graphItems = graph.path("targets")
        val owners = graph.path("ownership").associate { ownership ->
            ownership.path("repoPath").asText() to ownership.path("targetId").asText()
        }
        for (item in graphItems) {
            val targetId = item.path("spec").path("id").asText().trim()
            if (targetId.isBlank()) continue
            val targetRoots = linkedSetOf<Path>()
            val excludedRoots = linkedSetOf<Path>()
            val targetLibraries = linkedSetOf<Path>()
            collectRepoRoots(item.path("sourceRoots"), workspaceRoot, targetRoots, skipped, targetId, owners)
            collectRepoRoots(item.path("sources"), workspaceRoot, targetRoots, skipped, targetId, owners)
            collectExcludedRepoRoots(item.path("excludedSourceRoots"), workspaceRoot, excludedRoots, skipped)
            collectLibraries(item.path("libraries"), workspaceRoot, outRoot, targetLibraries, skipped)
            applyProjection(
                targetModule(targetId), targetRoots.toList(), targetLibraries.toList(),
                emptyList(), excludedRoots.toList(),
            )
            targetIds += targetId
            targetDependencies[targetId] = item.path("targetDependencies").map { it.asText() }
            sourceRoots.addAll(targetRoots)
            libraries.addAll(targetLibraries)
        }
        val generatedParams = mapper.writeValueAsString(
            mapOf("outRoot" to outRoot, "archives" to archiveRefs, "cacheRoot" to cacheRoot.toString()),
        )
        val generatedResult = mapper.readTree(engine.generatedMaterialize(generatedParams)).path("result")
        collectDiagnostics(generatedResult.path("diagnostics"), diagnostics)
        val generatedRoots = generatedResult.path("sourceRoots").mapNotNull { node ->
            node.path("root").asText().takeIf { it.isNotBlank() }?.let(Paths::get)
        }.filter { Files.isDirectory(it) }
        val generated = generatedModule()
        applyProjection(generated, generatedRoots, emptyList(), emptyList(), emptyList())
        for (targetId in targetIds) {
            val module = targetModule(targetId)
            addModuleDependency(module, generated)
            targetDependencies[targetId].orEmpty().forEach { dependencyId ->
                addModuleDependency(module, targetModule(dependencyId))
            }
        }
        sourceRoots.addAll(generatedRoots)
        return BuildFactsProjectionResult(
            sourceRoots.toList(), generatedRoots, libraries.toList(), targetIds.distinct().sorted(),
            skipped[0], diagnostics,
        )
    }

    private fun emptyResult(message: String) = BuildFactsProjectionResult(
        emptyList(), emptyList(), emptyList(), emptyList(), 0, listOf(message),
    )

    private fun collectRepoRoots(
        refs: JsonNode,
        workspaceRoot: Path,
        roots: MutableSet<Path>,
        skipped: IntArray,
        targetId: String,
        owners: Map<String, String>,
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
            if (owners[repoPath]?.let { it != targetId } == true) {
                continue
            }
            val candidate = workspaceRoot.resolve(repoPath.replace('/', java.io.File.separatorChar)).normalize()
            if (candidate.startsWith(workspaceRoot)) roots.add(candidate) else skipped[0]++
        }
    }

    private fun collectLibraries(
        refs: JsonNode,
        workspaceRoot: Path,
        outRoot: String,
        libraries: MutableSet<Path>,
        skipped: IntArray,
    ) {
        for (ref in refs) {
            val kind = ref.path("kind").asText()
            val raw = ref.path("raw").asText()
            val candidate = when (kind) {
                "repo-relative" -> workspaceRoot.resolve(ref.path("repoPath").asText().replace('/', java.io.File.separatorChar))
                "generated" -> Paths.get(outRoot)
                    .resolve(raw.replace('\\', '/').removePrefix("out/").replace('/', java.io.File.separatorChar))
                "external-absolute" -> Paths.get(raw)
                else -> null
            }?.normalize()
            val relocated = candidate
            val localOut = if (outRoot.isNotBlank() && !Paths.get(raw).isAbsolute) {
                val normalized = raw.replace('\\', '/')
                val suffix = normalized.removePrefix("out/")
                Paths.get(outRoot).resolve(suffix.replace('/', java.io.File.separatorChar)).normalize()
            } else null
            val selected = when {
                relocated != null && Files.isRegularFile(relocated) -> relocated
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

    private fun collectExcludedRepoRoots(
        refs: JsonNode,
        workspaceRoot: Path,
        roots: MutableSet<Path>,
        skipped: IntArray,
    ) {
        for (ref in refs) {
            val repoPath = ref.path("repoPath").asText()
            if (ref.path("kind").asText() != "repo-relative" || repoPath.isBlank()) {
                skipped[0]++
                continue
            }
            val candidate = workspaceRoot.resolve(repoPath.replace('/', java.io.File.separatorChar)).normalize()
            if (candidate.startsWith(workspaceRoot)) roots.add(candidate) else skipped[0]++
        }
    }

    private fun applyProjection(
        module: Module,
        roots: List<Path>,
        libraries: List<Path>,
        moduleDependencies: List<Module>,
        excludedRoots: List<Path>,
    ) {
        WriteAction.compute<Unit, RuntimeException> {
            val model = ModuleRootManager.getInstance(module).modifiableModel
            model.clear()
            roots.distinct().forEach { root ->
                val content = model.addContentEntry(root.toString())
                content.addSourceFolder(root.toString(), false)
                excludedRoots.filter { it != root && it.startsWith(root) }
                    .forEach { content.addExcludeFolder(it.toString()) }
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
            moduleDependencies.distinct().filter { it != module }.forEach(model::addModuleOrderEntry)
            model.commit()
        }
    }

    private fun addModuleDependency(module: Module, dependency: Module) {
        WriteAction.compute<Unit, RuntimeException> {
            val model = ModuleRootManager.getInstance(module).modifiableModel
            if (model.orderEntries.none { it is com.intellij.openapi.roots.ModuleOrderEntry && it.module == dependency }) {
                model.addModuleOrderEntry(dependency)
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
