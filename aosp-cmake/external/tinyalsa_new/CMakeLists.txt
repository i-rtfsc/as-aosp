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

project(tinyplay2)

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

#################### libtinyalsav2 ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa_new/src/mixer.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/mixer_hw.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/mixer_plugin.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/pcm.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/pcm_hw.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/pcm_plugin.c
        ${ANDROID_ROOT}/external/tinyalsa_new/src/snd_card_plugin.c
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
    add_library(libtinyalsav2 SHARED ${SOURCE_FILES})

    target_include_directories(libtinyalsav2
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa_new/include
            ${ANDROID_ROOT}/external/tinyalsa_new
    )

    # Add c/c++ flags
    target_compile_options(libtinyalsav2 PRIVATE
            -Werror
    )

    target_link_libraries(libtinyalsav2
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libdl
    )
endif ()
#################### libtinyalsav2 ####################

#################### tinycap2 ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa_new/utils/tinycap.c
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
    add_executable(tinycap2 ${SOURCE_FILES})

    target_include_directories(tinycap2
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa_new
            ${ANDROID_ROOT}/external/tinyalsa_new/include
    )

    # Add c/c++ flags
    target_compile_options(tinycap2 PRIVATE
            -Werror
    )

    target_link_libraries(tinycap2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libdl
            libm
            libtinyalsav2
    )
endif ()
#################### tinycap2 ####################

#################### tinymix2 ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa_new/utils/tinymix.c
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
    add_executable(tinymix2 ${SOURCE_FILES})

    target_include_directories(tinymix2
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa_new
            ${ANDROID_ROOT}/external/tinyalsa_new/include
    )

    # Add c/c++ flags
    target_compile_options(tinymix2 PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(tinymix2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libdl
            libm
            libtinyalsav2
    )
endif ()
#################### tinymix2 ####################

#################### tinypcminfo2 ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa_new/utils/tinypcminfo.c
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
    add_executable(tinypcminfo2 ${SOURCE_FILES})

    target_include_directories(tinypcminfo2
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa_new
            ${ANDROID_ROOT}/external/tinyalsa_new/include
    )

    # Add c/c++ flags
    target_compile_options(tinypcminfo2 PRIVATE
            -Werror
    )

    target_link_libraries(tinypcminfo2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libdl
            libm
            libtinyalsav2
    )
endif ()
#################### tinypcminfo2 ####################

#################### tinyplay2 ####################
set(FILES
        ${ANDROID_ROOT}/external/tinyalsa_new/utils/tinyplay.c
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
    add_executable(tinyplay2 ${SOURCE_FILES})

    target_include_directories(tinyplay2
            PUBLIC
            ${ANDROID_ROOT}/external/tinyalsa_new
            ${ANDROID_ROOT}/external/tinyalsa_new/include
    )

    # Add c/c++ flags
    target_compile_options(tinyplay2 PRIVATE
            -Werror
    )

    target_link_libraries(tinyplay2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libdl
            libm
            libtinyalsav2
    )
endif ()
#################### tinyplay2 ####################
