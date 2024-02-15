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

project(libmusicbundle)

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

#################### libmusicbundle ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_BypassMix.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_Control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_Equaliser.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_ReverbGenerator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_StereoEnhancer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src/LVCS_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/src/LVDBE_Control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/src/LVDBE_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/src/LVDBE_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/src/LVDBE_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_API_Specials.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_Buffers.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src/LVM_Control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_Control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_QPD_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_QPD_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src/LVPSA_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src/LVEQNB_CalcCoef.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src/LVEQNB_Control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src/LVEQNB_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src/LVEQNB_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src/LVEQNB_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/DC_2I_D16_TRC_WRA_01.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/DC_2I_D16_TRC_WRA_01_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Copy_16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MonoTo2I_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/dB_to_Lin32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Shift_Sat_v16xv16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Shift_Sat_v32xv32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/From2iToMono_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Mult3s_32x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/NonLinComp_D16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/DelayMix_16x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MSTo2i_Sat_16x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/From2iToMS_16x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Mac3s_Sat_32x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Add2_Sat_32x32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_MixSoft_1St_2i_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_MixSoft_1St_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_VarSlope_SetTimeConstant.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_SetTimeConstant.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_SetTarget.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_GetTarget.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_Init.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Core_MixHard_1St_2i_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Core_MixSoft_1St_2i_D16C31_WRA.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Core_MixInSoft_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Mixer_GetCurrent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_MixSoft_2St_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Core_MixSoft_1St_D16C31_WRA.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_Core_MixHard_2St_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVC_MixInSoft_D16C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/AGC_MIX_VOL_2St1Mon_D32_WRA.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_Timer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_Timer_Init.cpp
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
    add_library(libmusicbundle SHARED ${SOURCE_FILES})

    target_include_directories(libmusicbundle
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Eq/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bass/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Bundle/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/SpectrumAnalyzer/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/StereoWidening/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib
    )

    target_link_libraries(libmusicbundle
            libaudioutils
            libaudioutils
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libhardware_headers
            libhardware_headers
            liblog
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libmusicbundle ####################

#################### libreverb ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_ApplyNewSettings.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_ClearAudioBuffers.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_GetControlParameters.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_GetInstanceHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_Process.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_SetControlParameters.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src/LVREV_Tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/From2iToMono_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Mult3s_32x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Copy_16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Mac3s_Sat_32x16.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Shift_Sat_v32xv32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Add2_Sat_32x32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/JoinTo2i_32x32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MonoTo2I_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_FO_HPF.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_FO_LPF.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_Polynomial.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_Power10.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_GetOmega.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MixSoft_2St_D32C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MixSoft_1St_D32C31_WRA.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/MixInSoft_D32C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/LVM_Mixer_TimeConstant.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Core_MixHard_2St_D32C31_SAT.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Core_MixSoft_1St_D32C31_WRA.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src/Core_MixInSoft_D32C31_SAT.cpp
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
    add_library(libreverb SHARED ${SOURCE_FILES})

    target_include_directories(libreverb
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/lib
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Reverb/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib/Common/src
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/lvm/lib
    )

    target_link_libraries(libreverb
            libaudioutils
            libaudioutils
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libreverb ####################
