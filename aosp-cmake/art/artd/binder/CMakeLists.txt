# Copyright (c) 2024 anqi.huang@outlook.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use self file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required(VERSION 3.5)

project(artd-aidl-ndk_platform)

# Check if the major CMakeLists.txt file has correctly set
# the ANDROID_ROOT directory.
if (IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else ()
    if (IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
        set(OUT_INTERMEDIATES_ROOT ${ANDROID_ROOT}/out/soong/.intermediates)
    else ()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif ()
endif ()

#################### artd-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/ArtConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/ArtdDexoptResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/ArtifactsPath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/DexMetadataPath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/DexoptOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/DexoptTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/FileVisibility.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/FsPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/GetDexoptNeededResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/GetDexoptStatusResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/IArtd.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/IArtdCancellationSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/MergeProfileOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/OutputArtifacts.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/OutputProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/PriorityClass.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/ProfilePath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/com/android/server/art/VdexPath.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(artd-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(artd-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(artd-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(artd-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### artd-aidl-ndk_platform ####################

#################### artd-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/ArtConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/ArtdDexoptResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/ArtifactsPath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/DexMetadataPath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/DexoptOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/DexoptTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/FileVisibility.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/FsPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/GetDexoptNeededResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/GetDexoptStatusResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/IArtd.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/IArtdCancellationSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/MergeProfileOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/OutputArtifacts.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/OutputProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/PriorityClass.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/ProfilePath.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/com/android/server/art/VdexPath.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(artd-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(artd-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/art/artd/binder/artd-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(artd-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(artd-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### artd-aidl-ndk ####################
