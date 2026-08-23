package io.github.irtfsc.aospide.profile

import com.intellij.openapi.components.PersistentStateComponent
import com.intellij.openapi.components.Service
import com.intellij.openapi.components.State
import com.intellij.openapi.components.Storage
import com.intellij.openapi.components.StoragePathMacros

@State(
    name = "AospIdeEngineProfile",
    storages = [Storage(StoragePathMacros.WORKSPACE_FILE)],
)
@Service(Service.Level.PROJECT)
internal class EngineProfileSettings : PersistentStateComponent<EngineProfileState> {
    private var state = EngineProfileState()

    val current: EngineProfileState
        get() = state

    override fun getState(): EngineProfileState = state

    override fun loadState(state: EngineProfileState) {
        this.state = state
    }

    fun update(transform: (EngineProfileState) -> Unit) {
        transform(state)
    }
}
