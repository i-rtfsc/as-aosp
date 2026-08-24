package com.rtfsc.aosp.aidl

import com.intellij.codeInsight.navigation.actions.GotoDeclarationHandler
import com.intellij.openapi.editor.Editor
import com.intellij.openapi.fileTypes.FileTypeManager
import com.intellij.openapi.project.Project
import com.intellij.openapi.vfs.VirtualFile
import com.intellij.psi.PsiElement
import com.intellij.psi.PsiFile
import com.intellij.psi.PsiManager
import com.intellij.psi.search.FilenameIndex
import com.intellij.psi.search.GlobalSearchScope
import com.fasterxml.jackson.databind.ObjectMapper
import com.rtfsc.aosp.engine.ProjectEngineService
import com.rtfsc.aosp.profile.EngineProfileSettings
import java.nio.file.Path

private data class AidlDeclaration(
    val qualifiedName: String,
    val nameStart: Int,
)

private object AidlSourceResolver {
    private val packagePattern = Regex("(?m)\\bpackage\\s+([A-Za-z_]\\w*(?:\\.[A-Za-z_]\\w*)*)\\s*;")
    private val importPattern = Regex("(?m)\\bimport\\s+([A-Za-z_]\\w*(?:\\.[A-Za-z_]\\w*)*)\\s*;")
    private val declarationPattern = Regex("\\b(?:interface|parcelable|enum|union)\\s+([A-Za-z_]\\w*)")
    private val identifierPattern = Regex("[A-Za-z_]\\w*")

    fun resolve(source: PsiFile, offset: Int): PsiElement? {
        val sourceText = source.text
        val token = identifierAt(sourceText, offset) ?: return null
        resolveWithEngine(source, token)?.let { return it }
        val packageName = packagePattern.find(sourceText)?.groupValues?.get(1).orEmpty()
        val imports = importPattern.findAll(sourceText).map { it.groupValues[1] }.toList()
        val project = source.project
        val files = FilenameIndex.getAllFilesByExt(
            project,
            AidlFileType.defaultExtension,
            GlobalSearchScope.projectScope(project),
        )
        val candidates = files.mapNotNull { file -> declaration(file, project) }
        val ranked = candidates.filter { target ->
            val imported = imports.any { it == target.declaration.qualifiedName && it.substringAfterLast('.') == token }
            val samePackage = packageName.isNotEmpty() &&
                target.declaration.qualifiedName == "$packageName.$token"
            val exact = target.declaration.qualifiedName == token
            imported || samePackage || exact || target.declaration.qualifiedName.substringAfterLast('.') == token
        }.distinctBy { it.declaration.qualifiedName to it.file }
        val ordered = ranked.sortedWith(
            compareBy<AidlTarget> { target ->
                when {
                    imports.any { imported -> imported == target.declaration.qualifiedName } -> 0
                    packageName.isNotEmpty() &&
                        target.declaration.qualifiedName == "$packageName.$token" -> 1
                    target.declaration.qualifiedName == token -> 2
                    else -> 3
                }
            }.thenBy { target -> target.declaration.qualifiedName }
                .thenBy { target -> target.file.path },
        )
        if (ordered.size != 1) return null
        val target = ordered.single()
        val psiFile = PsiManager.getInstance(project).findFile(target.file) ?: return null
        return psiFile.findElementAt(target.declaration.nameStart)
    }

    private fun resolveWithEngine(source: PsiFile, token: String): PsiElement? {
        val project = source.project
        val settings = project.getService(EngineProfileSettings::class.java).current
        if (settings.engineExecutable.isBlank() || settings.workspaceViewRoot.isBlank()) return null
        val root = runCatching { Path.of(settings.workspaceViewRoot).toAbsolutePath().normalize() }.getOrNull() ?: return null
        val currentPath = runCatching {
            val filePath = source.virtualFile.toNioPath().toAbsolutePath().normalize()
            if (!filePath.startsWith(root)) return null
            root.relativize(filePath).toString().replace('\\', '/')
        }.getOrNull() ?: return null
        val request = ObjectMapper().createObjectNode().apply {
            put("workspaceViewRoot", root.toString())
            put("currentPath", currentPath)
            put("name", token)
        }
        val response = runCatching {
            project.getService(ProjectEngineService::class.java).aidlResolve(request.toString())
        }.getOrNull() ?: return null
        val resolution = runCatching { ObjectMapper().readTree(response).path("result").path("resolution") }.getOrNull()
            ?: return null
        if (resolution.path("state").asText() != "found") return null
        val symbols = resolution.path("symbols")
        if (!symbols.isArray || symbols.size() != 1) return null
        val symbol = symbols.single()
        val targetPath = symbol.path("path").asText()
        val targetFile = root.resolve(targetPath).normalize().toFile()
        val targetVirtual = com.intellij.openapi.vfs.LocalFileSystem.getInstance()
            .refreshAndFindFileByIoFile(targetFile) ?: return null
        val targetPsi = PsiManager.getInstance(project).findFile(targetVirtual) ?: return null
        val offset = symbol.path("range").path("start").path("offset").asInt(-1)
        if (offset < 0) return targetPsi
        return targetPsi.findElementAt(offset) ?: targetPsi
    }

    private data class AidlTarget(
        val file: VirtualFile,
        val declaration: AidlDeclaration,
    )

    private fun declaration(file: VirtualFile, project: Project): AidlTarget? {
        if (file.fileType != AidlFileType) return null
        val psi = PsiManager.getInstance(project).findFile(file) ?: return null
        val text = psi.text
        val packageName = packagePattern.find(text)?.groupValues?.get(1).orEmpty()
        val match = declarationPattern.find(text) ?: return null
        val name = match.groupValues[1]
        val qualified = if (packageName.isEmpty()) name else "$packageName.$name"
        return AidlTarget(file, AidlDeclaration(qualified, match.groups[1]!!.range.first))
    }

    private fun identifierAt(text: String, offset: Int): String? {
        if (offset < 0 || offset > text.length) return null
        val position = offset.coerceAtMost(text.length - 1)
        val match = identifierPattern.findAll(text).firstOrNull {
            position in it.range
        } ?: return null
        return match.value
    }
}

class AidlGotoDeclarationHandler : GotoDeclarationHandler {
    override fun getGotoDeclarationTargets(
        sourceElement: PsiElement?,
        offset: Int,
        editor: Editor,
    ): Array<PsiElement>? {
        val file = sourceElement?.containingFile ?: return null
        if (FileTypeManager.getInstance().getFileTypeByFile(file.virtualFile) != AidlFileType) return null
        return AidlSourceResolver.resolve(file, offset)?.let { target ->
            arrayOf<PsiElement>(target)
        }
    }
}
