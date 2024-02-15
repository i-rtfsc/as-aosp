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

project(vendor.qti.hardware.display.allocator-service)

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

#################### android.hardware.graphics.mapper-4.0-impl-qti-display ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/QtiMapper4.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/QtiMapperExtensions.cpp
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
    add_library(android.hardware.graphics.mapper-4.0-impl-qti-display SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.mapper-4.0-impl-qti-display
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/gralloc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.mapper-4.0-impl-qti-display PRIVATE
            -Wall
            -Werror
            -D__QTI_DISPLAY_GRALLOC__
    )

    target_link_libraries(android.hardware.graphics.mapper-4.0-impl-qti-display
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-2.1
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            display_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdmabufheap
            libgralloccore
            libgralloctypes
            libgrallocutils
            libhidlbase
            liblog
            libm
            libsync
            libutils
            vendor.qti.hardware.display.mapper-3.0
            vendor.qti.hardware.display.mapper-4.0
            vendor.qti.hardware.display.mapperextensions-1.0
            vendor.qti.hardware.display.mapperextensions-1.1
            vendor.qti.hardware.display.mapperextensions-1.2
            vendor.qti.hardware.display.mapperextensions-1.3
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.graphics.mapper-4.0-impl-qti-display ####################

#################### libgralloc.qti ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/QtiGralloc.cpp
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
    add_library(libgralloc.qti SHARED ${SOURCE_FILES})

    target_include_directories(libgralloc.qti
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/gralloc
    )

    # Add c/c++ flags
    target_compile_options(libgralloc.qti PRIVATE
            -Werror
            -D__QTI_DISPLAY_GRALLOC__
    )

    target_link_libraries(libgralloc.qti
            android.hardware.graphics.mapper-4.0
            display_headers
            display_intf_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgralloctypes
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libgralloc.qti ####################

#################### libgrallocutils ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_utils.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_adreno_info.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_camera_info.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_ubwcp_utils.cpp
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
    add_library(libgrallocutils SHARED ${SOURCE_FILES})

    target_include_directories(libgrallocutils
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/gralloc
    )

    # Add c/c++ flags
    target_compile_options(libgrallocutils PRIVATE
            -Wall
            -Werror
            -D__QTI_DISPLAY_GRALLOC__
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libgrallocutils
            android.hardware.graphics.common-1.2
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-2.1
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            device_kernel_headers
            display_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgralloctypes
            libhidlbase
            liblog
            libm
            libubwcp_headers
            libutils
            qti_display_kernel_headers
            qti_kernel_headers
            libc
            libdl
            libm
    )
endif ()
#################### libgrallocutils ####################

#################### libgralloccore ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_allocator.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_buf_mgr.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_dma_legacy_mgr.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_dma_mgr.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/gr_alloc_interface.cpp
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
    add_library(libgralloccore SHARED ${SOURCE_FILES})

    target_include_directories(libgralloccore
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/gralloc
    )

    # Add c/c++ flags
    target_compile_options(libgralloccore PRIVATE
            -Wall
            -Werror
            -D__QTI_DISPLAY_GRALLOC__
            -Wno-unused-parameter
    )

    target_link_libraries(libgralloccore
            android.hardware.graphics.mapper-2.1
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            device_kernel_headers
            display_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdmabufheap
            libgralloc.qti
            libgralloctypes
            libgrallocutils
            libhidlbase
            libion
            liblog
            libm
            libubwcp_headers
            libutils
            libvmmem
            libvmmem_headers
            qti_display_kernel_headers
            qti_kernel_headers
            smmu_proxy_uapi_header
            libc
            libdl
            libm
    )
endif ()
#################### libgralloccore ####################

#################### vendor.qti.hardware.display.allocator-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/QtiAllocatorAIDL.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/gralloc/service.cpp
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
    add_executable(vendor.qti.hardware.display.allocator-service ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.display.allocator-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/gralloc
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.display.allocator-service PRIVATE
            -Wall
            -Werror
            -D__QTI_DISPLAY_GRALLOC__
    )

    target_link_libraries(vendor.qti.hardware.display.allocator-service
            android.hardware.graphics.allocator-V1-ndk
            android.hardware.graphics.mapper-2.1
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            display_headers
            libaidlcommonsupport
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdmabufheap
            libgralloccore
            libgralloctypes
            libgrallocutils
            libhidlbase
            liblog
            libm
            libutils
            vendor.qti.hardware.display.mapper-3.0
            vendor.qti.hardware.display.mapper-4.0
            vendor.qti.hardware.display.mapperextensions-1.0
            vendor.qti.hardware.display.mapperextensions-1.1
            vendor.qti.hardware.display.mapperextensions-1.2
            vendor.qti.hardware.display.mapperextensions-1.3
            libc
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.display.allocator-service ####################
