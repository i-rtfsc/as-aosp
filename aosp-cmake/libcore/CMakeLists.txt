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

project(libjavacore)

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

#################### libjavacore ####################
set(FILES
        ${ANDROID_ROOT}/libcore/luni/src/main/native/ExecStrings.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/IcuUtilities.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/JniConstants.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/JniException.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/NetworkUtilities.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/Register.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/ZipUtilities.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/android_system_OsConstants.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/cbigint.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/java_lang_StringToReal.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/java_lang_invoke_MethodHandle.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/java_lang_invoke_VarHandle.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_math_NativeBN.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_icu_ICU.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_io_AsynchronousCloseMonitor.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_io_Linux.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_io_Memory.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/libcore_util_NativeAllocationRegistry.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/org_apache_harmony_xml_ExpatParser.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/sun_misc_Unsafe.cpp
        ${ANDROID_ROOT}/libcore/luni/src/main/native/valueOf.cpp
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
    add_library(libjavacore SHARED ${SOURCE_FILES})

    target_include_directories(libjavacore
            PUBLIC
            ${ANDROID_ROOT}/libcore
    )

    # Add c/c++ flags
    target_compile_options(libjavacore PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libjavacore
            jni_headers
            libandroidio
            libbase
            libc++
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libexpat
            libicu
            liblog
            libnativehelper
            libnativehelper_compat_libc++
            libz
            libziparchive
    )
endif ()
#################### libjavacore ####################

#################### libopenjdkd ####################
set(FILES
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ZipFile.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inflater.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Deflater.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/CRC32.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Adler32.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/zip_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jni_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jni_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/io_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/canonicalize_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileDescriptor_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/DatagramChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/DatagramDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Console_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/IOUtil.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/PollArrayWrapper.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileDispatcherImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileSystemPreferences.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/EPoll.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/EPollPort.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixAsynchronousServerSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixAsynchronousSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/io_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/NativeThread.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileKey.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixFileSystem_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectStreamClass.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectOutputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/LinuxNativeDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/LinuxWatchService.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixCopyFile.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixNativeDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/InetAddress.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/net_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/net_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Net.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/MappedByteBuffer.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inet6Address.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inet4Address.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/linux_close.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ServerSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketOutputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Float.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Double.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/StrictMath.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Math.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ProcessEnvironment_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/System.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Runtime.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UNIXProcess_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Character.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/socket_tagger_util.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/OnLoad.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/JniConstants.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixDomainSockets.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jdk_internal_misc_VM.cpp
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
    add_library(libopenjdkd SHARED ${SOURCE_FILES})

    target_include_directories(libopenjdkd
            PUBLIC
            ${ANDROID_ROOT}/libcore/luni/src/main/native
            ${ANDROID_ROOT}/libcore
    )

    # Add c/c++ flags
    target_compile_options(libopenjdkd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libopenjdkd
            jni_headers
            libandroidio
            libbase
            libc++
            libclang_rt.builtins
            libcrypto
            libfdlibm
            libicu
            liblog
            libnativehelper
            libnativehelper_compat_libc++
            libopenjdkjvmd
            libz
    )
endif ()
#################### libopenjdkd ####################

#################### libopenjdk ####################
set(FILES
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ZipFile.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inflater.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Deflater.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/CRC32.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Adler32.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/zip_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jni_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jni_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/io_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/canonicalize_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileDescriptor_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/DatagramChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/DatagramDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Console_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/IOUtil.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/PollArrayWrapper.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileDispatcherImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileSystemPreferences.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/EPoll.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/EPollPort.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixAsynchronousServerSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixAsynchronousSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/io_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/NativeThread.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/FileKey.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixFileSystem_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectStreamClass.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectOutputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ObjectInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/LinuxNativeDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/LinuxWatchService.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixCopyFile.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixNativeDispatcher.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/InetAddress.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/net_util.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/net_util_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Net.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/MappedByteBuffer.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inet6Address.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Inet4Address.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/linux_close.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ServerSocketChannelImpl.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketInputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/SocketOutputStream.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Float.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Double.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/StrictMath.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Math.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/ProcessEnvironment_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/System.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Runtime.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UNIXProcess_md.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/Character.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/socket_tagger_util.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/OnLoad.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/JniConstants.cpp
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/UnixDomainSockets.c
        ${ANDROID_ROOT}/libcore/ojluni/src/main/native/jdk_internal_misc_VM.cpp
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
    add_library(libopenjdk SHARED ${SOURCE_FILES})

    target_include_directories(libopenjdk
            PUBLIC
            ${ANDROID_ROOT}/libcore/luni/src/main/native
            ${ANDROID_ROOT}/libcore
    )

    # Add c/c++ flags
    target_compile_options(libopenjdk PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libopenjdk
            jni_headers
            libandroidio
            libbase
            libc++
            libclang_rt.builtins
            libcrypto
            libfdlibm
            libicu
            liblog
            libnativehelper
            libnativehelper_compat_libc++
            libopenjdkjvm
            libz
    )
endif ()
#################### libopenjdk ####################

#################### libandroidio ####################
set(FILES
        ${ANDROID_ROOT}/libcore/luni/src/main/native/AsynchronousCloseMonitor.cpp
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
    add_library(libandroidio SHARED ${SOURCE_FILES})

    target_include_directories(libandroidio
            PUBLIC
            ${ANDROID_ROOT}/libcore
    )

    # Add c/c++ flags
    target_compile_options(libandroidio PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libandroidio
            libc
            libc++
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libandroidio ####################
