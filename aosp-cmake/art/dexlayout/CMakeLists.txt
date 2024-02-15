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

project(dexlayoutd)

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

#################### libart-dexlayout ####################
set(FILES
        ${ANDROID_ROOT}/art/dexlayout/compact_dex_writer.cc
        ${ANDROID_ROOT}/art/dexlayout/dexlayout.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_ir.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_ir_builder.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_verify.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_visualize.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_writer.cc
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
    add_library(libart-dexlayout SHARED ${SOURCE_FILES})

    target_include_directories(libart-dexlayout
            PUBLIC
            ${ANDROID_ROOT}/art/dexlayout
    )

    # Add c/c++ flags
    target_compile_options(libart-dexlayout PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libart-dexlayout
            libartbase
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile
            libdl
            libm
            libprofile
            libwinpthread
            libz
    )
endif ()
#################### libart-dexlayout ####################

#################### dexlayout ####################
set(FILES
        ${ANDROID_ROOT}/art/dexlayout/dexlayout_main.cc
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
    add_executable(dexlayout ${SOURCE_FILES})

    target_include_directories(dexlayout
            PUBLIC
            ${ANDROID_ROOT}/art/dexlayout
    )

    # Add c/c++ flags
    target_compile_options(dexlayout PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(dexlayout
            libart-dexlayout
            libartbase
            libbase
            libc++
            libdexfile
            libprofile
    )
endif ()
#################### dexlayout ####################

#################### dexdiag ####################
set(FILES
        ${ANDROID_ROOT}/art/dexlayout/dexdiag.cc
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
    add_executable(dexdiag ${SOURCE_FILES})

    target_include_directories(dexdiag
            PUBLIC
            ${ANDROID_ROOT}/art/dexlayout
    )

    # Add c/c++ flags
    target_compile_options(dexdiag PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wall
    )

    target_link_libraries(dexdiag
            libart
            libart-dexlayout
            libartbase
            libbase
            libc++
            libdexfile
    )
endif ()
#################### dexdiag ####################

#################### libartd-dexlayout ####################
set(FILES
        ${ANDROID_ROOT}/art/dexlayout/compact_dex_writer.cc
        ${ANDROID_ROOT}/art/dexlayout/dexlayout.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_ir.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_ir_builder.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_verify.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_visualize.cc
        ${ANDROID_ROOT}/art/dexlayout/dex_writer.cc
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
    add_library(libartd-dexlayout SHARED ${SOURCE_FILES})

    target_include_directories(libartd-dexlayout
            PUBLIC
            ${ANDROID_ROOT}/art/dexlayout
    )

    # Add c/c++ flags
    target_compile_options(libartd-dexlayout PRIVATE
            -DVIXL_DEBUG
            -UNDEBUG
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libartd-dexlayout
            libartbased
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfiled
            libdl
            libm
            libprofiled
            libz
    )
endif ()
#################### libartd-dexlayout ####################

#################### dexlayoutd ####################
set(FILES
        ${ANDROID_ROOT}/art/dexlayout/dexlayout_main.cc
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
    add_executable(dexlayoutd ${SOURCE_FILES})

    target_include_directories(dexlayoutd
            PUBLIC
            ${ANDROID_ROOT}/art/dexlayout
    )

    # Add c/c++ flags
    target_compile_options(dexlayoutd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(dexlayoutd
            libartbased
            libartd-dexlayout
            libbase
            libc++
            libdexfiled
            libprofiled
    )
endif ()
#################### dexlayoutd ####################
