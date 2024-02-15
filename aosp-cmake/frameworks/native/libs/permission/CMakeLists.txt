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

project(framework-permission-aidl-cpp-analyzer)

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

#################### libpermission ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/permission/AppOpsManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/permission/IAppOpsCallback.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/permission/IAppOpsService.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/permission/android/permission/PermissionChecker.cpp
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
    add_library(libpermission SHARED ${SOURCE_FILES})

    target_include_directories(libpermission
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/permission/include
            ${ANDROID_ROOT}/frameworks/native/libs/permission
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libpermission PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libpermission
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libpermission ####################

#################### framework-permission-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/android/content/AttributionSourceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/android/permission/IPermissionChecker.cpp
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
    add_library(framework-permission-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(framework-permission-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(framework-permission-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(framework-permission-aidl-cpp
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
#################### framework-permission-aidl-cpp ####################

#################### framework-permission-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk_platform-source/gen/android/content/AttributionSourceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk_platform-source/gen/android/permission/IPermissionChecker.cpp
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
    add_library(framework-permission-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(framework-permission-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(framework-permission-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(framework-permission-aidl-ndk_platform
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
#################### framework-permission-aidl-ndk_platform ####################

#################### framework-permission-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk-source/gen/android/content/AttributionSourceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk-source/gen/android/permission/IPermissionChecker.cpp
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
    add_library(framework-permission-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(framework-permission-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(framework-permission-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(framework-permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### framework-permission-aidl-ndk ####################

#################### framework-permission-aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-analyzer-source/gen/android/content/AttributionSourceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-analyzer-source/gen/android/permission/IPermissionChecker.cpp
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
    add_executable(framework-permission-aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(framework-permission-aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/permission
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/permission/framework-permission-aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(framework-permission-aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(framework-permission-aidl-cpp-analyzer
            aidl-analyzer-main
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### framework-permission-aidl-cpp-analyzer ####################
