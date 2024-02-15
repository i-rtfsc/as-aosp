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

project(servicedispatcher)

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

#################### packagemanager_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/android/content/pm/IPackageManagerNative.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/android/content/pm/IStagedApexObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/android/content/pm/ApexStagedEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/android/content/pm/StagedApexInfo.cpp
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
    add_library(packagemanager_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(packagemanager_aidl-cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(packagemanager_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(packagemanager_aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### packagemanager_aidl-cpp ####################

#################### packagemanager_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk_platform-source/gen/android/content/pm/IPackageManagerNative.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk_platform-source/gen/android/content/pm/IStagedApexObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk_platform-source/gen/android/content/pm/ApexStagedEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk_platform-source/gen/android/content/pm/StagedApexInfo.cpp
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
    add_library(packagemanager_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(packagemanager_aidl-ndk_platform
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(packagemanager_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(packagemanager_aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### packagemanager_aidl-ndk_platform ####################

#################### libactivitymanager_aidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ActivityManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IActivityManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IUidObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libactivitymanager_aidl/android_${OUT_ARCH_CPU}_static/gen/aidl/android/app/ProcessStateEnum.cpp
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
    add_library(libactivitymanager_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libactivitymanager_aidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libactivitymanager_aidl/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libactivitymanager_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libactivitymanager_aidl
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libactivitymanager_aidl ####################

#################### libbinder ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BufferedTextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IPCThreadState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IServiceManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ProcessState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Static.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/ConnectionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/IServiceManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/os/ServiceDebugInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IPermissionController.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/PermissionCache.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/PermissionController.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/OS.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportRaw.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Binder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BpBinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Debug.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/FdTrigger.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IResultReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelFileDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RecordedTransaction.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcSession.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcServer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/TextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Trace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Utils.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IMemory.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IShellCallback.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/LazyServiceRegistrar.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/MemoryBase.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/MemoryDealer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/MemoryHeapBase.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelableHolder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/PersistableBundle.cpp
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
    add_library(libbinder SHARED ${SOURCE_FILES})

    target_include_directories(libbinder
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
    )

    # Add c/c++ flags
    target_compile_options(libbinder PRIVATE
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libbinder
            jni_headers
            libbase
            libbinder_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder ####################

#################### libbinder_rpc_unstable ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/libbinder_rpc_unstable.cpp
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
    add_library(libbinder_rpc_unstable SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_rpc_unstable
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_rpc_unstable
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform
    )

    # Add c/c++ flags
    target_compile_options(libbinder_rpc_unstable PRIVATE
            -Werror
    )

    target_link_libraries(libbinder_rpc_unstable
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils_sockets
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder_rpc_unstable ####################

#################### libbinder_rpc_single_threaded_no_kernel ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/OS.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportRaw.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Binder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BpBinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Debug.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/FdTrigger.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IResultReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelFileDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RecordedTransaction.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcSession.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcServer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/TextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Trace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Utils.cpp
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
    add_library(libbinder_rpc_single_threaded_no_kernel SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_rpc_single_threaded_no_kernel
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
    )

    # Add c/c++ flags
    target_compile_options(libbinder_rpc_single_threaded_no_kernel PRIVATE
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libbinder_rpc_single_threaded_no_kernel
            jni_headers
            libbase
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder_rpc_single_threaded_no_kernel ####################

#################### packagemanager_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk-source/gen/android/content/pm/IPackageManagerNative.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk-source/gen/android/content/pm/IStagedApexObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk-source/gen/android/content/pm/ApexStagedEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk-source/gen/android/content/pm/StagedApexInfo.cpp
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
    add_library(packagemanager_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(packagemanager_aidl-ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(packagemanager_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(packagemanager_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### packagemanager_aidl-ndk ####################

#################### libbinder_rpc_single_threaded ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BufferedTextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IPCThreadState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IServiceManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ProcessState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Static.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/ConnectionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IClientCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IServiceManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/ServiceDebugInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IPermissionController.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/PermissionCache.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/PermissionController.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/OS.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportRaw.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Binder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BpBinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Debug.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/FdTrigger.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IResultReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelFileDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RecordedTransaction.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcSession.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcServer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/TextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Trace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Utils.cpp
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
    add_library(libbinder_rpc_single_threaded SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_rpc_single_threaded
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder_rpc_single_threaded/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
    )

    # Add c/c++ flags
    target_compile_options(libbinder_rpc_single_threaded PRIVATE
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libbinder_rpc_single_threaded
            jni_headers
            libbase
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder_rpc_single_threaded ####################

#################### libbatterystats_aidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IBatteryStats.cpp
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
    add_library(libbatterystats_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libbatterystats_aidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_batterystats
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libbatterystats_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libbatterystats_aidl
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### libbatterystats_aidl ####################

#################### packagemanager_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-analyzer-source/gen/android/content/pm/IPackageManagerNative.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-analyzer-source/gen/android/content/pm/IStagedApexObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-analyzer-source/gen/android/content/pm/ApexStagedEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-analyzer-source/gen/android/content/pm/StagedApexInfo.cpp
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
    add_executable(packagemanager_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(packagemanager_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/packagemanager_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(packagemanager_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(packagemanager_aidl-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            packagemanager_aidl-cpp
    )
endif ()
#################### packagemanager_aidl-cpp-analyzer ####################

#################### libprocessinfoservice_aidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IProcessInfoService.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ProcessInfoService.cpp
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
    add_library(libprocessinfoservice_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libprocessinfoservice_aidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_processinfo
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libprocessinfoservice_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libprocessinfoservice_aidl
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libprocessinfoservice_aidl ####################

#################### libbinder_tls ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportTls.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTlsUtils.cpp
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
    add_library(libbinder_tls SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_tls
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_tls
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libbinder_tls PRIVATE
            -Werror
    )

    target_link_libraries(libbinder_tls
            libbase
            libbinder
            libbinder_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            liblog
            libm
            libssl
            libutils
    )
endif ()
#################### libbinder_tls ####################

#################### servicedispatcher ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/servicedispatcher.cpp
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
    add_executable(servicedispatcher ${SOURCE_FILES})

    target_include_directories(servicedispatcher
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(servicedispatcher PRIVATE
            -Werror
    )

    target_link_libraries(servicedispatcher
            android.debug_aidl-cpp
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### servicedispatcher ####################

#################### libbinder_tls_static ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportTls.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTlsUtils.cpp
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
    add_library(libbinder_tls_static SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_tls_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_tls
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libbinder_tls_static PRIVATE
            -Werror
    )

    target_link_libraries(libbinder_tls_static
            libbase
            libbinder
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            liblog
            libm
            libssl
            libutils
    )
endif ()
#################### libbinder_tls_static ####################

#################### libbinder_trusty ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportTipcAndroid.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTrusty.cpp
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
    add_library(libbinder_trusty SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_trusty
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include_trusty
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/binder/libbinder/android_vendor.34_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libbinder_trusty PRIVATE
            -Werror
    )

    target_link_libraries(libbinder_trusty
            libbase
            libbinder
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libtrusty
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libbinder_trusty ####################

#################### libbinder_on_trusty_mock ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Binder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BpBinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Debug.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/FdTrigger.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IResultReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelFileDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RecordedTransaction.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcSession.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcServer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/TextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Trace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Utils.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/logging.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/OS.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/RpcServerTrusty.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/RpcTransportTipcTrusty.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/TrustyStatus.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/socket.cpp
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
    add_library(libbinder_on_trusty_mock SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_on_trusty_mock
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/trusty/include_mock
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
    )

    # Add c/c++ flags
    target_compile_options(libbinder_on_trusty_mock PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-label
    )

    target_link_libraries(libbinder_on_trusty_mock
            libbase
            libbinder_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
            trusty_mock_headers
    )
endif ()
#################### libbinder_on_trusty_mock ####################

#################### libbinder_rpc_no_kernel ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/OS.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcTransportRaw.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Binder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/BpBinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Debug.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/FdTrigger.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IInterface.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/IResultReceiver.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ParcelFileDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RecordedTransaction.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcSession.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcServer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/RpcState.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/TextOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Trace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/Utils.cpp
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
    add_library(libbinder_rpc_no_kernel SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_rpc_no_kernel
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder
            ${ANDROID_ROOT}/frameworks/native/libs/binder/include
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
    )

    # Add c/c++ flags
    target_compile_options(libbinder_rpc_no_kernel PRIVATE
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libbinder_rpc_no_kernel
            jni_headers
            libbase
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder_rpc_no_kernel ####################
