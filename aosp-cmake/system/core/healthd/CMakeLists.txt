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

project(charger)

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

#################### android.hardware.health-2.0-service ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/HealthServiceDefault.cpp
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
    add_executable(android.hardware.health-2.0-service ${SOURCE_FILES})

    target_include_directories(android.hardware.health-2.0-service
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd
            ${ANDROID_ROOT}/system/core/healthd/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-2.0-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-2.0-service
            android.hardware.health-1.0-convert
            android.hardware.health-2.0-impl
            android.hardware.health-2.0
            libbase
            libbatterymonitor
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthservice
            libhealthstoragedefault
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.health-2.0-service ####################

#################### libcharger_sysprop ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/charger.sysprop.cpp
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
    add_library(libcharger_sysprop SHARED ${SOURCE_FILES})

    target_include_directories(libcharger_sysprop
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libcharger_sysprop PRIVATE
            -Werror
    )

    target_link_libraries(libcharger_sysprop
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libcharger_sysprop ####################

#################### libhealthd_draw ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/healthd_draw.cpp
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
    add_library(libhealthd_draw SHARED ${SOURCE_FILES})

    target_include_directories(libhealthd_draw
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libhealthd_draw PRIVATE
            -Werror
    )

    target_link_libraries(libhealthd_draw
            libbase
            libbatteryservice_headers
            libc
            libc++_static
            libc++demangle
            libcharger_sysprop
            libclang_rt.builtins
            libdl
            libm
            libminui
    )
endif ()
#################### libhealthd_draw ####################

#################### libhealthd_charger_nops ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/healthd_mode_charger_nops.cpp
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
    add_library(libhealthd_charger_nops SHARED ${SOURCE_FILES})

    target_include_directories(libhealthd_charger_nops
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd
            ${ANDROID_ROOT}/system/core/healthd/include
    )

    # Add c/c++ flags
    target_compile_options(libhealthd_charger_nops PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libhealthd_charger_nops
            android.hardware.health-2.1
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhealth2impl
            libhealthd_headers
            libhealthloop
            libm
            libutils
    )
endif ()
#################### libhealthd_charger_nops ####################

#################### libbatterymonitor ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/BatteryMonitor.cpp
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
    add_library(libbatterymonitor SHARED ${SOURCE_FILES})

    target_include_directories(libbatterymonitor
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include
            ${ANDROID_ROOT}/system/core/healthd
    )

    # Add c/c++ flags
    target_compile_options(libbatterymonitor PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libbatterymonitor
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-2.1
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhealthd_headers
            libm
            libutils
    )
endif ()
#################### libbatterymonitor ####################

#################### libbatterymonitor-V1 ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/BatteryMonitor_v1.cpp
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
    add_library(libbatterymonitor-V1 SHARED ${SOURCE_FILES})

    target_include_directories(libbatterymonitor-V1
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include
            ${ANDROID_ROOT}/system/core/healthd
    )

    # Add c/c++ flags
    target_compile_options(libbatterymonitor-V1 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libbatterymonitor-V1
            android.hardware.health-V1-ndk
            android.hardware.health-translate-V1-ndk
            android.hardware.health-2.1
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhealthd_headers
            libm
            libutils
    )
endif ()
#################### libbatterymonitor-V1 ####################

#################### libhealthd_charger_ui-V1 ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/healthd_mode_charger.cpp
        ${ANDROID_ROOT}/system/core/healthd/AnimationParser.cpp
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
    add_library(libhealthd_charger_ui-V1 SHARED ${SOURCE_FILES})

    target_include_directories(libhealthd_charger_ui-V1
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include
            ${ANDROID_ROOT}/system/core/healthd/include_charger
            ${ANDROID_ROOT}/system/core/healthd
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libhealthd_charger_ui-V1 PRIVATE
            -Werror
    )

    target_link_libraries(libhealthd_charger_ui-V1
            android.hardware.health-V1-ndk
            android.hardware.health-translate-V1-ndk
            libbase
            libc
            libc++_static
            libc++demangle
            libcharger_sysprop
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthd_draw
            libhealthd_headers
            libhealthloop
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
    )
endif ()
#################### libhealthd_charger_ui-V1 ####################

#################### libhealthd_charger ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/healthd_mode_charger_hidl.cpp
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
    add_library(libhealthd_charger SHARED ${SOURCE_FILES})

    target_include_directories(libhealthd_charger
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include
            ${ANDROID_ROOT}/system/core/healthd/include_charger
            ${ANDROID_ROOT}/system/core/healthd
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libhealthd_charger PRIVATE
            -Werror
    )

    target_link_libraries(libhealthd_charger
            android.hardware.health-1.0-convert
            android.hardware.health-2.1
            libbase
            libc
            libc++_static
            libc++demangle
            libcharger_sysprop
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealth2impl
            libhealthd_charger_ui
            liblog
            libm
            libutils
    )
endif ()
#################### libhealthd_charger ####################

#################### charger ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/charger.cpp
        ${ANDROID_ROOT}/system/core/healthd/charger_utils.cpp
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
    add_executable(charger ${SOURCE_FILES})

    target_include_directories(charger
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include_charger
            ${ANDROID_ROOT}/system/core/healthd
            ${ANDROID_ROOT}/system/core/healthd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(charger PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(charger
            android.hardware.health-V2-ndk
            android.hardware.health-1.0-convert
            android.hardware.health-2.0
            android.hardware.health-2.1
            libbase
            libbatterymonitor
            libc
            libc++
            libc++demangle
            libcharger_sysprop
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealth2impl
            libhealthd_charger
            libhealthd_charger_nops
            libhealthd_charger_ui
            libhealthd_draw
            libhealthloop
            libhidlbase
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
    )
endif ()
#################### charger ####################

#################### libhealthd_charger_ui ####################
set(FILES
        ${ANDROID_ROOT}/system/core/healthd/healthd_mode_charger.cpp
        ${ANDROID_ROOT}/system/core/healthd/AnimationParser.cpp
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
    add_library(libhealthd_charger_ui SHARED ${SOURCE_FILES})

    target_include_directories(libhealthd_charger_ui
            PUBLIC
            ${ANDROID_ROOT}/system/core/healthd/include
            ${ANDROID_ROOT}/system/core/healthd/include_charger
            ${ANDROID_ROOT}/system/core/healthd
            ${OUT_INTERMEDIATES_ROOT}/system/core/healthd/libcharger_sysprop/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libhealthd_charger_ui PRIVATE
            -Werror
    )

    target_link_libraries(libhealthd_charger_ui
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            libbase
            libc
            libc++_static
            libc++demangle
            libcharger_sysprop
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthd_draw
            libhealthd_headers
            libhealthloop
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
    )
endif ()
#################### libhealthd_charger_ui ####################
