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

project(sqlite3)

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

#################### libsqlite_static_noicu ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/dist/sqlite3.c
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
    add_library(libsqlite_static_noicu SHARED ${SOURCE_FILES})

    target_include_directories(libsqlite_static_noicu
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/dist
    )

    # Add c/c++ flags
    target_compile_options(libsqlite_static_noicu PRIVATE
            -DNDEBUG=1
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsqlite_static_noicu
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libsqlite3_android_noicu
    )
endif ()
#################### libsqlite_static_noicu ####################

#################### libsqlite ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/dist/sqlite3.c
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
    add_library(libsqlite SHARED ${SOURCE_FILES})

    target_include_directories(libsqlite
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/dist
    )

    # Add c/c++ flags
    target_compile_options(libsqlite PRIVATE
            -DNDEBUG=1
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsqlite
            libandroidicu
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            liblog
            libm
            libsqlite3_android
            libwinpthread
    )
endif ()
#################### libsqlite ####################

#################### libsqlite_static_minimal ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/dist/sqlite3.c
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
    add_library(libsqlite_static_minimal SHARED ${SOURCE_FILES})

    target_include_directories(libsqlite_static_minimal
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/dist
    )

    # Add c/c++ flags
    target_compile_options(libsqlite_static_minimal PRIVATE
            -DNDEBUG=1
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsqlite_static_minimal
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libsqlite_static_minimal ####################

#################### sqlite3 ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/dist/shell.c
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
    add_executable(sqlite3 ${SOURCE_FILES})

    target_include_directories(sqlite3
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/dist
    )

    # Add c/c++ flags
    target_compile_options(sqlite3 PRIVATE
            -DNDEBUG=1
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(sqlite3
            libandroidicu
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libsqlite
            libutils
            libwinpthread
    )
endif ()
#################### sqlite3 ####################

#################### sqlite_ext_percentile ####################
set(FILES
        ${ANDROID_ROOT}/external/sqlite/dist/ext/misc/percentile.c
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
    add_library(sqlite_ext_percentile SHARED ${SOURCE_FILES})

    target_include_directories(sqlite_ext_percentile
            PUBLIC
            ${ANDROID_ROOT}/external/sqlite/dist
    )

    # Add c/c++ flags
    target_compile_options(sqlite_ext_percentile PRIVATE
            -DNDEBUG=1
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(sqlite_ext_percentile
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### sqlite_ext_percentile ####################
