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

project(apex_aidl_interface-cpp-analyzer)

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

#################### apexd ####################
set(FILES
        ${ANDROID_ROOT}/system/apex/apexd/apexd_main.cpp
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
    add_executable(apexd ${SOURCE_FILES})

    target_include_directories(apexd
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/sysprop/libcom.android.sysprop.apex/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(apexd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(apexd
            apex_aidl_interface-cpp
            lib_apex_manifest_proto
            lib_apex_session_state_proto
            lib_microdroid_metadata_proto
            libapex
            libapexd
            libapexd_checkpoint_vold
            libapexservice
            libavb
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcrypto
            libcutils
            libdl
            libdm
            libext2_uuid
            liblog
            liblogwrap
            libm
            libprotobuf-cpp-full
            libselinux
            libsigningutils
            libtinyxml2
            libutils
            libverity_tree
            libvold_binder
            libziparchive
    )
endif ()
#################### apexd ####################

#################### libapex ####################
set(FILES
        ${ANDROID_ROOT}/system/apex/apexd/apex_file.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apex_file_repository.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apex_manifest.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apex_shim.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_verity.cpp
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
    add_library(libapex SHARED ${SOURCE_FILES})

    target_include_directories(libapex
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
    )

    # Add c/c++ flags
    target_compile_options(libapex PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libapex
            lib_apex_manifest_proto
            lib_apex_session_state_proto
            lib_microdroid_metadata_proto
            libavb
            libbase
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libm
            libprotobuf-cpp-full
            libselinux
            libutils_headers
            libverity_tree
            libziparchive
    )
endif ()
#################### libapex ####################

#################### libapexd_checkpoint_vold ####################
set(FILES
        ${ANDROID_ROOT}/system/apex/apexd/apexd_checkpoint_vold.cpp
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
    add_library(libapexd_checkpoint_vold SHARED ${SOURCE_FILES})

    target_include_directories(libapexd_checkpoint_vold
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
    )

    # Add c/c++ flags
    target_compile_options(libapexd_checkpoint_vold PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libapexd_checkpoint_vold
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
            libvold_binder
    )
endif ()
#################### libapexd_checkpoint_vold ####################

#################### libapexservice ####################
set(FILES
        ${ANDROID_ROOT}/system/apex/apexd/apexservice.cpp
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
    add_library(libapexservice SHARED ${SOURCE_FILES})

    target_include_directories(libapexservice
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/sysprop/libcom.android.sysprop.apex/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libapexservice PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libapexservice
            apex_aidl_interface-cpp
            lib_apex_manifest_proto
            lib_apex_session_state_proto
            lib_microdroid_metadata_proto
            libapex
            libapexd
            libavb
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcrypto
            libcutils
            libdl
            libdm
            libext2_uuid
            liblog
            liblogwrap
            libm
            libprotobuf-cpp-full
            libselinux
            libsigningutils
            libtinyxml2
            libutils
            libverity_tree
            libvold_binder
            libziparchive
    )
endif ()
#################### libapexservice ####################

#################### libapexd ####################
set(FILES
        ${ANDROID_ROOT}/system/apex/apexd/apex_classpath.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apex_database.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_lifecycle.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_loop.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_private.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_session.cpp
        ${ANDROID_ROOT}/system/apex/apexd/apexd_verity.cpp
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
    add_library(libapexd SHARED ${SOURCE_FILES})

    target_include_directories(libapexd
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/sysprop/libcom.android.sysprop.apex/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list-tinyxml/gen/cpp/include
    )

    # Add c/c++ flags
    target_compile_options(libapexd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libapexd
            lib_apex_manifest_proto
            lib_apex_session_state_proto
            lib_microdroid_metadata_proto
            libapex
            libavb
            libbase
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcrypto
            libcutils
            libdl
            libdm
            libext2_uuid
            liblog
            liblogwrap
            libm
            libprotobuf-cpp-full
            libselinux
            libsigningutils
            libtinyxml2
            libverity_tree
            libvold_binder
            libziparchive
    )
endif ()
#################### libapexd ####################

#################### apex_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/ApexInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/ApexInfoList.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/ApexSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/ApexSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/CompressedApexInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/CompressedApexInfoList.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/android/apex/IApexService.cpp
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
    add_executable(apex_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(apex_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/apex/apexd
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(apex_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(apex_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            apex_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### apex_aidl_interface-cpp-analyzer ####################

#################### apex_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/ApexInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/ApexInfoList.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/ApexSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/ApexSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/CompressedApexInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/CompressedApexInfoList.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/android/apex/IApexService.cpp
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
    add_library(apex_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(apex_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(apex_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(apex_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### apex_aidl_interface-cpp ####################
