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


import os
import shutil


def delete_build_folders(root_dir):
    for dirpath, dirnames, filenames in os.walk(root_dir):
        if 'build' in dirnames:
            build_dir = os.path.join(dirpath, 'build')
            try:
                shutil.rmtree(build_dir)
                print(f"Deleted {build_dir}")
            except OSError as e:
                print(f"Error deleting {build_dir}: {e}")


if __name__ == "__main__":
    pwd = os.getcwd()
    root_directory = os.path.dirname(pwd)

    delete_build_folders(root_directory)
