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

project(qrtr-lookup)

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

#################### libqrtr ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/lib/qrtr.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/lib/qmi.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/logging.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/qrtr.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/qmi.c
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
    add_library(libqrtr SHARED ${SOURCE_FILES})

    target_include_directories(libqrtr
            PUBLIC
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr
    )

    target_link_libraries(libqrtr
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libqrtr ####################

#################### qrtr-cfg ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/cfg.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/logging.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/addr.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/cfg.c
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
    add_executable(qrtr-cfg ${SOURCE_FILES})

    target_include_directories(qrtr-cfg
            PUBLIC
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr
    )

    target_link_libraries(qrtr-cfg
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libqrtr
            libc
            libdl
            libm
    )
endif ()
#################### qrtr-cfg ####################

#################### qrtr-ns ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/ns.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/map.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/hash.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/waiter.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/util.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/logging.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/addr.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/ns.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/map.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/hash.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/waiter.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/util.c
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
    add_executable(qrtr-ns ${SOURCE_FILES})

    target_include_directories(qrtr-ns
            PUBLIC
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr
    )

    target_link_libraries(qrtr-ns
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libqrtr
            libc
            libdl
            libm
    )
endif ()
#################### qrtr-ns ####################

#################### qrtr-lookup ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/lookup.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/qmi-framework/qrtr/src/util.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib/logging.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/lookup.c
        ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/src/util.c
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
    add_executable(qrtr-lookup ${SOURCE_FILES})

    target_include_directories(qrtr-lookup
            PUBLIC
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr/lib
            ${ANDROID_ROOT}/device/linaro/dragonboard/qcom/qrtr
    )

    target_link_libraries(qrtr-lookup
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libqrtr
            libc
            libdl
            libm
    )
endif ()
#################### qrtr-lookup ####################
