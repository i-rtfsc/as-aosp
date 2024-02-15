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

project(libservices.core)

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

#################### libservices.core ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/JavaRef.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/NativeCallbackThread.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/BroadcastRadioService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/Tuner.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/TunerCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/convert.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/BroadcastRadio/regions.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/stats/SurfaceFlingerPuller.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/tvinput/BufferProducerThread.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/tvinput/JTvInputHal.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/tvinput/TvInputHal_hidl.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_adb_AdbDebuggingManager.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_am_BatteryStatsService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_biometrics_SurfaceToNativeHandleConverter.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_ConsumerIrService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_companion_virtual_InputController.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_devicepolicy_CryptoTestHelper.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_display_DisplayControl.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_connectivity_Vpn.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_gpu_GpuService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_HardwarePropertiesManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_input_InputManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_lights_LightsService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_location_GnssLocationProvider.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_locksettings_SyntheticPasswordManager.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_power_PowerManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_powerstats_PowerStatsService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_hint_HintManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_SerialService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_soundtrigger_middleware_AudioSessionProviderImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_soundtrigger_middleware_ExternalCaptureStateTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_stats_pull_StatsPullAtomService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_storage_AppFuseBridge.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_SystemClockTime.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_SystemServer.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_tv_TvUinputBridge.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_tv_TvInputHal.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_vr_VrManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_UsbAlsaJackDetector.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_UsbAlsaMidiDevice.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_UsbDeviceManager.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_UsbDescriptorParser.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_UsbHostManager.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_vibrator_VibratorController.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_vibrator_VibratorManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_PersistentDataBlockService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_am_LowMemDetector.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_pm_PackageManagerShellCommandDataLoader.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_TiktapService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_sensor_SensorService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_wm_TaskFpsCallbackController.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_activityTriggerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_SmartSceneManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/onload.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_am_CachedAppOptimizer.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_app_GameManagerService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_DeviceStateService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/core/jni/com_android_server_DeviceControl.cpp
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
    add_library(libservices.core SHARED ${SOURCE_FILES})

    target_include_directories(libservices.core
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/services/core/jni
    )

    # Add c/c++ flags
    target_compile_options(libservices.core PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -DFLYME_SUPPORT_LTPO_V2
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libservices.core
            android.frameworks.schedulerservice-1.0
            android.frameworks.sensorservice-V1-ndk
            android.frameworks.sensorservice-1.0
            android.frameworks.stats-V2-ndk
            android.frameworks.stats-1.0
            android.hardware.audio.common-2.0
            android.hardware.broadcastradio-1.0
            android.hardware.broadcastradio-1.1
            android.hardware.broadcastradio-common-utils-1x-lib
            android.hardware.common.fmq-V1-ndk
            android.hardware.contexthub-1.0
            android.hardware.gnss-V3-cpp
            android.hardware.gnss.measurement_corrections-1.0
            android.hardware.gnss-1.0
            android.hardware.gnss-1.1
            android.hardware.gnss-2.0
            android.hardware.gnss-2.1
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.2
            android.hardware.graphics.mapper-4.0
            android.hardware.input.processor-V1-ndk
            android.hardware.ir-1.0
            android.hardware.light-2.0
            android.hardware.memtrack-V1-ndk
            android.hardware.power-V4-cpp
            android.hardware.power.stats-V1-ndk
            android.hardware.power.stats-1.0
            android.hardware.power-1.0
            android.hardware.power-1.1
            android.hardware.power-1.2
            android.hardware.power-1.3
            android.hardware.thermal-V1-ndk
            android.hardware.thermal-1.0
            android.hardware.tv.input-V1-ndk
            android.hardware.tv.input-1.0
            android.hardware.vibrator-V2-cpp
            android.hardware.vibrator-1.0
            android.hardware.vibrator-1.1
            android.hardware.vibrator-1.2
            android.hardware.vibrator-1.3
            android.hardware.vr-1.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            android.hidl.token-1.0-utils
            android.media.audio.common.types-V1-ndk
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-cpp
            android.system.suspend.control.internal-cpp
            bionic_libc_platform_headers
            libEGL
            libGLESv2
            libPlatformProperties
            libadb_pairing_connection
            libadb_pairing_server
            libaidlcommonsupport
            libandroid
            libandroid_runtime
            libandroidfw
            libappfuse
            libaudioclient
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdataloader
            libdl
            libfmq
            libgralloctypes
            libgraphicsenv
            libgui
            libhardware
            libhardware_legacy
            libhidlbase
            libhidlmemory
            libhwui
            libincfs
            libinput
            libinputflinger
            libinputflinger_base
            libinputservice
            liblog
            libm
            libmeminfo
            libmemtrackproxy
            libmtp
            libnativehelper
            libnativewindow
            libnetutils
            libpowermanager
            libprocessgroup
            libpsi
            libschedulerservicehidl
            libsensorservice
            libsensorserviceaidl
            libsensorservicehidl
            libservices.core-gnss
            libstatshidl
            libstatslog
            libstatspull
            libstatssocket
            libtimestats_atoms_proto
            libtinyalsa
            libui
            libusbhost
            libutils
            libvibratorservice
            server_configurable_flags
            service.incremental
            vendor.meizu.hardware.tiktap_hidl-1.0
    )
endif ()
#################### libservices.core ####################
