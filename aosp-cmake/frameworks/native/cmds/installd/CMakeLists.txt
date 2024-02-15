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

project(otapreopt_chroot)

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

#################### libotapreoptparameters ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/otapreopt_parameters.cpp
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
    add_library(libotapreoptparameters SHARED ${SOURCE_FILES})

    target_include_directories(libotapreoptparameters
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/installd
    )

    # Add c/c++ flags
    target_compile_options(libotapreoptparameters PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libotapreoptparameters
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libprocessgroup
            libutils
    )
endif ()
#################### libotapreoptparameters ####################

#################### libinstalld ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CacheItem.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CacheTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CrateManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/InstalldNativeService.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/QuotaUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/SysTrace.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/dexopt.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/execv_helper.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/globals.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/restorable_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/run_dex2oat.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/unique_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/slice.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/SliceLogger.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils_default.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/view_compiler.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzFileDedup.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzSoDedupSqliteManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzAppCompress.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzAppCompressSqliteManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzCowHelper.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/mzcow.c
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/android/os/IInstalld.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/android/os/IMzSpaceOptListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/libinstalld/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/CreateAppDataArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/libinstalld/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/CreateAppDataResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/libinstalld/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/ReconcileSdkDataArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/libinstalld/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/storage/CrateMetadata.cpp
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
    add_library(libinstalld SHARED ${SOURCE_FILES})

    target_include_directories(libinstalld
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/installd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/libinstalld/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libinstalld PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libinstalld
            libasync_safe
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libext2_uuid
            liblog
            liblogwrap
            libm
            libprocessgroup
            libselinux
            libsqlite
            libssl
            libutils
            libziparchive
            server_configurable_flags
    )
endif ()
#################### libinstalld ####################

#################### installd ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CacheItem.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CacheTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/CrateManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/InstalldNativeService.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/QuotaUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/SysTrace.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/dexopt.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/execv_helper.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/globals.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/restorable_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/run_dex2oat.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/unique_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/slice.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/SliceLogger.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils_default.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/view_compiler.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzFileDedup.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzSoDedupSqliteManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzAppCompress.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzAppCompressSqliteManager.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/MzCowHelper.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/mzcow.c
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/android/os/IInstalld.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/android/os/IMzSpaceOptListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/installd/android_${OUT_ARCH_CPU}/gen/aidl/android/os/CreateAppDataArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/installd/android_${OUT_ARCH_CPU}/gen/aidl/android/os/CreateAppDataResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/installd/android_${OUT_ARCH_CPU}/gen/aidl/android/os/ReconcileSdkDataArgs.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/installd/android_${OUT_ARCH_CPU}/gen/aidl/android/os/storage/CrateMetadata.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/installd.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/roothandler.cpp
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
    add_executable(installd ${SOURCE_FILES})

    target_include_directories(installd
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/installd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/cmds/installd/installd/android_${OUT_ARCH_CPU}/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(installd PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(installd
            libasync_safe
            libavb_user
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdiskusage
            libdl
            libext2_uuid
            libfbase_defrag
            libfbase_defrag_service
            libfs_mgr
            liblog
            liblogwrap
            libm
            libprocessgroup
            libselinux
            libsqlite
            libssl
            libutils
            libziparchive
            server_configurable_flags
    )
endif ()
#################### installd ####################

#################### otapreopt ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/dexopt.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/execv_helper.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/globals.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/otapreopt.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/otapreopt_utils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/restorable_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/run_dex2oat.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/unique_file.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/slice.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/SliceLogger.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/utils_default.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/view_compiler.cpp
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
    add_executable(otapreopt ${SOURCE_FILES})

    target_include_directories(otapreopt
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/installd
    )

    # Add c/c++ flags
    target_compile_options(otapreopt PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(otapreopt
            libasync_safe
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdiskusage
            libdl
            libext2_uuid
            liblog
            liblogwrap
            libm
            libotapreoptparameters
            libprocessgroup
            libselinux
            libutils
            libziparchive
            server_configurable_flags
    )
endif ()
#################### otapreopt ####################

#################### otapreopt_chroot ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/otapreopt_chroot.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/otapreopt_utils.cpp
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
    add_executable(otapreopt_chroot ${SOURCE_FILES})

    target_include_directories(otapreopt_chroot
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/installd
    )

    # Add c/c++ flags
    target_compile_options(otapreopt_chroot PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(otapreopt_chroot
            lib_apex_manifest_proto
            lib_apex_session_state_proto
            lib_microdroid_metadata_proto
            libapex
            libavb
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcrypto
            libcutils
            libdl
            libdm
            libext2_uuid
            liblog
            liblogwrap
            libm
            libprotobuf-cpp-full
            libselinux
            libsigningutils
            libtinyxml2
            libutils
            libverity_tree
            libvold_binder
            libziparchive
    )
endif ()
#################### otapreopt_chroot ####################
