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

project(vendor.aac.hardware.richtap.vibrator-cpp-analyzer)

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

#################### vendor.aac.hardware.richtap.vibrator-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-cpp-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-cpp-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V1-cpp ####################

#################### vendor.aac.hardware.richtap.vibrator-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V2-ndk ####################

#################### vendor.aac.hardware.richtap.vibrator-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V1-ndk ####################

#################### vendor.aac.hardware.richtap.vibrator-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-cpp-analyzer-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-cpp-analyzer-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_executable(vendor.aac.hardware.richtap.vibrator-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            vendor.aac.hardware.richtap.vibrator-V2-cpp
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-cpp-analyzer ####################

#################### vendor.aac.hardware.richtap.vibrator-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk_platform-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk_platform-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V2-ndk_platform ####################

#################### vendor.aac.hardware.richtap.vibrator-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-cpp-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-cpp-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V2-cpp ####################

#################### vendor.aac.hardware.richtap.vibrator-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk_platform-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk_platform-source/gen/vendor/aac/hardware/richtap/vibrator/IRichtapVibrator.cpp
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
    add_library(vendor.aac.hardware.richtap.vibrator-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.aac.hardware.richtap.vibrator-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/meizu/aac/interface/richtap/vibrator/aidl/vendor.aac.hardware.richtap.vibrator-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.aac.hardware.richtap.vibrator-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.aac.hardware.richtap.vibrator-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### vendor.aac.hardware.richtap.vibrator-V1-ndk_platform ####################
