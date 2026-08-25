package com.rtfsc.aosp.profile

import com.fasterxml.jackson.databind.ObjectMapper

internal object ProfileJson {
    private val mapper = ObjectMapper()

    fun validationParams(state: EngineProfileState): String {
        return mapper.writeValueAsString(
            mapOf(
                "profile" to mapOf(
                    "schemaVersion" to 1,
                    "id" to state.profileId,
                    "workspaceViewRoot" to state.workspaceViewRoot,
                    "scope" to mapOf(
                        "mode" to "selected-repositories",
                        "include" to state.focusedSourcePaths.split(',')
                            .map(String::trim).filter(String::isNotEmpty),
                    ),
                    "canonicalRepoRoot" to state.canonicalRepoRoot,
                    "outRoot" to state.outRoot,
                    "metadataPackRoot" to state.metadataPackRoot,
                    "visibilityPolicy" to mapOf("mode" to "strict-view"),
                    "buildIntegration" to mapOf("mode" to "read-only"),
                ),
            ),
        )
    }

    fun diagnosticCount(response: ByteArray): Int {
        val root = mapper.readTree(response)
        return root.path("result").path("diagnostics").size()
    }
}
