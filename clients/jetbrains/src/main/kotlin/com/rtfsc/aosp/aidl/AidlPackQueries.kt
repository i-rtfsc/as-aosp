package com.rtfsc.aosp.aidl

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.rtfsc.aosp.engine.ProjectEngineService

internal object AidlPackQueries {
    private val mapper = ObjectMapper()

    fun resolve(engine: ProjectEngineService, packRoot: String, name: String): JsonNode? {
        if (packRoot.isBlank()) return null
        val response = engine.aidlPackResolve(params(packRoot, name))
        val result = mapper.readTree(response).path("result")
        return result.takeIf {
            it.path("diagnostics").isEmpty && it.path("state").asText() == "found"
        }
    }

    fun references(engine: ProjectEngineService, packRoot: String, name: String): JsonNode? {
        if (packRoot.isBlank()) return null
        val response = engine.aidlPackReferences(params(packRoot, name))
        val result = mapper.readTree(response).path("result")
        return result.takeIf { it.path("diagnostics").isEmpty }
    }

    fun peers(engine: ProjectEngineService, packRoot: String, name: String, cacheRoot: String): JsonNode? {
        if (packRoot.isBlank()) return null
        val response = engine.aidlPackPeers(mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot, "name" to name, "cacheRoot" to cacheRoot)))
        val result = mapper.readTree(response).path("result")
        return result.takeIf {
            it.path("diagnostics").isEmpty && it.path("peers").any { peer ->
                peer.path("confidence").asText() == "exact" &&
                    peer.path("exists").asBoolean(false)
            }
        }
    }

    private fun params(packRoot: String, name: String): String {
        return mapper.writeValueAsString(mapOf("metadataPackRoot" to packRoot, "name" to name))
    }
}
