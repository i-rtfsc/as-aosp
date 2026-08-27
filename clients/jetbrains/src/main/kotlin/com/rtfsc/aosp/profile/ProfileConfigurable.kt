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
import com.intellij.ide.plugins.PluginManagerCore
import com.intellij.openapi.extensions.PluginId
import com.rtfsc.aosp.project.WorkspaceRootPlanner
import com.rtfsc.aosp.project.FocusedRootsService
import java.awt.BorderLayout
import java.awt.FlowLayout
import java.awt.event.ActionListener
import javax.swing.JButton
import javax.swing.JComponent
import javax.swing.JLabel
import javax.swing.JPanel

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
    private val targetEditor = TargetEditorPanel(project, { metadataPackRoot.text }, { outRoot.text })
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
        val content = JPanel(BorderLayout())
        content.add(form, BorderLayout.NORTH)
        content.add(targetEditor, BorderLayout.CENTER)
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
        targetEditor.load(state.targetsJson)
        engineExecutable.text = state.engineExecutable
    }

    override fun disposeUIResources() {
        validateButton.removeActionListener(validateAction)
    }

    private fun draftState(): EngineProfileState {
        return EngineProfileState(
            profileId = profileId.text,
            workspaceViewRoot = workspaceViewRoot.text,
            focusedSourcePaths = focusedSourcePaths.text,
            canonicalRepoRoot = canonicalRepoRoot.text,
            outRoot = outRoot.text,
            metadataPackRoot = metadataPackRoot.text,
            pathMappings = pathMappings.text,
            targetsJson = targetEditor.serialize(),
            engineExecutable = engineExecutable.text,
        )
    }

    private fun validateInBackground() {
        val draft = draftState()
        val errors = draft.localValidationErrors()
        if (errors.isNotEmpty()) {
            Messages.showWarningDialog(panel, errors.joinToString("\n"), "AOSP Workspace")
            return
        }
        val plan = WorkspaceRootPlanner.plan(draft)
        if (plan.diagnostics.isNotEmpty()) {
            Messages.showWarningDialog(panel, plan.diagnostics.joinToString("\n"), "AOSP Workspace")
            return
        }
        validateButton.isEnabled = false
        object : Task.Backgroundable(project, "Validating AOSP Workspace profile", true) {
            private var diagnosticCount: Int = 0

            override fun run(indicator: ProgressIndicator) {
                val process = EngineProcess()
                try {
                    indicator.text = "Starting engine"
                    val executable = EngineCommand.resolveExecutable(
                        draft.engineExecutable,
                        System.getProperty("user.home"),
                        PluginManagerCore.getPlugin(PluginId.getId("com.rtfsc.aosp"))?.pluginPath
                            ?.let(EngineCommand::bundledCandidate),
                    )
                    EngineCommand.start(process, executable)
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
                    "AOSP Workspace",
                )
            }

            override fun onThrowable(error: Throwable) {
                Messages.showErrorDialog(
                    panel,
                    error.message ?: error.javaClass.simpleName,
                    "AOSP Workspace profile validation failed",
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
            Messages.showWarningDialog(panel, errors.joinToString("\n"), "AOSP Workspace")
            return
        }
        val plan = WorkspaceRootPlanner.plan(draftState())
        if (!plan.isUsable) {
            Messages.showWarningDialog(panel, plan.diagnostics.joinToString("\n"), "AOSP Workspace")
            return
        }
        val applied = project.getService(FocusedRootsService::class.java).apply()
        Messages.showInfoMessage(
            panel,
            "Applied ${applied.sourceRoots.size} focused source root(s).",
            "AOSP Workspace",
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
                Messages.showInfoMessage(panel, message, "AOSP Workspace")
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
