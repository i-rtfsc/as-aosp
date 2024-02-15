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

project(libsonivoxwithoutjet)

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

#################### libsonivoxwithoutjet ####################
set(FILES
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_data.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_dlssynth.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_flog.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_imelody.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_imelodydata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_math.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mdls.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_midi.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mididata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mixbuf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mixer.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_ota.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_otadata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pan.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pcm.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pcmdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_public.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_reverb.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_reverbdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_rtttl.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_rtttldata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_smf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_smfdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_voicemgt.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_wtengine.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_wtsynth.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_xmf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_xmfdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/wt_22khz.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_config.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_hostmm.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_report.c
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
    add_library(libsonivoxwithoutjet SHARED ${SOURCE_FILES})

    target_include_directories(libsonivoxwithoutjet
            PUBLIC
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/include
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k
    )

    # Add c/c++ flags
    target_compile_options(libsonivoxwithoutjet PRIVATE
            -DUNIFIED_DEBUG_MESSAGES
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsonivoxwithoutjet
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libsonivoxwithoutjet ####################

#################### libsonivox ####################
set(FILES
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_data.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_dlssynth.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_flog.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_imelody.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_imelodydata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_math.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mdls.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_midi.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mididata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mixbuf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_mixer.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_ota.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_otadata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pan.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pcm.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_pcmdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_public.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_reverb.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_reverbdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_rtttl.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_rtttldata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_smf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_smfdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_voicemgt.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_wtengine.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_wtsynth.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_xmf.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/eas_xmfdata.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/wt_22khz.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_config.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_hostmm.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src/eas_report.c
        ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src/jet.c
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
    add_library(libsonivox SHARED ${SOURCE_FILES})

    target_include_directories(libsonivox
            PUBLIC
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/include
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/host_src
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k/lib_src
            ${ANDROID_ROOT}/external/sonivox/arm-wt-22k
    )

    # Add c/c++ flags
    target_compile_options(libsonivox PRIVATE
            -DUNIFIED_DEBUG_MESSAGES
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsonivox
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libsonivox ####################
