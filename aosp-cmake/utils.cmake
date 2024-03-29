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

# Check if the specified directory exists and contains the CMakeLists.txt file.
function(add_subdirectory_safe dir_path)
    if (IS_DIRECTORY ${dir_path})
        file(GLOB CMAKELISTS ${dir_path}/CMakeLists.txt)
        if (CMAKELISTS)
            add_subdirectory(${dir_path})
        else ()
            message(WARNING "CMakeLists.txt does not exist in directory: ${dir_path}")
        endif ()
    else ()
        message(WARNING "Directory does not exist: ${dir_path}")
    endif ()
endfunction()