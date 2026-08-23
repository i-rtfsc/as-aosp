package io.github.irtfsc.aospide.profile

data class EngineProfileState(
    var profileId: String = "",
    var workspaceViewRoot: String = "",
    var canonicalRepoRoot: String = "",
    var outRoot: String = "",
    var engineExecutable: String = "",
) {
    fun localValidationErrors(): List<String> {
        return buildList {
            if (profileId.isBlank()) add("Profile ID is required.")
            if (workspaceViewRoot.isBlank()) add("Workspace View Root is required.")
            if (engineExecutable.isBlank()) add("Engine executable is required.")
        }
    }
}
