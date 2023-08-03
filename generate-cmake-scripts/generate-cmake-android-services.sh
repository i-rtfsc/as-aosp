#!/bin/bash
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

# -d 可以通过“,”输入多目录
# (libservices.core) frameworks/base/services/core/jni
# (libservices) frameworks/base/libs/services
# (service.incremental) frameworks/base/services/incremental
python generate-cmake.py -p android-services -r /home/solo/code/aosp -d frameworks/base/services/core/jni,frameworks/base/libs/services,frameworks/base/services/incremental
