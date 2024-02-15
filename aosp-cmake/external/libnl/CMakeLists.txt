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

project(libnl)

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

#################### libnl ####################
set(FILES
        ${ANDROID_ROOT}/external/libnl/lib/addr.c
        ${ANDROID_ROOT}/external/libnl/lib/attr.c
        ${ANDROID_ROOT}/external/libnl/lib/cache.c
        ${ANDROID_ROOT}/external/libnl/lib/cache_mngr.c
        ${ANDROID_ROOT}/external/libnl/lib/cache_mngt.c
        ${ANDROID_ROOT}/external/libnl/lib/data.c
        ${ANDROID_ROOT}/external/libnl/lib/error.c
        ${ANDROID_ROOT}/external/libnl/lib/handlers.c
        ${ANDROID_ROOT}/external/libnl/lib/hash.c
        ${ANDROID_ROOT}/external/libnl/lib/hashtable.c
        ${ANDROID_ROOT}/external/libnl/lib/mpls.c
        ${ANDROID_ROOT}/external/libnl/lib/msg.c
        ${ANDROID_ROOT}/external/libnl/lib/nl.c
        ${ANDROID_ROOT}/external/libnl/lib/object.c
        ${ANDROID_ROOT}/external/libnl/lib/socket.c
        ${ANDROID_ROOT}/external/libnl/lib/utils.c
        ${ANDROID_ROOT}/external/libnl/lib/version.c
        ${ANDROID_ROOT}/external/libnl/lib/fib_lookup/lookup.c
        ${ANDROID_ROOT}/external/libnl/lib/fib_lookup/request.c
        ${ANDROID_ROOT}/external/libnl/lib/genl/ctrl.c
        ${ANDROID_ROOT}/external/libnl/lib/genl/family.c
        ${ANDROID_ROOT}/external/libnl/lib/genl/genl.c
        ${ANDROID_ROOT}/external/libnl/lib/genl/mngt.c
        ${ANDROID_ROOT}/external/libnl/lib/netfilter/nfnl.c
        ${ANDROID_ROOT}/external/libnl/lib/route/link.c
        ${ANDROID_ROOT}/external/libnl/lib/route/link/api.c
        ${ANDROID_ROOT}/external/libnl/lib/route/link/macsec.c
        ${ANDROID_ROOT}/external/libnl/lib/route/link/sriov.c
        ${ANDROID_ROOT}/external/libnl/lib/route/link/vlan.c
        ${ANDROID_ROOT}/external/libnl/lib/route/nexthop.c
        ${ANDROID_ROOT}/external/libnl/lib/route/nexthop_encap.c
        ${ANDROID_ROOT}/external/libnl/lib/route/nh_encap_mpls.c
        ${ANDROID_ROOT}/external/libnl/lib/route/route.c
        ${ANDROID_ROOT}/external/libnl/lib/route/route_obj.c
        ${ANDROID_ROOT}/external/libnl/lib/route/route_utils.c
        ${ANDROID_ROOT}/external/libnl/lib/route/rtnl.c
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
    add_library(libnl SHARED ${SOURCE_FILES})

    target_include_directories(libnl
            PUBLIC
            ${ANDROID_ROOT}/external/libnl/include
            ${ANDROID_ROOT}/external/libnl
    )

    # Add c/c++ flags
    target_compile_options(libnl PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -UNDEBUG
    )

    target_link_libraries(libnl
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libnl ####################
