package com.rtfsc.aosp.profile

import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.application.PathManager
import com.intellij.ui.components.JBTextField
import com.intellij.ui.components.JBLabel
import com.rtfsc.aosp.engine.ProjectEngineService
import java.awt.BorderLayout
import java.awt.FlowLayout
import javax.swing.BorderFactory
import javax.swing.JButton
import javax.swing.JCheckBox
import javax.swing.JPanel
import javax.swing.JScrollPane
import java.security.MessageDigest
import java.nio.file.Paths

internal class SourceModulePanel(
    private val project: Project,
    private val workspaceRoot: () -> String,
    private val onApply: (String) -> Unit,
) : JPanel(BorderLayout()) {
    private data class Entry(val name: String, val type: String, val status: String, val recommended: Boolean)
    private val mapper = ObjectMapper()
    private val query = JBTextField()
    private val summary = JBLabel("Not loaded")
    private val list = JPanel()
    private var entries = emptyList<Entry>()
    private val selected = linkedSetOf<String>()

    init {
        val toolbar = JPanel(BorderLayout()).apply {
            add(query, BorderLayout.CENTER)
            add(JButton("Refresh").apply { addActionListener { refresh() } }, BorderLayout.EAST)
        }
        list.layout = javax.swing.BoxLayout(list, javax.swing.BoxLayout.Y_AXIS)
        list.border = BorderFactory.createEmptyBorder(4, 4, 4, 4)
        val footer = JPanel(FlowLayout(FlowLayout.LEFT)).apply {
            add(summary)
            add(JButton("Recommended").apply { addActionListener { selectRecommended() } })
            add(JButton("Clear").apply { addActionListener { selected.clear(); render() } })
            add(JButton("Apply Modules").apply { addActionListener { onApply(encodeTargets()) } })
        }
        add(toolbar, BorderLayout.NORTH)
        add(JScrollPane(list), BorderLayout.CENTER)
        add(footer, BorderLayout.SOUTH)
        query.addActionListener { render() }
    }

    fun loadSelected(raw: String) {
        runCatching { mapper.readTree(raw).map { it.path("selectors").path("modules").firstOrNull()?.asText() }.filterNotNull() }
            .getOrNull()?.let { selected.addAll(it) }
    }

    fun refresh() {
        val root = workspaceRoot().trim()
        if (root.isBlank()) { summary.text = "Workspace View Root is required."; return }
        summary.text = "Discovering Android.bp modules..."
        object : Task.Backgroundable(project, "Discovering AOSP source modules", true) {
            private var loaded = emptyList<Entry>()
            private var message = ""
            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Parsing Android.bp and Android.mk"
                val projectKey = project.basePath.orEmpty().toByteArray()
                    .let { MessageDigest.getInstance("SHA-256").digest(it) }
                    .joinToString("") { "%02x".format(it) }.take(24)
                val response = mapper.readTree(project.getService(ProjectEngineService::class.java).sourceModules(
                    mapper.writeValueAsString(mapOf(
                        "workspaceViewRoot" to root,
                        "cacheRoot" to Paths.get(PathManager.getSystemPath(), "aosp-workspace", projectKey).toString(),
                    )),
                )).path("result")
                loaded = response.path("modules").map { node ->
                    Entry(node.path("name").asText(), node.path("type").asText(), node.path("status").asText(), node.path("recommended").asBoolean())
                }
                message = "${loaded.size} modules · ${response.path("filesScanned").asInt()} files · " +
                    if (response.path("cacheHit").asBoolean()) "cache reused" else "scanned"
            }
            override fun onSuccess() { entries = loaded; summary.text = message; if (selected.isEmpty()) selectRecommended() else render() }
            override fun onThrowable(error: Throwable) { summary.text = "Discovery failed: ${error.message ?: error.javaClass.simpleName}" }
        }.queue()
    }

    private fun selectRecommended() { selected.clear(); selected.addAll(entries.filter { it.recommended }.map { it.name }); render() }

    private fun render() {
        val value = query.text.trim().lowercase()
        list.removeAll()
        entries.filter { value.isBlank() || it.name.lowercase().contains(value) || it.type.lowercase().contains(value) }
            .forEach { entry ->
                val check = JCheckBox("${entry.name} · ${entry.type} · ${entry.status}", selected.contains(entry.name))
                check.toolTipText = if (entry.recommended) "Recommended module" else "Source module"
                check.addActionListener { if (check.isSelected) selected.add(entry.name) else selected.remove(entry.name) }
                list.add(check)
            }
        list.revalidate(); list.repaint()
    }

    private fun encodeTargets(): String = mapper.writeValueAsString(selected.sorted().map { name ->
        mapOf("id" to name, "displayName" to name, "kind" to "java", "selectors" to mapOf("modules" to listOf(name)), "includeDependencies" to true, "includeGeneratedSources" to true)
    })
}
