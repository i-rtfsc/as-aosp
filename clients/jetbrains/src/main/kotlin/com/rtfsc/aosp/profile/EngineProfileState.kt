package com.rtfsc.aosp.profile

data class EngineProfileState(
    var profileId: String = "",
    var workspaceViewRoot: String = "",
    var focusedSourcePaths: String = "",
    var canonicalRepoRoot: String = "",
    var outRoot: String = "",
    var metadataPackRoot: String = "",
    var pathMappings: String = "",
    var targetsJson: String = "",
    var engineExecutable: String = "",
) {
    fun localValidationErrors(): List<String> {
        return buildList {
            if (profileId.isBlank()) add("Profile ID is required.")
            if (workspaceViewRoot.isBlank()) add("Workspace View Root is required.")
            if (targetsJson.isBlank()) add("At least one TargetSpec is required.")
        }
    }
}
