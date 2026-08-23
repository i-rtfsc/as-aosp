import org.jetbrains.kotlin.gradle.dsl.JvmTarget

plugins {
    kotlin("jvm") version "2.1.21"
    id("org.jetbrains.intellij.platform") version "2.18.1"
}

group = "io.github.irtfsc.aospide"
version = "0.1.0-dev"

val localPlatformPath = providers.gradleProperty("intellijPlatformPath")

repositories {
    mavenCentral()
    intellijPlatform {
        defaultRepositories()
    }
}

dependencies {
    intellijPlatform {
        if (localPlatformPath.isPresent) {
            local(localPlatformPath.get())
        } else {
            intellijIdeaCommunity("2025.1.7.2")
        }
    }
    testImplementation("junit:junit:4.13.2")
}

intellijPlatform {
    buildSearchableOptions = false
    pluginConfiguration {
        name = "AOSP IDE"
        version = project.version.toString()
        ideaVersion {
            sinceBuild = "251"
            untilBuild = "251.*"
        }
        description = "Focused AOSP source navigation and cross-language indexing."
    }
}

kotlin {
    jvmToolchain(21)
}

tasks {
    test {
        useJUnit()
    }
    withType<org.jetbrains.kotlin.gradle.tasks.KotlinJvmCompile>().configureEach {
        compilerOptions.jvmTarget.set(JvmTarget.JVM_21)
    }
}
