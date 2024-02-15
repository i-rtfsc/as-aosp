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

project(libjni_legacymosaic)

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

#################### libjni_legacymosaic ####################
set(FILES
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos_jni.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/mosaic_renderer_jni.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/trsMatrix.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/AlignFeatures.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/Blend.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/Delaunay.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/ImageUtils.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/Mosaic.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic/Pyramid.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic_renderer/Renderer.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic_renderer/WarpRenderer.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic_renderer/SurfaceTextureRenderer.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic_renderer/YVURenderer.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic_renderer/FrameBuffer.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_feature_detection.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_feature_matching.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_framestitching.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_image_homography.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_rob_image_homography.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_utilities.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_utilities_camera.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_utilities_indexing.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_utilities_linalg.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm/db_utilities_poly.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/src/dbreg/dbreg.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/src/dbreg/dbstabsmooth.cpp
        ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/src/dbreg/vp_motionmodel.c
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
    add_library(libjni_legacymosaic SHARED ${SOURCE_FILES})

    target_include_directories(libjni_legacymosaic
            PUBLIC
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/db_vlvm
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/src
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_stab/src/dbreg
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni/feature_mos/src/mosaic
            ${ANDROID_ROOT}/packages/apps/LegacyCamera/jni
    )

    # Add c/c++ flags
    target_compile_options(libjni_legacymosaic PRIVATE
            -DNDEBUG
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
            -Wno-unused-function
    )

    target_link_libraries(libjni_legacymosaic
            libGLESv2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativehelper
    )
endif ()
#################### libjni_legacymosaic ####################
