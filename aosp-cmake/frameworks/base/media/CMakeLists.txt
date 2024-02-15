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

project(media_permission-aidl-cpp-analyzer)

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

#################### soundtrigger_middleware-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/IAcknowledgeEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/IInjectGlobalEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/IInjectModelEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/IInjectRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/ISoundTriggerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/ISoundTriggerInjection.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/ISoundTriggerMiddlewareService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/ISoundTriggerModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/PhraseRecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/RecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/android/media/soundtrigger_middleware/SoundTriggerModuleDescriptor.cpp
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
    add_library(soundtrigger_middleware-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(soundtrigger_middleware-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(soundtrigger_middleware-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(soundtrigger_middleware-aidl-ndk
            android.media.audio.common.types-V2-ndk
            android.media.soundtrigger.types-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            media_permission-aidl-ndk
    )
endif ()
#################### soundtrigger_middleware-aidl-ndk ####################

#################### media_permission-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-source/gen/android/media/permission/Identity.cpp
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
    add_library(media_permission-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(media_permission-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(media_permission-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(media_permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### media_permission-aidl-cpp ####################

#################### soundtrigger_middleware-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/IAcknowledgeEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/IInjectGlobalEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/IInjectModelEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/IInjectRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/ISoundTriggerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/ISoundTriggerInjection.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/ISoundTriggerMiddlewareService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/ISoundTriggerModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/PhraseRecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/RecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/android/media/soundtrigger_middleware/SoundTriggerModuleDescriptor.cpp
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
    add_library(soundtrigger_middleware-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(soundtrigger_middleware-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(soundtrigger_middleware-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(soundtrigger_middleware-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            android.media.soundtrigger.types-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            media_permission-aidl-ndk_platform
    )
endif ()
#################### soundtrigger_middleware-aidl-ndk_platform ####################

#################### media_permission-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk_platform-source/gen/android/media/permission/Identity.cpp
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
    add_library(media_permission-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(media_permission-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(media_permission-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(media_permission-aidl-ndk_platform
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
#################### media_permission-aidl-ndk_platform ####################

#################### soundtrigger_middleware-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/IAcknowledgeEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/IInjectGlobalEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/IInjectModelEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/IInjectRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/ISoundTriggerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/ISoundTriggerInjection.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/ISoundTriggerMiddlewareService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/ISoundTriggerModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/PhraseRecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/RecognitionEventSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/android/media/soundtrigger_middleware/SoundTriggerModuleDescriptor.cpp
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
    add_library(soundtrigger_middleware-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(soundtrigger_middleware-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/soundtrigger_middleware-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(soundtrigger_middleware-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(soundtrigger_middleware-aidl-cpp
            android.media.audio.common.types-V2-cpp
            android.media.soundtrigger.types-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            media_permission-aidl-cpp
    )
endif ()
#################### soundtrigger_middleware-aidl-cpp ####################

#################### media_permission-aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-analyzer-source/gen/android/media/permission/Identity.cpp
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
    add_executable(media_permission-aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(media_permission-aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/media
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(media_permission-aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(media_permission-aidl-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            media_permission-aidl-cpp
    )
endif ()
#################### media_permission-aidl-cpp-analyzer ####################

#################### media_permission-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk-source/gen/android/media/permission/Identity.cpp
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
    add_library(media_permission-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(media_permission-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/media_permission-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(media_permission-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(media_permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### media_permission-aidl-ndk ####################
