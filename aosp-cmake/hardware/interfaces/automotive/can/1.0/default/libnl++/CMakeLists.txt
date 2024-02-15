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

project(libnl++)

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

#################### libnl++ ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/common/Empty.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/common/Error.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/Ctrl.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/FamilyTracker.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/Generic.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/GenericMessageBase.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/Unknown.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/families/Mac80211hwsim.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/generic/families/Nl80211.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/route/Addr.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/route/Link.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/route/Route.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/route/attributes.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/route/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/MessageDefinition.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/NetlinkProtocol.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/all.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/protocols/structs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/Attributes.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/MessageFactory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/MessageMutator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/Socket.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/common.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/printer.cpp
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
    add_library(libnl++ SHARED ${SOURCE_FILES})

    target_include_directories(libnl++
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++/include
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/1.0/default/libnl++
    )

    # Add c/c++ flags
    target_compile_options(libnl++ PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libnl++
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
    )
endif ()
#################### libnl++ ####################
