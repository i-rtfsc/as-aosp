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

cmake_minimum_required(VERSION 3.5)

project(libcompiler_rt-extras)

# Check if the major CMakeLists.txt file has correctly set
# the ANDROID_ROOT directory.
if (IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else ()
    if (IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
        set(OUT_INTERMEDIATES_ROOT ${ANDROID_ROOT}/out/soong/.intermediates)
    else ()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif ()
endif ()

#################### libcompiler_rt-extras ####################
set(FILES
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulodi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/comparedf2.c
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libcompiler_rt-extras SHARED ${SOURCE_FILES})

    target_include_directories(libcompiler_rt-extras
            PUBLIC
            ${ANDROID_ROOT}/external/compiler-rt
    )

    # Add c/c++ flags
    target_compile_options(libcompiler_rt-extras PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libcompiler_rt-extras
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libcompiler_rt-extras ####################

#################### libcompiler_rt ####################
set(FILES
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/enable_execute_stack.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/absvdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/absvsi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/absvti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/adddf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/addsf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/addtf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/addvdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/addvsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/addvti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ashldi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ashlti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ashrdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ashrti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/clear_cache.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/clzdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/clzsi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/clzti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/cmpdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/cmpti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/comparedf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/comparesf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/comparetf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ctzdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ctzsi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ctzti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divdc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divdf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divmoddi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divmodsi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divsc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divsf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divtf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/divxc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/eprintf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/extenddftf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/extendhfsf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/extendsfdf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/extendsftf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ffsdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ffsti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixdfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixdfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixdfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixsfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixsfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixsfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixtfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixtfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixtfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsdfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsdfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsdfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunssfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunssfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunssfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunstfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunstfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunstfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsxfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsxfsi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixunsxfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixxfdi.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/fixxfti.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatdidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatdisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatdixf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatsidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatsisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatsitf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floattidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floattisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floattixf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatundidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatundisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatundixf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatunsidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatunsisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatunsitf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatuntidf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatuntisf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/floatuntixf.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/gcc_personality_v0.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/int_util.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/lshrdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/lshrti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/moddi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/modsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/modti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/muldc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/muldf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/muldi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulodi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulosi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/muloti4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulsc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulsf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/multf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/multi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulvdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulvsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulvti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/mulxc3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negdf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negsf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negvdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negvsi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/negvti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/paritydi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/paritysi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/parityti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/popcountdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/popcountsi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/popcountti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/powidf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/powisf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/powitf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/powixf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subdf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subsf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subtf3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subvdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subvsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/subvti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/trampoline_setup.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/truncdfhf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/truncdfsf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/truncsfhf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/trunctfdf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/trunctfsf2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ucmpdi2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/ucmpti2.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivdi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivmoddi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivmodsi4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivmodti4.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/udivti3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/umoddi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/umodsi3.c
        ${ANDROID_ROOT}/external/compiler-rt/lib/builtins/umodti3.c
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libcompiler_rt SHARED ${SOURCE_FILES})

    target_include_directories(libcompiler_rt
            PUBLIC
            ${ANDROID_ROOT}/external/compiler-rt
    )

    # Add c/c++ flags
    target_compile_options(libcompiler_rt PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libcompiler_rt
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
            libunwind
            libwinpthread
    )
endif ()
#################### libcompiler_rt ####################
