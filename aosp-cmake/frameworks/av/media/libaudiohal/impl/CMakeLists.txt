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

project(libaudiohal@aidl)

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

#################### libaudiohal-aidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectConversionHelperAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectBufferHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionAec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionAgc1.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionAgc2.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionBassBoost.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionDownmix.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionDynamicsProcessing.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionEnvReverb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionEq.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionHapticGenerator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionLoudnessEnhancer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionNoiseSuppression.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionPresetReverb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionSpatializer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionVendorExtension.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionVirtualizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/effectsAidlConversion/AidlConversionVisualizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalAidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectProxy.cpp
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
    add_library(libaudiohal-aidl SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-aidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-aidl PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-aidl
            android.hardware.audio.common-V2-ndk
            android.hardware.audio.core-V1-ndk
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-cpp
            android.media.audio.common.types-V2-ndk
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudio_aidl_conversion_common_cpp
            libaudio_aidl_conversion_common_ndk
            libaudio_aidl_conversion_common_ndk_cpp
            libaudio_aidl_conversion_core_ndk
            libaudio_aidl_conversion_effect_ndk
            libaudio_system_headers
            libaudioaidlcommon
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libeffectsconfig_headers
            libfmq
            libhardware
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-aidl ####################

#################### libaudiohal-6.0 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/CoreConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectBufferHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
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
    add_library(libaudiohal-6.0 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-6.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-6.0 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-6.0
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-6.0
            android.hardware.audio.common-6.0-util
            android.hardware.audio.effect-6.0
            android.hardware.audio.effect-6.0-util
            android.hardware.audio-6.0
            android.hardware.audio-6.0-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-6.0 ####################

#################### libaudiohal.effect-7.0 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectBufferHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalHidl.cpp
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
    add_library(libaudiohal.effect-7.0 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal.effect-7.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal.effect-7.0 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal.effect-7.0
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-7.0
            android.hardware.audio.common-7.0-util
            android.hardware.audio.effect-7.0
            android.hardware.audio.effect-7.0-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal.effect-7.0 ####################

#################### libaudiohal-7.0 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/CoreConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
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
    add_library(libaudiohal-7.0 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-7.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-7.0 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-7.0
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-7.0
            android.hardware.audio.common-7.0-enums
            android.hardware.audio.common-7.0-util
            android.hardware.audio.effect-7.0
            android.hardware.audio.effect-7.0-util
            android.hardware.audio-7.0
            android.hardware.audio-7.0-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal.effect-7.0
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-7.0 ####################

#################### libaudiohal-5.0 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/CoreConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectBufferHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
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
    add_library(libaudiohal-5.0 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-5.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-5.0 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-5.0
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-5.0
            android.hardware.audio.common-5.0-util
            android.hardware.audio.effect-5.0
            android.hardware.audio.effect-5.0-util
            android.hardware.audio-5.0
            android.hardware.audio-5.0-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-5.0 ####################

#################### libaudiohal-7.1 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/CoreConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
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
    add_library(libaudiohal-7.1 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-7.1
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-7.1 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-7.1
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-7.0
            android.hardware.audio.common-7.0-util
            android.hardware.audio.common-7.1-enums
            android.hardware.audio.common-7.1-util
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.audio.effect-7.0
            android.hardware.audio.effect-7.0-util
            android.hardware.audio.sounddose-V1-ndk
            android.hardware.audio-7.0
            android.hardware.audio-7.1
            android.hardware.audio-7.1-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal.effect-7.0
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-7.1 ####################

#################### libaudiohal-4.0 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/CoreConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DeviceHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/DevicesFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/StreamHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectBufferHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectConversionHelperHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl/EffectsFactoryHalEntry.cpp
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
    add_library(libaudiohal-4.0 SHARED ${SOURCE_FILES})

    target_include_directories(libaudiohal-4.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudiohal/impl
    )

    # Add c/c++ flags
    target_compile_options(libaudiohal-4.0 PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libaudiohal-4.0
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-4.0
            android.hardware.audio.common-4.0-util
            android.hardware.audio.effect-4.0
            android.hardware.audio.effect-4.0-util
            android.hardware.audio-4.0
            android.hardware.audio-4.0-util
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiohal_deathhandler
            libaudiohal_headers
            libaudioutils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libeffectsconfig
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmediautils
            libutils
    )
endif ()
#################### libaudiohal-4.0 ####################
