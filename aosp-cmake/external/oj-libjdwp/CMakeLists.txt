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

project(libjdwp)

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

#################### libjdwp ####################
set(FILES
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ArrayReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ArrayTypeImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ClassLoaderReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ClassObjectReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ClassTypeImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/DDMImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/EventRequestImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/FieldImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/FrameID.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/InterfaceTypeImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/MethodImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ObjectReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ReferenceTypeImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/SDE.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/StackFrameImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/StringReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ThreadGroupReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/ThreadReferenceImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/VirtualMachineImpl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/bag.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/classTrack.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/commonRef.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/debugDispatch.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/debugInit.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/debugLoop.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/error_messages.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/eventFilter.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/eventHandler.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/eventHelper.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/inStream.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/invoker.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/log_messages.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/outStream.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/standardHandlers.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/stepControl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/stream.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/threadControl.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/timing.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/transport.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/util.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/vmDebug.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/back/exec_md.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/back/linker_md.c
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
    add_library(libjdwp SHARED ${SOURCE_FILES})

    target_include_directories(libjdwp
            PUBLIC
            ${ANDROID_ROOT}/external/oj-libjdwp
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/back
            ${OUT_INTERMEDIATES_ROOT}/external/oj-libjdwp/jdwp_generated_headers/gen
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/npt
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/npt
    )

    # Add c/c++ flags
    target_compile_options(libjdwp PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
            -Wno-unused-function
    )

    target_link_libraries(libjdwp
            javavm_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjdwp_headers
            libm
            libnpt_headers
    )
endif ()
#################### libjdwp ####################

#################### libdt_socket ####################
set(FILES
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/transport/socket/socketTransport.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/transport/socket/socket_md.c
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
    add_library(libdt_socket SHARED ${SOURCE_FILES})

    target_include_directories(libdt_socket
            PUBLIC
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/transport/socket
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/transport/socket
            ${ANDROID_ROOT}/external/oj-libjdwp
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/back/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/back
            ${OUT_INTERMEDIATES_ROOT}/external/oj-libjdwp/jdwp_generated_headers/gen
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/npt
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/npt
    )

    # Add c/c++ flags
    target_compile_options(libdt_socket PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libdt_socket
            javavm_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjdwp_headers
            libm
            libnpt_headers
    )
endif ()
#################### libdt_socket ####################

#################### libnpt ####################
set(FILES
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/npt/npt.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/share/npt/utf.c
        ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/npt/utf_md.c
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
    add_library(libnpt SHARED ${SOURCE_FILES})

    target_include_directories(libnpt
            PUBLIC
            ${ANDROID_ROOT}/external/oj-libjdwp
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/javavm/export
            ${ANDROID_ROOT}/external/oj-libjdwp/src/share/npt
            ${ANDROID_ROOT}/external/oj-libjdwp/src/solaris/npt
    )

    # Add c/c++ flags
    target_compile_options(libnpt PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libnpt
            javavm_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnpt_headers
    )
endif ()
#################### libnpt ####################
