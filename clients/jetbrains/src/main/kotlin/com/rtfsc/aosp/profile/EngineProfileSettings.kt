package com.rtfsc.aosp.profile

import com.intellij.openapi.components.PersistentStateComponent
import com.intellij.openapi.components.Service
import com.intellij.openapi.components.State
import com.intellij.openapi.components.Storage
import com.intellij.openapi.components.StoragePathMacros
import java.util.concurrent.CopyOnWriteArrayList

@State(
    name = "AospIdeEngineProfile",
    storages = [Storage(StoragePathMacros.WORKSPACE_FILE)],
)
@Service(Service.Level.PROJECT)
internal class EngineProfileSettings : PersistentStateComponent<EngineProfileState> {
    private var state = EngineProfileState()
    private val listeners = CopyOnWriteArrayList<(EngineProfileState) -> Unit>()

    val current: EngineProfileState
        get() = state

    override fun getState(): EngineProfileState = state

    override fun loadState(state: EngineProfileState) {
        this.state = state
        notifyListeners()
    }

    fun update(transform: (EngineProfileState) -> Unit) {
        transform(state)
        notifyListeners()
    }

    fun addListener(listener: (EngineProfileState) -> Unit): AutoCloseable {
        listeners += listener
        listener(state)
        return AutoCloseable { listeners -= listener }
    }

    private fun notifyListeners() {
        listeners.forEach { it(state) }
    }
}
