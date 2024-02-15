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

project(py3-launcher-autorun)

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

#################### py3-launcher-lib ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython3/Modules/getpath.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/main.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/gcmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/getbuildinfo.c
        ${ANDROID_ROOT}/external/python/cpython3/android/bionic/config.c
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
    add_library(py3-launcher-lib SHARED ${SOURCE_FILES})

    target_include_directories(py3-launcher-lib
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython3/android/bionic/pyconfig
            ${ANDROID_ROOT}/external/python/cpython3/Include
            ${ANDROID_ROOT}/external/python/cpython3/Include/internal
            ${ANDROID_ROOT}/external/python/cpython3
    )

    # Add c/c++ flags
    target_compile_options(py3-launcher-lib PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py3-launcher-lib
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libexpat
            libffi
            liblog
            libm
            libsqlite
            libz
            py3-c-modules
            py3-interp
    )
endif ()
#################### py3-launcher-lib ####################

#################### py3-launcher ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython3/android/launcher_main.cpp
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
    add_executable(py3-launcher ${SOURCE_FILES})

    target_include_directories(py3-launcher
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython3/android/bionic/pyconfig
            ${ANDROID_ROOT}/external/python/cpython3/Include
            ${ANDROID_ROOT}/external/python/cpython3/Include/internal
            ${ANDROID_ROOT}/external/python/cpython3
    )

    # Add c/c++ flags
    target_compile_options(py3-launcher PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py3-launcher
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libexpat
            libffi
            liblog
            libm
            libsqlite
            libz
            py3-launcher-lib
    )
endif ()
#################### py3-launcher ####################

#################### py3-launcher-autorun ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython3/android/launcher_main.cpp
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
    add_executable(py3-launcher-autorun ${SOURCE_FILES})

    target_include_directories(py3-launcher-autorun
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython3/android/bionic/pyconfig
            ${ANDROID_ROOT}/external/python/cpython3/Include
            ${ANDROID_ROOT}/external/python/cpython3/Include/internal
            ${ANDROID_ROOT}/external/python/cpython3
    )

    # Add c/c++ flags
    target_compile_options(py3-launcher-autorun PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py3-launcher-autorun
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libexpat
            libffi
            liblog
            libm
            libsqlite
            libz
            py3-launcher-lib
    )
endif ()
#################### py3-launcher-autorun ####################

#################### py3-interp ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython3/Parser/pegen.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/parser.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/string_parser.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/peg_api.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/myreadline.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/token.c
        ${ANDROID_ROOT}/external/python/cpython3/Parser/tokenizer.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/abstract.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/accu.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/boolobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/bytes_methods.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/bytearrayobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/bytesobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/call.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/capsule.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/cellobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/classobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/codeobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/complexobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/descrobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/enumobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/exceptions.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/genericaliasobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/genobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/fileobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/floatobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/frameobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/funcobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/interpreteridobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/iterobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/listobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/longobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/dictobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/odictobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/memoryobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/methodobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/moduleobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/namespaceobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/object.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/obmalloc.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/picklebufobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/rangeobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/setobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/sliceobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/structseq.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/tupleobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/typeobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/unicodeobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/unicodectype.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/unionobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Objects/weakrefobject.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/_warnings.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/Python-ast.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/asdl.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/ast.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/ast_opt.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/ast_unparse.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/bltinmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/ceval.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/codecs.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/compile.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/context.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/dynamic_annotations.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/errors.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/frozen.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/frozenmain.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/future.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getargs.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getcompiler.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getcopyright.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getplatform.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getversion.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/hamt.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/hashtable.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/import.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/importdl.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/initconfig.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/marshal.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/modsupport.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/mysnprintf.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/mystrtoul.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pathconfig.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/preconfig.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pyarena.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pyctype.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pyfpe.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pyhash.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pylifecycle.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pymath.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pystate.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pythonrun.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pytime.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/bootstrap_hash.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/structmember.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/symtable.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/sysmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/thread.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/traceback.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/getopt.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pystrcmp.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pystrtod.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/pystrhex.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/dtoa.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/formatter_unicode.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/fileutils.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/suggestions.c
        ${ANDROID_ROOT}/external/python/cpython3/Python/dynload_shlib.c
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
    add_library(py3-interp SHARED ${SOURCE_FILES})

    target_include_directories(py3-interp
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython3/android/bionic/pyconfig
            ${ANDROID_ROOT}/external/python/cpython3/Include
            ${ANDROID_ROOT}/external/python/cpython3/Include/internal
            ${ANDROID_ROOT}/external/python/cpython3
    )

    # Add c/c++ flags
    target_compile_options(py3-interp PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py3-interp
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libffi
            libm
    )
endif ()
#################### py3-interp ####################

#################### py3-c-modules ####################
set(FILES
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_abc.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/arraymodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_asynciomodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/atexitmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/binascii.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_bisectmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_blake2/blake2b_impl.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_blake2/blake2module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_blake2/blake2s_impl.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_cn.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_hk.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_iso2022.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_jp.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_kr.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/_codecs_tw.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cjkcodecs/multibytecodec.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/cmathmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_codecsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_collectionsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_contextvarsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_csv.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_ctypes/callbacks.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_ctypes/callproc.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_ctypes/cfield.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_ctypes/_ctypes.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_ctypes/stgdict.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_datetimemodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/errnomodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/faulthandler.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/fcntlmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_functoolsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/grpmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_heapqmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/bufferedio.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/bytesio.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/fileio.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/iobase.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/_iomodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/stringio.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_io/textio.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/itertoolsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_json.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_localemodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_lsprof.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_math.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/mathmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/md5module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/mmapmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_multiprocessing/multiprocessing.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_multiprocessing/semaphore.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_opcode.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_operator.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_pickle.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/posixmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_posixsubprocess.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/pwdmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/pyexpat.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_queuemodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_randommodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/resource.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/rotatingtree.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/selectmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/sha1module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/sha256module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sha3/sha3module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/sha512module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/signalmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/socketmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/cache.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/connection.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/cursor.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/microprotocols.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/module.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/prepare_protocol.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/row.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/statement.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sqlite/util.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_sre.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_stat.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_statisticsmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_struct.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/symtablemodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/syslogmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/termios.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_threadmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/timemodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_tracemalloc.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/unicodedata.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_weakref.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_xxsubinterpretersmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/xxsubtype.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/zlibmodule.c
        ${ANDROID_ROOT}/external/python/cpython3/Modules/_zoneinfo.c
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
    add_library(py3-c-modules SHARED ${SOURCE_FILES})

    target_include_directories(py3-c-modules
            PUBLIC
            ${ANDROID_ROOT}/external/python/cpython3/android/bionic/pyconfig
            ${ANDROID_ROOT}/external/python/cpython3/Include
            ${ANDROID_ROOT}/external/python/cpython3/Include/internal
            ${ANDROID_ROOT}/external/python/cpython3
    )

    # Add c/c++ flags
    target_compile_options(py3-c-modules PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(py3-c-modules
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libexpat
            libffi
            libm
            libsqlite
            libz
    )
endif ()
#################### py3-c-modules ####################
