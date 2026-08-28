package com.rtfsc.aosp.profile

import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.application.PathManager
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.ui.components.JBLabel
import com.intellij.ui.components.JBTextField
import com.rtfsc.aosp.AospWorkspaceBundle.message
import com.rtfsc.aosp.engine.ProjectEngineService
import java.awt.BorderLayout
import java.awt.FlowLayout
import java.nio.file.Paths
import java.security.MessageDigest
import javax.swing.BorderFactory
import javax.swing.BoxLayout
import javax.swing.JButton
import javax.swing.JCheckBox
import javax.swing.JPanel
import javax.swing.JScrollPane

internal data class ModuleSelection(
    val targetsJson: String,
    val presetIdsJson: String,
    val manualModulesJson: String,
)

internal class SourceModulePanel(
    private val project: Project,
    private val workspaceRoot: () -> String,
    private val outRoot: () -> String,
    private val onApply: (ModuleSelection) -> Unit,
    private val onCancel: () -> Unit,
    private val onPresetPreview: (List<String>) -> Unit,
) : JPanel(BorderLayout()) {
    private data class Preset(
        val id: String, val name: String, val modules: List<String>,
        val recommended: Boolean,
    )
    private data class Entry(
        val name: String, val type: String, val status: String,
        val path: String, val capabilities: List<String>,
    )

    private val mapper = ObjectMapper()
    private val query = JBTextField().apply { emptyText.text = message("modules.search") }
    private val summary = JBLabel(message("modules.searchHint"))
    private val presetList = JPanel().apply { layout = BoxLayout(this, BoxLayout.Y_AXIS) }
    private val moduleList = JPanel().apply { layout = BoxLayout(this, BoxLayout.Y_AXIS) }
    private var presets = emptyList<Preset>()
    private var entries = emptyList<Entry>()
    private val selectedPresets = linkedSetOf<String>()
    private val selectedModules = linkedSetOf<String>()

    init {
        border = BorderFactory.createEmptyBorder(8, 8, 8, 8)
        val toolbar = JPanel(BorderLayout(5, 0)).apply {
            add(query, BorderLayout.CENTER)
            add(JButton(message("modules.refresh")).apply {
                addActionListener { search() }
            }, BorderLayout.EAST)
        }
        val content = JPanel().apply {
            layout = BoxLayout(this, BoxLayout.Y_AXIS)
            add(JBLabel(message("modules.presets")))
            add(presetList)
            add(JBLabel(message("modules.results")))
            add(summary)
            add(moduleList)
        }
        val footer = JPanel(FlowLayout(FlowLayout.RIGHT)).apply {
            add(JButton(message("modules.cancel")).apply { addActionListener { onCancel() } })
            add(JButton(message("modules.apply")).apply { addActionListener { applySelection() } })
        }
        add(toolbar, BorderLayout.NORTH)
        add(JScrollPane(content), BorderLayout.CENTER)
        add(footer, BorderLayout.SOUTH)
        query.addActionListener { search() }
    }

    fun load(state: EngineProfileState) {
        selectedPresets.clear()
        selectedModules.clear()
        selectedPresets.addAll(jsonStrings(state.selectedPresetsJson))
        selectedModules.addAll(jsonStrings(state.selectedModulesJson))
        loadPresets(selectedPresets.isNotEmpty())
    }

    private fun loadPresets(hasSavedSelection: Boolean) {
        object : Task.Backgroundable(project, message("modules.discovering"), true) {
            private var loaded = emptyList<Preset>()
            private var defaults = emptyList<String>()
            override fun run(indicator: ProgressIndicator) {
                val result = mapper.readTree(
                    project.getService(ProjectEngineService::class.java).selectionCatalog(),
                ).path("result")
                loaded = result.path("presets").map { node ->
                    Preset(
                        node.path("id").asText(), node.path("name").asText(),
                        node.path("modules").map { it.asText() },
                        node.path("recommended").asBoolean(),
                    )
                }
                defaults = result.path("defaultPresetIds").map { it.asText() }
            }
            override fun onSuccess() {
                presets = loaded
                if (!hasSavedSelection) selectedPresets.addAll(defaults)
                onPresetPreview(selectedPresets.toList())
                renderPresets()
            }
            override fun onThrowable(error: Throwable) {
                summary.text = message("modules.discoveryFailed", error.message ?: error.javaClass.simpleName)
            }
        }.queue()
    }

    private fun search() {
        val value = query.text.trim()
        if (value.length < 2) {
            entries = emptyList()
            summary.text = message("modules.searchHint")
            renderModules()
            return
        }
        summary.text = message("modules.discovering")
        object : Task.Backgroundable(project, message("modules.discovering"), true) {
            private var loaded = emptyList<Entry>()
            private var total = 0
            private var buildBacked = 0
            override fun run(indicator: ProgressIndicator) {
                val response = mapper.readTree(
                    project.getService(ProjectEngineService::class.java).sourceModules(
                        mapper.writeValueAsString(mapOf(
                            "workspaceViewRoot" to workspaceRoot().trim(),
                            "outRoot" to outRoot().trim(),
                            "cacheRoot" to cacheRoot(),
                            "query" to value,
                            "limit" to 50,
                        )),
                    ),
                ).path("result")
                total = response.path("total").asInt()
                buildBacked = response.path("buildBacked").asInt()
                loaded = response.path("modules").map { node ->
                    Entry(
                        node.path("name").asText(), node.path("type").asText(),
                        node.path("status").asText(), node.path("blueprint").asText(),
                        node.path("capabilities").map { it.asText() },
                    )
                }
            }
            override fun onSuccess() {
                entries = loaded
                summary.text = message("modules.summary", total, buildBacked, loaded.size)
                renderModules()
            }
            override fun onThrowable(error: Throwable) {
                summary.text = message("modules.discoveryFailed", error.message ?: error.javaClass.simpleName)
            }
        }.queue()
    }

    private fun applySelection() {
        object : Task.Backgroundable(project, message("workbench.apply"), true) {
            private var selection: ModuleSelection? = null
            override fun run(indicator: ProgressIndicator) {
                val response = mapper.readTree(
                    project.getService(ProjectEngineService::class.java).selectionPlan(
                        mapper.writeValueAsString(mapOf(
                            "workspaceViewRoot" to workspaceRoot().trim(),
                            "outRoot" to outRoot().trim(),
                            "cacheRoot" to cacheRoot(),
                            "presetIds" to selectedPresets.sorted(),
                            "manualModules" to selectedModules.sorted(),
                        )),
                    ),
                ).path("result")
                selection = ModuleSelection(
                    mapper.writeValueAsString(response.path("targets")),
                    mapper.writeValueAsString(response.path("presetIds")),
                    mapper.writeValueAsString(selectedModules.sorted()),
                )
            }
            override fun onSuccess() { selection?.let(onApply) }
            override fun onThrowable(error: Throwable) {
                summary.text = message("modules.discoveryFailed", error.message ?: error.javaClass.simpleName)
            }
        }.queue()
    }

    private fun renderPresets() {
        presetList.removeAll()
        presets.forEach { preset ->
            val label = buildString {
                append(preset.name).append("  ·  ").append(preset.modules.size)
                if (preset.recommended) append("  ·  Recommended")
            }
            presetList.add(JCheckBox(label, preset.id in selectedPresets).apply {
                toolTipText = preset.modules.joinToString(", ")
                addActionListener {
                    if (isSelected) selectedPresets.add(preset.id)
                    else selectedPresets.remove(preset.id)
                }
            })
        }
        presetList.revalidate(); presetList.repaint()
    }

    private fun renderModules() {
        moduleList.removeAll()
        entries.forEach { entry ->
            moduleList.add(JCheckBox(
                "${entry.name}  ·  ${entry.type}  ·  ${entry.status}",
                entry.name in selectedModules,
            ).apply {
                toolTipText = listOf(entry.path, entry.capabilities.joinToString(", "))
                    .filter(String::isNotBlank).joinToString(" · ")
                addActionListener {
                    if (isSelected) selectedModules.add(entry.name)
                    else selectedModules.remove(entry.name)
                }
            })
        }
        moduleList.revalidate(); moduleList.repaint()
    }

    private fun jsonStrings(raw: String): List<String> = runCatching {
        mapper.readTree(raw).takeIf { it.isArray }?.map { it.asText() }.orEmpty()
    }.getOrDefault(emptyList())

    private fun cacheRoot(): String {
        val key = project.basePath.orEmpty().toByteArray()
            .let { MessageDigest.getInstance("SHA-256").digest(it) }
            .joinToString("") { "%02x".format(it) }.take(24)
        return Paths.get(PathManager.getSystemPath(), "aosp-workspace", key).toString()
    }
}
