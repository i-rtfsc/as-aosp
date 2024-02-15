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

project(android.hardware.audio@2.0-impl)

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

#################### android.hardware.audio-2.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-2.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-2.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-2.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-2.0-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-2.0-util
            android.hardware.audio.common-2.0
            android.hardware.audio-2.0-util
            android.hardware.audio-2.0
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-2.0-impl ####################

#################### android.hardware.audio-7.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-7.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-7.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-7.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-7.0-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-7.0-enums
            android.hardware.audio.common-7.0-util
            android.hardware.audio.common-7.0
            android.hardware.audio-7.0-util
            android.hardware.audio-7.0
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-7.0-impl ####################

#################### android.hardware.audio-6.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-6.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-6.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-6.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-6.0-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-6.0-util
            android.hardware.audio.common-6.0
            android.hardware.audio-6.0-util
            android.hardware.audio-6.0
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-6.0-impl ####################

#################### android.hardware.audio-5.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-5.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-5.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-5.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-5.0-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-5.0-util
            android.hardware.audio.common-5.0
            android.hardware.audio-5.0-util
            android.hardware.audio-5.0
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-5.0-impl ####################

#################### android.hardware.audio-7.1-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-7.1-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-7.1-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-7.1-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-7.1-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-7.0
            android.hardware.audio.common-7.1-enums
            android.hardware.audio.common-7.1-util
            android.hardware.audio-7.0
            android.hardware.audio-7.1-util
            android.hardware.audio-7.1
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-7.1-impl ####################

#################### android.hardware.audio-4.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/DevicesFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/ParametersUtil.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/PrimaryDevice.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/Stream.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamIn.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/StreamOut.cpp
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
    add_library(android.hardware.audio-4.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio-4.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/include
            ${ANDROID_ROOT}/hardware/interfaces/audio/core/all-versions/default/util/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio-4.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.audio-4.0-impl
            android.hardware.audio-impl_headers
            android.hardware.audio.common-util
            android.hardware.audio.common.util-all-versions
            android.hardware.audio.common-4.0-util
            android.hardware.audio.common-4.0
            android.hardware.audio-4.0-util
            android.hardware.audio-4.0
            libaudio_system_headers
            libaudiofoundation
            libaudioutils_headers
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediautils_headers
            libmediautils_vendor
            libmemunreachable
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio-4.0-impl ####################
