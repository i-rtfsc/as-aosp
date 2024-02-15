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

project(libincident)

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

#################### libincident ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/libs/incident/src/incident_report.cpp
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
    add_library(libincident SHARED ${SOURCE_FILES})

    target_include_directories(libincident
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/incident/include
            ${ANDROID_ROOT}/frameworks/base/libs/incident
            ${ANDROID_ROOT}/frameworks/base/libs/incident/include_priv
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/incident/libincidentpriv/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libincident PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libincident
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libincidentpriv
            liblog
            libm
            libutils
    )
endif ()
#################### libincident ####################

#################### libincidentpriv ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/incident/libincidentpriv/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IIncidentDumpCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/incident/libincidentpriv/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IIncidentManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/incident/libincidentpriv/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IIncidentReportStatusListener.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/incident/src/IncidentReportArgs.cpp
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
    add_library(libincidentpriv SHARED ${SOURCE_FILES})

    target_include_directories(libincidentpriv
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/incident/include_priv
            ${ANDROID_ROOT}/frameworks/base/libs/incident
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libs/incident/libincidentpriv/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libincidentpriv PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libincidentpriv
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libincidentcompanion
            liblog
            libm
            libplatformprotos
            libprotobuf-cpp-lite
            libutils
    )
endif ()
#################### libincidentpriv ####################
