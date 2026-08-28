package com.rtfsc.aosp.workspace

import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.options.ShowSettingsUtil
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.openapi.wm.ToolWindowFactory
import com.intellij.ui.components.JBLabel
import com.intellij.ui.content.ContentFactory
import com.rtfsc.aosp.AospWorkspaceBundle.message
import com.rtfsc.aosp.profile.EngineProfileSettings
import com.rtfsc.aosp.profile.SourceModulePanel
import java.awt.BorderLayout
import java.awt.CardLayout
import java.awt.Color
import java.awt.FlowLayout
import java.awt.GridLayout
import java.nio.file.Files
import java.nio.file.Path
import javax.swing.BorderFactory
import javax.swing.BoxLayout
import javax.swing.JButton
import javax.swing.JPanel
import javax.swing.SwingUtilities

internal class AospWorkspaceToolWindowFactory : ToolWindowFactory {
    override fun createToolWindowContent(project: Project, toolWindow: ToolWindow) {
        val coordinator = project.getService(WorkspaceCoordinatorService::class.java)
        val settings = project.getService(EngineProfileSettings::class.java)
        val cards = CardLayout()
        val content = JPanel(cards)
        val state = JBLabel(message("workbench.idle"))
        val counts = JBLabel("0 source roots · 0 generated · 0 libraries")
        val diagnostics = JBLabel("0 diagnostics")
        val workspace = JBLabel()
        val out = JBLabel()
        val presetList = JPanel().apply { layout = BoxLayout(this, BoxLayout.Y_AXIS) }
        val compdb = JBLabel("CompDB not refreshed")

        fun renderPresetIds(ids: List<String>) {
            presetList.removeAll()
            presetNames(ids).forEach { presetList.add(JBLabel("●  $it")) }
            if (presetList.componentCount == 0) presetList.add(JBLabel(message("workbench.idle")))
            presetList.revalidate(); presetList.repaint()
        }

        fun renderProfile() {
            val profile = settings.current
            workspace.text = profile.workspaceViewRoot.ifBlank { message("workbench.notConfigured") }
            workspace.toolTipText = profile.workspaceViewRoot
            out.text = profile.outRoot.ifBlank { message("workbench.notConfigured") }
            out.toolTipText = profile.outRoot
            renderPresetIds(jsonStrings(profile.selectedPresetsJson))
        }

        val modulePanel = SourceModulePanel(
            project,
            { settings.current.workspaceViewRoot },
            { settings.current.outRoot },
            { selection ->
                settings.update {
                    it.targetsJson = selection.targetsJson
                    it.selectedPresetsJson = selection.presetIdsJson
                    it.selectedModulesJson = selection.manualModulesJson
                }
                renderProfile()
                cards.show(content, "workbench")
                coordinator.start(false)
            },
            { cards.show(content, "workbench") },
            { renderPresetIds(it) },
        )

        val accent = JPanel().apply {
            preferredSize = java.awt.Dimension(1, 3)
            background = Color(105, 164, 255)
        }
        val header = JPanel(BorderLayout()).apply {
            border = BorderFactory.createEmptyBorder(9, 10, 7, 10)
            add(JBLabel("AOSP Workspace"), BorderLayout.WEST)
            add(JPanel(FlowLayout(FlowLayout.RIGHT, 3, 0)).apply {
                add(JButton("↻").apply {
                    toolTipText = message("modules.refresh")
                    addActionListener { modulePanel.load(settings.current) }
                })
                add(JButton("⚙").apply {
                    toolTipText = message("workbench.settings")
                    addActionListener {
                        ShowSettingsUtil.getInstance().showSettingsDialog(project, "AOSP Workspace")
                    }
                })
            }, BorderLayout.EAST)
        }
        val overview = JPanel().apply {
            layout = BoxLayout(this, BoxLayout.Y_AXIS)
            border = BorderFactory.createEmptyBorder(5, 10, 8, 10)
            add(section(message("workbench.workspace"), workspace))
            add(section(message("workbench.outRoot"), out))
            add(section(message("workbench.index"), JPanel(GridLayout(0, 1, 0, 3)).apply {
                add(state); add(counts); add(diagnostics); add(compdb)
            }))
            add(section(message("workbench.enabledPresets"), presetList))
            add(JButton(message("workbench.manageModules")).apply {
                addActionListener {
                    modulePanel.load(settings.current)
                    cards.show(content, "modules")
                }
            })
        }
        val actions = JPanel(GridLayout(1, 0, 6, 0)).apply {
            border = BorderFactory.createEmptyBorder(8, 10, 10, 10)
            add(JButton(message("workbench.apply")).apply {
                addActionListener { coordinator.start(false) }
            })
            add(JButton(message("workbench.rebuild")).apply {
                addActionListener { coordinator.start(true) }
            })
            add(JButton("CompDB").apply {
                addActionListener { refreshCompdb(project, settings, compdb) }
            })
        }
        val workbench = JPanel(BorderLayout()).apply {
            add(JPanel(BorderLayout()).apply {
                add(accent, BorderLayout.NORTH)
                add(header, BorderLayout.CENTER)
            }, BorderLayout.NORTH)
            add(overview, BorderLayout.CENTER)
            add(actions, BorderLayout.SOUTH)
        }
        content.add(workbench, "workbench")
        content.add(modulePanel, "modules")
        renderProfile()
        if (settings.current.workspaceViewRoot.isNotBlank()) modulePanel.load(settings.current)

        val listener = coordinator.addListener { value ->
            SwingUtilities.invokeLater {
                state.text = "${value.state.uppercase()} · ${value.message}"
                counts.text = "${value.sourceRoots} source roots · ${value.generatedRoots} generated · ${value.libraries} libraries"
                diagnostics.text = "${value.diagnostics} diagnostics · ${value.skipped} skipped"
            }
        }
        val settingsListener = settings.addListener { SwingUtilities.invokeLater { renderProfile() } }
        toolWindow.disposable?.let {
            com.intellij.openapi.util.Disposer.register(it) {
                listener.close(); settingsListener.close()
            }
        }
        toolWindow.contentManager.addContent(
            ContentFactory.getInstance().createContent(content, "Overview", false),
        )
    }

    private fun section(title: String, component: java.awt.Component): JPanel =
        JPanel(BorderLayout()).apply {
            border = BorderFactory.createCompoundBorder(
                BorderFactory.createMatteBorder(1, 0, 0, 0, Color(90, 90, 90, 80)),
                BorderFactory.createEmptyBorder(8, 0, 8, 0),
            )
            add(JBLabel(title.uppercase()), BorderLayout.NORTH)
            add(component, BorderLayout.CENTER)
        }

    private fun presetNames(ids: List<String>): List<String> {
        val names = mapOf(
            "java-foundation" to "Java Foundation",
            "system-server-wms" to "System Server & WMS",
            "system-ui" to "System UI",
            "native-foundation" to "Native Foundation",
            "graphics-surfaceflinger" to "Graphics & SurfaceFlinger",
            "input-stack" to "Input Stack",
            "automotive" to "Automotive",
        )
        return ids.mapNotNull(names::get)
    }

    private fun jsonStrings(raw: String): List<String> = runCatching {
        ObjectMapper().readTree(raw).takeIf { it.isArray }?.map { it.asText() }.orEmpty()
    }.getOrDefault(emptyList())

    private fun refreshCompdb(project: Project, settings: EngineProfileSettings, label: JBLabel) {
        object : Task.Backgroundable(project, "Refreshing AOSP compilation database", true) {
            private var summary = "CompDB refresh failed"
            override fun run(indicator: ProgressIndicator) {
                val mapper = ObjectMapper()
                val profile = settings.current
                val mappings = runCatching { mapper.readTree(profile.pathMappings) }.getOrNull()
                    ?: mapper.createArrayNode()
                val result = mapper.readTree(
                    project.getService(com.rtfsc.aosp.engine.ProjectEngineService::class.java)
                        .compdbOverlay(mapper.writeValueAsString(mapOf(
                            "outRoot" to profile.outRoot,
                            "workspaceViewRoot" to profile.workspaceViewRoot,
                            "pathMappings" to mappings,
                        ))),
                ).path("result")
                val entries = result.path("entries")
                val target = Path.of(
                    project.basePath ?: System.getProperty("java.io.tmpdir"),
                    ".aosp-ide", "compdb-overlay", "compile_commands.json",
                )
                Files.createDirectories(target.parent)
                Files.writeString(target, mapper.writerWithDefaultPrettyPrinter().writeValueAsString(entries))
                summary = "${entries.size()} entries · ${result.path("diagnostics").size()} diagnostics"
            }
            override fun onSuccess() { label.text = summary }
            override fun onThrowable(error: Throwable) {
                label.text = "CompDB error · ${error.message ?: error.javaClass.simpleName}"
            }
        }.queue()
    }
}
