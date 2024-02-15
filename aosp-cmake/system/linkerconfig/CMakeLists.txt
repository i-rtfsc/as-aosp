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

project(linkerconfig)

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

#################### linkerconfig ####################
set(FILES
        ${ANDROID_ROOT}/system/linkerconfig/main.cc
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
    add_executable(linkerconfig ${SOURCE_FILES})

    target_include_directories(linkerconfig
            PUBLIC
            ${ANDROID_ROOT}/system/linkerconfig
            ${ANDROID_ROOT}/system/linkerconfig/modules/include
            ${ANDROID_ROOT}/system/linkerconfig/contents/include
            ${ANDROID_ROOT}/system/linkerconfig/generator/include
    )

    # Add c/c++ flags
    target_compile_options(linkerconfig PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(linkerconfig
            lib_apex_manifest_proto_lite
            lib_linker_config_proto_lite
            libapexutil
            libbase
            libc++_static
            libclang_rt.builtins
            libicui18n
            libicuuc
            libicuuc_stubdata
            liblog_for_runtime_apex
            libprotobuf-cpp-lite
            libtinyxml2
            linkerconfig_contents
            linkerconfig_generator
            linkerconfig_modules
    )
endif ()
#################### linkerconfig ####################

#################### linkerconfig_contents ####################
set(FILES
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/apexdefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/apexemptydefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/apexplatform.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/isolateddefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/postinstall.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/productdefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/recoverydefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/rs.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/sphal.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/system.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/systemdefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/unrestricteddefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/vendordefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/vndk.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/namespace/vndkinsystem.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/apexdefault.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/isolated.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/legacy.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/postinstall.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/product.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/recovery.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/sectionbuilder.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/system.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/unrestricted.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/section/vendor.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/configuration/apexconfig.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/configuration/baseconfig.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/configuration/legacy.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/configuration/recovery.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/context/context.cc
        ${ANDROID_ROOT}/system/linkerconfig/contents/common/system_links.cc
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
    add_library(linkerconfig_contents SHARED ${SOURCE_FILES})

    target_include_directories(linkerconfig_contents
            PUBLIC
            ${ANDROID_ROOT}/system/linkerconfig/contents/include
            ${ANDROID_ROOT}/system/linkerconfig
            ${ANDROID_ROOT}/system/linkerconfig/modules/include
    )

    # Add c/c++ flags
    target_compile_options(linkerconfig_contents PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(linkerconfig_contents
            lib_apex_manifest_proto_lite
            lib_linker_config_proto_lite
            libapexutil
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libicuuc_stubdata
            liblog_for_runtime_apex
            libm
            libprotobuf-cpp-lite
            libtinyxml2
            linkerconfig_modules
    )
endif ()
#################### linkerconfig_contents ####################

#################### linkerconfig_modules ####################
set(FILES
        ${ANDROID_ROOT}/system/linkerconfig/modules/apex.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/basecontext.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/configparser.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/configuration.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/configwriter.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/environment.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/link.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/namespace.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/section.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/stringutil.cc
        ${ANDROID_ROOT}/system/linkerconfig/modules/variables.cc
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
    add_library(linkerconfig_modules SHARED ${SOURCE_FILES})

    target_include_directories(linkerconfig_modules
            PUBLIC
            ${ANDROID_ROOT}/system/linkerconfig/modules/include
            ${ANDROID_ROOT}/system/linkerconfig
    )

    # Add c/c++ flags
    target_compile_options(linkerconfig_modules PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(linkerconfig_modules
            lib_apex_manifest_proto_lite
            lib_linker_config_proto_lite
            libapexutil
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libicuuc_stubdata
            liblog_for_runtime_apex
            libm
            libprotobuf-cpp-lite
            libtinyxml2
    )
endif ()
#################### linkerconfig_modules ####################

#################### linkerconfig_generator ####################
set(FILES
        ${ANDROID_ROOT}/system/linkerconfig/generator/librarylistloader.cc
        ${ANDROID_ROOT}/system/linkerconfig/generator/variableloader.cc
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
    add_library(linkerconfig_generator SHARED ${SOURCE_FILES})

    target_include_directories(linkerconfig_generator
            PUBLIC
            ${ANDROID_ROOT}/system/linkerconfig/generator/include
            ${ANDROID_ROOT}/system/linkerconfig
            ${ANDROID_ROOT}/system/linkerconfig/modules/include
            ${ANDROID_ROOT}/system/linkerconfig/contents/include
    )

    # Add c/c++ flags
    target_compile_options(linkerconfig_generator PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(linkerconfig_generator
            lib_apex_manifest_proto_lite
            lib_linker_config_proto_lite
            libapexutil
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libicuuc_stubdata
            liblog_for_runtime_apex
            libm
            libprotobuf-cpp-lite
            libtinyxml2
            linkerconfig_contents
            linkerconfig_modules
    )
endif ()
#################### linkerconfig_generator ####################
