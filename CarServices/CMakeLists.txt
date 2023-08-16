cmake_minimum_required(VERSION 3.5)
project(libcar_service)

#set(ANDROID_ROOT "/home/solo/code/flyme")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB
        SOURCE_FILES
        ${ANDROID_ROOT}/packages/services/Car/service/jni/**/*.cpp
        )

add_executable(
        libcar_service
        ${SOURCE_FILES}
)