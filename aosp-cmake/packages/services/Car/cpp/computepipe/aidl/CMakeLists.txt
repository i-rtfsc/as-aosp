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

project(android.automotive.computepipe.runner-V2-ndk_platform)

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

#################### android.automotive.computepipe.runner-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeDebugger.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeRunner.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeStateCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PacketDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PacketDescriptorPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputSourceDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfigOutputDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfigPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeProfilingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/android/automotive/computepipe/runner/ProfilingData.cpp
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
    add_library(android.automotive.computepipe.runner-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.runner-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.runner-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.runner-V2-ndk_platform
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
#################### android.automotive.computepipe.runner-V2-ndk_platform ####################

#################### android.automotive.computepipe.runner-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/IPipeDebugger.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/IPipeRunner.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/IPipeStateCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/IPipeStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PacketDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PacketDescriptorPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputSourceDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfigOutputDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfigPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeProfilingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/android/automotive/computepipe/runner/ProfilingData.cpp
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
    add_library(android.automotive.computepipe.runner-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.runner-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.runner-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.runner-V1-ndk
            android.hardware.graphics.common-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.computepipe.runner-V1-ndk ####################

#################### android.automotive.computepipe.runner-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/IPipeDebugger.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/IPipeRunner.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/IPipeStateCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/IPipeStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PacketDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PacketDescriptorPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputSourceDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfigOutputDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeOutputConfigPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeProfilingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/android/automotive/computepipe/runner/ProfilingData.cpp
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
    add_library(android.automotive.computepipe.runner-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.runner-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.runner-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.runner-V2-ndk
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
#################### android.automotive.computepipe.runner-V2-ndk ####################

#################### android.automotive.computepipe.registry-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk_platform-source/gen/android/automotive/computepipe/registry/IClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk_platform-source/gen/android/automotive/computepipe/registry/IPipeQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk_platform-source/gen/android/automotive/computepipe/registry/IPipeRegistration.cpp
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
    add_library(android.automotive.computepipe.registry-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.registry-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.registry-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.registry-V1-ndk_platform
            android.automotive.computepipe.runner-V1-ndk_platform
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
#################### android.automotive.computepipe.registry-V1-ndk_platform ####################

#################### android.automotive.computepipe.runner-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeDebugger.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeRunner.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeStateCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/IPipeStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PacketDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PacketDescriptorPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigCameraType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigImageFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputSourceDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeInputConfigVideoFileType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOffloadConfigOffloadType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfigOutputDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeOutputConfigPacketType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeProfilingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/PipeTerminationConfigTerminationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/android/automotive/computepipe/runner/ProfilingData.cpp
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
    add_library(android.automotive.computepipe.runner-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.runner-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.runner-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.runner-V1-ndk_platform
            android.hardware.graphics.common-V1-ndk_platform
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
#################### android.automotive.computepipe.runner-V1-ndk_platform ####################

#################### android.automotive.computepipe.registry-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk-source/gen/android/automotive/computepipe/registry/IClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk-source/gen/android/automotive/computepipe/registry/IPipeQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk-source/gen/android/automotive/computepipe/registry/IPipeRegistration.cpp
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
    add_library(android.automotive.computepipe.registry-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.registry-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.registry-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.registry-V1-ndk
            android.automotive.computepipe.runner-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.computepipe.registry-V1-ndk ####################

#################### android.automotive.computepipe.registry-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk-source/gen/android/automotive/computepipe/registry/IClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk-source/gen/android/automotive/computepipe/registry/IPipeQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk-source/gen/android/automotive/computepipe/registry/IPipeRegistration.cpp
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
    add_library(android.automotive.computepipe.registry-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.registry-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.registry-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.registry-V2-ndk
            android.automotive.computepipe.runner-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.computepipe.registry-V2-ndk ####################

#################### android.automotive.computepipe.registry-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk_platform-source/gen/android/automotive/computepipe/registry/IClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk_platform-source/gen/android/automotive/computepipe/registry/IPipeQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk_platform-source/gen/android/automotive/computepipe/registry/IPipeRegistration.cpp
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
    add_library(android.automotive.computepipe.registry-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.computepipe.registry-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.runner-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/aidl/android.automotive.computepipe.registry-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.computepipe.registry-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.computepipe.registry-V2-ndk_platform
            android.automotive.computepipe.runner-V2-ndk_platform
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
#################### android.automotive.computepipe.registry-V2-ndk_platform ####################
