package com.rtfsc.aosp.workspace

import com.intellij.openapi.components.Service
import com.intellij.openapi.Disposable
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.rtfsc.aosp.profile.EngineProfileSettings
import com.rtfsc.aosp.project.FocusedRootsService
import java.util.concurrent.CopyOnWriteArrayList
import java.util.concurrent.atomic.AtomicLong
import java.util.concurrent.ScheduledExecutorService
import java.util.concurrent.Executors
import java.util.concurrent.TimeUnit

internal data class WorkspaceStatus(
    val state: String = "idle",
    val sourceRoots: Int = 0,
    val generatedRoots: Int = 0,
    val libraries: Int = 0,
    val skipped: Int = 0,
    val diagnostics: Int = 0,
    val message: String = "Not started",
)

@Service(Service.Level.PROJECT)
internal class WorkspaceCoordinatorService(
    private val project: Project,
) : Disposable {
    private val generation = AtomicLong()
    private val listeners = CopyOnWriteArrayList<(WorkspaceStatus) -> Unit>()
    private val scheduler: ScheduledExecutorService = Executors.newSingleThreadScheduledExecutor()
    private var settingsSubscription: AutoCloseable? = null
    private var pending: java.util.concurrent.ScheduledFuture<*>? = null

    @Volatile
    var status: WorkspaceStatus = WorkspaceStatus()
        private set

    init {
        var initial = true
        settingsSubscription = project.getService(EngineProfileSettings::class.java).addListener {
            if (initial) {
                initial = false
                return@addListener
            }
            pending?.cancel(false)
            pending = scheduler.schedule({ start(false) }, 300, TimeUnit.MILLISECONDS)
        }
    }

    fun addListener(listener: (WorkspaceStatus) -> Unit): AutoCloseable {
        listeners += listener
        listener(status)
        return AutoCloseable { listeners -= listener }
    }

    fun start(rebuild: Boolean) {
        val state = project.getService(EngineProfileSettings::class.java).current
        if (state.workspaceViewRoot.isBlank()) return
        val current = generation.incrementAndGet()
        publish(WorkspaceStatus(state = "running", message = if (rebuild) "Rebuilding Build Facts" else "Applying Build Facts"))
        object : Task.Backgroundable(project, "AOSP Workspace indexing", true) {
            private var completed = WorkspaceStatus(state = "error", message = "Indexing failed")

            override fun run(indicator: ProgressIndicator) {
                indicator.text = "Loading module and generated source facts"
                val result = project.getService(FocusedRootsService::class.java).applyBuildFacts()
                val stateName = if (result.diagnostics.isEmpty()) "ready" else "partial"
                completed = WorkspaceStatus(
                    state = stateName,
                    sourceRoots = result.sourceRoots.size,
                    generatedRoots = result.generatedRoots.size,
                    libraries = result.libraries.size,
                    skipped = result.skipped,
                    diagnostics = result.diagnostics.size,
                    message = "${result.sourceRoots.size} source roots, ${result.generatedRoots.size} generated",
                )
            }

            override fun onSuccess() {
                if (generation.get() == current) publish(completed)
            }

            override fun onThrowable(error: Throwable) {
                if (generation.get() == current) {
                    publish(WorkspaceStatus(state = "error", diagnostics = 1, message = error.message ?: error.javaClass.simpleName))
                }
            }
        }.queue()
    }

    override fun dispose() {
        settingsSubscription?.close()
        pending?.cancel(false)
        scheduler.shutdownNow()
    }

    private fun publish(value: WorkspaceStatus) {
        status = value
        listeners.forEach { it(value) }
    }
}
