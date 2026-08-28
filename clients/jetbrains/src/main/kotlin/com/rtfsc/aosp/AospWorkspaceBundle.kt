package com.rtfsc.aosp

import com.intellij.DynamicBundle
import org.jetbrains.annotations.PropertyKey

internal object AospWorkspaceBundle : DynamicBundle("messages.AospWorkspaceBundle") {
    fun message(
        @PropertyKey(resourceBundle = "messages.AospWorkspaceBundle") key: String,
        vararg params: Any,
    ): String = getMessage(key, *params)
}
