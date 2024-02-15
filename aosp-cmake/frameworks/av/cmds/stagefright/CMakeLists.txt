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

project(audioloop)

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

#################### recordvideo ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/AudioPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/recordvideo.cpp
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
    add_executable(recordvideo ${SOURCE_FILES})

    target_include_directories(recordvideo
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(recordvideo PRIVATE
            -Werror
    )

    target_link_libraries(recordvideo
            framework-permission-aidl-cpp
            libaudioclient
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmedia
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libutils
    )
endif ()
#################### recordvideo ####################

#################### muxer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/muxer.cpp
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
    add_executable(muxer ${SOURCE_FILES})

    target_include_directories(muxer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(muxer PRIVATE
            -Werror
    )

    target_link_libraries(muxer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libutils
    )
endif ()
#################### muxer ####################

#################### record ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/AudioPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/SineSource.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/record.cpp
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
    add_executable(record ${SOURCE_FILES})

    target_include_directories(record
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(record PRIVATE
            -Werror
    )

    target_link_libraries(record
            camera_headers
            framework-permission-aidl-cpp
            libaudioclient
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdatasource
            libdl
            liblog
            libm
            libmedia
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libutils
    )
endif ()
#################### record ####################

#################### stream ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/stream.cpp
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
    add_executable(stream ${SOURCE_FILES})

    target_include_directories(stream
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(stream PRIVATE
            -Werror
    )

    target_link_libraries(stream
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdatasource
            libdl
            libgui
            liblog
            libm
            libmedia
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libui
            libutils
    )
endif ()
#################### stream ####################

#################### codec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/codec.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/SimplePlayer.cpp
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
    add_executable(codec ${SOURCE_FILES})

    target_include_directories(codec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(codec PRIVATE
            -Werror
    )

    target_link_libraries(codec
            libaudioclient
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgui
            liblog
            libm
            libmedia
            libmedia_omx
            libmediadrm_headers
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libui
            libutils
    )
endif ()
#################### codec ####################

#################### stagefright ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/AudioPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/stagefright.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/jpeg.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/SineSource.cpp
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
    add_executable(stagefright ${SOURCE_FILES})

    target_include_directories(stagefright
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(stagefright PRIVATE
            -Werror
    )

    target_link_libraries(stagefright
            android.hardware.media.omx-1.0
            framework-permission-aidl-cpp
            libaudioclient
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            libjpeg
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libui
            libutils
    )
endif ()
#################### stagefright ####################

#################### audioloop ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/AudioPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/SineSource.cpp
        ${ANDROID_ROOT}/frameworks/av/cmds/stagefright/audioloop.cpp
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
    add_executable(audioloop ${SOURCE_FILES})

    target_include_directories(audioloop
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/cmds/stagefright
    )

    # Add c/c++ flags
    target_compile_options(audioloop PRIVATE
            -Werror
    )

    target_link_libraries(audioloop
            framework-permission-aidl-cpp
            libaudioclient
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmedia
            libmediametrics_headers
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libutils
    )
endif ()
#################### audioloop ####################
