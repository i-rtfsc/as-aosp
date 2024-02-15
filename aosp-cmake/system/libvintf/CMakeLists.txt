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

project(vintf)

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

#################### libvintf ####################
set(FILES
        ${ANDROID_ROOT}/system/libvintf/parse_string.cpp
        ${ANDROID_ROOT}/system/libvintf/parse_xml.cpp
        ${ANDROID_ROOT}/system/libvintf/Apex.cpp
        ${ANDROID_ROOT}/system/libvintf/CompatibilityMatrix.cpp
        ${ANDROID_ROOT}/system/libvintf/FileSystem.cpp
        ${ANDROID_ROOT}/system/libvintf/FQName.cpp
        ${ANDROID_ROOT}/system/libvintf/FqInstance.cpp
        ${ANDROID_ROOT}/system/libvintf/HalManifest.cpp
        ${ANDROID_ROOT}/system/libvintf/HalInterface.cpp
        ${ANDROID_ROOT}/system/libvintf/KernelConfigTypedValue.cpp
        ${ANDROID_ROOT}/system/libvintf/KernelConfigParser.cpp
        ${ANDROID_ROOT}/system/libvintf/KernelInfo.cpp
        ${ANDROID_ROOT}/system/libvintf/RuntimeInfo.cpp
        ${ANDROID_ROOT}/system/libvintf/ManifestHal.cpp
        ${ANDROID_ROOT}/system/libvintf/ManifestInstance.cpp
        ${ANDROID_ROOT}/system/libvintf/MatrixHal.cpp
        ${ANDROID_ROOT}/system/libvintf/MatrixInstance.cpp
        ${ANDROID_ROOT}/system/libvintf/MatrixKernel.cpp
        ${ANDROID_ROOT}/system/libvintf/PropertyFetcher.cpp
        ${ANDROID_ROOT}/system/libvintf/Regex.cpp
        ${ANDROID_ROOT}/system/libvintf/SystemSdk.cpp
        ${ANDROID_ROOT}/system/libvintf/TransportArch.cpp
        ${ANDROID_ROOT}/system/libvintf/VintfObject.cpp
        ${ANDROID_ROOT}/system/libvintf/XmlFile.cpp
        ${ANDROID_ROOT}/system/libvintf/utils.cpp
        ${ANDROID_ROOT}/system/libvintf/RuntimeInfo-target.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list-tinyxml/gen/cpp/com_android_apex.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list-tinyxml/gen/cpp/com_android_apex_enums.cpp
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
    add_library(libvintf SHARED ${SOURCE_FILES})

    target_include_directories(libvintf
            PUBLIC
            ${ANDROID_ROOT}/system/libvintf/include
            ${ANDROID_ROOT}/system/libvintf
            ${ANDROID_ROOT}/system/libvintf/include/vintf
    )

    # Add c/c++ flags
    target_compile_options(libvintf PRIVATE
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libvintf
            libaidlmetadata_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlmetadata_headers
            libkver
            liblog
            libm
            libselinux
            libtinyxml2
            libz
    )
endif ()
#################### libvintf ####################

#################### vintf ####################
set(FILES
        ${ANDROID_ROOT}/system/libvintf/main.cpp
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
    add_executable(vintf ${SOURCE_FILES})

    target_include_directories(vintf
            PUBLIC
            ${ANDROID_ROOT}/system/libvintf
            ${ANDROID_ROOT}/system/libvintf/include
    )

    # Add c/c++ flags
    target_compile_options(vintf PRIVATE
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vintf
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libvintf
    )
endif ()
#################### vintf ####################

#################### libassemblevintf ####################
set(FILES
        ${ANDROID_ROOT}/system/libvintf/AssembleVintf.cpp
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
    add_library(libassemblevintf SHARED ${SOURCE_FILES})

    target_include_directories(libassemblevintf
            PUBLIC
            ${ANDROID_ROOT}/system/libvintf/include-test
            ${ANDROID_ROOT}/system/libvintf
            ${ANDROID_ROOT}/system/libvintf/include
    )

    # Add c/c++ flags
    target_compile_options(libassemblevintf PRIVATE
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(libassemblevintf
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libvintf
            libvts_vintf_test_common
    )
endif ()
#################### libassemblevintf ####################
