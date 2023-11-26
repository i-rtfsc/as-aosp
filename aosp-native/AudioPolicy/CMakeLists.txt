cmake_minimum_required(VERSION 3.5)
project(AudioPolicy)

#set(CMAKE_CXX_STANDARD 20)
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_FWK_BASE true)
set(AOSP_FWK_NATIVE true)
set(AOSP_SYSTEM_COMMON true)
set(AOSP_FWK_AV true)
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
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/src/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/common/src/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/src/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/parameter-framework/plugin/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/src/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/wrapper/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/enginedefault/src/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/fuzzer/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/managerdefault/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/service/*.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audiopolicy/tests/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/common"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/common/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/tests"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/interface"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/interface"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/parameter-framework/plugin"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/wrapper"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/wrapper/include"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/enginedefault"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/fuzzer"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/managerdefault"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/service"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/tests"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/common/managerdefinitions/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/common/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/common/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/config/tests/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engine/interface/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/interface/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/parameter-framework/plugin/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/wrapper/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/engineconfigurable/wrapper/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/enginedefault/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/fuzzer/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/managerdefault/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/service/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/services/audiopolicy/tests/**/*.h"
        )


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin"
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_batterystats"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include"
            "${ANDROID_ROOT}/frameworks/native/libs/sensorprivacy/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_batterystats/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/sensorprivacy/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
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
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include"
            "${ANDROID_ROOT}/system/libhwbinder/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
    )

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/system/libhwbinder/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
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
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiohal/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiousecasevalidation/include"
            "${ANDROID_ROOT}/frameworks/av/media/libeffects/config/include"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include"
            "${ANDROID_ROOT}/frameworks/av/media/libheadtracking/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include"
            "${ANDROID_ROOT}/frameworks/av/media/libshmem/include"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include"
            "${ANDROID_ROOT}/frameworks/av/media/utils/include"
            "${ANDROID_ROOT}/frameworks/av/services/audioflinger"
    )

    # add aosp fwk_av header file
    file(GLOB_RECURSE HEADERS_FWK_AV
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiohal/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiousecasevalidation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libeffects/config/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libheadtracking/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libshmem/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/utils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/services/audioflinger/**/*.h"
            )
    # append aosp fwk_av header
    list(APPEND HEADERS ${HEADERS_FWK_AV})
endif ()


# add aosp system
if (${AOSP_SYSTEM})
    message("${PROJECT_NAME}, enable aosp system")

    # add aosp system header dir
    include_directories(
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/media/audio_utils/include"
            "${ANDROID_ROOT}/system/tools/xsdc/utils/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio_utils/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/xsdc/utils/include/**/*.h"
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
            "${ANDROID_ROOT}/external/eigen"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libxml2/include"
            "${ANDROID_ROOT}/external/parameter-framework/support/android/parameter"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/include"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/log/include"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/utility"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/eigen/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libxml2/include/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/support/android/parameter/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/include/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/log/include/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/utility/**/*.h"
            "${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()


# add aosp packages
if (${AOSP_PACKAGES})
    message("${PROJECT_NAME}, enable aosp packages")

    # add aosp packages header dir
    include_directories(
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
            )
    # append aosp packages header
    list(APPEND HEADERS ${HEADERS_PACKAGES})
endif ()


# add aosp hardware
if (${AOSP_HARDWARE})
    message("${PROJECT_NAME}, enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/7.0/enums/include"
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/all-versions/default/include"
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/all-versions/util/include"
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/7.0/enums/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/all-versions/default/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/audio/common/all-versions/util/include/**/*.h"
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

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/av-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libshmem/shared-file-region-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/sensorprivacy/libsensorprivacy/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/audio/7.0/config/audio_policy_configuration_V7_0_enums/gen/cpp/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/av-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libshmem/shared-file-region-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/sensorprivacy/libsensorprivacy/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/audio/7.0/config/audio_policy_configuration_V7_0_enums/gen/cpp/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        AudioPolicy
        ${SOURCE_FILES}
        ${HEADERS}
)
