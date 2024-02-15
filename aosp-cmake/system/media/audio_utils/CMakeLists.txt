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

project(libaudioutils_ext)

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

#################### libaudioutils_ext ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/Balance.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/channels.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/ErrorLog.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo_index.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo_writer_T.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/format.c
        ${ANDROID_ROOT}/system/media/audio_utils/limiter.c
        ${ANDROID_ROOT}/system/media/audio_utils/Metadata.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/minifloat.c
        ${ANDROID_ROOT}/system/media/audio_utils/mono_blend.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/power.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/PowerLog.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/primitives.c
        ${ANDROID_ROOT}/system/media/audio_utils/roundup.c
        ${ANDROID_ROOT}/system/media/audio_utils/sample.c
        ${ANDROID_ROOT}/system/media/audio_utils/resampler.c
        ${ANDROID_ROOT}/system/media/audio_utils/echo_reference.c
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
    add_library(libaudioutils_ext SHARED ${SOURCE_FILES})

    target_include_directories(libaudioutils_ext
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudioutils_ext PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioutils_ext
            libaudio_system_headers
            libaudioutils_fixedfft
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libspeexresampler
            libutils_headers
            libc
            libdl
            libm
    )
endif ()
#################### libaudioutils_ext ####################

#################### libaudiospdif ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/spdif/BitFieldParser.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/spdif/FrameScanner.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/spdif/AC3FrameScanner.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/spdif/DTSFrameScanner.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/spdif/SPDIFEncoder.cpp
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
    add_library(libaudiospdif SHARED ${SOURCE_FILES})

    target_include_directories(libaudiospdif
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudiospdif PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudiospdif
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libaudiospdif ####################

#################### libaudioutils_fixedfft ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/fixedfft.cpp
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
    add_library(libaudioutils_fixedfft SHARED ${SOURCE_FILES})

    target_include_directories(libaudioutils_fixedfft
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudioutils_fixedfft PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioutils_fixedfft
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libaudioutils_fixedfft ####################

#################### libsndfile ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/tinysndfile.c
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
    add_library(libsndfile SHARED ${SOURCE_FILES})

    target_include_directories(libsndfile
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libsndfile PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libsndfile
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libsndfile ####################

#################### libaudioext ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/AudioDumpUtil.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/readproc.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/AudioDeadTrackCheck.cpp
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
    add_library(libaudioext SHARED ${SOURCE_FILES})

    target_include_directories(libaudioext
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudioext PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioext
            libaudioutils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libaudioext ####################

#################### libaudioutils ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/Balance.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/channels.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/ErrorLog.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo_index.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo_writer_T.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/format.c
        ${ANDROID_ROOT}/system/media/audio_utils/limiter.c
        ${ANDROID_ROOT}/system/media/audio_utils/MelAggregator.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/MelProcessor.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/Metadata.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/minifloat.c
        ${ANDROID_ROOT}/system/media/audio_utils/mono_blend.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/power.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/PowerLog.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/primitives.c
        ${ANDROID_ROOT}/system/media/audio_utils/roundup.c
        ${ANDROID_ROOT}/system/media/audio_utils/sample.c
        ${ANDROID_ROOT}/system/media/audio_utils/hal_smoothness.c
        ${ANDROID_ROOT}/system/media/audio_utils/resampler.c
        ${ANDROID_ROOT}/system/media/audio_utils/echo_reference.c
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
    add_library(libaudioutils SHARED ${SOURCE_FILES})

    target_include_directories(libaudioutils
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudioutils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioutils
            libaudio_system_headers
            libaudioutils_fastmath
            libaudioutils_fixedfft
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libspeexresampler
            libutils
            libutils_headers
    )
endif ()
#################### libaudioutils ####################

#################### libfifo ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/fifo.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/fifo_index.cpp
        ${ANDROID_ROOT}/system/media/audio_utils/primitives.c
        ${ANDROID_ROOT}/system/media/audio_utils/roundup.c
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
    add_library(libfifo SHARED ${SOURCE_FILES})

    target_include_directories(libfifo
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libfifo PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libfifo
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libfifo ####################

#################### libaudioutils_fastmath ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_utils/ChannelMix.cpp
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
    add_library(libaudioutils_fastmath SHARED ${SOURCE_FILES})

    target_include_directories(libaudioutils_fastmath
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_utils/include
            ${ANDROID_ROOT}/system/media/audio_utils
    )

    # Add c/c++ flags
    target_compile_options(libaudioutils_fastmath PRIVATE
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libaudioutils_fastmath
            libaudio_system_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
            libutils_headers
    )
endif ()
#################### libaudioutils_fastmath ####################
