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

project(libdynamic_depth)

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

#################### libdynamic_depth ####################
set(FILES
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/app_info.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/camera.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/cameras.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/container.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/depth_jpeg.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/depth_map.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/device.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/dynamic_depth.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/earth_pose.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/image.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/imaging_model.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/item.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/light_estimate.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/plane.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/planes.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/point_cloud.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/pose.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/profile.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/profiles.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/vendor_info.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/ascii_ctype.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/case.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/escaping.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/numbers.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/base64.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/file.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/jpeg_io.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/md5.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_data.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_parser.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_writer.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/deserializer_impl.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/search.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/serializer_impl.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/utils.cc
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
    add_library(libdynamic_depth SHARED ${SOURCE_FILES})

    target_include_directories(libdynamic_depth
            PUBLIC
            ${ANDROID_ROOT}/external/dynamic_depth
            ${ANDROID_ROOT}/external/dynamic_depth/includes
            ${ANDROID_ROOT}/external/dynamic_depth/internal
    )

    # Add c/c++ flags
    target_compile_options(libdynamic_depth PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libdynamic_depth
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libdynamic_depth-internal_headers
            libdynamic_depth-public_headers
            libimage_io
            libm
            libxml2
    )
endif ()
#################### libdynamic_depth ####################

#################### libdynamic_depth_ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/app_info.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/camera.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/cameras.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/container.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/depth_jpeg.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/depth_map.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/device.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/dynamic_depth.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/earth_pose.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/image.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/imaging_model.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/item.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/light_estimate.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/plane.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/planes.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/point_cloud.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/pose.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/profile.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/profiles.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/dynamic_depth/vendor_info.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/ascii_ctype.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/case.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/escaping.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/strings/numbers.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/base64.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/file.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/jpeg_io.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/md5.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_data.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_parser.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xmp_writer.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/const.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/deserializer_impl.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/search.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/serializer_impl.cc
        ${ANDROID_ROOT}/external/dynamic_depth/internal/xmpmeta/xml/utils.cc
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
    add_library(libdynamic_depth_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libdynamic_depth_ndk
            PUBLIC
            ${ANDROID_ROOT}/external/dynamic_depth/includes
            ${ANDROID_ROOT}/external/dynamic_depth/internal
            ${ANDROID_ROOT}/external/dynamic_depth
    )

    # Add c/c++ flags
    target_compile_options(libdynamic_depth_ndk PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libdynamic_depth_ndk
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libimage_io_ndk
            liblog
            libm
            libxml2_ndk
    )
endif ()
#################### libdynamic_depth_ndk ####################
