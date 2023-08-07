cmake_minimum_required(VERSION 3.5)
project(android-services)

# set(CMAKE_CXX_STANDARD 20)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_SYSTEM_COMMON true)
set(AOSP_OUT false)
set(AOSP_AV false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_HARDWARE false)

file(GLOB SOURCE_FILES
		${ANDROID_ROOT}/frameworks/base/libs/services/src/content/*.cpp
		${ANDROID_ROOT}/frameworks/base/libs/services/src/os/*.cpp
		)

include_directories(
		"${ANDROID_ROOT}/frameworks/base/libs/services"
		"${ANDROID_ROOT}/frameworks/base/libs/services/include"
		"${ANDROID_ROOT}/frameworks/native/include"
		"${ANDROID_ROOT}/frameworks/native/libs/binder/include"
		"${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
		"${ANDROID_ROOT}/frameworks/native/opengl/include"
)

file(GLOB_RECURSE HEADERS
		"${ANDROID_ROOT}/frameworks/base/libs/services/**/*.h"
		"${ANDROID_ROOT}/frameworks/base/libs/services/include/**/*.h"
		"${ANDROID_ROOT}/frameworks/native/include/**/*.h"
		"${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
		"${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
		"${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
		)

# add aosp system common
if (${AOSP_SYSTEM_COMMON})
	message("enable aosp system common")

	# add aosp system common header dir
	include_directories(
			"${ANDROID_ROOT}/system/core/include"
			"${ANDROID_ROOT}/system/core/libcutils/include"
			"${ANDROID_ROOT}/system/core/libprocessgroup/include"
			"${ANDROID_ROOT}/system/core/libsystem/include"
			"${ANDROID_ROOT}/system/core/libutils/include"
			"${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
			"${ANDROID_ROOT}/system/libbase/include"
			"${ANDROID_ROOT}/system/logging/liblog/include"
	)

	# add aosp system common header file
	file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
			"${ANDROID_ROOT}/system/core/include/**/*.h"
			"${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
			"${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
			"${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
			"${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
			"${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
			"${ANDROID_ROOT}/system/libbase/include/**/*.h"
			"${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
			)
	# append aosp system common header
	list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()

# add aosp system
if (${AOSP_SYSTEM})
	message("enable aosp system")

	# add aosp system header dir
	include_directories(
			"${ANDROID_ROOT}/system/media/audio/include"
	)

	# add aosp system header file
	file(GLOB_RECURSE HEADERS_SYSTEM
			"${ANDROID_ROOT}/system/media/audio/include/**/*.h"
			)
	# append aosp system header
	list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()

# add aosp av
if (${AOSP_AV})
	message("enable aosp av")

	# add aosp av header dir
	include_directories(
			"${ANDROID_ROOT}/frameworks/av/include"
	)

	# add aosp av header file
	file(GLOB_RECURSE HEADERS_AV
			"${ANDROID_ROOT}/frameworks/av/include/**/*.h"
			)
	# append aosp av header
	list(APPEND HEADERS ${HEADERS_AV})
endif ()

# add aosp out
if (${AOSP_OUT})
	message("enable aosp out")

	# add aosp out header dir
	include_directories(
			"${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/services/libservices/android_arm64_armv8-a-branchprot_shared/gen/aidl"
			"${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
	)

	# add aosp out header file
	file(GLOB_RECURSE HEADERS_OUT
			"${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/services/libservices/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
			"${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
			)
	# append aosp out header
	list(APPEND HEADERS ${HEADERS_OUT})
endif ()


# add aosp external
if (${AOSP_BIONIC})
	message("enable aosp bionic")

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
	message("enable aosp external")

	# add aosp external header dir
	include_directories(
			"${ANDROID_ROOT}/external/fmtlib/include"
			"${ANDROID_ROOT}/external/libcxx/include"
			"${ANDROID_ROOT}/external/libcxxabi/include"
	)

	# add aosp external header file
	file(GLOB_RECURSE HEADERS_EXTERNAL
			"${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
			"${ANDROID_ROOT}/external/libcxx/include/**/*.h"
			"${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
			)
	# append aosp external header
	list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()

# add aosp hardware
if (${AOSP_HARDWARE})
	message("enable aosp hardware")

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

add_executable(
		android-services
		${SOURCE_FILES}
		${HEADERS}
)
