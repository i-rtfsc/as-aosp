cmake_minimum_required(VERSION 3.5)
project(Installd)

#set(CMAKE_CXX_STANDARD 20)
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_FWK_BASE true)
set(AOSP_FWK_NATIVE true)
set(AOSP_SYSTEM_COMMON true)
set(AOSP_FWK_AV false)
set(AOSP_SYSTEM false)
set(AOSP_ART false)
set(AOSP_BIONIC false)
set(AOSP_EXTERNAL false)
set(AOSP_PACKAGES false)
set(AOSP_BOOTABLE false)
set(AOSP_HARDWARE false)
set(AOSP_VENDOR false)
set(AOSP_OUT false)
set(AOSP_OTHER false)

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/*.c
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/*.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/android/os/*.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/installd/tests/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/native/cmds/installd"
        "${ANDROID_ROOT}/frameworks/native/cmds/installd/tests"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/native/cmds/installd/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/cmds/installd/tests/**/*.h"
        )


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/cmds/fbase_defrag/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            )
    # append aosp fwk_native header
    list(APPEND HEADERS ${HEADERS_FWK_NATIVE})
endif ()


# add aosp system_common
if (${AOSP_SYSTEM_COMMON})
    message("${PROJECT_NAME}, enable aosp system_common")

    # add aosp system_common header dir
    include_directories(
            "${ANDROID_ROOT}/system/core/fs_mgr/include"
            "${ANDROID_ROOT}/system/core/fs_mgr/include_fstab"
            "${ANDROID_ROOT}/system/core/fs_mgr/libdm/include"
            "${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/include"
            "${ANDROID_ROOT}/system/core/fs_mgr/libfs_avb/include"
            "${ANDROID_ROOT}/system/core/fs_mgr/liblp/include"
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/logging/logwrapper/include"
    )

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/fs_mgr/include/**/*.h"
            "${ANDROID_ROOT}/system/core/fs_mgr/include_fstab/**/*.h"
            "${ANDROID_ROOT}/system/core/fs_mgr/libdm/include/**/*.h"
            "${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/include/**/*.h"
            "${ANDROID_ROOT}/system/core/fs_mgr/libfs_avb/include/**/*.h"
            "${ANDROID_ROOT}/system/core/fs_mgr/liblp/include/**/*.h"
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/logwrapper/include/**/*.h"
            )
    # append aosp system_common header
    list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()


# add aosp fwk_av
if (${AOSP_FWK_AV})
    message("${PROJECT_NAME}, enable aosp fwk_av")

    # add aosp fwk_av header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/av/include"
    )

    # add aosp fwk_av header file
    file(GLOB_RECURSE HEADERS_FWK_AV
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            )
    # append aosp fwk_av header
    list(APPEND HEADERS ${HEADERS_FWK_AV})
endif ()


# add aosp system
if (${AOSP_SYSTEM})
    message("${PROJECT_NAME}, enable aosp system")

    # add aosp system header dir
    include_directories(
            "${ANDROID_ROOT}/system/apex/apexd"
            "${ANDROID_ROOT}/system/extras/libfscrypt/include"
            "${ANDROID_ROOT}/system/extras/verity/include"
            "${ANDROID_ROOT}/system/libziparchive/include"
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/security/ondevice-signing/include"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/apex/apexd/**/*.h"
            "${ANDROID_ROOT}/system/extras/libfscrypt/include/**/*.h"
            "${ANDROID_ROOT}/system/extras/verity/include/**/*.h"
            "${ANDROID_ROOT}/system/libziparchive/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/security/ondevice-signing/include/**/*.h"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include/**/*.h"
            )
    # append aosp system header
    list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp bionic
if (${AOSP_BIONIC})
    message("${PROJECT_NAME}, enable aosp bionic")

    # add aosp bionic header dir
    include_directories(
            "${ANDROID_ROOT}/bionic/libc/async_safe/include"
            "${ANDROID_ROOT}/bionic/libc/include"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64"
            "${ANDROID_ROOT}/bionic/libc/system_properties/include"
    )

    # add aosp bionic header file
    file(GLOB_RECURSE HEADERS_BIONIC
            "${ANDROID_ROOT}/bionic/libc/async_safe/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/system_properties/include/**/*.h"
            )
    # append aosp bionic header
    list(APPEND HEADERS ${HEADERS_BIONIC})
endif ()


# add aosp external
if (${AOSP_EXTERNAL})
    message("${PROJECT_NAME}, enable aosp external")

    # add aosp external header dir
    include_directories(
            "${ANDROID_ROOT}/external/avb"
            "${ANDROID_ROOT}/external/boringssl/src/include"
            "${ANDROID_ROOT}/external/e2fsprogs/lib"
            "${ANDROID_ROOT}/external/e2fsprogs/lib/uuid"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/selinux/libselinux/include"
            "${ANDROID_ROOT}/external/sqlite/android"
            "${ANDROID_ROOT}/external/sqlite/dist"
            "${ANDROID_ROOT}/external/tinyxml2"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/avb/**/*.h"
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/e2fsprogs/lib/**/*.h"
            "${ANDROID_ROOT}/external/e2fsprogs/lib/uuid/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            "${ANDROID_ROOT}/external/selinux/libselinux/include/**/*.h"
            "${ANDROID_ROOT}/external/sqlite/android/**/*.h"
            "${ANDROID_ROOT}/external/sqlite/dist/**/*.h"
            "${ANDROID_ROOT}/external/tinyxml2/**/*.h"
            "${ANDROID_ROOT}/external/zlib/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()


# add aosp packages
if (${AOSP_PACKAGES})
    message("${PROJECT_NAME}, enable aosp packages")

    # add aosp packages header dir
    include_directories(
            "${ANDROID_ROOT}/packages/modules/Virtualization/microdroid/payload/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Virtualization/microdroid/payload/include/**/*.h"
            )
    # append aosp packages header
    list(APPEND HEADERS ${HEADERS_PACKAGES})
endif ()


# add aosp hardware
if (${AOSP_HARDWARE})
    message("${PROJECT_NAME}, enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
            "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


# add aosp out
if (${AOSP_OUT})
    message("${PROJECT_NAME}, enable aosp out")

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/installd/android_arm64_armv8-a-branchprot/gen/aidl/android/os/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/installd/android_arm64_armv8-a-branchprot/gen/aidl/android/os/storage/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/libinstalld/android_arm64_armv8-a-branchprot_static/gen/aidl/android/os/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/libinstalld/android_arm64_armv8-a-branchprot_static/gen/aidl/android/os/storage/*.cpp
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/java/libincremental_aidl-cpp/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/installd/android_arm64_armv8-a-branchprot/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/libinstalld/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Virtualization/microdroid/payload/lib_microdroid_metadata_proto/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Virtualization/microdroid/payload/lib_microdroid_metadata_proto/android_arm64_armv8-a-branchprot_static/gen/proto/packages/modules/Virtualization/microdroid/payload"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/apexd/sysprop/libcom.android.sysprop.apex/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_manifest_proto/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_manifest_proto/android_arm64_armv8-a-branchprot_static/gen/proto/system/apex/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_session_state_proto/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_session_state_proto/android_arm64_armv8-a-branchprot_static/gen/proto/system/apex/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/vold/libvold_binder/android_arm64_armv8-a-branchprot_static/gen/aidl"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/java/libincremental_aidl-cpp/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/installd/android_arm64_armv8-a-branchprot/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/cmds/installd/libinstalld/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Virtualization/microdroid/payload/lib_microdroid_metadata_proto/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Virtualization/microdroid/payload/lib_microdroid_metadata_proto/android_arm64_armv8-a-branchprot_static/gen/proto/packages/modules/Virtualization/microdroid/payload/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/apexd/sysprop/libcom.android.sysprop.apex/android_arm64_armv8-a-branchprot_static/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_manifest_proto/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_manifest_proto/android_arm64_armv8-a-branchprot_static/gen/proto/system/apex/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_session_state_proto/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/apex/proto/lib_apex_session_state_proto/android_arm64_armv8-a-branchprot_static/gen/proto/system/apex/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/vold/libvold_binder/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        Installd
        ${SOURCE_FILES}
        ${HEADERS}
)
