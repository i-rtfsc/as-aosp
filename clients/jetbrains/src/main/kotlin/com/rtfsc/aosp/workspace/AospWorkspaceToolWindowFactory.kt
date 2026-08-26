package com.rtfsc.aosp.workspace

import com.intellij.openapi.options.ShowSettingsUtil
import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.openapi.wm.ToolWindowFactory
import com.intellij.ui.content.ContentFactory
import com.intellij.ui.components.JBLabel
import com.intellij.ui.components.JBPanel
import com.rtfsc.aosp.profile.EngineProfileSettings
import java.awt.BorderLayout
import java.awt.Color
import java.awt.FlowLayout
import java.awt.Font
import java.awt.GridLayout
import javax.swing.BorderFactory
import javax.swing.JButton
import javax.swing.JPanel

internal class AospWorkspaceToolWindowFactory : ToolWindowFactory {
    override fun createToolWindowContent(project: Project, toolWindow: ToolWindow) {
        val coordinator = project.getService(WorkspaceCoordinatorService::class.java)
        val settings = project.getService(EngineProfileSettings::class.java)
        val root = JBLabel(settings.current.workspaceViewRoot.ifBlank { "Workspace not configured" })
        val pack = JBLabel(settings.current.metadataPackRoot.ifBlank { "Metadata Pack not configured" })
        val state = JBLabel("Idle")
        val counts = JBLabel("0 source roots · 0 generated · 0 libraries")
        val diagnostics = JBLabel("0 diagnostics")

        val header = JBPanel<JBPanel<*>>(BorderLayout()).apply {
            border = BorderFactory.createEmptyBorder(12, 12, 10, 12)
            background = Color(62, 49, 122)
            add(JBLabel("AOSP Workspace").apply { foreground = Color.WHITE; font = font.deriveFont(Font.BOLD, 18f) }, BorderLayout.NORTH)
            add(JBLabel("Build-backed navigation").apply { foreground = Color(188, 230, 255) }, BorderLayout.SOUTH)
        }
        val statusPanel = JPanel(GridLayout(0, 1, 0, 6)).apply {
            border = BorderFactory.createEmptyBorder(12, 12, 12, 12)
            add(JBLabel("WORKSPACE")); add(root)
            add(JBLabel("METADATA PACK")); add(pack)
            add(JBLabel("INDEX STATUS")); add(state); add(counts); add(diagnostics)
        }
        val actions = JPanel(FlowLayout(FlowLayout.LEFT)).apply {
            add(JButton("Apply").apply { addActionListener { coordinator.start(false) } })
            add(JButton("Rebuild").apply { addActionListener { coordinator.start(true) } })
            add(JButton("Settings").apply { addActionListener { ShowSettingsUtil.getInstance().showSettingsDialog(project, "AOSP Workspace") } })
        }
        val panel = JPanel(BorderLayout()).apply {
            add(header, BorderLayout.NORTH)
            add(statusPanel, BorderLayout.CENTER)
            add(actions, BorderLayout.SOUTH)
        }
        val listener = coordinator.addListener { value ->
            state.text = "${value.state.uppercase()} · ${value.message}"
            counts.text = "${value.sourceRoots} source roots · ${value.generatedRoots} generated · ${value.libraries} libraries"
            diagnostics.text = "${value.diagnostics} diagnostics · ${value.skipped} skipped"
        }
        toolWindow.disposable?.let { com.intellij.openapi.util.Disposer.register(it) { listener.close() } }
        toolWindow.contentManager.addContent(ContentFactory.getInstance().createContent(panel, "Overview", false))
    }
}
