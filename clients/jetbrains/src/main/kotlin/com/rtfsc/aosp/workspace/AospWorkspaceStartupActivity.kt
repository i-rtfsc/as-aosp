package com.rtfsc.aosp.workspace

import com.intellij.openapi.project.Project
import com.intellij.openapi.startup.ProjectActivity
import com.rtfsc.aosp.profile.EngineProfileSettings

internal class AospWorkspaceStartupActivity : ProjectActivity {
    override suspend fun execute(project: Project) {
        val state = project.getService(EngineProfileSettings::class.java).current
        if (state.workspaceViewRoot.isNotBlank() && state.outRoot.isNotBlank()) {
            project.getService(WorkspaceCoordinatorService::class.java).start(false)
        }
    }
}
