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

project(computepipe_face_runner)

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

#################### computepipe_face_tracker ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example/FaceTracker.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example/ClientSvc.cpp
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
    add_executable(computepipe_face_tracker ${SOURCE_FILES})

    target_include_directories(computepipe_face_tracker
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/example/libcomputepipefaceproto/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/computepipe/example
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/example/libcomputepipefaceproto/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(computepipe_face_tracker PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(computepipe_face_tracker
            android.automotive.computepipe.registry-V2-ndk
            android.automotive.computepipe.runner-V2-ndk
            android.hardware.automotive.occupant_awareness-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcomputepipefaceproto
            libdl
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### computepipe_face_tracker ####################

#################### computepipe_face_runner ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example/Runner.cpp
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
    add_executable(computepipe_face_runner ${SOURCE_FILES})

    target_include_directories(computepipe_face_runner
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example
    )

    # Add c/c++ flags
    target_compile_options(computepipe_face_runner PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(computepipe_face_runner
            android.automotive.computepipe.registry-V2-ndk
            android.automotive.computepipe.runner-V2-ndk
            computepipe_client_interface
            computepipe_prebuilt_graph
            computepipe_runner_engine
            computepipe_runner_includes
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcomputepipeprotos
            libdl
            libfacegraph
            liblog
            libm
            libnativewindow
            libprotobuf-cpp-lite
            libutils
    )
endif ()
#################### computepipe_face_runner ####################

#################### libcomputepipefaceproto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/example/libcomputepipefaceproto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/computepipe/example/FaceOutput.pb.cc
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
    add_library(libcomputepipefaceproto SHARED ${SOURCE_FILES})

    target_include_directories(libcomputepipefaceproto
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/computepipe/example
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/example/libcomputepipefaceproto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/services/Car/cpp/computepipe/example
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/computepipe/example/libcomputepipefaceproto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libcomputepipefaceproto PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libcomputepipefaceproto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libcomputepipefaceproto ####################
