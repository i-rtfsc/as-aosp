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

project(mediatuner)

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

#################### tv_tuner_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerLnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerLnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/android/media/tv/tuner/ITunerTimeFilter.cpp
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
    add_library(tv_tuner_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(tv_tuner_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(tv_tuner_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(tv_tuner_aidl_interface-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.tv.tuner-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tv_tuner_aidl_interface-ndk ####################

#################### mediatuner ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/tuner/main_tunerservice.cpp
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
    add_executable(mediatuner ${SOURCE_FILES})

    target_include_directories(mediatuner
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/tuner
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediatuner PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(mediatuner
            android.hardware.tv.tuner-V2-ndk
            android.hardware.tv.tuner-1.0
            android.hardware.tv.tuner-1.1
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libfmq
            libhidlbase
            liblog
            libm
            libtunerservice
            libutils
            tv_tuner_aidl_interface-ndk
            tv_tuner_resource_manager_aidl_interface-ndk
    )
endif ()
#################### mediatuner ####################

#################### tv_tuner_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerLnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerLnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/android/media/tv/tuner/ITunerTimeFilter.cpp
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
    add_library(tv_tuner_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(tv_tuner_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(tv_tuner_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(tv_tuner_aidl_interface-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.hardware.tv.tuner-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### tv_tuner_aidl_interface-ndk_platform ####################

#################### libtunerservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerDemux.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerDescrambler.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerDvr.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerFilter.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerFrontend.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerHelper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerLnb.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/TunerTimeFilter.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlDemux.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlDescrambler.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlDvr.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlFilter.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlFrontend.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlLnb.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/tuner/hidl/TunerHidlTimeFilter.cpp
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
    add_library(libtunerservice SHARED ${SOURCE_FILES})

    target_include_directories(libtunerservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/tuner
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libtunerservice PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libtunerservice
            android.hardware.common.fmq-V1-ndk
            android.hardware.tv.tuner-V2-ndk
            android.hardware.tv.tuner-1.0
            android.hardware.tv.tuner-1.1
            libaidlcommonsupport
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhidlbase
            liblog
            libm
            libmedia
            libutils
            packagemanager_aidl-cpp
            tv_tuner_aidl_interface-ndk
            tv_tuner_resource_manager_aidl_interface-ndk
    )
endif ()
#################### libtunerservice ####################
