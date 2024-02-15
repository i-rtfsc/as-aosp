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

project(dnsresolver_aidl_interface-V2-ndk)

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

#################### dnsresolver_aidl_interface-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V2-ndk ####################

#################### dnsresolver_aidl_interface-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V2-ndk_platform ####################

#################### dnsresolver_aidl_interface-V7-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V7-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V7-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V7-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V7-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V7-ndk ####################

#################### dnsresolver_aidl_interface-V5-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V5-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V5-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V5-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V5-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V5-cpp ####################

#################### dnsresolver_aidl_interface-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V3-ndk ####################

#################### stats_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/stats_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/DnsResolver/stats.pb.cc
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
    add_library(stats_proto SHARED ${SOURCE_FILES})

    target_include_directories(stats_proto
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/DnsResolver
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/stats_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/DnsResolver
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/stats_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(stats_proto PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(stats_proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### stats_proto ####################

#################### dnsresolver_aidl_interface-V12-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V12-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V12-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V12-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V12-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V12-ndk_platform ####################

#################### dnsresolver_aidl_interface-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V1-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V1-cpp ####################

#################### dnsresolver_aidl_interface-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V5-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V5-ndk_platform ####################

#################### dnsresolver_aidl_interface-V9-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V9-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V9-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V9-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V9-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V9-ndk_platform ####################

#################### dnsresolver_aidl_interface-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V6-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V6-ndk ####################

#################### dnsresolver_aidl_interface-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V4-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V4-cpp ####################

#################### dnsresolver_aidl_interface-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V2-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V2-cpp ####################

#################### dnsresolver_aidl_interface-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V5-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V5-ndk ####################

#################### dnsresolver_aidl_interface-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V4-ndk_platform ####################

#################### dnsresolver_aidl_interface-V8-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V8-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V8-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V8-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V8-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V8-ndk_platform ####################

#################### dnsresolver_aidl_interface-V7-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V7-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V7-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V7-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V7-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V7-cpp ####################

#################### dnsresolver_aidl_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V1-ndk ####################

#################### dnsresolver_aidl_interface-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V3-ndk_platform ####################

#################### dnsresolver_aidl_interface-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V3-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V3-cpp ####################

#################### dnsresolver_aidl_interface-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V4-ndk ####################

#################### dnsresolver_aidl_interface-V6-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-cpp-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V6-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V6-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V6-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V6-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V6-cpp ####################

#################### libstatslog_resolv ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/statslog_resolv.cpp/gen/statslog_resolv.cpp
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
    add_library(libstatslog_resolv SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_resolv
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/DnsResolver
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/statslog_resolv.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/statslog_resolv.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libstatslog_resolv PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(libstatslog_resolv
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libm
            libstatspush_compat
    )
endif ()
#################### libstatslog_resolv ####################

#################### dnsresolver_aidl_interface-V11-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V11-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V11-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V11-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V11-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V11-ndk_platform ####################

#################### dnsresolver_aidl_interface-V11-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V11-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V11-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V11-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V11-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V11-ndk ####################

#################### dnsresolver_aidl_interface-V8-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V8-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V8-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V8-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V8-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V8-ndk ####################

#################### dnsresolver_aidl_interface-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V6-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V6-ndk_platform ####################

#################### dnsresolver_aidl_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V1-ndk_platform ####################

#################### dnsresolver_aidl_interface-V12-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V12-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V12-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V12-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V12-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V12-cpp ####################

#################### dnsresolver_aidl_interface-V10-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V10-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V10-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V10-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V10-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V10-ndk ####################

#################### dnsresolver_aidl_interface-V9-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V9-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V9-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V9-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V9-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V9-ndk ####################

#################### libnetd_resolv ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/DnsResolver/getaddrinfo.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/gethnamaddr.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/sethostent.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_cache.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/resolv_filter.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_comp.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_debug.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_mkquery.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_query.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_send.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/res_stats.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/util.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/Dns64Configuration.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsProxyListener.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsQueryLog.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsResolver.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsResolverService.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsStats.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsDispatcher.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsQueryMap.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsTransport.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsServer.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsSessionCache.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/DnsTlsSocket.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/Experiments.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/PrivateDnsConfiguration.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/ResolverController.cpp
        ${ANDROID_ROOT}/packages/modules/DnsResolver/ResolverEventReporter.cpp
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
    add_library(libnetd_resolv SHARED ${SOURCE_FILES})

    target_include_directories(libnetd_resolv
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/DnsResolver/include
            ${ANDROID_ROOT}/packages/modules/DnsResolver
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/statslog_resolv.h/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/stats_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/DnsResolver
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/stats_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libnetd_resolv PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(libnetd_resolv
            dnsresolver_aidl_interface-lateststable-ndk
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libdoh_ffi
            liblog
            libm
            libmodules-utils-build
            libnetdbinder_utils_headers
            libnetdutils
            libprotobuf-cpp-lite
            libssl
            libstatslog_resolv
            libstatspush_compat
            libsysutils
            libutils
            netd_event_listener_interface-lateststable-ndk
            server_configurable_flags
            stats_proto
    )
endif ()
#################### libnetd_resolv ####################

#################### dnsresolver_aidl_interface-V11-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V11-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V11-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V11-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V11-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V11-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V11-cpp ####################

#################### dnsresolver_aidl_interface-V8-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V8-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V8-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V8-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V8-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V8-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V8-cpp ####################

#################### dnsresolver_aidl_interface-V10-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V10-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V10-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V10-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V10-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V10-cpp ####################

#################### dnsresolver_aidl_interface-V10-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V10-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V10-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V10-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V10-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V10-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V10-ndk_platform ####################

#################### dnsresolver_aidl_interface-V12-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V12-ndk SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V12-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V12-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V12-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V12-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk
    )
endif ()
#################### dnsresolver_aidl_interface-V12-ndk ####################

#################### dnsresolver_aidl_interface-V7-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V7-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V7-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V7-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V7-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V7-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            netd_event_listener_interface-V1-ndk_platform
    )
endif ()
#################### dnsresolver_aidl_interface-V7-ndk_platform ####################

#################### dnsresolver_aidl_interface-V9-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/IDnsResolver.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/ResolverHostsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/ResolverOptionsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/ResolverParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/resolv/aidl/DnsHealthEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/resolv/aidl/IDnsResolverUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/resolv/aidl/Nat64PrefixEventParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/android/net/resolv/aidl/PrivateDnsValidationEventParcel.cpp
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
    add_library(dnsresolver_aidl_interface-V9-cpp SHARED ${SOURCE_FILES})

    target_include_directories(dnsresolver_aidl_interface-V9-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/DnsResolver/dnsresolver_aidl_interface-V9-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(dnsresolver_aidl_interface-V9-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(dnsresolver_aidl_interface-V9-cpp
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
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### dnsresolver_aidl_interface-V9-cpp ####################
