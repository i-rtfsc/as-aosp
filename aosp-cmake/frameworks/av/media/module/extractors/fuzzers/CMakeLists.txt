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

project(mpeg2ts_extractor_fuzzer)

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

#################### libextractorfuzzerbase ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/ExtractorFuzzerBase.cpp
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
    add_library(libextractorfuzzerbase SHARED ${SOURCE_FILES})

    target_include_directories(libextractorfuzzerbase
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libextractorfuzzerbase PRIVATE
            -Werror
    )

    target_link_libraries(libextractorfuzzerbase
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libutils
    )
endif ()
#################### libextractorfuzzerbase ####################

#################### aac_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/aac_extractor_fuzzer.cpp
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
    add_executable(aac_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(aac_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(aac_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(aac_extractor_fuzzer
            libaacextractor
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_metadatautils
            libutils
    )
endif ()
#################### aac_extractor_fuzzer ####################

#################### flac_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/flac_extractor_fuzzer.cpp
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
    add_executable(flac_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(flac_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(flac_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(flac_extractor_fuzzer
            libFLAC
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            libflacextractor
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_metadatautils
            libutils
    )
endif ()
#################### flac_extractor_fuzzer ####################

#################### ogg_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/ogg_extractor_fuzzer.cpp
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
    add_executable(ogg_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(ogg_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(ogg_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ogg_extractor_fuzzer
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            liboggextractor
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_metadatautils
            libutils
            libvorbisidec
    )
endif ()
#################### ogg_extractor_fuzzer ####################

#################### mpeg2ps_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/mpeg2_extractor_fuzzer.cpp
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
    add_executable(mpeg2ps_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(mpeg2ps_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mpeg2ps_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mpeg2ps_extractor_fuzzer
            android.hardware.cas.native-1.0
            android.hardware.cas-1.0
            android.hidl.allocator-1.0
            android.hidl.token-1.0-utils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libextractorfuzzerbase
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmedia_ndkformatpriv
            libmediandk_format
            libmpeg2extractor
            libstagefright_esds
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_foundation_without_imemory
            libstagefright_mpeg2extractor
            libstagefright_mpeg2support
            libutils
    )
endif ()
#################### mpeg2ps_extractor_fuzzer ####################

#################### midi_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/midi_extractor_fuzzer.cpp
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
    add_executable(midi_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(midi_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(midi_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(midi_extractor_fuzzer
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_midiiowrapper
            libmedia_ndkformatpriv
            libmediandk_format
            libmidiextractor
            libsonivox
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libutils
            libwatchdog
    )
endif ()
#################### midi_extractor_fuzzer ####################

#################### mkv_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/mkv_extractor_fuzzer.cpp
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
    add_executable(mkv_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(mkv_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mkv_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mkv_extractor_fuzzer
            libFLAC
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libmkvextractor
            libstagefright_flacdec
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_metadatautils
            libutils
            libwebm_mkvparser
    )
endif ()
#################### mkv_extractor_fuzzer ####################

#################### wav_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/wav_extractor_fuzzer.cpp
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
    add_executable(wav_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(wav_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(wav_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(wav_extractor_fuzzer
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            libfifo
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libutils
            libwavextractor
    )
endif ()
#################### wav_extractor_fuzzer ####################

#################### mp3_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/mp3_extractor_fuzzer.cpp
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
    add_executable(mp3_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(mp3_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mp3_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mp3_extractor_fuzzer
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            libfifo
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libmp3extractor
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_id3
            libutils
    )
endif ()
#################### mp3_extractor_fuzzer ####################

#################### amr_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/amr_extractor_fuzzer.cpp
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
    add_executable(amr_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(amr_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(amr_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(amr_extractor_fuzzer
            libamrextractor
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libutils
    )
endif ()
#################### amr_extractor_fuzzer ####################

#################### mp4_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/mp4_extractor_fuzzer.cpp
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
    add_executable(mp4_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(mp4_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mp4_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mp4_extractor_fuzzer
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libextractorfuzzerbase
            liblog
            libm
            libmedia_ndkformatpriv
            libmediandk_format
            libmp4extractor
            libstagefright_esds
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_id3
            libutils
    )
endif ()
#################### mp4_extractor_fuzzer ####################

#################### mpeg2ts_extractor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/mpeg2_extractor_fuzzer.cpp
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
    add_executable(mpeg2ts_extractor_fuzzer ${SOURCE_FILES})

    target_include_directories(mpeg2ts_extractor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/module/extractors/fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mpeg2ts_extractor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mpeg2ts_extractor_fuzzer
            android.hardware.cas.native-1.0
            android.hardware.cas-1.0
            android.hidl.allocator-1.0
            android.hidl.token-1.0-utils
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libextractorfuzzerbase
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_helper
            libmedia_ndkformatpriv
            libmediandk_format
            libmpeg2extractor
            libstagefright_esds
            libstagefright_foundation
            libstagefright_foundation_colorutils_ndk
            libstagefright_foundation_without_imemory
            libstagefright_mpeg2extractor
            libstagefright_mpeg2support
            libutils
    )
endif ()
#################### mpeg2ts_extractor_fuzzer ####################
