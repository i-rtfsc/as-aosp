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

project(libharfbuzz_subset)

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

#################### libharfbuzz_subset ####################
set(FILES
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-number.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-cff1-table.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-cff2-table.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-static.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-cff-common.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-cff1.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-cff2.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-input.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-plan.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset-repacker.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-subset.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/graph/gsubgpos-context.cc
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
    add_library(libharfbuzz_subset SHARED ${SOURCE_FILES})

    target_include_directories(libharfbuzz_subset
            PUBLIC
            ${ANDROID_ROOT}/external/harfbuzz_ng/src
            ${ANDROID_ROOT}/external/harfbuzz_ng
    )

    # Add c/c++ flags
    target_compile_options(libharfbuzz_subset PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libharfbuzz_subset
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libharfbuzz_ng
            libm
    )
endif ()
#################### libharfbuzz_subset ####################

#################### libharfbuzz_ng ####################
set(FILES
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-aat-layout.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-aat-map.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-blob.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-buffer-serialize.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-buffer-verify.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-buffer.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-common.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-draw.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-face.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-fallback-shape.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-font.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-map.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-number.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-cff1-table.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-cff2-table.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-color.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-face.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-font.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-layout.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-map.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-math.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-meta.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-metrics.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-name.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shape-fallback.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shape-normalize.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shape.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-arabic.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-default.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-hangul.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-hebrew.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-indic-table.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-indic.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-khmer.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-myanmar.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-syllabic.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-thai.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-use.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-shaper-vowel-constraints.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-tag.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ot-var.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-set.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-shape-plan.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-shape.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-shaper.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-static.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-style.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-ucd.cc
        ${ANDROID_ROOT}/external/harfbuzz_ng/src/hb-unicode.cc
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
    add_library(libharfbuzz_ng SHARED ${SOURCE_FILES})

    target_include_directories(libharfbuzz_ng
            PUBLIC
            ${ANDROID_ROOT}/external/harfbuzz_ng/src
            ${ANDROID_ROOT}/external/harfbuzz_ng
    )

    # Add c/c++ flags
    target_compile_options(libharfbuzz_ng PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libharfbuzz_ng
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libicu
            liblog
            libm
            libutils
            libwinpthread
    )
endif ()
#################### libharfbuzz_ng ####################
