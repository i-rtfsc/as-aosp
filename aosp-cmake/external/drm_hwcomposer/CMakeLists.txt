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

project(libdrmhwc_utils)

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

#################### libdrmhwc_utils ####################
set(FILES
        ${ANDROID_ROOT}/external/drm_hwcomposer/utils/Worker.cpp
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
    add_library(libdrmhwc_utils SHARED ${SOURCE_FILES})

    target_include_directories(libdrmhwc_utils
            PUBLIC
            ${ANDROID_ROOT}/external/drm_hwcomposer
    )

    # Add c/c++ flags
    target_compile_options(libdrmhwc_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libdrmhwc_utils
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libdrmhwc_utils ####################

#################### hwcomposer.drm_minigbm ####################
set(FILES
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoGetter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoMapperMetadata.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/compositor/DrmKmsPlan.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmAtomicStateManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmConnector.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmCrtc.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDevice.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDisplayPipeline.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmEncoder.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmFbImporter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmMode.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmPlane.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmProperty.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/ResourceManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/UEventListener.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/VSyncWorker.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/Backend.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendClient.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/DrmHwcTwo.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplay.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplayConfigs.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcLayer.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/hwc2_device.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/legacy/BufferInfoMinigbm.cpp
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
    add_library(hwcomposer.drm_minigbm SHARED ${SOURCE_FILES})

    target_include_directories(hwcomposer.drm_minigbm
            PUBLIC
            ${ANDROID_ROOT}/external/drm_hwcomposer
    )

    # Add c/c++ flags
    target_compile_options(hwcomposer.drm_minigbm PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(hwcomposer.drm_minigbm
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdrm
            libdrmhwc_utils
            libhardware
            libhidlbase
            liblog
            libm
            libsync
            libui
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### hwcomposer.drm_minigbm ####################

#################### drm_hwcomposer ####################
set(FILES
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoGetter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoMapperMetadata.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/compositor/DrmKmsPlan.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmAtomicStateManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmConnector.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmCrtc.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDevice.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDisplayPipeline.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmEncoder.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmFbImporter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmMode.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmPlane.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmProperty.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/ResourceManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/UEventListener.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/VSyncWorker.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/Backend.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendClient.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/DrmHwcTwo.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplay.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplayConfigs.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcLayer.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/hwc2_device.cpp
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
    add_library(drm_hwcomposer SHARED ${SOURCE_FILES})

    target_include_directories(drm_hwcomposer
            PUBLIC
            ${ANDROID_ROOT}/external/drm_hwcomposer
    )

    # Add c/c++ flags
    target_compile_options(drm_hwcomposer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(drm_hwcomposer
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdrm
            libdrmhwc_utils
            libhardware
            libhidlbase
            liblog
            libm
            libsync
            libui
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### drm_hwcomposer ####################

#################### hwcomposer.drm ####################
set(FILES
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoGetter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/BufferInfoMapperMetadata.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/compositor/DrmKmsPlan.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmAtomicStateManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmConnector.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmCrtc.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDevice.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmDisplayPipeline.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmEncoder.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmFbImporter.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmMode.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmPlane.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/DrmProperty.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/ResourceManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/UEventListener.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/drm/VSyncWorker.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/Backend.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendClient.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/backend/BackendManager.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/DrmHwcTwo.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplay.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcDisplayConfigs.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/HwcLayer.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/hwc2_device/hwc2_device.cpp
        ${ANDROID_ROOT}/external/drm_hwcomposer/bufferinfo/legacy/BufferInfoLibdrm.cpp
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
    add_library(hwcomposer.drm SHARED ${SOURCE_FILES})

    target_include_directories(hwcomposer.drm
            PUBLIC
            ${ANDROID_ROOT}/external/drm_hwcomposer
    )

    # Add c/c++ flags
    target_compile_options(hwcomposer.drm PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(hwcomposer.drm
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdrm
            libdrmhwc_utils
            libhardware
            libhidlbase
            liblog
            libm
            libsync
            libui
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### hwcomposer.drm ####################
