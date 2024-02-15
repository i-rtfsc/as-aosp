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

project(libxml2_schema_fuzzer)

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

#################### libxml2_xpath_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/fuzz/fuzz.c
        ${ANDROID_ROOT}/external/libxml2/fuzz/xpath.c
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
    add_executable(libxml2_xpath_fuzzer ${SOURCE_FILES})

    target_include_directories(libxml2_xpath_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2
            ${ANDROID_ROOT}/external/libxml2/include
    )

    # Add c/c++ flags
    target_compile_options(libxml2_xpath_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libxml2_xpath_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libxml2
    )
endif ()
#################### libxml2_xpath_fuzzer ####################

#################### libxml2_regexp_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/fuzz/fuzz.c
        ${ANDROID_ROOT}/external/libxml2/fuzz/regexp.c
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
    add_executable(libxml2_regexp_fuzzer ${SOURCE_FILES})

    target_include_directories(libxml2_regexp_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2
            ${ANDROID_ROOT}/external/libxml2/include
    )

    # Add c/c++ flags
    target_compile_options(libxml2_regexp_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libxml2_regexp_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libxml2
    )
endif ()
#################### libxml2_regexp_fuzzer ####################

#################### libxml2_ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/SAX.c
        ${ANDROID_ROOT}/external/libxml2/entities.c
        ${ANDROID_ROOT}/external/libxml2/encoding.c
        ${ANDROID_ROOT}/external/libxml2/error.c
        ${ANDROID_ROOT}/external/libxml2/parserInternals.c
        ${ANDROID_ROOT}/external/libxml2/parser.c
        ${ANDROID_ROOT}/external/libxml2/tree.c
        ${ANDROID_ROOT}/external/libxml2/hash.c
        ${ANDROID_ROOT}/external/libxml2/list.c
        ${ANDROID_ROOT}/external/libxml2/xmlIO.c
        ${ANDROID_ROOT}/external/libxml2/xmlmemory.c
        ${ANDROID_ROOT}/external/libxml2/uri.c
        ${ANDROID_ROOT}/external/libxml2/valid.c
        ${ANDROID_ROOT}/external/libxml2/xlink.c
        ${ANDROID_ROOT}/external/libxml2/debugXML.c
        ${ANDROID_ROOT}/external/libxml2/xpath.c
        ${ANDROID_ROOT}/external/libxml2/xpointer.c
        ${ANDROID_ROOT}/external/libxml2/xinclude.c
        ${ANDROID_ROOT}/external/libxml2/catalog.c
        ${ANDROID_ROOT}/external/libxml2/globals.c
        ${ANDROID_ROOT}/external/libxml2/threads.c
        ${ANDROID_ROOT}/external/libxml2/c14n.c
        ${ANDROID_ROOT}/external/libxml2/xmlstring.c
        ${ANDROID_ROOT}/external/libxml2/buf.c
        ${ANDROID_ROOT}/external/libxml2/xmlregexp.c
        ${ANDROID_ROOT}/external/libxml2/xmlschemas.c
        ${ANDROID_ROOT}/external/libxml2/xmlschemastypes.c
        ${ANDROID_ROOT}/external/libxml2/xmlunicode.c
        ${ANDROID_ROOT}/external/libxml2/xmlreader.c
        ${ANDROID_ROOT}/external/libxml2/relaxng.c
        ${ANDROID_ROOT}/external/libxml2/dict.c
        ${ANDROID_ROOT}/external/libxml2/SAX2.c
        ${ANDROID_ROOT}/external/libxml2/xmlwriter.c
        ${ANDROID_ROOT}/external/libxml2/legacy.c
        ${ANDROID_ROOT}/external/libxml2/chvalid.c
        ${ANDROID_ROOT}/external/libxml2/pattern.c
        ${ANDROID_ROOT}/external/libxml2/xmlsave.c
        ${ANDROID_ROOT}/external/libxml2/xmlmodule.c
        ${ANDROID_ROOT}/external/libxml2/schematron.c
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
    add_library(libxml2_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libxml2_ndk
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2/include
            ${ANDROID_ROOT}/external/libxml2
    )

    # Add c/c++ flags
    target_compile_options(libxml2_ndk PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libxml2_ndk
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libxml2_ndk ####################

#################### libxml2 ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/SAX.c
        ${ANDROID_ROOT}/external/libxml2/entities.c
        ${ANDROID_ROOT}/external/libxml2/encoding.c
        ${ANDROID_ROOT}/external/libxml2/error.c
        ${ANDROID_ROOT}/external/libxml2/parserInternals.c
        ${ANDROID_ROOT}/external/libxml2/parser.c
        ${ANDROID_ROOT}/external/libxml2/tree.c
        ${ANDROID_ROOT}/external/libxml2/hash.c
        ${ANDROID_ROOT}/external/libxml2/list.c
        ${ANDROID_ROOT}/external/libxml2/xmlIO.c
        ${ANDROID_ROOT}/external/libxml2/xmlmemory.c
        ${ANDROID_ROOT}/external/libxml2/uri.c
        ${ANDROID_ROOT}/external/libxml2/valid.c
        ${ANDROID_ROOT}/external/libxml2/xlink.c
        ${ANDROID_ROOT}/external/libxml2/debugXML.c
        ${ANDROID_ROOT}/external/libxml2/xpath.c
        ${ANDROID_ROOT}/external/libxml2/xpointer.c
        ${ANDROID_ROOT}/external/libxml2/xinclude.c
        ${ANDROID_ROOT}/external/libxml2/catalog.c
        ${ANDROID_ROOT}/external/libxml2/globals.c
        ${ANDROID_ROOT}/external/libxml2/threads.c
        ${ANDROID_ROOT}/external/libxml2/c14n.c
        ${ANDROID_ROOT}/external/libxml2/xmlstring.c
        ${ANDROID_ROOT}/external/libxml2/buf.c
        ${ANDROID_ROOT}/external/libxml2/xmlregexp.c
        ${ANDROID_ROOT}/external/libxml2/xmlschemas.c
        ${ANDROID_ROOT}/external/libxml2/xmlschemastypes.c
        ${ANDROID_ROOT}/external/libxml2/xmlunicode.c
        ${ANDROID_ROOT}/external/libxml2/xmlreader.c
        ${ANDROID_ROOT}/external/libxml2/relaxng.c
        ${ANDROID_ROOT}/external/libxml2/dict.c
        ${ANDROID_ROOT}/external/libxml2/SAX2.c
        ${ANDROID_ROOT}/external/libxml2/xmlwriter.c
        ${ANDROID_ROOT}/external/libxml2/legacy.c
        ${ANDROID_ROOT}/external/libxml2/chvalid.c
        ${ANDROID_ROOT}/external/libxml2/pattern.c
        ${ANDROID_ROOT}/external/libxml2/xmlsave.c
        ${ANDROID_ROOT}/external/libxml2/xmlmodule.c
        ${ANDROID_ROOT}/external/libxml2/schematron.c
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
    add_library(libxml2 SHARED ${SOURCE_FILES})

    target_include_directories(libxml2
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2/include
            ${ANDROID_ROOT}/external/libxml2
    )

    # Add c/c++ flags
    target_compile_options(libxml2 PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libxml2
            libandroidicu
            libc
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libm
    )
endif ()
#################### libxml2 ####################

#################### libxml2_uri_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/fuzz/fuzz.c
        ${ANDROID_ROOT}/external/libxml2/fuzz/uri.c
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
    add_executable(libxml2_uri_fuzzer ${SOURCE_FILES})

    target_include_directories(libxml2_uri_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2
            ${ANDROID_ROOT}/external/libxml2/include
    )

    # Add c/c++ flags
    target_compile_options(libxml2_uri_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libxml2_uri_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libxml2
    )
endif ()
#################### libxml2_uri_fuzzer ####################

#################### libxml2_xml_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/fuzz/fuzz.c
        ${ANDROID_ROOT}/external/libxml2/fuzz/xml.c
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
    add_executable(libxml2_xml_fuzzer ${SOURCE_FILES})

    target_include_directories(libxml2_xml_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2
            ${ANDROID_ROOT}/external/libxml2/include
    )

    # Add c/c++ flags
    target_compile_options(libxml2_xml_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libxml2_xml_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libxml2
    )
endif ()
#################### libxml2_xml_fuzzer ####################

#################### libxml2_schema_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libxml2/fuzz/fuzz.c
        ${ANDROID_ROOT}/external/libxml2/fuzz/schema.c
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
    add_executable(libxml2_schema_fuzzer ${SOURCE_FILES})

    target_include_directories(libxml2_schema_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libxml2
            ${ANDROID_ROOT}/external/libxml2/include
    )

    # Add c/c++ flags
    target_compile_options(libxml2_schema_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libxml2_schema_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libxml2
    )
endif ()
#################### libxml2_schema_fuzzer ####################
