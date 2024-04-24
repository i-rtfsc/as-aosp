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


import glob
import optparse
import os
import zipfile


def parseargs():
    usage = "usage: %prog [options] arg1 arg2"
    parser = optparse.OptionParser(usage=usage)

    buildoptiongroup = optparse.OptionGroup(parser, "copy aosp scrjar(aidl、proto、R...) to build/aosp_generated_dir ")

    buildoptiongroup.add_option("-m", "--module", dest="module",
                                help="module dir",
                                default="/Users/solo/code/github/as-aosp/aosp-modules/Framework/build/aosp_srcjars")
    buildoptiongroup.add_option("-s", "--srcjars", dest="srcjars",
                                help="srcjars",
                                default="/Users/solo/code/aosp/out/soong/.intermediates/frameworks/base/*-java/android_common/gen/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/frameworks/base/core/res/framework-res/android_common/gen/android/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/frameworks/base/framework-javastream-protos/**/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/frameworks/base/framework-minus-apex/**/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/frameworks/av/media/aconfig/aconfig_mediacodec_flags_java_lib/android_common/gen/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/hardware/interfaces/audio/7.0/config/audio_policy_configuration_V7_0/gen/java/*.srcjar#/Users/solo/code/aosp/out/soong/.intermediates/system/apex/apexd/apex-info-list/gen/java/*.srcjar")

    parser.add_option_group(buildoptiongroup)

    (options, args) = parser.parse_args()

    return (options, args)


def find_srcjar_files(input_path):
    # 去除通配符路径中的 "**/" 部分
    base_dir = os.path.dirname(input_path)
    pattern = os.path.basename(input_path)

    # 使用 glob.glob 函数递归查找符合条件的文件
    srcjar_files = glob.glob(os.path.join(base_dir, "**", pattern), recursive=True)
    srcjar_files = set(srcjar_files)
    return srcjar_files


def work(module, srcjars):
    build_dir = module
    if not os.path.exists(build_dir):
        os.makedirs(build_dir)

    for srcjar in srcjars:
        srcjar_files = find_srcjar_files(srcjar)

        for srcjar_file in srcjar_files:
            # print(srcjar_file)
            # 打开.srcjar文件

            if not os.path.exists(srcjar_file):
                print("file = " + srcjar_file + " not exists")
                continue

            with zipfile.ZipFile(srcjar_file, 'r') as zip_ref:
                # 解压文件到目标目录
                zip_ref.extractall(build_dir)


def main():
    (options, args) = parseargs()
    module = options.module.strip()

    srcjars = options.srcjars.strip()
    srcjars = srcjars.split("#")

    work(module, srcjars)

    return 0


if __name__ == "__main__":
    main()
