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

project(libFFTEm)

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

#################### libFFTEm ####################
set(FILES
        ${ANDROID_ROOT}/external/neven/FaceDetector_jni.cpp
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_APIEm/DCR.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_APIEm/BFFaceFinder.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_APIEm/FaceFinder.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_APIEm/FaceFinderRef.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_APIEm/Functions.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/APh.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/APhArr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Complex.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/ComplexArr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Context.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/DynMemManager.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Functions.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Int16Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Int32Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Int8Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Math.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/MemSeg.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/MemTbl.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Memory.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/Phase.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/String.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/UInt16Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/UInt32Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BasicEm/UInt8Arr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/BitParam.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/Feature.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/Functions.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/I04Dns2x2Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/I04Dns2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/I04Tld2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L01Dns2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L01Tld1x1Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L01Tld2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L04Dns2x2Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L04Dns2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L04Dns3x3Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L04Tld2x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L06Dns3x3Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L06Dns4x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/L06DnsNx4x4Ftr.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/LocalScanDetector.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/LocalScanner.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/ScanDetector.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/Scanner.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_BitFeatureEm/Sequence.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/APhImage.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/ComplexImage.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/Flt16Image.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/Functions.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/HistoEq.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/UInt16ByteImage.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/UInt16BytePyrImage.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/UInt8Image.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/UInt32Image.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_ImageEm/UInt8PyramidalImage.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Alt.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Cluster2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Cluster3D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/CompactAlt.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/CompactMat.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Alt2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Alt3D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Mat2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Mat3D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Vec.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Vec2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Flt16Vec3D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Functions.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/IdCluster2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Int16Mat2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Int16Rect.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Int16Vec2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Int16Vec3D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Int32Mat.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/MapSequence.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Mat.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Normalizer.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/RBFMap2D.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/SubVecMap.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/Uint32Rect.c
        ${ANDROID_ROOT}/external/neven/Embedded/common/src/b_TensorEm/VectorMap.c
        ${ANDROID_ROOT}/external/neven/FaceRecEm/common/src/b_FDSDK/DCR.c
        ${ANDROID_ROOT}/external/neven/FaceRecEm/common/src/b_FDSDK/FaceFinder.c
        ${ANDROID_ROOT}/external/neven/FaceRecEm/common/src/b_FDSDK/SDK.c
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
    add_library(libFFTEm SHARED ${SOURCE_FILES})

    target_include_directories(libFFTEm
            PUBLIC
            ${ANDROID_ROOT}/external/neven/FaceRecEm/common/src/b_FDSDK
            ${ANDROID_ROOT}/external/neven/FaceRecEm/common/src
            ${ANDROID_ROOT}/external/neven/Embedded/common/conf
            ${ANDROID_ROOT}/external/neven/Embedded/common/src
            ${ANDROID_ROOT}/external/neven
    )

    # Add c/c++ flags
    target_compile_options(libFFTEm PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libFFTEm
            jni_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libjnigraphics
            liblog
            libm
            libnativehelper
            libutils
    )
endif ()
#################### libFFTEm ####################
