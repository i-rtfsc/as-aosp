package com.rtfsc.aosp.aidl

import com.intellij.testFramework.fixtures.BasePlatformTestCase
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull

class AidlGotoDeclarationPlatformTest : BasePlatformTestCase() {
    fun testResolvesImportedSourceDeclarationWithoutOut() {
        myFixture.addFileToProject(
            "android/demo/IFoo.aidl",
            "package android.demo; interface IFoo {}",
        )
        val source = myFixture.configureByText(
            "Client.aidl",
            "package android.client; import android.demo.IFoo; interface Client { IFoo value(); }",
        )
        assertEquals(AidlFileType, source.virtualFile.fileType)

        val offset = source.text.indexOf("IFoo value")
        val element = source.findElementAt(offset)
        val targets = AidlGotoDeclarationHandler().getGotoDeclarationTargets(
            element,
            offset,
            myFixture.editor,
        )

        assertNotNull(targets)
        assertEquals(1, targets!!.size)
        assertEquals("IFoo.aidl", targets.single().containingFile.name)
        assertEquals("package android.demo; interface IFoo {}", targets.single().text)
    }
}
