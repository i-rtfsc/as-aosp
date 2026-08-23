package io.github.irtfsc.aospide.profile

data class EngineProfileState(
    var profileId: String = "",
    var workspaceViewRoot: String = "",
    var canonicalRepoRoot: String = "",
    var outRoot: String = "",
    var engineExecutable: String = "",
)
