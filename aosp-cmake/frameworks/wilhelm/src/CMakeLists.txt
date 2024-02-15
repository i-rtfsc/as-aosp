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

project(libwilhelm)

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

#################### libwilhelm ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/wilhelm/src/OpenSLES_IID.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/assert.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/classes.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/data.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/devices.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/entry.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/handler_bodies.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/trace.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/locks.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/sles.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/sl_iid.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/sllog.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/ThreadPool.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/AudioPlayer_to_android.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/AudioRecorder_to_android.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/MediaPlayer_to_android.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/OutputMix_to_android.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/VideoCodec_to_android.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/BufferQueueSource.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/CallbackProtector.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/AacBqToPcmCbRenderer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_AudioSfDecoder.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_AudioToCbRenderer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_GenericMediaPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_GenericPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_LocAVPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_StreamPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/android_Effect.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/util/AacAdtsExtractor.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/android/channels.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/autogen/IID_to_MPH.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/C3DGroup.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/CAudioPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/CAudioRecorder.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/CEngine.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/COutputMix.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/objects/CMediaPlayer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAndroidBufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAndroidConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAndroidEffect.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAndroidEffectCapabilities.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAndroidEffectSend.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAcousticEchoCancellation.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IAutomaticGainControl.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IBassBoost.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IBufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IDynamicInterfaceManagement.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IEffectSend.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IEngine.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IEngineCapabilities.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IEnvironmentalReverb.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IEqualizer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IMetadataExtraction.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/INoiseSuppression.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IMuteSolo.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IObject.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IOutputMix.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IPlay.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IPlaybackRate.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IPrefetchStatus.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IPresetReverb.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IRecord.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/ISeek.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IStreamInformation.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IVideoDecoderCapabilities.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IVirtualizer.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/itf/IVolume.cpp
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
    add_library(libwilhelm SHARED ${SOURCE_FILES})

    target_include_directories(libwilhelm
            PUBLIC
            ${ANDROID_ROOT}/frameworks/wilhelm/src
            ${ANDROID_ROOT}/frameworks/wilhelm/src/ut
    )

    # Add c/c++ flags
    target_compile_options(libwilhelm PRIVATE
            -UNDEBUG
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libwilhelm
            framework-permission-aidl-cpp
            jni_headers
            libOpenSLESUT
            libOpenSLES_headers
            libandroid_runtime
            libaudioclient
            libaudioclient_aidl_conversion
            libaudiofoundation
            libaudiomanager
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdatasource
            libdl
            libgui
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmediametrics_headers
            libopensles_helper
            libstagefright
            libstagefright_foundation
            libstagefright_http_support
            libstagefright_metadatautils
            libstagefright_mpeg2support_headers
            libutils
            media_plugin_headers
    )
endif ()
#################### libwilhelm ####################

#################### libOpenMAXAL ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/wilhelm/src/xa_entry.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/xa_iid.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/assert.cpp
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
    add_library(libOpenMAXAL SHARED ${SOURCE_FILES})

    target_include_directories(libOpenMAXAL
            PUBLIC
            ${ANDROID_ROOT}/frameworks/wilhelm/src
    )

    # Add c/c++ flags
    target_compile_options(libOpenMAXAL PRIVATE
            -UNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libOpenMAXAL
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediametrics_headers
            libnativewindow
            libwilhelm
    )
endif ()
#################### libOpenMAXAL ####################

#################### libopensles_helper ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/wilhelm/src/assert.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/MPH_to.c
        ${ANDROID_ROOT}/frameworks/wilhelm/src/handlers.c
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
    add_library(libopensles_helper SHARED ${SOURCE_FILES})

    target_include_directories(libopensles_helper
            PUBLIC
            ${ANDROID_ROOT}/frameworks/wilhelm/src
    )

    # Add c/c++ flags
    target_compile_options(libopensles_helper PRIVATE
            -UNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libopensles_helper
            libOpenSLES_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libopensles_helper ####################

#################### libOpenSLESUT ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/wilhelm/src/assert.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/ut/OpenSLESUT.c
        ${ANDROID_ROOT}/frameworks/wilhelm/src/ut/slesutResult.c
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
    add_library(libOpenSLESUT SHARED ${SOURCE_FILES})

    target_include_directories(libOpenSLESUT
            PUBLIC
            ${ANDROID_ROOT}/frameworks/wilhelm/src/ut
            ${ANDROID_ROOT}/frameworks/wilhelm/src
    )

    # Add c/c++ flags
    target_compile_options(libOpenSLESUT PRIVATE
            -UNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libOpenSLESUT
            jni_headers
            libOpenSLES_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libOpenSLESUT ####################

#################### libOpenSLES ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/wilhelm/src/sl_entry.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/sl_iid.cpp
        ${ANDROID_ROOT}/frameworks/wilhelm/src/assert.cpp
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
    add_library(libOpenSLES SHARED ${SOURCE_FILES})

    target_include_directories(libOpenSLES
            PUBLIC
            ${ANDROID_ROOT}/frameworks/wilhelm/src
    )

    # Add c/c++ flags
    target_compile_options(libOpenSLES PRIVATE
            -UNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libOpenSLES
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmediametrics_headers
            libnativewindow
            libwilhelm
    )
endif ()
#################### libOpenSLES ####################
