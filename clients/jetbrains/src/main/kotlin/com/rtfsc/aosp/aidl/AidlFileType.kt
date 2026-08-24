package com.rtfsc.aosp.aidl

import com.intellij.openapi.fileTypes.LanguageFileType
import javax.swing.Icon

object AidlFileType : LanguageFileType(AidlLanguage) {
    override fun getName(): String = "AIDL"

    override fun getDescription(): String = "Android Interface Definition Language"

    override fun getDefaultExtension(): String = "aidl"

    override fun getIcon(): Icon? = null
}
