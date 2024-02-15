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

project(android.hardware.graphics.allocator@4.0-service.minigbm_msm)

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

#################### android.hardware.graphics.allocator-4.0-service.minigbm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Allocator.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4AllocatorService.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_executable(android.hardware.graphics.allocator-4.0-service.minigbm ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.allocator-4.0-service.minigbm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.allocator-4.0-service.minigbm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.allocator-4.0-service.minigbm
            android.hardware.graphics.allocator-4.0
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.allocator-4.0-service.minigbm ####################

#################### android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Mapper.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_library(android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc_arcvm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.mapper-4.0-impl.minigbm_arcvm ####################

#################### libminigbm_gralloc4_utils ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_library(libminigbm_gralloc4_utils SHARED ${SOURCE_FILES})

    target_include_directories(libminigbm_gralloc4_utils
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(libminigbm_gralloc4_utils PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libminigbm_gralloc4_utils
            android.hardware.graphics.mapper-4.0
            libarect
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
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
#################### libminigbm_gralloc4_utils ####################

#################### android.hardware.graphics.mapper-4.0-impl.minigbm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Mapper.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_library(android.hardware.graphics.mapper-4.0-impl.minigbm SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.mapper-4.0-impl.minigbm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.mapper-4.0-impl.minigbm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.mapper-4.0-impl.minigbm
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.mapper-4.0-impl.minigbm ####################

#################### android.hardware.graphics.mapper-4.0-impl.minigbm_msm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Mapper.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_library(android.hardware.graphics.mapper-4.0-impl.minigbm_msm SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.mapper-4.0-impl.minigbm_msm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.mapper-4.0-impl.minigbm_msm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.mapper-4.0-impl.minigbm_msm
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc_msm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.mapper-4.0-impl.minigbm_msm ####################

#################### android.hardware.graphics.allocator-4.0-service.minigbm_arcvm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Allocator.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4AllocatorService.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_executable(android.hardware.graphics.allocator-4.0-service.minigbm_arcvm ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.allocator-4.0-service.minigbm_arcvm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.allocator-4.0-service.minigbm_arcvm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.allocator-4.0-service.minigbm_arcvm
            android.hardware.graphics.allocator-4.0
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc_arcvm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.allocator-4.0-service.minigbm_arcvm ####################

#################### android.hardware.graphics.allocator-4.0-service.minigbm_msm ####################
set(FILES
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Allocator.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4AllocatorService.cc
        ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4/CrosGralloc4Utils.cc
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
    add_executable(android.hardware.graphics.allocator-4.0-service.minigbm_msm ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.allocator-4.0-service.minigbm_msm
            PUBLIC
            ${ANDROID_ROOT}/external/minigbm/cros_gralloc/gralloc4
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.allocator-4.0-service.minigbm_msm PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.graphics.allocator-4.0-service.minigbm_msm
            android.hardware.graphics.allocator-4.0
            android.hardware.graphics.mapper-4.0
            libarect
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libdrm
            libgralloctypes
            libhardware_headers
            libhidlbase
            liblog
            libm
            libminigbm_gralloc_msm
            libnativebase_headers
            libnativewindow
            libsync
            libsystem_headers
            libutils
            minigbm_headers
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.allocator-4.0-service.minigbm_msm ####################
