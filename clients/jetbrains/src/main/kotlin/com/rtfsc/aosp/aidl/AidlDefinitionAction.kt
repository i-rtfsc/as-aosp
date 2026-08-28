package com.rtfsc.aosp.aidl

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.actionSystem.CommonDataKeys
import com.intellij.openapi.fileEditor.OpenFileDescriptor
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import com.intellij.openapi.vfs.LocalFileSystem
import com.intellij.openapi.vfs.VirtualFile
import com.rtfsc.aosp.engine.ProjectEngineService
import com.rtfsc.aosp.profile.EngineProfileSettings
import java.nio.charset.StandardCharsets
import java.nio.file.Path

class AidlDefinitionAction : AnAction("Go to AIDL Definition") {
    private val objectMapper = ObjectMapper()

    override fun update(event: AnActionEvent) {
        val file = event.getData(CommonDataKeys.VIRTUAL_FILE)
        event.presentation.isEnabledAndVisible = file?.fileType == AidlFileType
    }

    override fun actionPerformed(event: AnActionEvent) {
        val project = event.project ?: return
        val editor = event.getData(CommonDataKeys.EDITOR) ?: return
        val file = event.getData(CommonDataKeys.VIRTUAL_FILE) ?: return
        val token = identifierAt(editor.document.text, editor.caretModel.offset) ?: run {
            showError(project, "Place the caret on an AIDL type name.")
            return
        }
        val settings = project.getService(EngineProfileSettings::class.java).current
        val root = runCatching {
            Path.of(settings.workspaceViewRoot).toAbsolutePath().normalize()
        }.getOrNull() ?: run {
            showError(project, "Workspace View Root is not configured.")
            return
        }
        val currentPath = repoPath(root, file) ?: run {
            showError(project, "The current file is outside the configured Workspace View Root.")
            return
        }
        val request = objectMapper.createObjectNode().apply {
            put("workspaceViewRoot", root.toString())
            put("currentPath", currentPath)
            put("name", token)
        }
        object : Task.Backgroundable(project, "Find AIDL Definition", true) {
            private var target: Target? = null
            private var failureMessage: String? = null

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Resolving AIDL source symbol"
                val engine = project.getService(ProjectEngineService::class.java)
                val response = try {
                    engine.aidlResolve(request.toString())
                } catch (failure: Throwable) {
                    failureMessage = "AIDL engine query failed: ${failure.message ?: failure.javaClass.simpleName}"
                    return
                }
                target = selectTarget(objectMapper.readTree(response))
                if (target == null) failureMessage = "No unique AIDL definition was found."
            }

            override fun onSuccess() {
                val resolved = target ?: run {
                    showError(project, failureMessage ?: "AIDL definition query failed.")
                    return
                }
                val targetFile = findVisibleFile(root, resolved.path)
                if (targetFile == null) {
                    showError(project, "AIDL definition is not visible: ${resolved.path}")
                    return
                }
                val offset = byteOffsetToCharOffset(targetFile, resolved.start)
                OpenFileDescriptor(project, targetFile, offset).navigate(true)
            }
        }.queue()
    }

    private fun selectTarget(response: JsonNode): Target? {
        val result = response.path("result")
        val resolution = if (result.has("resolution")) result.path("resolution") else result
        if (resolution.path("state").asText() != "found") return null
        val symbols = resolution.path("symbols")
        if (!symbols.isArray || symbols.size() != 1) return null
        val symbol = symbols.single()
        return Target(
            symbol.path("path").asText(),
            symbol.path("range").path("start").path("offset").asInt(),
        )
    }

    private fun identifierAt(text: String, offset: Int): String? {
        if (text.isEmpty()) return null
        val position = offset.coerceIn(0, text.lastIndex)
        return Regex("[A-Za-z_]\\w*").findAll(text)
            .firstOrNull { position in it.range || offset == it.range.last + 1 }
            ?.value
    }

    private fun repoPath(root: Path, file: VirtualFile): String? {
        val path = runCatching { file.toNioPath().toAbsolutePath().normalize() }.getOrNull()
            ?: return null
        if (!path.startsWith(root) || path == root) return null
        return root.relativize(path).toString().replace('\\', '/')
    }

    private fun findVisibleFile(root: Path, repoPath: String): VirtualFile? {
        return LocalFileSystem.getInstance().refreshAndFindFileByIoFile(
            root.resolve(repoPath).normalize().toFile(),
        )
    }

    private fun byteOffsetToCharOffset(file: VirtualFile, byteOffset: Int): Int {
        val bytes = file.contentsToByteArray()
        val safe = byteOffset.coerceIn(0, bytes.size)
        return String(bytes.copyOf(safe), StandardCharsets.UTF_8).length
    }

    private fun showError(project: Project, message: String) {
        Messages.showWarningDialog(project, message, "AOSP AIDL Navigation")
    }

    private data class Target(val path: String, val start: Int)
}
