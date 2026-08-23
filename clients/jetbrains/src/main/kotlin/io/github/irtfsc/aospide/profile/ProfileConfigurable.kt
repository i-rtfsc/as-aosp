package io.github.irtfsc.aospide.profile

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
import io.github.irtfsc.aospide.engine.EngineProcess
import java.awt.BorderLayout
import java.awt.FlowLayout
import java.awt.event.ActionListener
import java.nio.file.Path
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
    private val canonicalRepoRoot = TextFieldWithBrowseButton()
    private val outRoot = TextFieldWithBrowseButton()
    private val engineExecutable = TextFieldWithBrowseButton()
    private val validateButton = JButton("Validate")
    private val panel = JPanel(BorderLayout())
    private val validateAction = ActionListener { validateInBackground() }

    init {
        configureChooser(workspaceViewRoot, folderDescriptor())
        configureChooser(canonicalRepoRoot, folderDescriptor())
        configureChooser(outRoot, folderDescriptor())
        configureChooser(
            engineExecutable,
            FileChooserDescriptorFactory.createSingleFileOrExecutableAppDescriptor(),
        )
        validateButton.addActionListener(validateAction)
        val form = FormBuilder.createFormBuilder()
            .addLabeledComponent(JLabel("Profile ID"), profileId)
            .addLabeledComponent(JLabel("Workspace View Root"), workspaceViewRoot)
            .addLabeledComponent(JLabel("Canonical AOSP Root"), canonicalRepoRoot)
            .addLabeledComponent(JLabel("Out Root"), outRoot)
            .addLabeledComponent(JLabel("Engine Executable"), engineExecutable)
            .panel
        val actions = JPanel(FlowLayout(FlowLayout.LEFT))
        actions.add(validateButton)
        panel.add(form, BorderLayout.NORTH)
        panel.add(actions, BorderLayout.SOUTH)
    }

    override fun getDisplayName(): String = "AOSP IDE"

    override fun createComponent(): JComponent = panel

    override fun isModified(): Boolean = draftState() != settings.current

    override fun apply() {
        val draft = draftState()
        settings.update { state ->
            state.profileId = draft.profileId
            state.workspaceViewRoot = draft.workspaceViewRoot
            state.canonicalRepoRoot = draft.canonicalRepoRoot
            state.outRoot = draft.outRoot
            state.engineExecutable = draft.engineExecutable
        }
    }

    override fun reset() {
        val state = settings.current
        profileId.text = state.profileId
        workspaceViewRoot.text = state.workspaceViewRoot
        canonicalRepoRoot.text = state.canonicalRepoRoot
        outRoot.text = state.outRoot
        engineExecutable.text = state.engineExecutable
    }

    override fun disposeUIResources() {
        validateButton.removeActionListener(validateAction)
    }

    private fun draftState(): EngineProfileState {
        return EngineProfileState(
            profileId = profileId.text,
            workspaceViewRoot = workspaceViewRoot.text,
            canonicalRepoRoot = canonicalRepoRoot.text,
            outRoot = outRoot.text,
            engineExecutable = engineExecutable.text,
        )
    }

    private fun validateInBackground() {
        val draft = draftState()
        val errors = draft.localValidationErrors()
        if (errors.isNotEmpty()) {
            Messages.showWarningDialog(panel, errors.joinToString("\n"), "AOSP IDE")
            return
        }
        validateButton.isEnabled = false
        object : Task.Backgroundable(project, "Validating AOSP IDE profile", true) {
            private var diagnosticCount: Int = 0

            override fun run(indicator: ProgressIndicator) {
                val process = EngineProcess()
                try {
                    indicator.text = "Starting engine"
                    process.start(Path.of(draft.engineExecutable))
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

    private fun configureChooser(
        field: TextFieldWithBrowseButton,
        descriptor: FileChooserDescriptor,
    ) {
        field.addBrowseFolderListener(project, descriptor)
    }

    private fun folderDescriptor() = FileChooserDescriptorFactory.createSingleFolderDescriptor()
}
