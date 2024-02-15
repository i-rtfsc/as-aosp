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

project(com.google.hardware.pixel.display-V6-ndk_platform)

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

#################### com.google.hardware.pixel.display-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V6-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V6-ndk_platform ####################

#################### com.google.hardware.pixel.display-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V1-ndk_platform ####################

#################### com.google.hardware.pixel.display-V8-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V8-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V8-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V8-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V8-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V8-ndk ####################

#################### com.google.hardware.pixel.display-V7-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V7-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V7-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V7-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V7-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V7-ndk_platform ####################

#################### com.google.hardware.pixel.display-V9-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V9-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V9-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V9-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V9-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V9-ndk ####################

#################### com.google.hardware.pixel.display-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V5-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V5-ndk ####################

#################### com.google.hardware.pixel.display-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V2-ndk_platform ####################

#################### com.google.hardware.pixel.display-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V1-ndk ####################

#################### com.google.hardware.pixel.display-V9-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V9-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V9-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V9-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V9-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V9-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V9-ndk_platform ####################

#################### com.google.hardware.pixel.display-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V5-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V5-ndk_platform ####################

#################### com.google.hardware.pixel.display-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V4-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V4-ndk ####################

#################### com.google.hardware.pixel.display-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V2-ndk ####################

#################### com.google.hardware.pixel.display-V8-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V8-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V8-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V8-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V8-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V8-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V8-ndk_platform ####################

#################### com.google.hardware.pixel.display-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V4-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V4-ndk_platform ####################

#################### com.google.hardware.pixel.display-V7-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V7-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V7-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V7-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V7-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V7-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V7-ndk ####################

#################### com.google.hardware.pixel.display-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V3-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V3-ndk ####################

#################### com.google.hardware.pixel.display-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/HistogramErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/HistogramPos.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/LbeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/PanelCalibrationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/com/google/hardware/pixel/display/Weight.cpp
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
    add_library(com.google.hardware.pixel.display-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V6-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### com.google.hardware.pixel.display-V6-ndk ####################

#################### com.google.hardware.pixel.display-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk_platform-source/gen/com/google/hardware/pixel/display/HbmState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk_platform-source/gen/com/google/hardware/pixel/display/IDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk_platform-source/gen/com/google/hardware/pixel/display/LbeState.cpp
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
    add_library(com.google.hardware.pixel.display-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(com.google.hardware.pixel.display-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/google/interfaces/display/com.google.hardware.pixel.display-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(com.google.hardware.pixel.display-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(com.google.hardware.pixel.display-V3-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### com.google.hardware.pixel.display-V3-ndk_platform ####################
