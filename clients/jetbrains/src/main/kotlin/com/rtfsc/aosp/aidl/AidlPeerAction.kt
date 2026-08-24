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
import java.nio.file.Path

class AidlPeerAction : AnAction("Go to AIDL Generated Peer") {
    private val mapper = ObjectMapper()

    override fun update(event: AnActionEvent) {
        event.presentation.isEnabledAndVisible =
            event.getData(CommonDataKeys.VIRTUAL_FILE)?.fileType == AidlFileType
    }

    override fun actionPerformed(event: AnActionEvent) {
        val project = event.project ?: return
        val file = event.getData(CommonDataKeys.VIRTUAL_FILE) ?: return
        val root = runCatching {
            Path.of(project.getService(EngineProfileSettings::class.java).current.workspaceViewRoot)
                .toAbsolutePath().normalize()
        }.getOrNull() ?: run {
            showError(project, "Workspace View Root is not configured.")
            return
        }
        val currentPath = runCatching {
            val current = file.toNioPath().toAbsolutePath().normalize()
            if (!current.startsWith(root)) return
            root.relativize(current).toString().replace('\\', '/')
        }.getOrNull() ?: return
        val name = file.nameWithoutExtension
        val request = mapper.createObjectNode().apply {
            put("workspaceViewRoot", root.toString())
            put("currentPath", currentPath)
            put("name", name)
        }
        object : Task.Backgroundable(project, "Find AIDL Generated Peer", true) {
            private var target: Target? = null

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Resolving generated AIDL peers"
                val response = runCatching {
                    project.getService(ProjectEngineService::class.java).aidlPeers(request.toString())
                }.getOrNull() ?: return
                target = selectTarget(mapper.readTree(response))
            }

            override fun onSuccess() {
                val resolved = target ?: run {
                    showError(project, "No unique visible AIDL generated peer was found.")
                    return
                }
                val targetFile = LocalFileSystem.getInstance()
                    .refreshAndFindFileByIoFile(root.resolve(resolved.path).normalize().toFile())
                if (targetFile == null) {
                    showError(project, "Generated AIDL peer is not visible: ${resolved.path}")
                    return
                }
                OpenFileDescriptor(project, targetFile, resolved.offset).navigate(true)
            }
        }.queue()
    }

    private fun selectTarget(response: JsonNode): Target? {
        val peers = response.path("result").path("peers")
        val visible = peers.filter { peer ->
            peer.path("confidence").asText() == "exact" && peer.path("exists").asBoolean(false)
        }.mapNotNull { peer ->
            val path = peer.path("path").path("repoPath").asText(peer.path("path").path("raw").asText())
            if (path.isBlank() || path.startsWith("/")) null else Target(path, 0)
        }.distinct()
        return visible.singleOrNull()
    }

    private fun showError(project: Project, message: String) {
        Messages.showWarningDialog(project, message, "AOSP AIDL Navigation")
    }

    private data class Target(val path: String, val offset: Int)
}
