package com.rtfsc.aosp.workspace

import com.intellij.openapi.options.ShowSettingsUtil
import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.openapi.wm.ToolWindowFactory
import com.intellij.ui.content.ContentFactory
import com.intellij.ui.components.JBLabel
import com.intellij.ui.components.JBPanel
import com.rtfsc.aosp.profile.EngineProfileSettings
import com.rtfsc.aosp.profile.SourceModulePanel
import java.awt.BorderLayout
import java.awt.Color
import java.awt.FlowLayout
import java.awt.Font
import java.awt.GridLayout
import javax.swing.BorderFactory
import javax.swing.JButton
import javax.swing.JPanel
import javax.swing.SwingUtilities
import com.fasterxml.jackson.databind.ObjectMapper
import java.nio.file.Files
import java.nio.file.Path
import com.intellij.openapi.progress.Task
import com.intellij.openapi.progress.ProgressIndicator

internal class AospWorkspaceToolWindowFactory : ToolWindowFactory {
    override fun createToolWindowContent(project: Project, toolWindow: ToolWindow) {
        val coordinator = project.getService(WorkspaceCoordinatorService::class.java)
        val settings = project.getService(EngineProfileSettings::class.java)
        val root = JBLabel(settings.current.workspaceViewRoot.ifBlank { "Workspace not configured" })
        val state = JBLabel("Idle")
        val counts = JBLabel("0 source roots · 0 generated · 0 libraries")
        val diagnostics = JBLabel("0 diagnostics")
        val compdb = JBLabel("CompDB not refreshed")
        val modulePanel = SourceModulePanel(
            project,
            { settings.current.workspaceViewRoot },
            { targets -> settings.update { it.targetsJson = targets } },
        )
        modulePanel.loadSelected(settings.current.targetsJson)
        modulePanel.refresh()

        val header = JBPanel<JBPanel<*>>(BorderLayout()).apply {
            border = BorderFactory.createEmptyBorder(12, 12, 10, 12)
            background = Color(62, 49, 122)
            add(JBLabel("AOSP Workspace").apply { foreground = Color.WHITE; font = font.deriveFont(Font.BOLD, 18f) }, BorderLayout.NORTH)
            add(JBLabel("Build-backed navigation").apply { foreground = Color(188, 230, 255) }, BorderLayout.SOUTH)
        }
        val statusPanel = JPanel(GridLayout(0, 1, 0, 6)).apply {
            border = BorderFactory.createEmptyBorder(12, 12, 12, 12)
            add(JBLabel("WORKSPACE")); add(root)
            add(JBLabel("OUT ROOT")); add(root)
            add(JBLabel("INDEX STATUS")); add(state); add(counts); add(diagnostics)
            add(JBLabel("C/C++ COMPDB")); add(compdb)
        }
        val actions = JPanel(FlowLayout(FlowLayout.LEFT)).apply {
            add(JButton("Apply").apply { addActionListener { coordinator.start(false) } })
            add(JButton("Rebuild").apply { addActionListener { coordinator.start(true) } })
            add(JButton("Refresh CompDB").apply { addActionListener { refreshCompdb(project, settings, compdb) } })
            add(JButton("Settings").apply { addActionListener { ShowSettingsUtil.getInstance().showSettingsDialog(project, "AOSP Workspace") } })
        }
        val moduleSection = JPanel(BorderLayout()).apply {
            add(JBLabel("DISCOVERED MODULES"), BorderLayout.NORTH)
            add(modulePanel, BorderLayout.CENTER)
        }
        val center = JPanel(BorderLayout()).apply {
            add(statusPanel, BorderLayout.NORTH)
            add(moduleSection, BorderLayout.CENTER)
        }
        val panel = JPanel(BorderLayout()).apply {
            add(header, BorderLayout.NORTH)
            add(center, BorderLayout.CENTER)
            add(actions, BorderLayout.SOUTH)
        }
        val listener = coordinator.addListener { value ->
            SwingUtilities.invokeLater {
                state.text = "${value.state.uppercase()} · ${value.message}"
                counts.text = "${value.sourceRoots} source roots · ${value.generatedRoots} generated · ${value.libraries} libraries"
                diagnostics.text = "${value.diagnostics} diagnostics · ${value.skipped} skipped"
            }
        }
        val settingsListener = settings.addListener { value ->
            SwingUtilities.invokeLater {
                root.text = value.workspaceViewRoot.ifBlank { "Workspace not configured" }
            }
        }
        toolWindow.disposable?.let {
            com.intellij.openapi.util.Disposer.register(it) {
                listener.close()
                settingsListener.close()
            }
        }
        toolWindow.contentManager.addContent(ContentFactory.getInstance().createContent(panel, "Overview", false))
    }

    private fun refreshCompdb(project: Project, settings: EngineProfileSettings, label: JBLabel) {
        object : Task.Backgroundable(project, "Refreshing AOSP compilation database", true) {
            private var summary = "CompDB refresh failed"

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Mapping Native Pack compilation commands"
                val mapper = ObjectMapper()
                val profile = settings.current
                val mappings = runCatching { mapper.readTree(profile.pathMappings) }.getOrNull()
                    ?: mapper.createArrayNode()
                val params = mapper.writeValueAsString(
                    mapOf(
                        "outRoot" to profile.outRoot,
                        "workspaceViewRoot" to profile.workspaceViewRoot,
                        "pathMappings" to mappings,
                    ),
                )
                val result = mapper.readTree(
                    project.getService(com.rtfsc.aosp.engine.ProjectEngineService::class.java)
                        .compdbOverlay(params),
                ).path("result")
                val entries = result.path("entries")
                val target = Path.of(
                    project.basePath ?: System.getProperty("java.io.tmpdir"),
                    ".aosp-ide",
                    "compdb-overlay",
                    "compile_commands.json",
                )
                Files.createDirectories(target.parent)
                Files.writeString(target, mapper.writerWithDefaultPrettyPrinter().writeValueAsString(entries))
                summary = "${entries.size()} entries · ${result.path("diagnostics").size()} diagnostics"
            }

            override fun onSuccess() {
                label.text = summary
            }

            override fun onThrowable(error: Throwable) {
                label.text = "CompDB error · ${error.message ?: error.javaClass.simpleName}"
            }
        }.queue()
    }
}
