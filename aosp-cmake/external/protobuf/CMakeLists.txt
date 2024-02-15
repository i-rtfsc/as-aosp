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

project(libprotobuf-cpp-lite)

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

#################### libprotobuf-cpp-lite ####################
set(FILES
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format_lite.cc
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
    add_library(libprotobuf-cpp-lite SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-cpp-lite
            PUBLIC
            ${ANDROID_ROOT}/external/protobuf/src
            ${ANDROID_ROOT}/external/protobuf/android
            ${ANDROID_ROOT}/external/protobuf
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-cpp-lite PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-cpp-lite
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libprotobuf-cpp-lite ####################

#################### libprotobuf-cpp-lite_static ####################
set(FILES
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format_lite.cc
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
    add_library(libprotobuf-cpp-lite_static SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-cpp-lite_static
            PUBLIC
            ${ANDROID_ROOT}/external/protobuf/src
            ${ANDROID_ROOT}/external/protobuf/android
            ${ANDROID_ROOT}/external/protobuf
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-cpp-lite_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-cpp-lite_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libprotobuf-cpp-lite_static ####################

#################### libprotobuf-cpp-lite-ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format_lite.cc
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
    add_library(libprotobuf-cpp-lite-ndk SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-cpp-lite-ndk
            PUBLIC
            ${ANDROID_ROOT}/external/protobuf/src
            ${ANDROID_ROOT}/external/protobuf/android
            ${ANDROID_ROOT}/external/protobuf
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-cpp-lite-ndk PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-cpp-lite-ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libprotobuf-cpp-lite-ndk ####################

#################### libprotobuf-cpp-full ####################
set(FILES
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/api.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/compiler/importer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/compiler/parser.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor_database.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/duration.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/dynamic_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/empty.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set_heavy.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/field_mask.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_bases.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_reflection.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_full.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/gzip_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/printer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/tokenizer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/reflection_ops.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/service.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/source_context.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/struct.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/substitute.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/text_format.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/timestamp.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/type.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/unknown_field_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/delimited_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/field_comparator.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/field_mask_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/datapiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/default_value_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/error_listener.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/field_mask_utility.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_escaping.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_stream_parser.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/object_writer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/proto_writer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/protostream_objectsource.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/protostream_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/type_info.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/utility.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/json_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/message_differencer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/time_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/type_resolver_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wrappers.pb.cc
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
    add_library(libprotobuf-cpp-full SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-cpp-full
            PUBLIC
            ${ANDROID_ROOT}/external/protobuf/src
            ${ANDROID_ROOT}/external/protobuf/android
            ${ANDROID_ROOT}/external/protobuf
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-cpp-full PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-cpp-full
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libwinpthread
            libz
    )
endif ()
#################### libprotobuf-cpp-full ####################

#################### libprotobuf-cpp-full-ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arena.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenastring.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/arenaz_sampler.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_enum_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/implicit_weak_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/inlined_string_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/coded_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/io_win32.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/strtod.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/zero_copy_stream_impl_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/parse_context.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/repeated_ptr_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/bytestream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/common.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/int128.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/status.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/statusor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringpiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/stringprintf.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/structurally_valid.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/strutil.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/time.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format_lite.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/any.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/api.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/compiler/importer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/compiler/parser.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/descriptor_database.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/duration.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/dynamic_message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/empty.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/extension_set_heavy.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/field_mask.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_bases.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_reflection.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/generated_message_tctable_full.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/gzip_stream.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/printer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/io/tokenizer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/map_field.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/message.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/reflection_ops.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/service.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/source_context.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/struct.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/stubs/substitute.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/text_format.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/timestamp.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/type.pb.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/unknown_field_set.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/delimited_message_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/field_comparator.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/field_mask_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/datapiece.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/default_value_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/error_listener.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/field_mask_utility.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_escaping.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/json_stream_parser.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/object_writer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/proto_writer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/protostream_objectsource.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/protostream_objectwriter.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/type_info.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/internal/utility.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/json_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/message_differencer.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/time_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/util/type_resolver_util.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wire_format.cc
        ${ANDROID_ROOT}/external/protobuf/src/google/protobuf/wrappers.pb.cc
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
    add_library(libprotobuf-cpp-full-ndk SHARED ${SOURCE_FILES})

    target_include_directories(libprotobuf-cpp-full-ndk
            PUBLIC
            ${ANDROID_ROOT}/external/protobuf/src
            ${ANDROID_ROOT}/external/protobuf/android
            ${ANDROID_ROOT}/external/protobuf
    )

    # Add c/c++ flags
    target_compile_options(libprotobuf-cpp-full-ndk PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libprotobuf-cpp-full-ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libz
    )
endif ()
#################### libprotobuf-cpp-full-ndk ####################
