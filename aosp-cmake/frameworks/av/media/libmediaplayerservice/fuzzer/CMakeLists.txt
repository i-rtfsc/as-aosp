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

project(mediaplayer_fuzzer)

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

#################### mediarecorder_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer/mediarecorder_fuzzer.cpp
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
    add_executable(mediarecorder_fuzzer ${SOURCE_FILES})

    target_include_directories(mediarecorder_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(mediarecorder_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mediarecorder_fuzzer
            android.hardware.media.omx-1.0
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            libactivitymanager_aidl
            libandroid_net
            libaudioclient
            libaudioclient_aidl_conversion
            libaudioflinger
            libaudioflinger_headers
            libaudiohal_headers
            libaudioprocessing
            libaudioprocessing_base
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2_client
            libcrypto
            libcutils
            libdatasource
            libdl
            libdrmframework
            libfakeservicemanager
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_omx
            libmediadrm
            libmediametrics
            libmediametricsservice
            libmediaplayerservice
            libmediautils
            libmemunreachable
            libnbaio
            libnblog
            libnetd_client
            libplayerservice_datasource
            libpowermanager
            libresourcemanagerservice
            libstagefright
            libstagefright_foundation
            libstagefright_httplive
            libstagefright_nuplayer
            libstagefright_rtsp
            libstagefright_timedtext
            libutils
            libvibrator
            media_permission-aidl-cpp
            mediametricsservice-aidl-cpp
            packagemanager_aidl-cpp
    )
endif ()
#################### mediarecorder_fuzzer ####################

#################### metadataretriever_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer/metadataretriever_fuzzer.cpp
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
    add_executable(metadataretriever_fuzzer ${SOURCE_FILES})

    target_include_directories(metadataretriever_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(metadataretriever_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(metadataretriever_fuzzer
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libdrmframework
            liblog
            libm
            libmedia
            libmediaplayerservice
            libplayerservice_datasource
            libstagefright
            libstagefright_foundation
            libutils
    )
endif ()
#################### metadataretriever_fuzzer ####################

#################### mediaplayer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer/mediaplayer_fuzzer.cpp
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
    add_executable(mediaplayer_fuzzer ${SOURCE_FILES})

    target_include_directories(mediaplayer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(mediaplayer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mediaplayer_fuzzer
            android.hardware.media.c2-1.0
            android.hardware.media.omx-1.0
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            libactivitymanager_aidl
            libandroid_net
            libaudioclient
            libaudioclient_aidl_conversion
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2_client
            libcrypto
            libcutils
            libdatasource
            libdl
            libdrmframework
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_omx
            libmediadrm
            libmediametrics
            libmediaplayerservice
            libmediautils
            libmemunreachable
            libnetd_client
            libplayerservice_datasource
            libpowermanager
            libstagefright
            libstagefright_foundation
            libstagefright_httplive
            libstagefright_nuplayer
            libstagefright_rtsp
            libstagefright_timedtext
            libutils
    )
endif ()
#################### mediaplayer_fuzzer ####################
