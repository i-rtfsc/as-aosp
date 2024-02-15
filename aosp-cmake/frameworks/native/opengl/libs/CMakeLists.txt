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

project(libGLESv3)

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

#################### libGLESv3 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/GLES2/gl2.cpp
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
    add_library(libGLESv3 SHARED ${SOURCE_FILES})

    target_include_directories(libGLESv3
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include
    )

    # Add c/c++ flags
    target_compile_options(libGLESv3 PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libGLESv3
            bionic_libc_platform_headers
            gl_headers
            libEGL
            libarect
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnativebase_headers
            libsystem_headers
    )
endif ()
#################### libGLESv3 ####################

#################### libEGL_getProcAddress ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/getProcAddress.cpp
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
    add_library(libEGL_getProcAddress SHARED ${SOURCE_FILES})

    target_include_directories(libEGL_getProcAddress
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
    )

    # Add c/c++ flags
    target_compile_options(libEGL_getProcAddress PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libEGL_getProcAddress
            bionic_libc_platform_headers
            gl_headers
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgraphicsenv
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsystem_headers
    )
endif ()
#################### libEGL_getProcAddress ####################

#################### libGLESv1_CM ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/GLES_CM/gl.cpp
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
    add_library(libGLESv1_CM SHARED ${SOURCE_FILES})

    target_include_directories(libGLESv1_CM
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include
    )

    # Add c/c++ flags
    target_compile_options(libGLESv1_CM PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libGLESv1_CM
            bionic_libc_platform_headers
            gl_headers
            libEGL
            libarect
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnativebase_headers
            libsystem_headers
    )
endif ()
#################### libGLESv1_CM ####################

#################### libEGL_blobCache ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/BlobCache.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/FileBlobCache.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/MultifileBlobCache.cpp
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
    add_library(libEGL_blobCache SHARED ${SOURCE_FILES})

    target_include_directories(libEGL_blobCache
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
    )

    # Add c/c++ flags
    target_compile_options(libEGL_blobCache PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libEGL_blobCache
            bionic_libc_platform_headers
            gl_headers
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgraphicsenv
            liblog
            libm
            libnativebase_headers
            libnativewindow
            libsystem_headers
    )
endif ()
#################### libEGL_blobCache ####################

#################### libETC1 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/ETC1/etc1.cpp
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
    add_library(libETC1 SHARED ${SOURCE_FILES})

    target_include_directories(libETC1
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
    )

    # Add c/c++ flags
    target_compile_options(libETC1 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libETC1
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libETC1 ####################

#################### libEGL ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_tls.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_cache.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_display.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_object.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_layers.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/eglApi.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_platform_entries.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/Loader.cpp
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/egl_angle_platform.cpp
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
    add_library(libEGL SHARED ${SOURCE_FILES})

    target_include_directories(libEGL
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL
    )

    # Add c/c++ flags
    target_compile_options(libEGL PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libEGL
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            bionic_libc_platform_headers
            gl_headers
            libEGL_blobCache
            libEGL_getProcAddress
            libSurfaceFlingerProp
            libarect
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgraphicsenv
            libhidlbase
            liblog
            libm
            libnativebase_headers
            libnativebridge_lazy
            libnativeloader_lazy
            libnativewindow
            libsurfaceflinger_headers
            libsystem_headers
            libunwindstack
            libutils
            libvndksupport
    )
endif ()
#################### libEGL ####################

#################### libGLESv2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/opengl/libs/GLES2/gl2.cpp
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
    add_library(libGLESv2 SHARED ${SOURCE_FILES})

    target_include_directories(libGLESv2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/opengl/libs
            ${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include
    )

    # Add c/c++ flags
    target_compile_options(libGLESv2 PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libGLESv2
            bionic_libc_platform_headers
            gl_headers
            libEGL
            libarect
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnativebase_headers
            libsystem_headers
    )
endif ()
#################### libGLESv2 ####################
