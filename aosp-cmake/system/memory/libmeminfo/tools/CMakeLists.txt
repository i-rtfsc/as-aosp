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

project(procrank)

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

#################### showmap ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/showmap.cpp
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
    add_executable(showmap ${SOURCE_FILES})

    target_include_directories(showmap
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(showmap PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(showmap
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libmeminfo
            libsmapinfo
    )
endif ()
#################### showmap ####################

#################### wsstop ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/wsstop.cpp
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
    add_executable(wsstop ${SOURCE_FILES})

    target_include_directories(wsstop
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(wsstop PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(wsstop
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libmeminfo
    )
endif ()
#################### wsstop ####################

#################### bugreport_procdump ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/bugreport_procdump.cpp
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
    add_executable(bugreport_procdump ${SOURCE_FILES})

    target_include_directories(bugreport_procdump
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(bugreport_procdump PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(bugreport_procdump
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libmeminfo
            libsmapinfo
    )
endif ()
#################### bugreport_procdump ####################

#################### librank ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/librank.cpp
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
    add_executable(librank ${SOURCE_FILES})

    target_include_directories(librank
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(librank PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librank
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libmeminfo
            libsmapinfo
    )
endif ()
#################### librank ####################

#################### procmem ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/procmem.cpp
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
    add_executable(procmem ${SOURCE_FILES})

    target_include_directories(procmem
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(procmem PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(procmem
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libmeminfo
    )
endif ()
#################### procmem ####################

#################### procrank ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/libmeminfo/tools/procrank.cpp
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
    add_executable(procrank ${SOURCE_FILES})

    target_include_directories(procrank
            PUBLIC
            ${ANDROID_ROOT}/system/memory/libmeminfo/tools
    )

    # Add c/c++ flags
    target_compile_options(procrank PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(procrank
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libmeminfo
            libprocinfo
            libsmapinfo
    )
endif ()
#################### procrank ####################
