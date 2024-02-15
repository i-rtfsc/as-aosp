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

project(hwservicemanager)

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

#################### hwservicemanager ####################
set(FILES
        ${ANDROID_ROOT}/system/hwservicemanager/service.cpp
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
    add_executable(hwservicemanager ${SOURCE_FILES})

    target_include_directories(hwservicemanager
            PUBLIC
            ${ANDROID_ROOT}/system/hwservicemanager
    )

    # Add c/c++ flags
    target_compile_options(hwservicemanager PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(hwservicemanager
            android.hidl.token-1.0
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhidl-gen-utils
            libhidlbase
            libhwservicemanager
            liblog
            libm
            libselinux
            libtokenmanager
            libutils
            libvintf
    )
endif ()
#################### hwservicemanager ####################

#################### libtokenmanager ####################
set(FILES
        ${ANDROID_ROOT}/system/hwservicemanager/TokenManager.cpp
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
    add_library(libtokenmanager SHARED ${SOURCE_FILES})

    target_include_directories(libtokenmanager
            PUBLIC
            ${ANDROID_ROOT}/system/hwservicemanager
    )

    # Add c/c++ flags
    target_compile_options(libtokenmanager PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libtokenmanager
            android.hidl.token-1.0
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libhidlbase
            liblog
            libm
    )
endif ()
#################### libtokenmanager ####################

#################### libhwservicemanager ####################
set(FILES
        ${ANDROID_ROOT}/system/hwservicemanager/AccessControl.cpp
        ${ANDROID_ROOT}/system/hwservicemanager/HidlService.cpp
        ${ANDROID_ROOT}/system/hwservicemanager/ServiceManager.cpp
        ${ANDROID_ROOT}/system/hwservicemanager/Vintf.cpp
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
    add_library(libhwservicemanager SHARED ${SOURCE_FILES})

    target_include_directories(libhwservicemanager
            PUBLIC
            ${ANDROID_ROOT}/system/hwservicemanager
    )

    # Add c/c++ flags
    target_compile_options(libhwservicemanager PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libhwservicemanager
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidl-gen-utils
            libhidlbase
            liblog
            libm
            libselinux
            libutils
            libvintf
    )
endif ()
#################### libhwservicemanager ####################
