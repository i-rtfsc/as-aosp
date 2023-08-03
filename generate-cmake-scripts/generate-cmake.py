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

    buildoptiongroup = optparse.OptionGroup(parser, "git push to gerrit options")

    buildoptiongroup.add_option("-p", "--project", dest="project",
                                help="project name", default="android-services")
    buildoptiongroup.add_option("-r", "--root", dest="root",
                                help="root dir", default="/home/solo/code/flyme")
    buildoptiongroup.add_option("-d", "--dir", dest="dir",
                                help="project dir", default="frameworks/base/services/core/jni,frameworks/base/libs/services,frameworks/base/services/incremental")

    parser.add_option_group(buildoptiongroup)

    (options, args) = parser.parse_args()

    return (options, args)


def file_is_code(file):
    # codes = ["c", "cpp", "cc", "c++"]
    codes = ["c", "cpp"]
    real = False
    for code in codes:
        if file.endswith(code):
            real = True
            break

    return real


def file_is_headers(file):
    # headers = ["h", "hpp", "h++"]
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
        project_dirs = fast_scandir(project_dir)
        project_dirs.append(project_dir)
        project_dirs.sort()

        for dir in project_dirs:
            # 代码文件夹
            for file in os.listdir(dir):
                if file_is_code(file):
                    code_dirs.append(dir)
                    include_dirs.append(dir)
                    break
            # 头文件文件夹
            for file in os.listdir(dir):
                if file_is_headers(file):
                    header_dirs.append(dir)
                    include_dirs.append(dir)
                    break

    include_dirs = list(dict.fromkeys(include_dirs))

    code_dirs.sort()
    header_dirs.sort()
    include_dirs.sort()

    # print("project_dirs=", project_dirs)
    # print("code_dirs=", code_dirs)
    # print("header_dirs=", header_dirs)

    code_files = "\n"
    for code in code_dirs:
        code_files += "        ${ANDROID_ROOT}/" + os.path.relpath(code, root) + "/*.c\n"
        code_files += "        ${ANDROID_ROOT}/" + os.path.relpath(code, root) + "/*.cpp\n"

    # print(code_files)

    header_files = "\n"
    for header in header_dirs:
        header_files += "        ${ANDROID_ROOT}/" + os.path.relpath(header, root) + "/*.h\n"
        header_files += "        ${ANDROID_ROOT}/" + os.path.relpath(header, root) + "/*.hpp\n"

    # print(header_files)

    include_directories = "\n"
    for project in include_dirs:
        include_directories += "        ${ANDROID_ROOT}/" + os.path.relpath(project, root) + "\n"

    # print(include_directories)

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

    gradle_file = os.path.join(os.path.dirname(__file__), "../", project_name, "build.gradle")
    gradle_file_text = Template(gradle_template).substitute({'rootDir': "${rootDir}"})
    write_text(gradle_file, gradle_file_text)


def main():
    (options, args) = parseargs()
    root = options.root.strip()
    project_dir = options.dir.strip()
    project = options.project.strip()

    project_list = []
    for dir in project_dir.split(","):
        print(dir)
        project_list.append(os.path.join(root, dir))

    work(project, root, project_list)

    return 0


if __name__ == "__main__":
    main()
