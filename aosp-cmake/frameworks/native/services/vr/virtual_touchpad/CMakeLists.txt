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

project(virtual_touchpad)

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

#################### libvirtualtouchpadclient ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/VirtualTouchpadClient.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/DvrVirtualTouchpadClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/vr/virtual_touchpad/libvirtualtouchpadclient/android_${OUT_ARCH_CPU}_static/gen/aidl/android/dvr/IVirtualTouchpadService.cpp
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
    add_library(libvirtualtouchpadclient SHARED ${SOURCE_FILES})

    target_include_directories(libvirtualtouchpadclient
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/include
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/vr/virtual_touchpad/libvirtualtouchpadclient/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libvirtualtouchpadclient PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libvirtualtouchpadclient
            jni_headers
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libdvr_headers
            liblog
            libm
            libutils
    )
endif ()
#################### libvirtualtouchpadclient ####################

#################### virtual_touchpad ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/main.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/VirtualTouchpadService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/vr/virtual_touchpad/virtual_touchpad/android_${OUT_ARCH_CPU}/gen/aidl/android/dvr/IVirtualTouchpadService.cpp
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
    add_executable(virtual_touchpad ${SOURCE_FILES})

    target_include_directories(virtual_touchpad
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/vr/virtual_touchpad/virtual_touchpad/android_${OUT_ARCH_CPU}/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/include
    )

    # Add c/c++ flags
    target_compile_options(virtual_touchpad PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(virtual_touchpad
            jni_headers
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdvr_headers
            liblog
            libm
            libutils
            libvirtualtouchpad
    )
endif ()
#################### virtual_touchpad ####################

#################### libvirtualtouchpad ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/EvdevInjector.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/VirtualTouchpadEvdev.cpp
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
    add_library(libvirtualtouchpad SHARED ${SOURCE_FILES})

    target_include_directories(libvirtualtouchpad
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad/include
            ${ANDROID_ROOT}/frameworks/native/services/vr/virtual_touchpad
    )

    # Add c/c++ flags
    target_compile_options(libvirtualtouchpad PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libvirtualtouchpad
            jni_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libdvr_headers
            liblog
            libm
            libutils
    )
endif ()
#################### libvirtualtouchpad ####################
