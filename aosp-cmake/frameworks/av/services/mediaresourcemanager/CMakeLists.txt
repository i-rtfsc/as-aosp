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

project(resourcemanager_aidl_interface-cpp-analyzer)

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

#################### resourceobserver_aidl_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/IResourceObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/IResourceObserverService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/MediaObservableEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/MediaObservableFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/MediaObservableParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/android/media/MediaObservableType.cpp
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
    add_library(resourceobserver_aidl_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(resourceobserver_aidl_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourceobserver_aidl_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourceobserver_aidl_interface-V1-ndk_platform
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
#################### resourceobserver_aidl_interface-V1-ndk_platform ####################

#################### libresourcemanagerservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/ResourceManagerMetrics.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/ResourceManagerService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/ResourceObserverService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/ServiceLog.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/UidObserver.cpp
        ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager/IMediaResourceMonitor.cpp
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
    add_library(libresourcemanagerservice SHARED ${SOURCE_FILES})

    target_include_directories(libresourcemanagerservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libresourcemanagerservice PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libresourcemanagerservice
            libactivitymanager_aidl
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libmedia
            libmediautils
            libplatformprotos
            libprotobuf-cpp-lite
            libstats_media_metrics
            libstatspull
            libstatssocket
            libutils
            resourceobserver_aidl_interface-V1-ndk
    )
endif ()
#################### libresourcemanagerservice ####################

#################### resourcemanager_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/IResourceManagerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/IResourceManagerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/MediaResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/MediaResourceSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/MediaResourceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/MediaResourcePolicyParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/ClientInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/android/media/ClientConfigParcel.cpp
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
    add_executable(resourcemanager_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(resourcemanager_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/mediaresourcemanager
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourcemanager_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourcemanager_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            resourcemanager_aidl_interface-cpp
    )
endif ()
#################### resourcemanager_aidl_interface-cpp-analyzer ####################

#################### resourceobserver_aidl_interface-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/IResourceObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/IResourceObserverService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/MediaObservableEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/MediaObservableFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/MediaObservableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/android/media/MediaObservableParcel.cpp
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
    add_library(resourceobserver_aidl_interface-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(resourceobserver_aidl_interface-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourceobserver_aidl_interface-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourceobserver_aidl_interface-V2-ndk_platform
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
#################### resourceobserver_aidl_interface-V2-ndk_platform ####################

#################### resourcemanager_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/IResourceManagerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/IResourceManagerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/MediaResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/MediaResourceSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/MediaResourceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/MediaResourcePolicyParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/ClientInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/android/media/ClientConfigParcel.cpp
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
    add_library(resourcemanager_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(resourcemanager_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourcemanager_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourcemanager_aidl_interface-cpp
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
#################### resourcemanager_aidl_interface-cpp ####################

#################### resourceobserver_aidl_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/IResourceObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/IResourceObserverService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/MediaObservableEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/MediaObservableFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/MediaObservableParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/android/media/MediaObservableType.cpp
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
    add_library(resourceobserver_aidl_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(resourceobserver_aidl_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourceobserver_aidl_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourceobserver_aidl_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### resourceobserver_aidl_interface-V1-ndk ####################

#################### resourcemanager_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/IResourceManagerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/IResourceManagerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/MediaResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/MediaResourceSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/MediaResourceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/MediaResourcePolicyParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/ClientInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/android/media/ClientConfigParcel.cpp
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
    add_library(resourcemanager_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(resourcemanager_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourcemanager_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourcemanager_aidl_interface-ndk_platform
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
#################### resourcemanager_aidl_interface-ndk_platform ####################

#################### resourceobserver_aidl_interface-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/IResourceObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/IResourceObserverService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/MediaObservableEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/MediaObservableFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/MediaObservableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/android/media/MediaObservableParcel.cpp
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
    add_library(resourceobserver_aidl_interface-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(resourceobserver_aidl_interface-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourceobserver_aidl_interface-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourceobserver_aidl_interface-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourceobserver_aidl_interface-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### resourceobserver_aidl_interface-V2-ndk ####################

#################### resourcemanager_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/IResourceManagerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/IResourceManagerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/MediaResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/MediaResourceSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/MediaResourceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/MediaResourcePolicyParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/ClientInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/android/media/ClientConfigParcel.cpp
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
    add_library(resourcemanager_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(resourcemanager_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(resourcemanager_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(resourcemanager_aidl_interface-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### resourcemanager_aidl_interface-ndk ####################
