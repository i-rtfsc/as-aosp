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

project(memcpy)

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

#################### memcpy-2048kb ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/memcpy-2048kb.cpp
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
    add_executable(memcpy-2048kb ${SOURCE_FILES})

    target_include_directories(memcpy-2048kb
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(memcpy-2048kb
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### memcpy-2048kb ####################

#################### memcpy-byte ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/memcpy-byte.cpp
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
    add_executable(memcpy-byte ${SOURCE_FILES})

    target_include_directories(memcpy-byte
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(memcpy-byte
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### memcpy-byte ####################

#################### simd ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/simd.cpp
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
    add_executable(simd ${SOURCE_FILES})

    target_include_directories(simd
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(simd
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libeigen
            libm
    )
endif ()
#################### simd ####################

#################### while-true ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/while-true.cpp
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
    add_executable(while-true ${SOURCE_FILES})

    target_include_directories(while-true
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(while-true
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### while-true ####################

#################### memcpy-16kb ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/memcpy-16kb.cpp
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
    add_executable(memcpy-16kb ${SOURCE_FILES})

    target_include_directories(memcpy-16kb
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(memcpy-16kb
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### memcpy-16kb ####################

#################### pss ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/pss.cpp
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
    add_executable(pss ${SOURCE_FILES})

    target_include_directories(pss
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(pss
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
#################### pss ####################

#################### memcpy ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/cpu_loads/memcpy.cpp
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
    add_executable(memcpy ${SOURCE_FILES})

    target_include_directories(memcpy
            PUBLIC
            ${ANDROID_ROOT}/system/extras/cpu_loads
    )

    target_link_libraries(memcpy
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### memcpy ####################
