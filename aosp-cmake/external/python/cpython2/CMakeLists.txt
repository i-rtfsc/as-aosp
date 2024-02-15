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

project(py2-launcher-autorun)

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

#################### py2-c-module-_sha ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/shamodule.c
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
    add_library(py2-c-module-_sha SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_sha
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_sha PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_sha
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_sha ####################

#################### py2-interp-parser ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Parser/acceler.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/bitset.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/firstsets.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/grammar.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/grammar1.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/listnode.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/metagrammar.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/myreadline.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/node.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/parser.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/parsetok.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/pgen.c
        ${ANDROID_ROOT}/external/python/cpython2/Parser/tokenizer.c
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
    add_library(py2-interp-parser SHARED ${SOURCE_FILES})

    target_include_directories(py2-interp-parser
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-interp-parser PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-interp-parser
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-interp-parser ####################

#################### py2-c-module-_struct ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_struct.c
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
    add_library(py2-c-module-_struct SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_struct
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_struct PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_struct
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_struct ####################

#################### py2-c-module-binascii ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/binascii.c
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
    add_library(py2-c-module-binascii SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-binascii
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-binascii PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-binascii
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-binascii ####################

#################### py2-c-module-pyexpat ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/pyexpat.c
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
    add_library(py2-c-module-pyexpat SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-pyexpat
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-pyexpat PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-pyexpat
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libexpat
            libm
    )
endif ()
#################### py2-c-module-pyexpat ####################

#################### py2-c-module-datetime ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/datetimemodule.c
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
    add_library(py2-c-module-datetime SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-datetime
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-datetime PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-datetime
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-datetime ####################

#################### py2-c-module-audioop ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/audioop.c
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
    add_library(py2-c-module-audioop SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-audioop
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-audioop PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-audioop
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-audioop ####################

#################### py2-launcher ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Launcher/launcher_main.cpp
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
    add_executable(py2-launcher ${SOURCE_FILES})

    target_include_directories(py2-launcher
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-launcher PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-launcher
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libexpat
            liblog
            libm
            libsqlite
            libssl
            libz
            py2-launcher-lib
    )
endif ()
#################### py2-launcher ####################

#################### py2-c-module-_codecs_hk ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_hk.c
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
    add_library(py2-c-module-_codecs_hk SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_hk
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_hk PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_hk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_hk ####################

#################### py2-c-module-syslog ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/syslogmodule.c
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
    add_library(py2-c-module-syslog SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-syslog
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-syslog PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-syslog
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-syslog ####################

#################### py2-c-module-_ctypes ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/_ctypes.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/callbacks.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/callproc.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/stgdict.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/cfield.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/src/prep_cif.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/src/closures.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/src/aarch64/ffi.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/src/aarch64/sysv.S
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
    add_library(py2-c-module-_ctypes SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_ctypes
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/libffi
            ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/src/aarch64
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2/Modules/_ctypes/libffi/include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_ctypes PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_ctypes
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_ctypes ####################

#################### py2-c-module-_hotshot ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_hotshot.c
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
    add_library(py2-c-module-_hotshot SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_hotshot
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_hotshot PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_hotshot
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_hotshot ####################

#################### py2-c-module-operator ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/operator.c
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
    add_library(py2-c-module-operator SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-operator
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-operator PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-operator
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-operator ####################

#################### py2-c-module-strop ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/stropmodule.c
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
    add_library(py2-c-module-strop SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-strop
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-strop PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-strop
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-strop ####################

#################### py2-c-module-_functools ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_functoolsmodule.c
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
    add_library(py2-c-module-_functools SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_functools
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_functools PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_functools
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_functools ####################

#################### py2-c-module-cStringIO ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cStringIO.c
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
    add_library(py2-c-module-cStringIO SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-cStringIO
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-cStringIO PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-cStringIO
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-cStringIO ####################

#################### py2-c-module-_codecs_jp ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_jp.c
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
    add_library(py2-c-module-_codecs_jp SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_jp
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_jp PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_jp
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_jp ####################

#################### py2-c-module-_codecs_iso2022 ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_iso2022.c
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
    add_library(py2-c-module-_codecs_iso2022 SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_iso2022
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_iso2022 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_iso2022
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_iso2022 ####################

#################### py2-c-module-zlib ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/zlibmodule.c
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
    add_library(py2-c-module-zlib SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-zlib
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-zlib PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-zlib
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### py2-c-module-zlib ####################

#################### py2-c-module-future_builtins ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/future_builtins.c
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
    add_library(py2-c-module-future_builtins SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-future_builtins
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-future_builtins PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-future_builtins
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-future_builtins ####################

#################### py2-launcher-lib ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/gcmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/getpath.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/config.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/getbuildinfo.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/threadmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/signalmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/posixmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/errnomodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/pwdmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sre.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_codecsmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_weakref.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/zipimport.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/symtablemodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/main.c
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
    add_library(py2-launcher-lib SHARED ${SOURCE_FILES})

    target_include_directories(py2-launcher-lib
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-launcher-lib PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-launcher-lib
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libexpat
            liblog
            libm
            libsqlite
            libssl
            libz
            py2-c-module-_bisect
            py2-c-module-_codecs_cn
            py2-c-module-_codecs_hk
            py2-c-module-_codecs_iso2022
            py2-c-module-_codecs_jp
            py2-c-module-_codecs_kr
            py2-c-module-_codecs_tw
            py2-c-module-_collections
            py2-c-module-_csv
            py2-c-module-_ctypes
            py2-c-module-_ctypes_test
            py2-c-module-_functools
            py2-c-module-_heapq
            py2-c-module-_hotshot
            py2-c-module-_io
            py2-c-module-_json
            py2-c-module-_lsprof
            py2-c-module-_md5
            py2-c-module-_multibytecodec
            py2-c-module-_multiprocessing
            py2-c-module-_random
            py2-c-module-_sha
            py2-c-module-_sha256
            py2-c-module-_sha512
            py2-c-module-_socket
            py2-c-module-_sqlite3
            py2-c-module-_ssl
            py2-c-module-_struct
            py2-c-module-_testcapi
            py2-c-module-array
            py2-c-module-audioop
            py2-c-module-binascii
            py2-c-module-cPickle
            py2-c-module-cStringIO
            py2-c-module-cmath
            py2-c-module-datetime
            py2-c-module-fcntl
            py2-c-module-future_builtins
            py2-c-module-grp
            py2-c-module-itertools
            py2-c-module-math
            py2-c-module-mmap
            py2-c-module-operator
            py2-c-module-parser
            py2-c-module-pyexpat
            py2-c-module-resource
            py2-c-module-select
            py2-c-module-strop
            py2-c-module-syslog
            py2-c-module-termios
            py2-c-module-time
            py2-c-module-unicodedata
            py2-c-module-xxsubtype
            py2-c-module-zlib
            py2-interp-object
            py2-interp-parser
            py2-interp-python
    )
endif ()
#################### py2-launcher-lib ####################

#################### py2-c-module-_md5 ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/md5module.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/md5.c
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
    add_library(py2-c-module-_md5 SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_md5
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_md5 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_md5
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_md5 ####################

#################### py2-c-module-array ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/arraymodule.c
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
    add_library(py2-c-module-array SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-array
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-array PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-array
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-array ####################

#################### py2-c-module-termios ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/termios.c
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
    add_library(py2-c-module-termios SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-termios
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-termios PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-termios
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-termios ####################

#################### py2-c-module-xxsubtype ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/xxsubtype.c
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
    add_library(py2-c-module-xxsubtype SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-xxsubtype
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-xxsubtype PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-xxsubtype
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-xxsubtype ####################

#################### py2-c-module-_heapq ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_heapqmodule.c
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
    add_library(py2-c-module-_heapq SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_heapq
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_heapq PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_heapq
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_heapq ####################

#################### py2-c-module-_multiprocessing ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_multiprocessing/multiprocessing.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_multiprocessing/semaphore.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_multiprocessing/socket_connection.c
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
    add_library(py2-c-module-_multiprocessing SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_multiprocessing
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_multiprocessing PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_multiprocessing
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_multiprocessing ####################

#################### py2-c-module-_bisect ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_bisectmodule.c
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
    add_library(py2-c-module-_bisect SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_bisect
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_bisect PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_bisect
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_bisect ####################

#################### py2-c-module-unicodedata ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/unicodedata.c
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
    add_library(py2-c-module-unicodedata SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-unicodedata
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-unicodedata PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-unicodedata
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-unicodedata ####################

#################### py2-c-module-_sha256 ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/sha256module.c
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
    add_library(py2-c-module-_sha256 SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_sha256
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_sha256 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_sha256
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_sha256 ####################

#################### py2-c-module-_ssl ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_ssl.c
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
    add_library(py2-c-module-_ssl SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_ssl
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_ssl PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_ssl
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_static
            libdl
            libm
            libssl
    )
endif ()
#################### py2-c-module-_ssl ####################

#################### py2-interp-python ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Python/_warnings.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/Python-ast.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/asdl.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/ast.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/bltinmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/ceval.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/compile.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/codecs.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/errors.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/frozen.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/frozenmain.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/future.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getargs.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getcompiler.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getcopyright.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getplatform.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getversion.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/graminit.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/import.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/importdl.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/marshal.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/modsupport.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/mystrtoul.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/mysnprintf.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/peephole.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pyarena.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pyctype.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pyfpe.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pymath.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pystate.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pythonrun.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/random.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/structmember.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/symtable.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/sysmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/traceback.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/getopt.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pystrcmp.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/pystrtod.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/dtoa.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/formatter_unicode.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/formatter_string.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/dynload_shlib.c
        ${ANDROID_ROOT}/external/python/cpython2/Python/thread.c
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
    add_library(py2-interp-python SHARED ${SOURCE_FILES})

    target_include_directories(py2-interp-python
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-interp-python PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-interp-python
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-interp-python ####################

#################### py2-interp-object ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Objects/abstract.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/boolobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/bufferobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/bytes_methods.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/bytearrayobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/capsule.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/cellobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/classobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/cobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/codeobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/complexobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/descrobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/enumobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/exceptions.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/genobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/fileobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/floatobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/frameobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/funcobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/intobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/iterobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/listobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/longobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/dictobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/memoryobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/methodobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/moduleobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/object.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/obmalloc.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/rangeobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/setobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/sliceobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/stringobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/structseq.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/tupleobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/typeobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/weakrefobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/unicodeobject.c
        ${ANDROID_ROOT}/external/python/cpython2/Objects/unicodectype.c
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
    add_library(py2-interp-object SHARED ${SOURCE_FILES})

    target_include_directories(py2-interp-object
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-interp-object PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-interp-object
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-interp-object ####################

#################### py2-c-module-time ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/timemodule.c
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
    add_library(py2-c-module-time SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-time
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-time PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-time
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-time ####################

#################### py2-c-module-cmath ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cmathmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_math.c
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
    add_library(py2-c-module-cmath SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-cmath
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-cmath PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-cmath
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-cmath ####################

#################### py2-c-module-_collections ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_collectionsmodule.c
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
    add_library(py2-c-module-_collections SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_collections
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_collections PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_collections
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_collections ####################

#################### py2-c-module-itertools ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/itertoolsmodule.c
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
    add_library(py2-c-module-itertools SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-itertools
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-itertools PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-itertools
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-itertools ####################

#################### py2-c-module-_codecs_kr ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_kr.c
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
    add_library(py2-c-module-_codecs_kr SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_kr
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_kr PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_kr
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_kr ####################

#################### py2-c-module-_multibytecodec ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/multibytecodec.c
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
    add_library(py2-c-module-_multibytecodec SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_multibytecodec
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_multibytecodec PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_multibytecodec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_multibytecodec ####################

#################### py2-c-module-parser ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/parsermodule.c
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
    add_library(py2-c-module-parser SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-parser
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-parser PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-parser
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-parser ####################

#################### py2-c-module-mmap ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/mmapmodule.c
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
    add_library(py2-c-module-mmap SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-mmap
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-mmap PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-mmap
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-mmap ####################

#################### py2-launcher-autorun ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Launcher/launcher_main.cpp
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
    add_executable(py2-launcher-autorun ${SOURCE_FILES})

    target_include_directories(py2-launcher-autorun
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-launcher-autorun PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-launcher-autorun
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libexpat
            liblog
            libm
            libsqlite
            libssl
            libz
            py2-launcher-lib
    )
endif ()
#################### py2-launcher-autorun ####################

#################### py2-c-module-cPickle ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cPickle.c
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
    add_library(py2-c-module-cPickle SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-cPickle
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-cPickle PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-cPickle
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-cPickle ####################

#################### py2-c-module-select ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/selectmodule.c
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
    add_library(py2-c-module-select SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-select
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-select PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-select
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-select ####################

#################### py2-c-module-fcntl ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/fcntlmodule.c
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
    add_library(py2-c-module-fcntl SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-fcntl
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-fcntl PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-fcntl
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-fcntl ####################

#################### py2-c-module-_random ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_randommodule.c
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
    add_library(py2-c-module-_random SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_random
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_random PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_random
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_random ####################

#################### py2-c-module-_csv ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_csv.c
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
    add_library(py2-c-module-_csv SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_csv
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_csv PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_csv
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_csv ####################

#################### py2-c-module-math ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/mathmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_math.c
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
    add_library(py2-c-module-math SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-math
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-math PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-math
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-math ####################

#################### py2-c-module-_socket ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/socketmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/timemodule.c
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
    add_library(py2-c-module-_socket SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_socket
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_socket PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_socket
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_socket ####################

#################### py2-c-module-_codecs_tw ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_tw.c
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
    add_library(py2-c-module-_codecs_tw SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_tw
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_tw PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_tw
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_tw ####################

#################### py2-c-module-_json ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_json.c
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
    add_library(py2-c-module-_json SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_json
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_json PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_json
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_json ####################

#################### py2-c-module-grp ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/grpmodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/rotatingtree.c
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
    add_library(py2-c-module-grp SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-grp
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-grp PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-grp
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-grp ####################

#################### py2-c-module-_lsprof ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_lsprof.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/rotatingtree.c
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
    add_library(py2-c-module-_lsprof SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_lsprof
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_lsprof PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_lsprof
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_lsprof ####################

#################### py2-c-module-_codecs_cn ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/cjkcodecs/_codecs_cn.c
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
    add_library(py2-c-module-_codecs_cn SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_codecs_cn
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_codecs_cn PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_codecs_cn
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_codecs_cn ####################

#################### py2-c-module-_sha512 ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/sha512module.c
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
    add_library(py2-c-module-_sha512 SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_sha512
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_sha512 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_sha512
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_sha512 ####################

#################### py2-c-module-resource ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/resource.c
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
    add_library(py2-c-module-resource SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-resource
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-resource PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-resource
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-resource ####################

#################### py2-c-module-_sqlite3 ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/cache.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/connection.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/cursor.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/microprotocols.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/module.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/prepare_protocol.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/row.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/statement.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite/util.c
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
    add_library(py2-c-module-_sqlite3 SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_sqlite3
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2/Modules/_sqlite
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_sqlite3 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_sqlite3
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libsqlite
    )
endif ()
#################### py2-c-module-_sqlite3 ####################

#################### py2-c-module-_io ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/bufferedio.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/bytesio.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/fileio.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/iobase.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/_iomodule.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/stringio.c
        ${ANDROID_ROOT}/external/python/cpython2/Modules/_io/textio.c
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
    add_library(py2-c-module-_io SHARED ${SOURCE_FILES})

    target_include_directories(py2-c-module-_io
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython2/android_arm64/pyconfig
            ${ANDROID_ROOT}/external/python/cpython2/Include
            ${ANDROID_ROOT}/external/python/cpython2/Modules/_io
            ${ANDROID_ROOT}/external/python/cpython2
    )

    # Add c/c++ flags
    target_compile_options(py2-c-module-_io PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py2-c-module-_io
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### py2-c-module-_io ####################
