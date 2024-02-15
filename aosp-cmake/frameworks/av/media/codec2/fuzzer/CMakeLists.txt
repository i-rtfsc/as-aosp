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

project(C2FuzzerFlacDec)

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

#################### C2FuzzerH263Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerH263Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerH263Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerH263Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerH263Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_h263dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libstagefright_m4vh263dec
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerH263Dec ####################

#################### C2FuzzerG711MlawDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerG711MlawDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerG711MlawDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerG711MlawDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerG711MlawDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            codecs_g711dec
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_g711mlawdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerG711MlawDec ####################

#################### C2FuzzerRawDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerRawDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerRawDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerRawDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerRawDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_rawdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerRawDec ####################

#################### C2FuzzerMp3Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerMp3Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerMp3Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerMp3Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerMp3Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_mp3dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libstagefright_mp3dec
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerMp3Dec ####################

#################### C2FuzzerAvcDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerAvcDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerAvcDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerAvcDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerAvcDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libavcdec
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_avcdec
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerAvcDec ####################

#################### C2FuzzerOpusDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerOpusDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerOpusDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerOpusDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerOpusDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_opusdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libopus
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerOpusDec ####################

#################### C2FuzzerG711AlawDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerG711AlawDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerG711AlawDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerG711AlawDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerG711AlawDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            codecs_g711dec
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_g711alawdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerG711AlawDec ####################

#################### C2FuzzerGsmDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerGsmDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerGsmDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerGsmDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerGsmDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_gsmdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libgsm
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerGsmDec ####################

#################### C2FuzzerMpeg2Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerMpeg2Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerMpeg2Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerMpeg2Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerMpeg2Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_mpeg2dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libmpeg2dec
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerMpeg2Dec ####################

#################### C2FuzzerXaacDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerXaacDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerXaacDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerXaacDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerXaacDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_xaacdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            libxaacdec
    )
endif ()
#################### C2FuzzerXaacDec ####################

#################### C2FuzzerAacDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerAacDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerAacDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerAacDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerAacDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libFraunhoferAAC
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_aacdec
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerAacDec ####################

#################### C2FuzzerVp8Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerVp8Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerVp8Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerVp8Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerVp8Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_vp8dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            libvpx
    )
endif ()
#################### C2FuzzerVp8Dec ####################

#################### C2FuzzerAV1Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerAV1Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerAV1Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerAV1Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerAV1Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_av1dec_gav1
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgav1
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            libyuv_static
    )
endif ()
#################### C2FuzzerAV1Dec ####################

#################### C2FuzzerAmrnbDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerAmrnbDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerAmrnbDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerAmrnbDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerAmrnbDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_amrnbdec
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_amrnb_common
            libstagefright_amrnbdec
            libstagefright_amrwbdec
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerAmrnbDec ####################

#################### C2FuzzerMpeg4Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerMpeg4Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerMpeg4Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerMpeg4Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerMpeg4Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_mpeg4dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libstagefright_m4vh263dec
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerMpeg4Dec ####################

#################### C2FuzzerHevcDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerHevcDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerHevcDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerHevcDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerHevcDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_hevcdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhevcdec
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerHevcDec ####################

#################### C2FuzzerVp9Dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerVp9Dec ${SOURCE_FILES})

    target_include_directories(C2FuzzerVp9Dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerVp9Dec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerVp9Dec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_vp9dec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            libvpx
    )
endif ()
#################### C2FuzzerVp9Dec ####################

#################### C2FuzzerAmrwbDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerAmrwbDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerAmrwbDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerAmrwbDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerAmrwbDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_amrwbdec
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_amrnb_common
            libstagefright_amrnbdec
            libstagefright_amrwbdec
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerAmrwbDec ####################

#################### C2FuzzerVorbisDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerVorbisDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerVorbisDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerVorbisDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerVorbisDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_vorbisdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            libvorbisidec
    )
endif ()
#################### C2FuzzerVorbisDec ####################

#################### C2FuzzerFlacDec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer/C2Fuzzer.cpp
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
    add_executable(C2FuzzerFlacDec ${SOURCE_FILES})

    target_include_directories(C2FuzzerFlacDec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(C2FuzzerFlacDec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(C2FuzzerFlacDec
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libFLAC
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcodec2
            libcodec2_soft_common
            libcodec2_soft_flacdec
            libcodec2_vndk
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsfplugin_ccodec_utils
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_flacdec
            libstagefright_foundation
            libui
            libutils
            libvndksupport
    )
endif ()
#################### C2FuzzerFlacDec ####################
