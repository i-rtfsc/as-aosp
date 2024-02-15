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

project(android.hardware.audio.effect.service-aidl.example)

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

#################### libaudioservicesounddoseimpl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/SoundDose.cpp
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
    add_library(libaudioservicesounddoseimpl SHARED ${SOURCE_FILES})

    target_include_directories(libaudioservicesounddoseimpl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(libaudioservicesounddoseimpl PRIVATE
            -Werror
    )

    target_link_libraries(libaudioservicesounddoseimpl
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.audio.sounddose-V1-ndk
            android.media.audio.common.types-V2-ndk
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libaudioservicesounddoseimpl ####################

#################### android.hardware.audio.service-aidl.example ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/main.cpp
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
    add_executable(android.hardware.audio.service-aidl.example ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.service-aidl.example
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/audio_policy_configuration_aidl_default/gen/cpp/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/engine/audio_policy_engine_configuration_aidl_default/gen/cpp/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.service-aidl.example PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.audio.service-aidl.example
            android.hardware.audio.core-V1-ndk
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
            libalsautilsv2
            libaudio_system_headers
            libaudioaidl_headers
            libaudioaidlcommon
            libaudioserviceexampleimpl
            libaudioutils
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libm
            libstagefright_foundation
            libtinyalsav2
            libutils
            libxml2
            libxsdc-utils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio.service-aidl.example ####################

#################### android.hardware.audio.effect.service-aidl.example ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EffectConfig.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EffectFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EffectMain.cpp
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
    add_executable(android.hardware.audio.effect.service-aidl.example ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.effect.service-aidl.example
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.effect.service-aidl.example PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.audio.effect.service-aidl.example
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
            libaudio_system_headers
            libaudioaidl_headers
            libaudioaidlcommon
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            liblog
            libm
            libsystem_headers
            libtinyxml2
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio.effect.service-aidl.example ####################

#################### libaudioserviceexampleimpl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/AudioPolicyConfigXmlConverter.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Bluetooth.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Config.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Configuration.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EngineConfigXmlConverter.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Module.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/SoundDose.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/StreamStub.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/Telephony.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/usb/ModuleUsb.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/usb/StreamUsb.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/usb/UsbAlsaMixerControl.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/usb/UsbAlsaUtils.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/audio_policy_configuration_aidl_default/gen/cpp/android_audio_policy_configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/audio_policy_configuration_aidl_default/gen/cpp/android_audio_policy_configuration_enums.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/engine/audio_policy_engine_configuration_aidl_default/gen/cpp/android_audio_policy_engine_configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/engine/audio_policy_engine_configuration_aidl_default/gen/cpp/android_audio_policy_engine_configuration_enums.cpp
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
    add_library(libaudioserviceexampleimpl SHARED ${SOURCE_FILES})

    target_include_directories(libaudioserviceexampleimpl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/audio_policy_configuration_aidl_default/gen/cpp/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/default/config/audioPolicy/engine/audio_policy_engine_configuration_aidl_default/gen/cpp/include
    )

    # Add c/c++ flags
    target_compile_options(libaudioserviceexampleimpl PRIVATE
            -Werror
    )

    target_link_libraries(libaudioserviceexampleimpl
            android.hardware.audio.core-V1-ndk
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
            libalsautilsv2
            libaudio_system_headers
            libaudioaidl_headers
            libaudioaidlcommon
            libaudioutils
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libm
            libstagefright_foundation
            libtinyalsav2
            libutils
            libxml2
            libxsdc-utils
            libc
            libdl
            libm
    )
endif ()
#################### libaudioserviceexampleimpl ####################
