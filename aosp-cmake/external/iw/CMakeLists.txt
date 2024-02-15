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

project(iw)

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

#################### iw_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/iw/iw.c
        ${ANDROID_ROOT}/external/iw/genl.c
        ${ANDROID_ROOT}/external/iw/event.c
        ${ANDROID_ROOT}/external/iw/info.c
        ${ANDROID_ROOT}/external/iw/phy.c
        ${ANDROID_ROOT}/external/iw/interface.c
        ${ANDROID_ROOT}/external/iw/ibss.c
        ${ANDROID_ROOT}/external/iw/station.c
        ${ANDROID_ROOT}/external/iw/survey.c
        ${ANDROID_ROOT}/external/iw/util.c
        ${ANDROID_ROOT}/external/iw/ocb.c
        ${ANDROID_ROOT}/external/iw/mesh.c
        ${ANDROID_ROOT}/external/iw/mpath.c
        ${ANDROID_ROOT}/external/iw/mpp.c
        ${ANDROID_ROOT}/external/iw/scan.c
        ${ANDROID_ROOT}/external/iw/reg.c
        ${ANDROID_ROOT}/external/iw/reason.c
        ${ANDROID_ROOT}/external/iw/status.c
        ${ANDROID_ROOT}/external/iw/connect.c
        ${ANDROID_ROOT}/external/iw/link.c
        ${ANDROID_ROOT}/external/iw/offch.c
        ${ANDROID_ROOT}/external/iw/ps.c
        ${ANDROID_ROOT}/external/iw/cqm.c
        ${ANDROID_ROOT}/external/iw/bitrate.c
        ${ANDROID_ROOT}/external/iw/wowlan.c
        ${ANDROID_ROOT}/external/iw/coalesce.c
        ${ANDROID_ROOT}/external/iw/roc.c
        ${ANDROID_ROOT}/external/iw/p2p.c
        ${ANDROID_ROOT}/external/iw/vendor.c
        ${ANDROID_ROOT}/external/iw/sections.c
        ${ANDROID_ROOT}/external/iw/ap.c
        ${ANDROID_ROOT}/external/iw/bloom.c
        ${ANDROID_ROOT}/external/iw/ftm.c
        ${ANDROID_ROOT}/external/iw/measurements.c
        ${ANDROID_ROOT}/external/iw/mgmt.c
        ${ANDROID_ROOT}/external/iw/nan.c
        ${ANDROID_ROOT}/external/iw/sar.c
        ${ANDROID_ROOT}/external/iw/sha256.c
        ${OUT_INTERMEDIATES_ROOT}/external/iw/iw_version/gen/version.c
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
    add_executable(iw_vendor ${SOURCE_FILES})

    target_include_directories(iw_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/iw
    )

    # Add c/c++ flags
    target_compile_options(iw_vendor PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(iw_vendor
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libnl
            libc
            libdl
            libm
    )
endif ()
#################### iw_vendor ####################

#################### iw ####################
set(FILES
        ${ANDROID_ROOT}/external/iw/iw.c
        ${ANDROID_ROOT}/external/iw/genl.c
        ${ANDROID_ROOT}/external/iw/event.c
        ${ANDROID_ROOT}/external/iw/info.c
        ${ANDROID_ROOT}/external/iw/phy.c
        ${ANDROID_ROOT}/external/iw/interface.c
        ${ANDROID_ROOT}/external/iw/ibss.c
        ${ANDROID_ROOT}/external/iw/station.c
        ${ANDROID_ROOT}/external/iw/survey.c
        ${ANDROID_ROOT}/external/iw/util.c
        ${ANDROID_ROOT}/external/iw/ocb.c
        ${ANDROID_ROOT}/external/iw/mesh.c
        ${ANDROID_ROOT}/external/iw/mpath.c
        ${ANDROID_ROOT}/external/iw/mpp.c
        ${ANDROID_ROOT}/external/iw/scan.c
        ${ANDROID_ROOT}/external/iw/reg.c
        ${ANDROID_ROOT}/external/iw/reason.c
        ${ANDROID_ROOT}/external/iw/status.c
        ${ANDROID_ROOT}/external/iw/connect.c
        ${ANDROID_ROOT}/external/iw/link.c
        ${ANDROID_ROOT}/external/iw/offch.c
        ${ANDROID_ROOT}/external/iw/ps.c
        ${ANDROID_ROOT}/external/iw/cqm.c
        ${ANDROID_ROOT}/external/iw/bitrate.c
        ${ANDROID_ROOT}/external/iw/wowlan.c
        ${ANDROID_ROOT}/external/iw/coalesce.c
        ${ANDROID_ROOT}/external/iw/roc.c
        ${ANDROID_ROOT}/external/iw/p2p.c
        ${ANDROID_ROOT}/external/iw/vendor.c
        ${ANDROID_ROOT}/external/iw/sections.c
        ${ANDROID_ROOT}/external/iw/ap.c
        ${ANDROID_ROOT}/external/iw/bloom.c
        ${ANDROID_ROOT}/external/iw/ftm.c
        ${ANDROID_ROOT}/external/iw/measurements.c
        ${ANDROID_ROOT}/external/iw/mgmt.c
        ${ANDROID_ROOT}/external/iw/nan.c
        ${ANDROID_ROOT}/external/iw/sar.c
        ${ANDROID_ROOT}/external/iw/sha256.c
        ${OUT_INTERMEDIATES_ROOT}/external/iw/iw_version/gen/version.c
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
    add_executable(iw ${SOURCE_FILES})

    target_include_directories(iw
            PUBLIC
            ${ANDROID_ROOT}/external/iw
    )

    # Add c/c++ flags
    target_compile_options(iw PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(iw
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnl
    )
endif ()
#################### iw ####################
