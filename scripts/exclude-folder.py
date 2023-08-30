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


def parseargs():
    usage = "usage: %prog [options] arg1 arg2"
    parser = optparse.OptionParser(usage=usage)

    buildoptiongroup = optparse.OptionGroup(parser, "exclude folder")

    buildoptiongroup.add_option("-p", "--project", dest="project",
                                help="project", default="aosp")
    buildoptiongroup.add_option("-f", "--file", dest="file",
                                help="file",
                                default="/home/solo/code/github/as-aosp/.idea/modules/aosp-native/flyme.aosp-native.main.iml")

    parser.add_option_group(buildoptiongroup)

    (options, args) = parser.parse_args()

    return (options, args)


def work(project, file):
    old = "      <sourceFolder url=\"file://$MODULE_DIR$/../../../../../{project}\" type=\"native-Source-root\"/>"

    new = "      <sourceFolder url=\"file://$MODULE_DIR$/../../../../../{project}\" type=\"native-Source-root\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/art\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/bionic\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/bootable\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/build\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/cts\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/dalvik\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/developers\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/development\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/device\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/disregard\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/external\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/hardware\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/kernel\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/libcore\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/libnativehelper\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/out\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/packages\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/pdk\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/platform\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/platform_testing\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/prebuilts\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/.repo\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/sdk\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/system\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/test\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/toolchain\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/tools\"/>\n" \
          "      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/vendor\"/>"

    check_line="      <excludeFolder url=\"file://$MODULE_DIR$/../../../../../{project}/art\"/>".format(project=project)

    old = old.format(project=project)
    new = new.format(project=project)

    with open(file, 'r') as f:
        lines = f.read()

        if check_line in lines:
            print("has't been exclude folder")
            return

        lines = lines.replace(old, new)

        with open(file, 'w') as f:
            f.write(lines)


def main():
    (options, args) = parseargs()
    project = options.project.strip()
    file = options.file.strip()

    work(project, file)

    return 0


if __name__ == "__main__":
    main()
