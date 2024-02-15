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

project(ndk_extractor_fuzzer)

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

#################### ndk_crypto_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_crypto_fuzzer.cpp
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
    add_executable(ndk_crypto_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_crypto_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_crypto_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_crypto_fuzzer
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_crypto_fuzzer ####################

#################### ndk_image_reader_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_image_reader_fuzzer.cpp
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
    add_executable(ndk_image_reader_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_image_reader_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_image_reader_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_image_reader_fuzzer
            android.hardware.graphics.bufferqueue-1.0
            android.hidl.token-1.0-utils
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_image_reader_fuzzer ####################

#################### ndk_drm_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_drm_fuzzer.cpp
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
    add_executable(ndk_drm_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_drm_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_drm_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_drm_fuzzer
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_drm_fuzzer ####################

#################### ndk_mediamuxer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_mediamuxer_fuzzer.cpp
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
    add_executable(ndk_mediamuxer_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_mediamuxer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_mediamuxer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_mediamuxer_fuzzer
            libandroid_runtime_lazy
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_mediamuxer_fuzzer ####################

#################### ndk_sync_codec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_sync_codec_fuzzer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/NdkMediaCodecFuzzerBase.cpp
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
    add_executable(ndk_sync_codec_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_sync_codec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_sync_codec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_sync_codec_fuzzer
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libnativewindow_headers
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_sync_codec_fuzzer ####################

#################### ndk_mediaformat_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_mediaformat_fuzzer.cpp
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
    add_executable(ndk_mediaformat_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_mediaformat_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_mediaformat_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_mediaformat_fuzzer
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_mediaformat_fuzzer ####################

#################### ndk_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer/ndk_extractor_fuzzer.cpp
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
    add_executable(ndk_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(ndk_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(ndk_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ndk_extractor_fuzzer
            libandroid_runtime_lazy
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediandk
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### ndk_extractor_fuzzer ####################
