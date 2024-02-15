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

project(lib-platform-compat-native-api)

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

#################### lib-platform-compat-native-api ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/native/android/lib-platform-compat-native-api/android_${OUT_ARCH_CPU}_shared/gen/aidl/com/android/internal/compat/IPlatformCompatNative.cpp
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
    add_library(lib-platform-compat-native-api SHARED ${SOURCE_FILES})

    target_include_directories(lib-platform-compat-native-api
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/android/aidl
            ${ANDROID_ROOT}/frameworks/base/native/android
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/native/android/lib-platform-compat-native-api/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(lib-platform-compat-native-api PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(lib-platform-compat-native-api
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
#################### lib-platform-compat-native-api ####################

#################### libandroid ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/android/activity_manager.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/asset_manager.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/choreographer.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/configuration.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/hardware_buffer_jni.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/input.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/looper.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/native_activity.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/native_window_jni.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/net.c
        ${ANDROID_ROOT}/frameworks/base/native/android/obb.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/permission_manager.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/performance_hint.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/sensor.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/sharedmem.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/storage_manager.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/surface_control.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/surface_texture.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/system_fonts.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/trace.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/thermal.cpp
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
    add_library(libandroid SHARED ${SOURCE_FILES})

    target_include_directories(libandroid
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/android/include_platform
            ${ANDROID_ROOT}/frameworks/base/native/android
    )

    # Add c/c++ flags
    target_compile_options(libandroid PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libandroid
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.power-V4-ndk
            libEGL
            libGLESv2
            libactivitymanager_aidl
            libandroid_headers_private
            libandroid_runtime
            libandroidfw
            libarect
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui
            libharfbuzz_ng
            libhidlbase
            libhwui
            libhwui_internal_headers
            libinput
            liblog
            libm
            libminikin
            libnativedisplay
            libnativewindow
            libnetd_client
            libpowermanager
            libsensor
            libstorage
            libui
            libutils
            libxml2
    )
endif ()
#################### libandroid ####################

#################### libandroid_net ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/android/net.c
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
    add_library(libandroid_net SHARED ${SOURCE_FILES})

    target_include_directories(libandroid_net
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/android
    )

    # Add c/c++ flags
    target_compile_options(libandroid_net PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libandroid_net
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnetd_client
    )
endif ()
#################### libandroid_net ####################
