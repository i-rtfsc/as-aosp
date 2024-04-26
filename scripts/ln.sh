projects=(
    "art"
    "bionic"
    "bootable/libbootloader"
    "bootable/recovery"
    "external/libcxx"
    "external/libcxxabi"
    "frameworks/av"
    "frameworks/base"
    "frameworks/hardware/interfaces"
    "frameworks/libs/modules-utils"
    "frameworks/libs/net"
    "frameworks/libs/systemui"
    "frameworks/native"
    "frameworks/opt/car/services"
    "hardware/interfaces"
    "hardware/libhardware"
    "hardware/libhardware_legacy"
    "hardware/qcom/display"
    "libcore"
    "libnativehelper"
    "packages/apps/Car/Settings"
    "packages/apps/Car/SystemUI"
    "packages/apps/Settings"
    "packages/modules/ConfigInfrastructure"
    "packages/modules/Connectivity"
    "packages/modules/ExtServices"
    "packages/modules/NeuralNetworks"
    "packages/modules/Permission"
    "packages/modules/Wifi"
    "packages/services/Car"
    "system/core"
    "system/gsid"
    "system/libbase"
    "system/libfmq"
    "system/libhidl"
    "system/libhwbinder"
    "system/libvintf"
    "system/logging"
    "system/media"
    "system/memory/lmkd"
    "system/server_configurable_flags"
    "system/tools/aidl"
    "system/tools/hidl"
    "tools/platform-compat"
)

SOURCE="/data/code/aosp"
DEST="/home/solo/code/aosp"

for project in "${projects[@]}"; do
    path=$DEST/$project
    if [ ! -d "$path" ]; then
        parent_dir=$(dirname "$path")
        if [ ! -d "$parent_dir" ]; then
            mkdir -p "$parent_dir"
        fi
        ln -sf $SOURCE/$project $path
    fi
done