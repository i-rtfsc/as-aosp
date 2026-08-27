package com.rtfsc.aosp.profile

import com.fasterxml.jackson.databind.ObjectMapper
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.ui.components.JBTextField
import com.intellij.util.ui.FormBuilder
import com.rtfsc.aosp.engine.ProjectEngineService
import java.awt.BorderLayout
import java.awt.FlowLayout
import javax.swing.JButton
import javax.swing.JCheckBox
import javax.swing.JComboBox
import javax.swing.JLabel
import javax.swing.JList
import javax.swing.JPanel
import javax.swing.JScrollPane
import javax.swing.DefaultListModel

internal data class TargetDraft(
    val id: String,
    val kind: String,
    val artifactKind: String,
    val artifact: String,
    val module: String,
    val sourceRoots: List<String>,
    val includeDependencies: Boolean,
    val includeGeneratedSources: Boolean,
    val includeResources: Boolean,
    val includeManifest: Boolean,
) {
    fun toTargetSpec(): Map<String, Any> = mapOf(
        "id" to id,
        "displayName" to id,
        "kind" to kind,
        "selectors" to mapOf(
            "modules" to if (module.isBlank()) emptyList<String>() else listOf(module),
            "artifacts" to if (artifact.isBlank()) emptyList<Any>() else listOf(
                mapOf("kind" to artifactKind, "name" to artifact),
            ),
            "sourceRoots" to sourceRoots,
        ),
        "includeDependencies" to includeDependencies,
        "includeGeneratedSources" to includeGeneratedSources,
        "includeResources" to includeResources,
        "includeManifest" to includeManifest,
    )
}

internal object TargetDraftCodec {
    private val mapper = ObjectMapper()

    fun encode(targets: List<TargetDraft>): String = mapper.writeValueAsString(targets.map(TargetDraft::toTargetSpec))

    fun decode(raw: String): List<TargetDraft> {
        val root = runCatching { mapper.readTree(raw) }.getOrNull() ?: return emptyList()
        if (!root.isArray) return emptyList()
        return root.map { target ->
            val selectors = target.path("selectors")
            val artifact = selectors.path("artifacts").firstOrNull()
            TargetDraft(
                id = target.path("id").asText(),
                kind = target.path("kind").asText("java"),
                artifactKind = artifact?.path("kind")?.asText("jar") ?: "jar",
                artifact = artifact?.path("name")?.asText().orEmpty(),
                module = selectors.path("modules").firstOrNull()?.asText().orEmpty(),
                sourceRoots = selectors.path("sourceRoots").map { it.asText() },
                includeDependencies = target.path("includeDependencies").asBoolean(true),
                includeGeneratedSources = target.path("includeGeneratedSources").asBoolean(true),
                includeResources = target.path("includeResources").asBoolean(false),
                includeManifest = target.path("includeManifest").asBoolean(false),
            )
        }
    }
}

internal class TargetEditorPanel(
    private val project: Project,
    private val metadataPackRoot: () -> String,
    private val metadataRoot: () -> String = { "" },
) : JPanel(BorderLayout()) {
    private data class ArtifactChoice(
        val id: String,
        val kind: String,
        val name: String,
        val path: String,
        val owners: List<String>,
    ) {
        override fun toString(): String = "$name · ${owners.joinToString(", ").ifBlank { "no owner" }}"
    }

    private val mapper = ObjectMapper()
    private val drafts = mutableListOf<TargetDraft>()
    private val targetListModel = DefaultListModel<String>()
    private val targetList = JList(targetListModel)
    private val targetId = JBTextField()
    private val targetKind = JComboBox(arrayOf("java", "android", "native", "mixed"))
    private val artifactKind = JComboBox(arrayOf("jar", "apk", "srcjar", "aar", "output"))
    private val artifactCatalog = JComboBox<ArtifactChoice>()
    private val targetArtifact = JBTextField()
    private val targetModule = JBTextField()
    private val targetSources = JBTextField()
    private val includeDependencies = JCheckBox("Dependencies", true)
    private val includeGenerated = JCheckBox("Generated sources", true)
    private val includeResources = JCheckBox("Resources")
    private val includeManifest = JCheckBox("Manifest")
    private val loadArtifactsButton = JButton("Load Artifacts")
    private val resolveTargetButton = JButton("Resolve Target")
    private val addTargetButton = JButton("Add Target")
    private val removeTargetButton = JButton("Remove Target")
    private val targetPreview = JLabel("Not resolved")
    private var artifactFacts = emptyList<ArtifactChoice>()

    init {
        val editor = FormBuilder.createFormBuilder()
            .addLabeledComponent(JLabel("Target Name"), targetId)
            .addLabeledComponent(JLabel("Kind"), targetKind)
            .addLabeledComponent(JLabel("Artifact Kind"), artifactKind)
            .addLabeledComponent(JLabel("Artifact From Pack"), artifactCatalog)
            .addLabeledComponent(JLabel("Artifact"), targetArtifact)
            .addLabeledComponent(JLabel("Soong Module"), targetModule)
            .addLabeledComponent(JLabel("Source Roots"), targetSources)
            .panel
        val options = JPanel(FlowLayout(FlowLayout.LEFT)).apply {
            add(includeDependencies)
            add(includeGenerated)
            add(includeResources)
            add(includeManifest)
        }
        val actions = JPanel(FlowLayout(FlowLayout.LEFT)).apply {
            add(loadArtifactsButton)
            add(resolveTargetButton)
            add(addTargetButton)
            add(removeTargetButton)
        }
        val controls = JPanel(BorderLayout()).apply {
            add(targetPreview, BorderLayout.NORTH)
            add(editor, BorderLayout.CENTER)
            add(JPanel(BorderLayout()).apply {
                add(options, BorderLayout.NORTH)
                add(actions, BorderLayout.SOUTH)
            }, BorderLayout.SOUTH)
        }
        add(JScrollPane(targetList), BorderLayout.CENTER)
        add(controls, BorderLayout.SOUTH)

        loadArtifactsButton.addActionListener { loadArtifacts() }
        resolveTargetButton.addActionListener { resolveTarget() }
        addTargetButton.addActionListener { addTarget() }
        removeTargetButton.addActionListener { removeTarget() }
        artifactKind.addActionListener { refreshArtifactChoices() }
        artifactCatalog.addActionListener { selectArtifact() }
    }

    fun load(raw: String) {
        drafts.clear()
        drafts.addAll(TargetDraftCodec.decode(raw))
        renderTargets()
    }

    fun serialize(): String = TargetDraftCodec.encode(drafts)

    private fun currentDraft(): TargetDraft = TargetDraft(
        id = targetId.text.trim(),
        kind = targetKind.selectedItem.toString(),
        artifactKind = artifactKind.selectedItem.toString(),
        artifact = targetArtifact.text.trim(),
        module = targetModule.text.trim(),
        sourceRoots = targetSources.text.split(',').map(String::trim).filter(String::isNotEmpty),
        includeDependencies = includeDependencies.isSelected,
        includeGeneratedSources = includeGenerated.isSelected,
        includeResources = includeResources.isSelected,
        includeManifest = includeManifest.isSelected,
    )

    private fun addTarget() {
        val draft = currentDraft()
        if (draft.id.isBlank()) {
            targetPreview.text = "Enter a target name first."
            return
        }
        drafts += draft
        renderTargets()
        targetId.text = ""
        targetArtifact.text = ""
        targetModule.text = ""
        targetSources.text = ""
        targetPreview.text = "Target added. Save the profile to apply it."
    }

    private fun removeTarget() {
        val index = targetList.selectedIndex
        if (index < 0) return
        drafts.removeAt(index)
        renderTargets()
    }

    private fun renderTargets() {
        targetListModel.clear()
        drafts.forEach { target ->
            val selector = target.artifact.ifBlank { target.module.ifBlank { target.sourceRoots.joinToString(", ") } }
            targetListModel.addElement("${target.id} · ${target.kind} · ${selector.ifBlank { "no selector" }}")
        }
    }

    private fun loadArtifacts() {
        val packRoot = metadataPackRoot().trim()
        val outRoot = metadataRoot().trim()
        if (packRoot.isBlank() && outRoot.isBlank()) {
            targetPreview.text = "Metadata Pack Root or Out Root is required."
            return
        }
        loadArtifactsButton.isEnabled = false
        object : Task.Backgroundable(project, "Loading AOSP artifacts", true) {
            private var choices = emptyList<ArtifactChoice>()

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Loading Artifact Catalog"
                val params = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot, "metadataRoot" to outRoot))
                val result = mapper.readTree(project.getService(ProjectEngineService::class.java).artifactsCatalog(params)).path("result")
                choices = result.path("artifacts").map { artifact ->
                    ArtifactChoice(
                        id = artifact.path("id").asText(),
                        kind = artifact.path("kind").asText(),
                        name = artifact.path("name").asText(),
                        path = artifact.path("path").path("raw").asText(),
                        owners = artifact.path("ownerModules").map { it.asText() },
                    )
                }
            }

            override fun onSuccess() {
                artifactFacts = choices
                refreshArtifactChoices()
                targetPreview.text = "Loaded ${choices.size} artifacts."
            }

            override fun onThrowable(error: Throwable) {
                targetPreview.text = "Artifact load failed: ${error.message ?: error.javaClass.simpleName}"
            }

            override fun onFinished() {
                loadArtifactsButton.isEnabled = true
            }
        }.queue()
    }

    private fun refreshArtifactChoices() {
        val kind = artifactKind.selectedItem?.toString().orEmpty()
        artifactCatalog.removeAllItems()
        artifactFacts.filter { it.kind == kind }.sortedBy { it.id }.forEach(artifactCatalog::addItem)
    }

    private fun selectArtifact() {
        val choice = artifactCatalog.selectedItem as? ArtifactChoice ?: return
        targetArtifact.text = choice.name
        choice.owners.firstOrNull()?.let { targetModule.text = it }
        targetPreview.text = "${choice.path} · ${choice.owners.joinToString(", ").ifBlank { "no owner" }}"
    }

    private fun resolveTarget() {
        val target = currentDraft()
        val packRoot = metadataPackRoot().trim()
        val outRoot = metadataRoot().trim()
        if (target.id.isBlank()) {
            targetPreview.text = "Enter a target name first."
            return
        }
        if (packRoot.isBlank() && outRoot.isBlank()) {
            targetPreview.text = "Metadata Pack Root or Out Root is required."
            return
        }
        resolveTargetButton.isEnabled = false
        object : Task.Backgroundable(project, "Resolving AOSP target", true) {
            private var preview = ""

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Resolving target through ProjectGraph"
                val params = mapper.writeValueAsString(
                    mapOf(
                        "metadataPackRoot" to packRoot,
                        "metadataRoot" to outRoot,
                        "targets" to listOf(target.toTargetSpec()),
                    ),
                )
                val response = mapper.readTree(project.getService(ProjectEngineService::class.java).projectGraph(params)).path("result")
                val graph = response.path("graph")
                val targetProject = graph.path("targets").firstOrNull()
                val diagnostics = sequenceOf(
                    response.path("diagnostics"),
                    graph.path("diagnostics"),
                    targetProject?.path("diagnostics"),
                ).filterNotNull().sumOf { it.size() }
                val state = when {
                    diagnostics > 0 -> "error"
                    targetProject == null || targetProject.isMissingNode -> "unresolved"
                    else -> "exact"
                }
                preview = "$state · ${targetProject?.path("resolvedModules")?.size() ?: 0} modules · " +
                    "${targetProject?.path("resolvedArtifacts")?.size() ?: 0} artifacts · " +
                    "${targetProject?.path("sourceRoots")?.size() ?: 0} sources · " +
                    "${targetProject?.path("generatedRoots")?.size() ?: 0} generated · " +
                    "${targetProject?.path("libraries")?.size() ?: 0} libraries · $diagnostics diagnostics"
            }

            override fun onSuccess() {
                targetPreview.text = preview
            }

            override fun onThrowable(error: Throwable) {
                targetPreview.text = "Resolve failed: ${error.message ?: error.javaClass.simpleName}"
            }

            override fun onFinished() {
                resolveTargetButton.isEnabled = true
            }
        }.queue()
    }
}
