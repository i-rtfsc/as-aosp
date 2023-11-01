SOURCE="/data/code/aosp"
DEST="/home/solo/code/aosp"

mkdir -p $DEST/
ln -s $SOURCE/art $DEST/

mkdir -p $DEST/
ln -s $SOURCE/bionic $DEST/

mkdir -p $DEST/bootable/
ln -s $SOURCE/bootable/libbootloader $DEST/bootable/

mkdir -p $DEST/bootable/
ln -s $SOURCE/bootable/recovery $DEST/bootable/

mkdir -p $DEST/external/
ln -s $SOURCE/external/tensorflow $DEST/external/

mkdir -p $DEST/frameworks/
ln -s $SOURCE/frameworks/av $DEST/frameworks/

mkdir -p $DEST/frameworks/
ln -s $SOURCE/frameworks/base $DEST/frameworks/

mkdir -p $DEST/frameworks/hardware/
ln -s $SOURCE/frameworks/hardware/interfaces $DEST/frameworks/hardware/

mkdir -p $DEST/frameworks/libs/
ln -s $SOURCE/frameworks/libs/modules-utils $DEST/frameworks/libs/

mkdir -p $DEST/frameworks/libs/
ln -s $SOURCE/frameworks/libs/net $DEST/frameworks/libs/

mkdir -p $DEST/frameworks/libs/
ln -s $SOURCE/frameworks/libs/systemui $DEST/frameworks/libs/

mkdir -p $DEST/frameworks/
ln -s $SOURCE/frameworks/native $DEST/frameworks/

mkdir -p $DEST/frameworks/opt/car/
ln -s $SOURCE/frameworks/opt/car/services $DEST/frameworks/opt/car/

mkdir -p $DEST/hardware/
ln -s $SOURCE/hardware/interfaces $DEST/hardware/

mkdir -p $DEST/hardware/
ln -s $SOURCE/hardware/libhardware $DEST/hardware/

mkdir -p $DEST/hardware/
ln -s $SOURCE/hardware/libhardware_legacy $DEST/hardware/

mkdir -p $DEST/hardware/qcom/
ln -s $SOURCE/hardware/qcom/display $DEST/hardware/qcom/

mkdir -p $DEST/
ln -s $SOURCE/libcore $DEST/

mkdir -p $DEST/
ln -s $SOURCE/libnativehelper $DEST/

mkdir -p $DEST/packages/apps/
ln -s $SOURCE/packages/apps/Settings $DEST/packages/apps/

mkdir -p $DEST/packages/modules/
ln -s $SOURCE/packages/modules/Connectivity $DEST/packages/modules/

mkdir -p $DEST/packages/modules/
ln -s $SOURCE/packages/modules/ExtServices $DEST/packages/modules/

mkdir -p $DEST/packages/modules/
ln -s $SOURCE/packages/modules/NeuralNetworks $DEST/packages/modules/

mkdir -p $DEST/packages/modules/
ln -s $SOURCE/packages/modules/Wifi $DEST/packages/modules/

mkdir -p $DEST/packages/services/
ln -s $SOURCE/packages/services/Car $DEST/packages/services/

mkdir -p $DEST/system/
ln -s $SOURCE/system/core $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/libbase $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/libfmq $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/libhidl $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/libhwbinder $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/libvintf $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/logging $DEST/system/

mkdir -p $DEST/system/
ln -s $SOURCE/system/media $DEST/system/

mkdir -p $DEST/system/memory/
ln -s $SOURCE/system/memory/lmkd $DEST/system/memory/

mkdir -p $DEST/system/
ln -s $SOURCE/system/server_configurable_flags $DEST/system/

mkdir -p $DEST/system/tools/
ln -s $SOURCE/system/tools/aidl $DEST/system/tools/

mkdir -p $DEST/system/tools/
ln -s $SOURCE/system/tools/hidl $DEST/system/tools/
