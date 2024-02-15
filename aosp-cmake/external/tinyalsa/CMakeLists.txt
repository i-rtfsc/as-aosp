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

project(tinycap)

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

#################### tinymix ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/tinymix.c
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
    add_executable(tinymix ${SOURCE_FILES})

    target_include_directories(tinymix
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa
            ${ANDROID_ROOT}/external/tinyalsa/include
    )

    # Add c/c++ flags
    target_compile_options(tinymix PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(tinymix
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtinyalsa
    )
endif ()
#################### tinymix ####################

#################### tinyplay ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/tinyplay.c
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
    add_executable(tinyplay ${SOURCE_FILES})

    target_include_directories(tinyplay
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa
            ${ANDROID_ROOT}/external/tinyalsa/include
    )

    # Add c/c++ flags
    target_compile_options(tinyplay PRIVATE
            -Werror
    )

    target_link_libraries(tinyplay
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtinyalsa
    )
endif ()
#################### tinyplay ####################

#################### tinyhostless ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/tinyhostless.c
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
    add_executable(tinyhostless ${SOURCE_FILES})

    target_include_directories(tinyhostless
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa
            ${ANDROID_ROOT}/external/tinyalsa/include
    )

    # Add c/c++ flags
    target_compile_options(tinyhostless PRIVATE
            -Werror
    )

    target_link_libraries(tinyhostless
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtinyalsa
    )
endif ()
#################### tinyhostless ####################

#################### tinypcminfo ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/tinypcminfo.c
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
    add_executable(tinypcminfo ${SOURCE_FILES})

    target_include_directories(tinypcminfo
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa
            ${ANDROID_ROOT}/external/tinyalsa/include
    )

    # Add c/c++ flags
    target_compile_options(tinypcminfo PRIVATE
            -Werror
    )

    target_link_libraries(tinypcminfo
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtinyalsa
    )
endif ()
#################### tinypcminfo ####################

#################### tinycap ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/tinycap.c
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
    add_executable(tinycap ${SOURCE_FILES})

    target_include_directories(tinycap
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa
            ${ANDROID_ROOT}/external/tinyalsa/include
    )

    # Add c/c++ flags
    target_compile_options(tinycap PRIVATE
            -Werror
    )

    target_link_libraries(tinycap
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtinyalsa
    )
endif ()
#################### tinycap ####################

#################### libtinyalsa ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa/mixer.c
        ${ANDROID_ROOT}/external/tinyalsa/mixer_hw.c
        ${ANDROID_ROOT}/external/tinyalsa/mixer_plugin.c
        ${ANDROID_ROOT}/external/tinyalsa/pcm.c
        ${ANDROID_ROOT}/external/tinyalsa/pcm_hw.c
        ${ANDROID_ROOT}/external/tinyalsa/pcm_plugin.c
        ${ANDROID_ROOT}/external/tinyalsa/snd_utils.c
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
    add_library(libtinyalsa SHARED ${SOURCE_FILES})

    target_include_directories(libtinyalsa
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa/include
            ${ANDROID_ROOT}/external/tinyalsa
    )

    # Add c/c++ flags
    target_compile_options(libtinyalsa PRIVATE
            -Werror
    )

    target_link_libraries(libtinyalsa
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
    )
endif ()
#################### libtinyalsa ####################
