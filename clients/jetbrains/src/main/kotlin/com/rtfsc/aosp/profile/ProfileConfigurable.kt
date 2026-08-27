package com.rtfsc.aosp.profile

import com.intellij.openapi.fileChooser.FileChooserDescriptor
import com.intellij.openapi.fileChooser.FileChooserDescriptorFactory
import com.intellij.openapi.options.Configurable
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import com.intellij.openapi.ui.TextFieldWithBrowseButton
import com.intellij.ui.components.JBTextField
import com.intellij.util.ui.FormBuilder
import com.rtfsc.aosp.engine.EngineProcess
import com.rtfsc.aosp.engine.EngineCommand
import com.rtfsc.aosp.project.WorkspaceRootPlanner
import com.rtfsc.aosp.project.FocusedRootsService
import java.awt.BorderLayout
import java.awt.FlowLayout
import java.awt.event.ActionListener
import javax.swing.JButton
import javax.swing.JComponent
import javax.swing.JLabel
import javax.swing.JPanel
import javax.swing.DefaultListModel
import javax.swing.JList
import javax.swing.JComboBox
import javax.swing.JScrollPane
import javax.swing.JCheckBox

internal class ProfileConfigurable(
    private val project: Project,
) : Configurable {
    private val settings = project.getService(EngineProfileSettings::class.java)
    private val profileId = JBTextField()
    private val workspaceViewRoot = TextFieldWithBrowseButton()
    private val focusedSourcePaths = JBTextField()
    private val canonicalRepoRoot = TextFieldWithBrowseButton()
    private val outRoot = TextFieldWithBrowseButton()
    private val metadataPackRoot = TextFieldWithBrowseButton()
    private val pathMappings = JBTextField()
    private val targetsJson = JBTextField()
    private val targetId = JBTextField()
    private val targetArtifact = JBTextField()
    private val targetModule = JBTextField()
    private val targetSources = JBTextField()
    private val targetKind = JComboBox(arrayOf("java", "android", "native", "mixed"))
    private val artifactKind = JComboBox(arrayOf("jar", "apk", "srcjar", "aar", "output"))
    private val artifactCatalog = JComboBox<String>()
    private val loadArtifactsButton = JButton("Load Artifacts")
    private val artifactFacts = mutableMapOf<String, Pair<String, String>>()
    private val includeDependencies = JCheckBox("Dependencies", true)
    private val includeGenerated = JCheckBox("Generated sources", true)
    private val includeResources = JCheckBox("Resources")
    private val includeManifest = JCheckBox("Manifest")
    private val targetListModel = DefaultListModel<String>()
    private val targetList = JList(targetListModel)
    private val addTargetButton = JButton("Add Target")
    private val removeTargetButton = JButton("Remove Target")
    private val engineExecutable = TextFieldWithBrowseButton()
    private val validateButton = JButton("Validate")
    private val applyRootsButton = JButton("Apply Focused Roots")
    private val applyBuildFactsButton = JButton("Apply Build Facts")
    private val panel = JPanel(BorderLayout())
    private val validateAction = ActionListener { validateInBackground() }

    init {
        configureChooser(workspaceViewRoot, folderDescriptor())
        configureChooser(canonicalRepoRoot, folderDescriptor())
        configureChooser(outRoot, folderDescriptor())
        configureChooser(metadataPackRoot, folderDescriptor())
        configureChooser(
            engineExecutable,
            FileChooserDescriptorFactory.createSingleFileOrExecutableAppDescriptor(),
        )
        validateButton.addActionListener(validateAction)
        applyRootsButton.addActionListener { applyFocusedRoots() }
        applyBuildFactsButton.addActionListener { applyBuildFacts() }
        addTargetButton.addActionListener { addTarget() }
        removeTargetButton.addActionListener { removeTarget() }
        loadArtifactsButton.addActionListener { loadArtifacts() }
        artifactKind.addActionListener { refreshArtifactChoices() }
        artifactCatalog.addActionListener { selectArtifact() }
        val targetsPanel = JPanel(BorderLayout()).apply {
            add(JScrollPane(targetList), BorderLayout.CENTER)
            val editor = FormBuilder.createFormBuilder()
                .addLabeledComponent(JLabel("Target Name"), targetId)
                .addLabeledComponent(JLabel("Kind"), targetKind)
                .addLabeledComponent(JLabel("Artifact Kind"), artifactKind)
                .addLabeledComponent(JLabel("Artifact From Pack"), artifactCatalog)
                .addLabeledComponent(JLabel("Artifact"), targetArtifact)
                .addLabeledComponent(JLabel("Soong Module"), targetModule)
                .addLabeledComponent(JLabel("Source Roots"), targetSources)
                .panel
            val buttons = JPanel(FlowLayout(FlowLayout.LEFT))
            buttons.add(includeDependencies)
            buttons.add(includeGenerated)
            buttons.add(includeResources)
            buttons.add(includeManifest)
            buttons.add(loadArtifactsButton)
            buttons.add(addTargetButton)
            buttons.add(removeTargetButton)
            val container = JPanel(BorderLayout())
            container.add(editor, BorderLayout.CENTER)
            container.add(buttons, BorderLayout.SOUTH)
            add(container, BorderLayout.SOUTH)
        }
        val form = FormBuilder.createFormBuilder()
            .addLabeledComponent(JLabel("Profile ID"), profileId)
            .addLabeledComponent(JLabel("Workspace View Root"), workspaceViewRoot)
            .addLabeledComponent(JLabel("Focused Source Paths"), focusedSourcePaths)
            .addLabeledComponent(JLabel("Canonical AOSP Root"), canonicalRepoRoot)
            .addLabeledComponent(JLabel("Out Root"), outRoot)
            .addLabeledComponent(JLabel("Metadata Pack Root"), metadataPackRoot)
            .addLabeledComponent(JLabel("Path Mappings JSON"), pathMappings)
            .addLabeledComponent(JLabel("Engine Executable"), engineExecutable)
            .panel
        val targetsContainer = JPanel(BorderLayout())
        targetsContainer.add(JLabel("Targets"), BorderLayout.NORTH)
        targetsContainer.add(targetsPanel, BorderLayout.CENTER)
        val content = JPanel(BorderLayout())
        content.add(form, BorderLayout.NORTH)
        content.add(targetsContainer, BorderLayout.CENTER)
        val actions = JPanel(FlowLayout(FlowLayout.LEFT))
        actions.add(validateButton)
        actions.add(applyRootsButton)
        actions.add(applyBuildFactsButton)
        panel.add(content, BorderLayout.CENTER)
        panel.add(actions, BorderLayout.SOUTH)
    }

    override fun getDisplayName(): String = "AOSP Workspace"

    override fun createComponent(): JComponent = panel

    override fun isModified(): Boolean = draftState() != settings.current

    override fun apply() {
        val draft = draftState()
        settings.update { state ->
            state.profileId = draft.profileId
            state.workspaceViewRoot = draft.workspaceViewRoot
            state.focusedSourcePaths = draft.focusedSourcePaths
            state.canonicalRepoRoot = draft.canonicalRepoRoot
            state.outRoot = draft.outRoot
            state.metadataPackRoot = draft.metadataPackRoot
            state.pathMappings = draft.pathMappings
            state.targetsJson = draft.targetsJson
            state.engineExecutable = draft.engineExecutable
        }
    }

    override fun reset() {
        val state = settings.current
        profileId.text = state.profileId
        workspaceViewRoot.text = state.workspaceViewRoot
        focusedSourcePaths.text = state.focusedSourcePaths
        canonicalRepoRoot.text = state.canonicalRepoRoot
        outRoot.text = state.outRoot
        metadataPackRoot.text = state.metadataPackRoot
        pathMappings.text = state.pathMappings
        targetsJson.text = state.targetsJson
        loadTargets(state.targetsJson)
        engineExecutable.text = state.engineExecutable
    }

    override fun disposeUIResources() {
        validateButton.removeActionListener(validateAction)
    }

    private fun draftState(): EngineProfileState {
        targetsJson.text = serializeTargets()
        return EngineProfileState(
            profileId = profileId.text,
            workspaceViewRoot = workspaceViewRoot.text,
            focusedSourcePaths = focusedSourcePaths.text,
            canonicalRepoRoot = canonicalRepoRoot.text,
            outRoot = outRoot.text,
            metadataPackRoot = metadataPackRoot.text,
            pathMappings = pathMappings.text,
            targetsJson = targetsJson.text,
            engineExecutable = engineExecutable.text,
        )
    }

    private fun addTarget() {
        val id = targetId.text.trim()
        if (id.isBlank()) return
        val kind = targetKind.selectedItem.toString()
        val artifact = targetArtifact.text.trim()
        val module = targetModule.text.trim()
        val sources = targetSources.text.trim()
        val summary = listOf(id, kind, artifactKind.selectedItem.toString(), artifact, module, sources, includeDependencies.isSelected.toString(), includeGenerated.isSelected.toString(), includeResources.isSelected.toString(), includeManifest.isSelected.toString()).joinToString("|")
        targetListModel.addElement(summary)
        targetId.text = ""
        targetArtifact.text = ""
        targetModule.text = ""
        targetSources.text = ""
    }

    private fun loadArtifacts() {
        val packRoot = metadataPackRoot.text.trim()
        if (packRoot.isBlank()) return
        val mapper = com.fasterxml.jackson.databind.ObjectMapper()
        val params = mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot))
        val response = runCatching {
            project.getService(com.rtfsc.aosp.engine.ProjectEngineService::class.java)
                .artifactsCatalog(params)
        }.getOrNull() ?: return
        artifactFacts.clear()
        mapper.readTree(response).path("result").path("artifacts").forEach { artifact ->
            val key = artifact.path("id").asText()
            val display = artifact.path("name").asText() + " · " +
                artifact.path("ownerModules").map { it.asText() }.joinToString(", " )
            artifactFacts[key] = artifact.path("kind").asText() to display
        }
        refreshArtifactChoices()
    }

    private fun refreshArtifactChoices() {
        val kind = artifactKind.selectedItem?.toString().orEmpty()
        artifactCatalog.removeAllItems()
        artifactCatalog.addItem("Custom artifact")
        artifactFacts.filterValues { it.first == kind }.toSortedMap().forEach { (id, value) ->
            artifactCatalog.addItem(id + "|" + value.second)
        }
    }

    private fun selectArtifact() {
        val selected = artifactCatalog.selectedItem?.toString().orEmpty()
        if (selected.isBlank() || selected == "Custom artifact") return
        val id = selected.substringBefore('|')
        targetArtifact.text = id.substringAfter(':').substringAfterLast('/')
        val owner = selected.substringAfter('|').substringAfter(" · " ).substringBefore(',').trim()
        if (owner.isNotBlank()) targetModule.text = owner
    }

    private fun removeTarget() {
        val index = targetList.selectedIndex
        if (index >= 0) targetListModel.remove(index)
    }

    private fun serializeTargets(): String {
        val targets = targetListModel.elements().toList().map { value ->
            val parts = value.split('|')
            val id = parts.getOrElse(0) { "" }
            val kind = parts.getOrElse(1) { "java" }
            val artifactType = parts.getOrElse(2) { "jar" }
            val artifact = parts.getOrElse(3) { "" }
            val module = parts.getOrElse(4) { "" }
            val sources = parts.getOrElse(5) { "" }.split(',').map(String::trim).filter(String::isNotEmpty)
            mutableMapOf<String, Any>(
                "id" to id, "displayName" to id, "kind" to kind,
                "selectors" to mapOf(
                    "modules" to if (module.isBlank()) emptyList<String>() else listOf(module),
                    "artifacts" to if (artifact.isBlank()) emptyList<Any>() else listOf(mapOf("kind" to artifactType, "name" to artifact)),
                    "sourceRoots" to sources,
                ),
                "includeDependencies" to parts.getOrElse(6) { "true" }.toBoolean(),
                "includeGeneratedSources" to parts.getOrElse(7) { "true" }.toBoolean(),
                "includeResources" to parts.getOrElse(8) { "false" }.toBoolean(),
                "includeManifest" to parts.getOrElse(9) { "false" }.toBoolean(),
            )
        }
        return com.fasterxml.jackson.databind.ObjectMapper().writeValueAsString(targets)
    }

    private fun loadTargets(raw: String) {
        targetListModel.clear()
        val node = runCatching { com.fasterxml.jackson.databind.ObjectMapper().readTree(raw) }.getOrNull() ?: return
        if (!node.isArray) return
        node.forEach { target ->
            val selectors = target.path("selectors")
            val artifactNode = selectors.path("artifacts").firstOrNull()
            val artifact = artifactNode?.path("name")?.asText().orEmpty()
            val artifactType = artifactNode?.path("kind")?.asText("jar") ?: "jar"
            val module = selectors.path("modules").firstOrNull()?.asText().orEmpty()
            val sources = selectors.path("sourceRoots").map { it.asText() }.joinToString(",")
            targetListModel.addElement(listOf(target.path("id").asText(), target.path("kind").asText("java"), artifactType, artifact, module, sources, target.path("includeDependencies").asBoolean(true).toString(), target.path("includeGeneratedSources").asBoolean(true).toString(), target.path("includeResources").asBoolean(false).toString(), target.path("includeManifest").asBoolean(false).toString()).joinToString("|"))
        }
    }

    private fun validateInBackground() {
        val draft = draftState()
        val errors = draft.localValidationErrors()
        if (errors.isNotEmpty()) {
            Messages.showWarningDialog(panel, errors.joinToString("\n"), "AOSP IDE")
            return
        }
        val plan = WorkspaceRootPlanner.plan(draft)
        if (plan.diagnostics.isNotEmpty()) {
            Messages.showWarningDialog(panel, plan.diagnostics.joinToString("\n"), "AOSP IDE")
            return
        }
        validateButton.isEnabled = false
        object : Task.Backgroundable(project, "Validating AOSP IDE profile", true) {
            private var diagnosticCount: Int = 0

            override fun run(indicator: ProgressIndicator) {
                val process = EngineProcess()
                try {
                    indicator.text = "Starting engine"
                    EngineCommand.start(process, draft.engineExecutable)
                    indicator.text = "Validating profile"
                    val response = process.request(
                        "aosp/profile/validate",
                        ProfileJson.validationParams(draft),
                    )
                    diagnosticCount = ProfileJson.diagnosticCount(response)
                    indicator.text = "Profile validation completed"
                } finally {
                    process.stop()
                }
            }

            override fun onSuccess() {
                Messages.showInfoMessage(
                    panel,
                    "Profile validation completed with $diagnosticCount diagnostic(s).",
                    "AOSP IDE",
                )
            }

            override fun onThrowable(error: Throwable) {
                Messages.showErrorDialog(
                    panel,
                    error.message ?: error.javaClass.simpleName,
                    "AOSP IDE profile validation failed",
                )
            }

            override fun onFinished() {
                validateButton.isEnabled = true
            }
        }.queue()
    }

    private fun applyFocusedRoots() {
        val errors = draftState().localValidationErrors()
        if (errors.isNotEmpty()) {
            Messages.showWarningDialog(panel, errors.joinToString("\n"), "AOSP IDE")
            return
        }
        val plan = WorkspaceRootPlanner.plan(draftState())
        if (!plan.isUsable) {
            Messages.showWarningDialog(panel, plan.diagnostics.joinToString("\n"), "AOSP IDE")
            return
        }
        val applied = project.getService(FocusedRootsService::class.java).apply()
        Messages.showInfoMessage(
            panel,
            "Applied ${applied.sourceRoots.size} focused source root(s).",
            "AOSP IDE",
        )
    }

    private fun applyBuildFacts() {
        apply()
        applyBuildFactsButton.isEnabled = false
        object : Task.Backgroundable(project, "Applying AOSP Build Facts", true) {
            private var message = ""

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Loading module catalog and generated sources"
                val result = project.getService(FocusedRootsService::class.java).applyBuildFacts()
                message = "Applied ${result.sourceRoots.size} source root(s), " +
                    "${result.generatedRoots.size} generated root(s), " +
                    "${result.libraries.size} library(s), skipped ${result.skipped}."
                if (result.diagnostics.isNotEmpty()) {
                    message += " Diagnostics: ${result.diagnostics.size}."
                }
            }

            override fun onSuccess() {
                Messages.showInfoMessage(panel, message, "AOSP IDE")
            }

            override fun onThrowable(error: Throwable) {
                Messages.showErrorDialog(panel, error.message ?: error.javaClass.simpleName, "AOSP Build Facts failed")
            }

            override fun onFinished() {
                applyBuildFactsButton.isEnabled = true
            }
        }.queue()
    }

    private fun configureChooser(
        field: TextFieldWithBrowseButton,
        descriptor: FileChooserDescriptor,
    ) {
        field.addBrowseFolderListener(project, descriptor)
    }

    private fun folderDescriptor() = FileChooserDescriptorFactory.createSingleFolderDescriptor()
}
