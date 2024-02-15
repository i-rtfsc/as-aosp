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

project(libandroid_runtime)

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

#################### libandroid_runtime ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_animation_PropertyValuesHolder.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_SystemClock.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_SystemProperties.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_Trace.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_text_AndroidCharacter.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_EventLog.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_Log.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_SeempLog.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_StringBlock.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_XmlBlock.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_jar_StrictJarFile.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_app_ActivityTrigger.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_util_VirtualRefBasePtr.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/core_jni_helpers.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/jni/libandroid_runtime/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/hardware/display/IDeviceProductInfoConstants.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/AndroidRuntime.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_content_F2fsUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_content_NativeLibraryHelper.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_google_android_gles_jni_EGLImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_google_android_gles_jni_GLImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_app_Activity.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_app_ActivityThread.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_app_NativeActivity.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_app_admin_SecurityLog.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_EGL14.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_EGL15.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_EGLExt.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES10.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES10Ext.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES11.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES11Ext.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES20.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES30.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES31.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES31Ext.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_opengl_GLES32.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_database_CursorWindow.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_database_SQLiteCommon.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_database_SQLiteConnection.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_database_SQLiteGlobal.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_database_SQLiteDebug.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_graphics_GraphicBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_graphics_SurfaceTexture.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_CompositionSamplingListener.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_DisplayEventReceiver.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_InputChannel.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_InputDevice.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_InputEventReceiver.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_InputEventSender.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_InputQueue.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_KeyCharacterMap.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_KeyEvent.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_MotionEvent.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_MotionPredictor.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_PointerIcon.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_Surface.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_SurfaceControl.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_SurfaceControlHdrLayerInfoListener.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_graphics_BLASTBufferQueue.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_SurfaceSession.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_TextureView.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_TunnelModeEnabledListener.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_VelocityTracker.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_VerifiedKeyEvent.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_view_VerifiedMotionEvent.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_text_Hyphenator.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_Debug.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_GraphicsEnvironment.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HidlMemory.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HidlSupport.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HwBinder.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HwBlob.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HwParcel.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_HwRemoteBinder.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_NativeHandle.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_MemoryFile.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_MessageQueue.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_Parcel.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_PerformanceHintManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_SELinux.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_ServiceManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_SharedMemory.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_storage_StorageManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_UEventObserver.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_VintfObject.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_VintfRuntimeInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_os_incremental_IncrementalManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_net_LocalSocketImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_service_DataLoaderService.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_AssetManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_Binder.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_CharsetUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_MemoryIntArray.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_Process.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_audio_common_AidlConversion.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioDeviceAttributes.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioEffectDescriptor.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioRecord.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioSystem.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioTrackCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioTrack.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioAttributes.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioProductStrategies.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioVolumeGroups.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_AudioVolumeGroupCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_DeviceCallback.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_MediaMetricsJNI.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_MicrophoneInfo.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_midi.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_RemoteDisplay.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_media_ToneGenerator.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_Camera.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_camera2_CameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_camera2_DngCreator.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_camera2_impl_CameraExtensionJpegProcessor.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_camera2_utils_SurfaceUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_display_DisplayManagerGlobal.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_display_DisplayViewport.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_HardwareBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_OverlayProperties.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_SensorManager.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_SerialPort.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_SyncFence.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_UsbDevice.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_UsbDeviceConnection.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_UsbRequest.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_location_ActivityRecognitionHardware.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_util_FileObserver.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/poly_clip.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/util.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_ddm_DdmHandleNativeHeap.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_backup_BackupDataInput.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_backup_BackupDataOutput.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_backup_FileBackupHelperBase.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_backup_BackupHelperDispatcher.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_app_backup_FullBackup.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_content_res_ApkAssets.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_content_res_ObbScanner.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_content_res_Configuration.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_content_res_ResourceTimer.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_security_Scrypt.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_content_om_OverlayConfig.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_content_om_OverlayManagerImpl.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_net_NetworkUtilsInternal.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_ClassLoaderFactory.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_FuseAppLoop.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelAllocationStats.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelCpuBpfTracking.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelCpuTotalBpfMapReader.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelCpuUidBpfMapReader.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelSingleProcessCpuThreadReader.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_KernelSingleUidTimeReader.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_LongArrayMultiStateCounter.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_LongMultiStateCounter.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_Zygote.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_ZygoteCommandBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_os_ZygoteInit.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/com_android_internal_security_VerityUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/hwbinder/EphemeralStorage.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/fd_utils.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_input_InputWindowHandle.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_hardware_input_InputApplicationHandle.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_window_WindowInfosListener.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android_window_ScreenCapture.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/jni_common.cpp
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
    add_library(libandroid_runtime SHARED ${SOURCE_FILES})

    target_include_directories(libandroid_runtime
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/jni
            ${ANDROID_ROOT}/frameworks/base/core/jni/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/jni/libandroid_runtime/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libandroid_runtime PRIVATE
            -Wno-unused-parameter
            -Wall
            -Werror
    )

    target_link_libraries(libandroid_runtime
            android.hardware.camera.device-3.2
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            audioflinger-aidl-cpp
            audiopolicy-types-aidl-cpp
            av-types-aidl-cpp
            bionic_libc_platform_headers
            dnsproxyd_protocol_headers
            flatbuffer_headers
            libEGL
            libETC1
            libGLESv1_CM
            libGLESv2
            libGLESv3
            libPlatformProperties
            libandroid_net
            libandroidfw
            libappfuse
            libasync_safe
            libaudioclient
            libaudioclient_aidl_conversion
            libaudiofoundation
            libaudiopolicy
            libbase
            libbattery
            libbinary_parse
            libbinder
            libbinder_ndk
            libbinderthreadstateutils
            libc
            libc++
            libc++demangle
            libcamera_client
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdataloader
            libdebuggerd_client
            libdl
            libdl_android
            libdmabufinfo
            libdng_sdk
            libexpresslog_jni
            libft2
            libgif
            libgrallocusage
            libgraphicsenv
            libgui
            libgui_window_info_static
            libhardware
            libhardware_legacy
            libharfbuzz_ng
            libhidlbase
            libhostgraphics
            libhwui
            libicu
            libicui18n
            libicuuc
            libimage_io
            libimage_type_recognition
            libimg_utils
            libincfs
            libinput
            libjpeg
            libjpegdecoder
            libjpegencoder
            liblog
            libm
            libmedia
            libmedia_helper
            libmediametrics
            libmediandk
            libmeminfo
            libmemtrack
            libmemunreachable
            libminikin
            libnativebridge_lazy
            libnativedisplay
            libnativehelper_lazy
            libnativeloader_lazy
            libnativewindow
            libnetd_client
            libnetdutils
            libpdfium
            libpermission
            libpiex
            libpng
            libprocessgroup
            libprocinfo
            libscrypt_static
            libseccomp_policy
            libselinux
            libsensor
            libskia
            libsqlite
            libstatssocket_lazy
            libtextclassifier_hash_headers
            libtextclassifier_hash_static
            libtiff_directory
            libtimeinstate
            libui
            libui-types
            libultrahdr
            libusbhost
            libutils
            libvintf
            libvulkan
            libwebp-decode
            libwebp-encode
            libwuffs_mirror_release_c
            libz
            libziparchive_for_incfs
            server_configurable_flags
            spatializer-aidl-cpp
            tensorflow_headers
    )
endif ()
#################### libandroid_runtime ####################
