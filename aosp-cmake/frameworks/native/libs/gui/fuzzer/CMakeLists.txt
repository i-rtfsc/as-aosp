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

project(libgui_surfaceComposer_fuzzer)

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

#################### libgui_displayEvent_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_displayEvent_fuzzer.cpp
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
    add_executable(libgui_displayEvent_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_displayEvent_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_displayEvent_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_displayEvent_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_displayEvent_fuzzer ####################

#################### libgui_bufferQueue_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_bufferQueue_fuzzer.cpp
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
    add_executable(libgui_bufferQueue_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_bufferQueue_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_bufferQueue_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_bufferQueue_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_bufferQueue_fuzzer ####################

#################### libgui_surfaceComposerClient_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_surfaceComposerClient_fuzzer.cpp
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
    add_executable(libgui_surfaceComposerClient_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_surfaceComposerClient_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_surfaceComposerClient_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_surfaceComposerClient_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_surfaceComposerClient_fuzzer ####################

#################### libgui_consumer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_consumer_fuzzer.cpp
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
    add_executable(libgui_consumer_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_consumer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_consumer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_consumer_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_consumer_fuzzer ####################

#################### libgui_parcelable_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_parcelable_fuzzer.cpp
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
    add_executable(libgui_parcelable_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_parcelable_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_parcelable_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_parcelable_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_parcelable_fuzzer ####################

#################### libgui_surfaceComposer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer/libgui_surfaceComposer_fuzzer.cpp
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
    add_executable(libgui_surfaceComposer_fuzzer ${SOURCE_FILES})

    target_include_directories(libgui_surfaceComposer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/gui/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(libgui_surfaceComposer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgui_surfaceComposer_fuzzer
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.power-V4-cpp
            android.hidl.token-1.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libbase
            libbinder
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine_mocks
            libcutils
            libdl
            libdvr_headers
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
            libgui_window_info_static
            libhidlbase
            libinput
            liblog
            libm
            libnativewindow
            libpdx
            librenderengine_mocks
            libui
            libui_fuzzableDataspaces_headers
            libutils
            libvndksupport
            perfetto_trace_protos
    )
endif ()
#################### libgui_surfaceComposer_fuzzer ####################
