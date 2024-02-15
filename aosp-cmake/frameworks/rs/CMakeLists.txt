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

project(libRSDriver)

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

#################### libRSDriver ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdBcc.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdCore.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdFrameBuffer.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdFrameBufferObj.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdGL.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdMesh.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdMeshObj.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdProgram.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdProgramRaster.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdProgramStore.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdRuntimeStubs.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdSampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdShader.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdShaderCache.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdType.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdVertexArray.cpp
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
    add_library(libRSDriver SHARED ${SOURCE_FILES})

    target_include_directories(libRSDriver
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs
    )

    # Add c/c++ flags
    target_compile_options(libRSDriver PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRSDriver
            libEGL
            libGLESv1_CM
            libGLESv2
            libRSCpuRef
            libRS_internal
            libarect
            libbcinfo
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libutils_headers
    )
endif ()
#################### libRSDriver ####################

#################### libRS_internal ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/rsApiAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiContext.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiDevice.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiFileA3D.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiMesh.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiType.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsAnimation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsComponent.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsContext.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsClosure.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsCppUtils.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsDevice.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsDriverLoader.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFBOCache.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFifoSocket.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFileA3D.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFont.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsGrallocConsumer.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsObjectBase.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix2x2.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix3x3.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix4x4.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMesh.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMutex.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsProgram.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsProgramFragment.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsProgramStore.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsProgramRaster.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsProgramVertex.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsSampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptC.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptC_Lib.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptC_LibGL.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptGroup2.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptIntrinsic.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsSignal.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsStream.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsThreadIO.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources/gen/gensrcs/frameworks/rs/rsgApi.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources/gen/gensrcs/frameworks/rs/rsgApiReplay.cpp
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
    add_library(libRS_internal SHARED ${SOURCE_FILES})

    target_include_directories(libRS_internal
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs
            ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources/gen/gensrcs
            ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_headers/gen/gensrcs/frameworks/rs
    )

    # Add c/c++ flags
    target_compile_options(libRS_internal PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRS_internal
            libEGL
            libGLESv1_CM
            libGLESv2
            libarect
            libbcinfo
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libft2
            liblog
            libm
            libmediandk
            libnativewindow
            libutils
            libutilscallstack
    )
endif ()
#################### libRS_internal ####################

#################### libRSSupport ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/rsAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiContext.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiDevice.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsApiType.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsClosure.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsComponent.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsContext.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsCppUtils.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsDevice.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsDriverLoader.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFifoSocket.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsObjectBase.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix2x2.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix3x3.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMatrix4x4.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsMutex.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsSampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptC.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptC_Lib.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptGroup2.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsScriptIntrinsic.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsSignal.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsStream.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsThreadIO.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsType.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdAllocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdBcc.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdCore.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdElement.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdRuntimeStubs.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdSampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/driver/rsdType.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuCore.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuExecutable.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuRuntimeMath.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScriptGroup2.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsic.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsic3DLUT.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBlend.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBlur.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBLAS.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicColorMatrix.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicConvolve3x3.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicConvolve5x5.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicHistogram.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicLUT.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicResize.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicYuvToRGB.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_3DLUT.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Blend.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Blur.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_ColorMatrix.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Convolve.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Resize.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_YuvToRGB.S
        ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources_support/gen/gensrcs/frameworks/rs/rsgApi.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources_support/gen/gensrcs/frameworks/rs/rsgApiReplay.cpp
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
    add_library(libRSSupport SHARED ${SOURCE_FILES})

    target_include_directories(libRSSupport
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs
            ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_sources_support/gen/gensrcs
            ${OUT_INTERMEDIATES_ROOT}/frameworks/rs/rs_generated_headers_support/gen/gensrcs/frameworks/rs
    )

    # Add c/c++ flags
    target_compile_options(libRSSupport PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRSSupport
            libbnnmlowpV8
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### libRSSupport ####################

#################### libRS ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/rsApiStubs.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsHidlAdaptation.cpp
        ${ANDROID_ROOT}/frameworks/rs/rsFallbackAdaptation.cpp
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
    add_library(libRS SHARED ${SOURCE_FILES})

    target_include_directories(libRS
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs
    )

    # Add c/c++ flags
    target_compile_options(libRS PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRS
            android.hardware.renderscript-1.0
            libRSDispatch
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libutilscallstack
    )
endif ()
#################### libRS ####################

#################### libRSCacheDir ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/rsCacheDir.cpp
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
    add_library(libRSCacheDir SHARED ${SOURCE_FILES})

    target_include_directories(libRSCacheDir
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs
    )

    # Add c/c++ flags
    target_compile_options(libRSCacheDir PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRSCacheDir
            libandroid_runtime
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libRSCacheDir ####################
