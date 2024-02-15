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

project(libplatformprotos-static)

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

#################### libplatformprotos-static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/proto/src/ipconnectivity.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/common/adservices_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/fledge/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/anr/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/measurement/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/media_output_enum.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/remoteprovisioner_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/settings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/tvsettings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/wearsettings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/job/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/wearservices/wearservices_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/autofill/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/contexthub/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/corenetworking/connectivity/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/debug/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/display/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/dnd/dnd_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hardware/biometrics/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hardware/sensor/assist/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/healthfitness/api/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/healthfitness/ui/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hotword/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/jank/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/audio/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/codec/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/drm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/hdr/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/midi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/outputswitcher/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/mms/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/nearby/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/net/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/neuralnetworks/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/nfc/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/os/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/display/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/job/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/location/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/service/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/service/procstats_enum.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/accessibility/accessibility_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/docsui/docsui_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/hdmi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/intelligence/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/location/location_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mediametrics/mediametrics.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mediaprovider/mediaprovider_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mobiledatadownload/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/otaupdate/updateengine_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/privacysignals/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/safetycenter/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/storage/storage_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/sysui/notification_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/tls/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/tv/tif_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/wm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/system/security/keystore2/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telecomm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/qns/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/satellite/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/uwb/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/view/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/view/inputmethod/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/wear/media/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/wifi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/connectivity/data_stall_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/devicepolicy/device_policy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/devicepolicy/device_policy_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/dnsresolver/dns_resolver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/launcher/launcher.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/network_stack.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/connectivity_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/a2dp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/hci/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/hfp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/le/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/smp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/leaudio/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/style/style_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/tethering.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/textclassifier/textclassifier_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/message/mediametrics_message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/privacy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/section.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/typedef.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/activitymanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/alarmmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/appexitinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/location_time_zone_manager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification_channel.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification_channel_group.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notificationmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/pendingintent.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/profilerinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/statusbarmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/time_zone_detector.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/window_configuration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/companion/context_sync_message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/companion/telecom.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/activityinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/clipdata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/clipdescription.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/component_name.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/configuration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/featureinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/intent.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/locale.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/locusid.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/package_item_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/pixelformat.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/point.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/rect.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/hardware/sensorprivacy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/hardware/location/context_hub_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/inputmethodservice/inputmethodservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/inputmethodservice/softinputwindow.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/binder_latency.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/locallog.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/powerprofile.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/processstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/protolog.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/media/audioattributes.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/network.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/networkcapabilities.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/networkrequest.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/aid_group.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/apdu_service_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/card_emulation.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/ndef.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/nfc_fservice_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/nfc_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfcstm/card_emulation_stm.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfcstm/nfc_service_stm.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/appbackgroundrestrictioninfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/appbatterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/backtrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batterytype.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batteryusagestats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/bundle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpu_usage.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpufreq.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpuinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/data.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/header.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/incident.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/kernelwake.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/looper.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/messagequeue.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/metadata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/pagetypeinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/patternmatcher.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/persistablebundle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/powermanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/procrank.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/ps.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/statsdata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/system_properties.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/tombstone.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/worksource.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/common.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/generation.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/global.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/secure.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/system.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/accessibility.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/accessibilitytrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/activitymanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/animationadapter.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/apphibernationservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/appstatetracker.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/background_install_control.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/biometrics.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/blobstoremanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/bluetooth_manager_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/face.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/fingerprint.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/intentresolver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/jobscheduler.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/notificationhistory.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/peopleservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/powermanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/powerstatsservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/statlogger.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/surfaceanimator.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/syncstorageengine.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/usagestatsservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/usagestatsservice_v2.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowcontainerthumbnail.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagertrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagertransitiontrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/wirelesschargerdetector.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/alarm/alarmmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/inputmethod/inputmethodmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/location/context_hub.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/vibrator/vibratormanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/adb.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/appwidget.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/battery.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/batterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/diskstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/dropbox.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/graphicsstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/network_watchlist.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/notification.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/package.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/print.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/procstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/restricted_image.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/runtime.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/sensor_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/usb.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/common.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/event_log_tags.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/log.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/quotatracker.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/textdump.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/display.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/displaycutout.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/displayinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/imefocuscontroller.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/imeinsetssourceconsumer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetsanimationcontrolimpl.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetscontroller.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssource.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssourceconsumer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssourcecontrol.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetsstate.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/remote_animation_target.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/surface.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/surfacecontrol.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/viewrootimpl.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/windowlayoutparams.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/editorinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputconnection.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputmethodeditortrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputmethodmanager.pb.cc
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
    add_library(libplatformprotos-static SHARED ${SOURCE_FILES})

    target_include_directories(libplatformprotos-static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos-static/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libplatformprotos-static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libplatformprotos-static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libplatformprotos-static ####################

#################### libplatformprotos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/proto/src/ipconnectivity.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/common/adservices_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/fledge/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/anr/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/adservices/measurement/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/media_output_enum.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/remoteprovisioner_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/settings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/tvsettings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/wearsettings_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/job/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/app/wearservices/wearservices_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/autofill/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/contexthub/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/corenetworking/connectivity/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/debug/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/display/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/dnd/dnd_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hardware/biometrics/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hardware/sensor/assist/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/healthfitness/api/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/healthfitness/ui/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/hotword/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/jank/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/audio/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/codec/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/drm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/hdr/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/midi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/media/outputswitcher/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/mms/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/nearby/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/net/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/neuralnetworks/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/nfc/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/os/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/display/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/job/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/location/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/service/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/service/procstats_enum.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/accessibility/accessibility_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/docsui/docsui_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/hdmi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/intelligence/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/location/location_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mediametrics/mediametrics.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mediaprovider/mediaprovider_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/mobiledatadownload/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/otaupdate/updateengine_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/privacysignals/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/safetycenter/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/storage/storage_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/sysui/notification_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/tls/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/tv/tif_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/wm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/system/security/keystore2/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telecomm/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/qns/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/telephony/satellite/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/uwb/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/view/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/view/inputmethod/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/wear/media/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/wifi/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/server/connectivity/data_stall_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/devicepolicy/device_policy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/devicepolicy/device_policy_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/dnsresolver/dns_resolver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/launcher/launcher.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/network_stack.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/connectivity_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/a2dp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/hci/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/hfp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/le/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/smp/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/bluetooth/leaudio/enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/style/style_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/connectivity/tethering.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/enums/stats/textclassifier/textclassifier_enums.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/proto_logging/stats/message/mediametrics_message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/privacy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/section.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/typedef.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/activitymanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/alarmmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/appexitinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/location_time_zone_manager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification_channel.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notification_channel_group.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/notificationmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/pendingintent.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/profilerinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/statusbarmanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/time_zone_detector.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/app/window_configuration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/companion/context_sync_message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/companion/telecom.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/activityinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/clipdata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/clipdescription.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/component_name.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/configuration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/featureinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/intent.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/locale.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/locusid.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/content/package_item_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/pixelformat.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/point.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/graphics/rect.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/hardware/sensorprivacy.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/hardware/location/context_hub_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/inputmethodservice/inputmethodservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/inputmethodservice/softinputwindow.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/binder_latency.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/locallog.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/powerprofile.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/processstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/internal/protolog.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/media/audioattributes.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/network.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/networkcapabilities.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/net/networkrequest.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/aid_group.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/apdu_service_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/card_emulation.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/ndef.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/nfc_fservice_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfc/nfc_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfcstm/card_emulation_stm.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/nfcstm/nfc_service_stm.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/appbackgroundrestrictioninfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/appbatterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/backtrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batterytype.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/batteryusagestats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/bundle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpu_usage.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpufreq.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/cpuinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/data.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/header.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/incident.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/kernelwake.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/looper.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/message.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/messagequeue.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/metadata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/pagetypeinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/patternmatcher.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/persistablebundle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/powermanager.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/procrank.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/ps.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/statsdata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/system_properties.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/tombstone.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/os/worksource.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/common.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/generation.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/global.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/secure.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/providers/settings/system.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/accessibility.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/accessibilitytrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/activitymanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/animationadapter.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/apphibernationservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/appstatetracker.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/background_install_control.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/biometrics.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/blobstoremanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/bluetooth_manager_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/face.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/fingerprint.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/intentresolver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/jobscheduler.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/notificationhistory.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/peopleservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/powermanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/powerstatsservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/statlogger.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/surfaceanimator.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/syncstorageengine.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/usagestatsservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/usagestatsservice_v2.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowcontainerthumbnail.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagertrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/windowmanagertransitiontrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/wirelesschargerdetector.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/alarm/alarmmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/inputmethod/inputmethodmanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/location/context_hub.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/server/vibrator/vibratormanagerservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/adb.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/appwidget.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/battery.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/batterystats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/diskstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/dropbox.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/graphicsstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/network_watchlist.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/notification.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/package.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/print.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/procstats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/restricted_image.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/runtime.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/sensor_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/service/usb.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/common.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/event_log_tags.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/log.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/quotatracker.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/util/textdump.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/display.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/displaycutout.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/displayinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/imefocuscontroller.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/imeinsetssourceconsumer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetsanimationcontrolimpl.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetscontroller.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssource.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssourceconsumer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetssourcecontrol.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/insetsstate.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/remote_animation_target.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/surface.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/surfacecontrol.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/viewrootimpl.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/windowlayoutparams.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/editorinfo.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputconnection.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputmethodeditortrace.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/core/proto/android/view/inputmethod/inputmethodmanager.pb.cc
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
    add_library(libplatformprotos SHARED ${SOURCE_FILES})

    target_include_directories(libplatformprotos
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/libplatformprotos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libplatformprotos PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libplatformprotos
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
    )
endif ()
#################### libplatformprotos ####################
