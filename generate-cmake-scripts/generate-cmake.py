#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# -*- encoding: utf-8 -*-
#
# Copyright (c) 2023 anqi.huang@outlook.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import optparse
import os

from string import Template

# 生成cmake文件模板
cmake_template = """cmake_minimum_required(VERSION 3.5)
project(${project_name})

set(CMAKE_CXX_STANDARD 17)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB SOURCE_FILES${code_files})

include_directories(${include_directories})

file (GLOB_RECURSE HEADERS${header_files})

add_executable(
        ${project_name}
        ${SOURCE_FILES}
        ${HEADERS}
)
"""

# 生成cmake文件模板
gradle_template = """def androidRoot = rootProject.ext.aosp.root

ext {
    BUILD_NATIVE_ROOT = androidRoot
}

apply from: "${rootDir}/scripts/asop-native-build.gradle"
"""

def parseargs():
    usage = "usage: %prog [options] arg1 arg2"
    parser = optparse.OptionParser(usage=usage)

    buildoptiongroup = optparse.OptionGroup(parser, "generate cmake file")

    buildoptiongroup.add_option("-r", "--root", dest="root",
                                help="root dir", default="/home/solo/code/flyme")
    buildoptiongroup.add_option("-p", "--project", dest="project",
                                help="project name(android_runtime,android_services,inputflinger,surfaceflinger,media), or aosp-native[all projects]",
                                default="aosp-native")
    buildoptiongroup.add_option("-f", "--full", dest="full",
                                help="make full code dir", default="0")


    parser.add_option_group(buildoptiongroup)

    (options, args) = parser.parse_args()

    return (options, args)


def file_is_code(file):
    codes = [".c", ".cpp", ".cc"]
    real = False
    for code in codes:
        if file.endswith(code):
            real = True
            break

    return real


def file_is_headers(file):
    headers = ["h", "hpp"]
    real = False
    for header in headers:
        if file.endswith(header):
            real = True
            break

    return real


def fast_scandir(dirname):
    folders = [f.path for f in os.scandir(dirname) if f.is_dir()]
    for dirname in list(folders):
        folders.extend(fast_scandir(dirname))

    return folders


def write_text(file, text):
    print(file)
    directory = os.path.dirname(file)
    if not os.path.exists(directory):
        os.makedirs(directory)

    with open(file, 'w') as f:
        print("[generate]: " + f.name)
        f.write(text)


def work(project_name, root, project_list):
    code_dirs = []
    header_dirs = []
    include_dirs = []

    for project_dir in project_list:
        project_dirs = fast_scandir(os.path.join(root, project_dir))
        project_dirs.append(os.path.join(root, project_dir))
        project_dirs.sort()

        for dir in project_dirs:
            if ".git" in dir or "x86" in dir:
                continue

            # 代码文件夹

            for file in os.listdir(dir):
                if file.endswith(".c"):
                    code_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "/*.c\n")
                    # include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")
                elif file.endswith(".cc"):
                    code_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "/*.cc\n")
                    # include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")
                elif file.endswith(".cpp"):
                    code_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "/*.cpp\n")
                    # include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")

            # 头文件文件夹
            for file in os.listdir(dir):

                if os.path.isdir(os.path.join(dir)):
                    if "include" in os.path.basename(dir):
                        include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")

                if file.endswith(".h"):
                    header_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "/*.h\n")
                    include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")
                elif file.endswith(".hpp"):
                    header_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "/*.hpp\n")
                    include_dirs.append("        ${ANDROID_ROOT}/" + os.path.relpath(dir, root) + "\n")


    # 去重
    code_dirs = list(dict.fromkeys(code_dirs))
    header_dirs = list(dict.fromkeys(header_dirs))
    include_dirs = list(dict.fromkeys(include_dirs))

    # 排序
    code_dirs.sort()
    header_dirs.sort()
    include_dirs.sort()

    # print("project_dirs=", project_dirs)
    # print("code_dirs=", code_dirs)
    # print("header_dirs=", header_dirs)

    code_files = "\n"
    for code in code_dirs:
        code_files += code

    # print(code_files)

    header_files = "\n"
    for header in header_dirs:
        header_files += header

    # print(header_files)

    include_directories = "\n"
    for project in include_dirs:
        include_directories += project

    # print(include_directories)

    # cmake_file = os.path.join(os.path.dirname(__file__), project_name, "CMakeLists.txt")
    cmake_file = os.path.join(os.path.dirname(__file__), "../", project_name, "CMakeLists.txt")
    cmake_file_text = Template(cmake_template).substitute({'project_name': project_name,
                                                           'code_files': code_files,
                                                           'include_directories': include_directories,
                                                           'header_files': header_files,
                                                           'BUILD_NATIVE_ROOT': "${BUILD_NATIVE_ROOT}",
                                                           'SOURCE_FILES': "${SOURCE_FILES}",
                                                           'HEADERS': "${HEADERS}"
                                                           })
    write_text(cmake_file, cmake_file_text)

    # gradle_file = os.path.join(os.path.dirname(__file__), project_name, "build.gradle")
    gradle_file = os.path.join(os.path.dirname(__file__), "../", project_name, "build.gradle")
    gradle_file_text = Template(gradle_template).substitute({'rootDir': "${rootDir}"})
    write_text(gradle_file, gradle_file_text)


# aosp_projects = [
#     "frameworks/av/camera",
#     "frameworks/av/media/img_utils",
#     "frameworks/av/media/libaudioclient",
#     "frameworks/av/media/libaudiofoundation",
#     "frameworks/av/media/liberror",
#     "frameworks/av/media/libmedia",
#     "frameworks/av/media/libmediahelper",
#     "frameworks/av/media/libmediametrics",
#     "frameworks/av/media/libstagefright",
#     "frameworks/av/media/ndk",
#     "frameworks/base/apex/jobscheduler/service/jni",
#     "frameworks/base/cmds/app_process/",
#     "frameworks/base/cmds/bootanimation",
#     "frameworks/base/cmds/uinput/jni",
#     "frameworks/base/core/jni",
#     "frameworks/base/drm/jni",
#     "frameworks/base/libs/androidfw",
#     "frameworks/base/libs/hostgraphics",
#     "frameworks/base/libs/hwui",
#     "frameworks/base/libs/incident",
#     "frameworks/base/libs/input",
#     "frameworks/base/libs/protoutil",
#     "frameworks/base/libs/services",
#     "frameworks/base/libs/storage",
#     "frameworks/base/media/jni",
#     "frameworks/base/services/core/jni",
#     "frameworks/base/services/incremental",
#     "frameworks/native/headers/media_plugin",
#     "frameworks/native/libs/arect",
#     "frameworks/native/libs/battery",
#     "frameworks/native/libs/binder",
#     "frameworks/native/libs/binderthreadstate",
#     "frameworks/native/libs/cputimeinstate",
#     "frameworks/native/libs/gralloc",
#     "frameworks/native/libs/graphicsenv",
#     "frameworks/native/libs/gui",
#     "frameworks/native/libs/math",
#     "frameworks/native/libs/nativebase",
#     "frameworks/native/libs/nativedisplay",
#     "frameworks/native/libs/nativewindow",
#     "frameworks/native/libs/permission",
#     "frameworks/native/libs/sensor",
#     "frameworks/native/libs/ui",
#     "frameworks/native/opengl",
#     "frameworks/native/services/inputflinger",
#     "frameworks/native/services/surfaceflinger",
#     "hardware/libhardware",
#     "libnativehelper",
#     "system/core/libprocessgroup",
#     "system/core/libsystem",
#     "system/core/libutils",
#     "system/core/libcutils",
#     "system/core/property_service/libpropertyinfoparser",
#     "system/libbase",
#     "system/libhidl",
#     "system/libhwbinder",
#     "system/logging/liblog",
# ]
def get_dirs(project, full):

    common_projects = [
        "frameworks/native/include",
        "frameworks/native/libs/arect",
        "frameworks/native/libs/battery",
        "frameworks/native/libs/binder",
        "frameworks/native/libs/binderthreadstate",
        "frameworks/native/libs/cputimeinstate",
        "frameworks/native/libs/gralloc",
        "frameworks/native/libs/graphicsenv",
        "frameworks/native/libs/gui",
        "frameworks/native/libs/math",
        "frameworks/native/libs/nativebase",
        "frameworks/native/libs/nativedisplay",
        "frameworks/native/libs/nativewindow",
        "frameworks/native/libs/permission",
        "frameworks/native/libs/sensor",
        "frameworks/native/libs/ui",
        "frameworks/native/opengl",
        "hardware/libhardware",
        "libnativehelper",
        "system/core/libprocessgroup",
        "system/core/libsystem",
        "system/core/libutils",
        "system/core/libcutils",
        "system/core/property_service/libpropertyinfoparser",
        "system/libbase",
        "system/libhidl",
        "system/libhwbinder",
        "system/logging/liblog",
    ]

    android_runtime = [
        "frameworks/base/apex/jobscheduler/service/jni",
        "frameworks/base/cmds/app_process/",
        "frameworks/base/cmds/bootanimation",
        "frameworks/base/core/jni",
        "frameworks/base/drm/jni",
        "frameworks/base/libs/androidfw",
        "frameworks/base/libs/hostgraphics",
        "frameworks/base/libs/hwui",
        "frameworks/base/libs/incident",
        "frameworks/base/libs/protoutil",
        "frameworks/base/libs/storage",
    ]

    android_services = [
        "frameworks/base/libs/services",
        "frameworks/base/services/core/jni",
        "frameworks/base/services/incremental",
    ]

    inputflinger = [
        "system/core/libcutils",
        "frameworks/native/services/inputflinger",
        "frameworks/base/libs/input",
        "frameworks/base/cmds/uinput/jni",
    ]

    surfaceflinger = [
        "frameworks/native/services/surfaceflinger",
    ]

    media = [
        "frameworks/av/camera",
        "frameworks/av/media/img_utils",
        "frameworks/av/media/libaudioclient",
        "frameworks/av/media/libaudiofoundation",
        "frameworks/av/media/liberror",
        "frameworks/av/media/libmedia",
        "frameworks/av/media/libmediahelper",
        "frameworks/av/media/libmediametrics",
        "frameworks/av/media/libstagefright",
        "frameworks/av/media/ndk",
        "frameworks/base/media/jni",
        "frameworks/native/headers/media_plugin",
    ]

    if full == "1":
        dirs = common_projects
    else:
        dirs = []

    if project == "android_runtime":
        dirs.extend(common_projects)
        dirs.extend(android_runtime)
    elif project == "android_services":
        dirs.extend(android_services)
    elif project == "inputflinger":
        dirs.extend(inputflinger)
    elif project == "surfaceflinger":
        dirs.extend(surfaceflinger)
    elif project == "media":
        dirs.extend(media)
    else:
        dirs.extend(common_projects)
        dirs.extend(android_runtime)
        dirs.extend(android_services)
        dirs.extend(inputflinger)
        dirs.extend(surfaceflinger)
        dirs.extend(media)

    # 删除重复并重新排序
    aosp_projects = list(dict.fromkeys(dirs))
    aosp_projects.sort()

    # return ["system/core/libcutils"]
    return aosp_projects

def main():
    (options, args) = parseargs()
    root = options.root.strip()
    project = options.project.strip()
    full = options.full.strip()

    aosp_projects = get_dirs(project, full)

    work(project, root, aosp_projects)

    return 0


if __name__ == "__main__":
    main()
