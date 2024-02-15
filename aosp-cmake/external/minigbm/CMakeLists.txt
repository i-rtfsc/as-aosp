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

project(libminigbm_gralloc_arcvm)

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

#################### libminigbm_gralloc_arcvm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/amdgpu.c
        ${ANDROID_ROOT}/external/minigbm/drv.c
        ${ANDROID_ROOT}/external/minigbm/drv_array_helpers.c
        ${ANDROID_ROOT}/external/minigbm/drv_helpers.c
        ${ANDROID_ROOT}/external/minigbm/dumb_driver.c
        ${ANDROID_ROOT}/external/minigbm/i915.c
        ${ANDROID_ROOT}/external/minigbm/mediatek.c
        ${ANDROID_ROOT}/external/minigbm/msm.c
        ${ANDROID_ROOT}/external/minigbm/rockchip.c
        ${ANDROID_ROOT}/external/minigbm/vc4.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_cross_domain.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_virgl.c
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_buffer.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_helpers.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_driver.cc
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
    add_library(libminigbm_gralloc_arcvm SHARED ${SOURCE_FILES})

    target_include_directories(libminigbm_gralloc_arcvm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(libminigbm_gralloc_arcvm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libminigbm_gralloc_arcvm
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### libminigbm_gralloc_arcvm ####################

#################### libminigbm_gralloc_meson ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/amdgpu.c
        ${ANDROID_ROOT}/external/minigbm/drv.c
        ${ANDROID_ROOT}/external/minigbm/drv_array_helpers.c
        ${ANDROID_ROOT}/external/minigbm/drv_helpers.c
        ${ANDROID_ROOT}/external/minigbm/dumb_driver.c
        ${ANDROID_ROOT}/external/minigbm/i915.c
        ${ANDROID_ROOT}/external/minigbm/mediatek.c
        ${ANDROID_ROOT}/external/minigbm/msm.c
        ${ANDROID_ROOT}/external/minigbm/rockchip.c
        ${ANDROID_ROOT}/external/minigbm/vc4.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_cross_domain.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_virgl.c
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_buffer.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_helpers.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_driver.cc
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
    add_library(libminigbm_gralloc_meson SHARED ${SOURCE_FILES})

    target_include_directories(libminigbm_gralloc_meson
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(libminigbm_gralloc_meson PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libminigbm_gralloc_meson
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### libminigbm_gralloc_meson ####################

#################### libgbm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/amdgpu.c
        ${ANDROID_ROOT}/external/minigbm/drv.c
        ${ANDROID_ROOT}/external/minigbm/drv_array_helpers.c
        ${ANDROID_ROOT}/external/minigbm/drv_helpers.c
        ${ANDROID_ROOT}/external/minigbm/dumb_driver.c
        ${ANDROID_ROOT}/external/minigbm/i915.c
        ${ANDROID_ROOT}/external/minigbm/mediatek.c
        ${ANDROID_ROOT}/external/minigbm/msm.c
        ${ANDROID_ROOT}/external/minigbm/rockchip.c
        ${ANDROID_ROOT}/external/minigbm/vc4.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_cross_domain.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_virgl.c
        ${ANDROID_ROOT}/external/minigbm/gbm.c
        ${ANDROID_ROOT}/external/minigbm/gbm_helpers.c
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
    add_library(libgbm SHARED ${SOURCE_FILES})

    target_include_directories(libgbm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(libgbm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libgbm
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libdrm
            libdrm_headers
            liblog
            libm
    )
endif ()
#################### libgbm ####################

#################### gralloc.minigbm_msm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc0/gralloc0.cc
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
    add_library(gralloc.minigbm_msm SHARED ${SOURCE_FILES})

    target_include_directories(gralloc.minigbm_msm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(gralloc.minigbm_msm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(gralloc.minigbm_msm
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libminigbm_gralloc_msm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### gralloc.minigbm_msm ####################

#################### gralloc.minigbm_arcvm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc0/gralloc0.cc
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
    add_library(gralloc.minigbm_arcvm SHARED ${SOURCE_FILES})

    target_include_directories(gralloc.minigbm_arcvm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(gralloc.minigbm_arcvm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(gralloc.minigbm_arcvm
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libminigbm_gralloc_arcvm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### gralloc.minigbm_arcvm ####################

#################### libminigbm_gralloc_msm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/amdgpu.c
        ${ANDROID_ROOT}/external/minigbm/drv.c
        ${ANDROID_ROOT}/external/minigbm/drv_array_helpers.c
        ${ANDROID_ROOT}/external/minigbm/drv_helpers.c
        ${ANDROID_ROOT}/external/minigbm/dumb_driver.c
        ${ANDROID_ROOT}/external/minigbm/i915.c
        ${ANDROID_ROOT}/external/minigbm/mediatek.c
        ${ANDROID_ROOT}/external/minigbm/msm.c
        ${ANDROID_ROOT}/external/minigbm/rockchip.c
        ${ANDROID_ROOT}/external/minigbm/vc4.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_cross_domain.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_virgl.c
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_buffer.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_helpers.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_driver.cc
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
    add_library(libminigbm_gralloc_msm SHARED ${SOURCE_FILES})

    target_include_directories(libminigbm_gralloc_msm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(libminigbm_gralloc_msm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libminigbm_gralloc_msm
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### libminigbm_gralloc_msm ####################

#################### gralloc.minigbm_meson ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc0/gralloc0.cc
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
    add_library(gralloc.minigbm_meson SHARED ${SOURCE_FILES})

    target_include_directories(gralloc.minigbm_meson
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(gralloc.minigbm_meson PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(gralloc.minigbm_meson
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libminigbm_gralloc_meson
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### gralloc.minigbm_meson ####################

#################### gralloc.minigbm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc0/gralloc0.cc
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
    add_library(gralloc.minigbm SHARED ${SOURCE_FILES})

    target_include_directories(gralloc.minigbm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(gralloc.minigbm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(gralloc.minigbm
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libminigbm_gralloc
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### gralloc.minigbm ####################

#################### libminigbm_gralloc ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/amdgpu.c
        ${ANDROID_ROOT}/external/minigbm/drv.c
        ${ANDROID_ROOT}/external/minigbm/drv_array_helpers.c
        ${ANDROID_ROOT}/external/minigbm/drv_helpers.c
        ${ANDROID_ROOT}/external/minigbm/dumb_driver.c
        ${ANDROID_ROOT}/external/minigbm/i915.c
        ${ANDROID_ROOT}/external/minigbm/mediatek.c
        ${ANDROID_ROOT}/external/minigbm/msm.c
        ${ANDROID_ROOT}/external/minigbm/rockchip.c
        ${ANDROID_ROOT}/external/minigbm/vc4.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_cross_domain.c
        ${ANDROID_ROOT}/external/minigbm/virtgpu_virgl.c
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_buffer.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_helpers.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/cros_gralloc_driver.cc
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
    add_library(libminigbm_gralloc SHARED ${SOURCE_FILES})

    target_include_directories(libminigbm_gralloc
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm
    )

    # Add c/c++ flags
    target_compile_options(libminigbm_gralloc PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libminigbm_gralloc
            libarect
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libhardware_headers
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### libminigbm_gralloc ####################
