package com.rtfsc.aosp.jni

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.fileEditor.OpenFileDescriptor
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import com.intellij.openapi.vfs.LocalFileSystem
import com.intellij.openapi.vfs.VirtualFile
import com.intellij.psi.search.FilenameIndex
import com.intellij.psi.search.GlobalSearchScope
import com.rtfsc.aosp.engine.ProjectEngineService
import com.rtfsc.aosp.profile.EngineProfileSettings
import java.nio.charset.StandardCharsets
import java.nio.file.Path

private data class SourceSet(
    val javaFiles: Map<String, String>,
    val cppFiles: Map<String, String>,
)

private val excludedSegments = setOf(".git", ".repo", "out", ".aosp-ide")

class JniPeerAction : AnAction("Go to JNI Peer") {
    private val objectMapper = ObjectMapper()

    override fun update(event: AnActionEvent) {
        val file = event.getData(com.intellij.openapi.actionSystem.CommonDataKeys.VIRTUAL_FILE)
        event.presentation.isEnabledAndVisible = file != null && isSupported(file)
    }

    override fun actionPerformed(event: AnActionEvent) {
        val project = event.project ?: return
        val file = event.getData(com.intellij.openapi.actionSystem.CommonDataKeys.VIRTUAL_FILE) ?: return
        val settings = project.getService(EngineProfileSettings::class.java).current
        val root = runCatching { Path.of(settings.workspaceViewRoot).toAbsolutePath().normalize() }.getOrNull()
        if (root == null) {
            showError(project, "Workspace View Root is not configured.")
            return
        }
        val currentPath = repoPath(root, file)
        if (currentPath == null) {
            showError(project, "The current file is outside the configured Workspace View Root.")
            return
        }
        val sources = collectSources(project, root)
        val request = objectMapper.createObjectNode().apply {
            set<JsonNode>("javaFiles", objectMapper.valueToTree(sources.javaFiles))
            set<JsonNode>("cppFiles", objectMapper.valueToTree(sources.cppFiles))
        }
        val response = try {
            project.getService(ProjectEngineService::class.java).jniPeers(request.toString())
        } catch (failure: Throwable) {
            showError(project, "JNI engine query failed: ${failure.message ?: failure.javaClass.simpleName}")
            return
        }
        val target = selectTarget(objectMapper.readTree(response), currentPath) ?: run {
            showError(project, "No JNI peer was found for this file.")
            return
        }
        val virtualTarget = findVisibleFile(project, root, target.path)
        if (virtualTarget == null) {
            showError(project, "JNI peer is not visible: ${target.path}")
            return
        }
        val offset = byteOffsetToCharOffset(virtualTarget, target.start)
        OpenFileDescriptor(project, virtualTarget, offset).navigate(true)
    }

    private fun collectSources(project: Project, root: Path): SourceSet {
        val java = linkedMapOf<String, String>()
        val cpp = linkedMapOf<String, String>()
        val extensions = listOf("java", "kt", "c", "cc", "cpp", "cxx", "h", "hh", "hpp")
        extensions.forEach { extension ->
            FilenameIndex.getAllFilesByExt(project, extension, GlobalSearchScope.projectScope(project))
                .forEach { file ->
                    val path = repoPath(root, file) ?: return@forEach
                    val source = runCatching {
                        String(file.contentsToByteArray(), StandardCharsets.UTF_8)
                    }.getOrNull() ?: return@forEach
                    if (extension == "java" || extension == "kt") java[path] = source else cpp[path] = source
                }
        }
        return SourceSet(java, cpp)
    }

    private fun repoPath(root: Path, file: VirtualFile): String? {
        val path = runCatching { file.toNioPath().toAbsolutePath().normalize() }.getOrNull() ?: return null
        if (!path.startsWith(root) || path == root) return null
        val relative = root.relativize(path)
        if (relative.any { segment -> segment.toString() in excludedSegments }) return null
        return relative.toString().replace('\\', '/')
    }

    private fun findVisibleFile(project: Project, root: Path, repoPath: String): VirtualFile? {
        val direct = LocalFileSystem.getInstance()
            .refreshAndFindFileByIoFile(root.resolve(repoPath).normalize().toFile())
        if (direct != null && repoPath(root, direct) == repoPath) return direct
        val extension = repoPath.substringAfterLast('.', "")
        if (extension.isEmpty()) return null
        return FilenameIndex.getAllFilesByExt(
            project,
            extension,
            GlobalSearchScope.projectScope(project),
        ).firstOrNull { file -> repoPath(root, file) == repoPath }
    }

    private fun selectTarget(response: JsonNode, currentPath: String): Target? {
        val result = response.path("result")
        if (result.isMissingNode || result.path("diagnostics").any()) return null
        val targets = buildList {
            result.path("peers").filter { peer ->
            peer.path("java").path("path").asText() == currentPath
            }.forEach { peer ->
                val native = peer.path("native")
                add(Target(native.path("path").asText(), native.path("range").path("start").asInt()))
            }
            result.path("peers").filter { peer ->
            peer.path("native").path("path").asText() == currentPath
            }.forEach { peer ->
                val java = peer.path("java")
                add(Target(java.path("path").asText(), java.path("range").path("start").asInt()))
            }
        }.distinct()
        return targets.singleOrNull()
    }

    private fun byteOffsetToCharOffset(file: VirtualFile, byteOffset: Int): Int {
        val bytes = file.contentsToByteArray()
        val safe = byteOffset.coerceIn(0, bytes.size)
        return String(bytes.copyOf(safe), StandardCharsets.UTF_8).length
    }

    private fun isSupported(file: VirtualFile): Boolean {
        return file.extension?.lowercase() in setOf("java", "kt", "c", "cc", "cpp", "cxx", "h", "hh", "hpp")
    }

    private fun showError(project: Project, message: String) {
        Messages.showWarningDialog(project, message, "AOSP JNI Navigation")
    }

    private data class Target(val path: String, val start: Int)
}
