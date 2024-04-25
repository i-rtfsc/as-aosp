#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# -*- encoding: utf-8 -*-
#
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

import argparse
import os
import platform
import shutil
import subprocess
import time

from tqdm import tqdm


def split_list_into_chunks(lst, num_chunks):
    chunk_size = len(lst) // num_chunks
    remainder = len(lst) % num_chunks
    chunks = [lst[i * chunk_size:(i + 1) * chunk_size] for i in range(num_chunks)]
    if remainder > 0:
        chunks[-1].extend(lst[-remainder:])
    return chunks


def find_files_with_glob(directory, extensions):
    start_time = time.time()

    # import glob
    # files = []
    # for extension in extensions:
    #     # files.extend(glob.glob(f"{directory}/**/*{extension}", recursive=True))
    #     found_files = [f for f in glob.glob(f"{directory}/**/*{extension}", recursive=True) if os.path.isfile(f)]
    #     files.extend(found_files)

    # 经过测试，使用 find 命令比用 glob 效率更高
    command = f"find {directory} -type f \( " + \
              " ".join([f"-name '*{ext}' -o" for ext in extensions])[:-3] + \
              " \)"
    print(command)
    output = subprocess.check_output(command, shell=True)
    files = output.decode('utf-8').splitlines()

    end_time = time.time()
    return files, end_time - start_time


def process_files(files, aosp_dir, target_dir, symbolic_link=True):
    # for file in files:
    for file in tqdm(files, leave=False):
        destination_path = os.path.join(target_dir, os.path.relpath(file, aosp_dir))
        # 如果目标路径存在，并且是一个文件或软链接，则跳过
        if os.path.exists(destination_path) and (os.path.isfile(destination_path) or os.path.islink(destination_path)):
            print(f"Skipping: {destination_path} already exists.")
            continue
        # 确保目录存在
        os.makedirs(os.path.dirname(destination_path), exist_ok=True)
        if symbolic_link:
            try:
                os.symlink(file, destination_path)
            except FileExistsError:
                print(f"Skipping: {destination_path} already exists.")
        else:
            shutil.copy2(file, destination_path)


def work(aosp_dir, target_dir, symbolic_link, thread=1):
    out_dir = os.path.join(aosp_dir, "out/soong/.intermediates")
    extensions = [
        '.c', '.c++',
        '.cc', '.cpp',
        '.h', '.hh',
        '.hpp', '.h++',
        '.srcjar'
    ]

    # java文件非常多，目前也不知道有没有用
    # 所有 copy 的场景只 copy 名字是 EventLogTags.java
    # 如果是软链接则无所谓
    if symbolic_link:
        extensions.append('.java')
    else:
        extensions.append('*LogTags.java')
        extensions.append('*StatsLog.java')

    files, elapsed_time = find_files_with_glob(out_dir, extensions)
    print(f"Time elapsed: {elapsed_time:.2f} seconds")

    start_time = time.time()

    if thread <= 1:
        process_files(files, aosp_dir, target_dir, symbolic_link)
    else:
        from concurrent.futures import ThreadPoolExecutor
        import random

        executor = ThreadPoolExecutor(max_workers=8)
        tasks = []
        random.shuffle(files)
        files_chunks = split_list_into_chunks(files, thread)
        for i, chunk in enumerate(files_chunks, start=1):
            args = (chunk, aosp_dir, target_dir, symbolic_link)
            task = executor.submit(process_files, *args)
            tasks.append(task)

    # 等待所有任务完成
    executor.shutdown()
    for task in tasks:
        task.result()  # 等待任务完成

    elapsed_time = time.time() - start_time
    print(f"Time elapsed: {elapsed_time:.2f} seconds")


def parseargs():
    parser = argparse.ArgumentParser(description="copy or link java/c/c++ files from out directory")
    parser.add_argument("--aosp", dest="aosp",
                        help="aosp dir", default="/Users/solo/workspace/code/aosp/")
    parser.add_argument("--target", dest="target",
                        help="target dir", default="/Users/solo/code/aosp/")
    parser.add_argument("--copy", dest="copy",
                        help="copy or link(1=copy, 0=link, default=link)", default=0)
    return parser.parse_args()


def main():
    args = parseargs()

    aosp_dir = args.aosp.strip()
    target_dir = args.target.strip()
    symbolic_link = (int(args.copy) == 0)
    print("link =", symbolic_link)

    if platform.system() != "Darwin":
        aosp_dir = aosp_dir.replace("Users", "home")
        target_dir = target_dir.replace("Users", "home")

    work(aosp_dir, target_dir, symbolic_link, 8)

    return 0


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        exit()
