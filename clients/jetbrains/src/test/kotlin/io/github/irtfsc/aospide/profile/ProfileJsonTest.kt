package io.github.irtfsc.aospide.profile

import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class ProfileJsonTest {
    @Test
    fun serializesWindowsPathsWithoutChangingSeparators() {
        val params = ProfileJson.validationParams(
            EngineProfileState(
                profileId = "framework",
                workspaceViewRoot = "C:\\aosp view",
                canonicalRepoRoot = "D:\\aosp",
                outRoot = "Z:\\out",
            ),
        )

        assertTrue(params.contains("C:\\\\aosp view"))
        assertTrue(params.contains("D:\\\\aosp"))
        assertTrue(params.contains("\"mode\":\"strict-view\""))
    }

    @Test
    fun readsDiagnosticCountFromValidationResponse() {
        val response = "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"diagnostics\":[{},{}]}}"

        assertEquals(2, ProfileJson.diagnosticCount(response.toByteArray()))
    }

    @Test
    fun reportsRequiredLocalFieldsBeforeStartingTheEngine() {
        val errors = EngineProfileState().localValidationErrors()

        assertEquals(3, errors.size)
    }
}
