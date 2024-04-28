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

import logging
import optparse
import os
import platform
import subprocess
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


def get_logger(log_file, level=logging.INFO):
    """Create a configured instance of logger."""
    fmt = '[%(asctime)s] %(levelname)s : %(message)s'
    date_fmt = '%Y-%m-%d %H:%M:%S'
    formatter = logging.Formatter(fmt, datefmt=date_fmt)

    logger = logging.getLogger()

    if not os.path.exists(log_file):
        pardir = os.path.abspath(os.path.join(log_file, os.pardir))
        if not os.path.exists(pardir):
            os.makedirs(pardir)
        file = open(log_file, 'w')
        file.close()
    fh = logging.FileHandler(filename=log_file, mode='w')
    fh.setFormatter(formatter)
    logger.addHandler(fh)

    logger.setLevel(level)
    logger.info("logger get or created.\n")

    return logger


def find_srcjar_files(input_path, logger):
    if os.path.exists(input_path):
        if platform.system() == 'Windows':
            logger.info("find srcjar on win, use glob = " + input_path)
            # win 没有 find 命令，换成 glob.glob
            import glob
            pattern = os.path.join(input_path, '**', '*.srcjar')
            files = glob.glob(pattern, recursive=True)
            return files
        else:
            command = f"find {input_path} -type f \( -name '*.srcjar' \)"
            logger.info("command = " + command)
            output = subprocess.check_output(command, shell=True)
            files = output.decode('utf-8').splitlines()
            return files
    else:
        logger.warning("dir = " + input_path + " not exists")
        return []


def work(module, srcjars):
    build_dir = module
    if not os.path.exists(build_dir):
        os.makedirs(build_dir)

    logger = get_logger(os.path.join(build_dir, 'copy.log'))

    for srcjar in srcjars:
        logger.info("srcjar = " + srcjar)
        srcjar_files = find_srcjar_files(srcjar, logger)

        for srcjar_file in srcjar_files:
            if not os.path.exists(srcjar_file):
                logger.warning("file = " + srcjar_file + " not exists")
                continue

            if srcjar_file.endswith("stubs.srcjar"):
                logger.warning("file = " + srcjar_file + " don't unzip")
                continue

            logger.info("file = " + srcjar_file)
            # 打开 .srcjar 文件
            with zipfile.ZipFile(srcjar_file, 'r') as zip_ref:
                # 解压文件到目标目录
                zip_ref.extractall(build_dir)


def main():
    (options, args) = parseargs()
    module = options.module.strip()

    srcjars = options.srcjars.strip()

    # gradle中调用必现得加 "" 才不报错
    # 加之后要去掉
    if srcjars.startswith('"') or srcjars.endswith('"'):
        srcjars = srcjars.replace('"', "")

    srcjars = srcjars.split("#")

    work(module, srcjars)

    return 0


if __name__ == "__main__":
    main()
