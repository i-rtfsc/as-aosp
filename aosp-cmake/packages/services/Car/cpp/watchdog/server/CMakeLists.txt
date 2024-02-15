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

project(carwatchdogd)

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

#################### libwatchdog_process_service ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/WatchdogProcessService.cpp
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
    add_library(libwatchdog_process_service SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_process_service
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_process_service PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_process_service
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libutils
            libvhalclient
            libwatchdog_binder_utils
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_process_service ####################

#################### libwatchdog_perf_service ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/WatchdogPerfService.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/IoOveruseConfigs.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/IoOveruseMonitor.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/PerformanceProfiler.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/OveruseConfigurationXmlHelper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/ProcDiskStatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/ProcStatCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/UidCpuStatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/UidIoStatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/UidProcStatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/UidStatsCollector.cpp
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
    add_library(libwatchdog_perf_service SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_perf_service
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_perf_service PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_perf_service
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libprocessgroup
            libtinyxml2
            libutils
            libvhalclient
            libwatchdog_binder_utils
            libwatchdog_package_info_resolver
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_perf_service ####################

#################### libwatchdog_package_info_resolver ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/PackageInfoResolver.cpp
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
    add_library(libwatchdog_package_info_resolver SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_package_info_resolver
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_package_info_resolver PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_package_info_resolver
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libutils
            libvhalclient
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_package_info_resolver ####################

#################### libwatchdog_binder_utils ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/AIBinderDeathRegistrationWrapper.cpp
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
    add_library(libwatchdog_binder_utils SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_binder_utils
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_binder_utils PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_binder_utils
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libutils
            libvhalclient
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_binder_utils ####################

#################### carwatchdogd ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/main.cpp
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
    add_executable(carwatchdogd ${SOURCE_FILES})

    target_include_directories(carwatchdogd
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src
    )

    # Add c/c++ flags
    target_compile_options(carwatchdogd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(carwatchdogd
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libprocessgroup
            libtinyxml2
            libutils
            libvhalclient
            libwatchdog_binder_mediator
            libwatchdog_binder_utils
            libwatchdog_package_info_resolver
            libwatchdog_perf_service
            libwatchdog_process_service
            libwatchdog_properties
            libwatchdog_service_manager
            packagemanager_aidl-cpp
    )
endif ()
#################### carwatchdogd ####################

#################### libwatchdog_service_manager ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/ServiceManager.cpp
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
    add_library(libwatchdog_service_manager SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_service_manager
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_service_manager PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_service_manager
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libprocessgroup
            libtinyxml2
            libutils
            libvhalclient
            libwatchdog_binder_mediator
            libwatchdog_binder_utils
            libwatchdog_package_info_resolver
            libwatchdog_perf_service
            libwatchdog_process_service
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_service_manager ####################

#################### libwatchdog_binder_mediator ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/ThreadPriorityController.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/WatchdogBinderMediator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/WatchdogInternalHandler.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src/WatchdogServiceHelper.cpp
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
    add_library(libwatchdog_binder_mediator SHARED ${SOURCE_FILES})

    target_include_directories(libwatchdog_binder_mediator
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/server
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/server/sysprop/libwatchdog_properties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_binder_mediator PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libwatchdog_binder_mediator
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libprocessgroup
            libtinyxml2
            libutils
            libvhalclient
            libwatchdog_binder_utils
            libwatchdog_perf_service
            libwatchdog_process_service
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_binder_mediator ####################
