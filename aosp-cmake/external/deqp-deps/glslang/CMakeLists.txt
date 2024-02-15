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

project(deqp_glslang_MachineIndependent)

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

#################### deqp_glslang_MachineIndependent ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/glslang_tab.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/attribute.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/Constant.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/iomapper.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/InfoSink.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/Initialize.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/IntermTraverse.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/Intermediate.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/ParseContextBase.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/ParseHelper.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/PoolAlloc.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/RemoveTree.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/Scan.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/ShaderLang.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/SpirvIntrinsics.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/SymbolTable.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/Versions.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/intermOut.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/limits.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/linkValidate.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/parseConst.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/reflection.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/preprocessor/Pp.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/preprocessor/PpAtom.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/preprocessor/PpContext.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/preprocessor/PpScanner.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/preprocessor/PpTokens.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent/propagateNoContraction.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslAttributes.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslParseHelper.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslScanContext.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslOpMap.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslTokenStream.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslGrammar.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL/hlslParseables.cpp
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
    add_library(deqp_glslang_MachineIndependent SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_MachineIndependent
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL
            ${ANDROID_ROOT}/external/deqp-deps/glslang/OGLCompilersDLL
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/OSDependent
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/GenericCodeGen
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_MachineIndependent PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_MachineIndependent
            deqp_glslang_GenericCodeGen
            deqp_glslang_OGLCompiler
            deqp_glslang_OSDependent
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_MachineIndependent ####################

#################### deqp_glslang_SPIRV ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/GlslangToSpv.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/InReadableOrder.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/Logger.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/SpvBuilder.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/SpvPostProcess.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/doc.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/SpvTools.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/disassemble.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/CInterface/spirv_c_interface.cpp
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
    add_library(deqp_glslang_SPIRV SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_SPIRV
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_SPIRV PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_SPIRV
            deqp_glslang_MachineIndependent
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_SPIRV ####################

#################### deqp_glslang_glslang ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/CInterface/glslang_c_interface.cpp
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
    add_library(deqp_glslang_glslang SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_glslang
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/OGLCompilersDLL
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/OSDependent
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/MachineIndependent
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/HLSL
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_glslang PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_glslang
            deqp_glslang_MachineIndependent
            deqp_glslang_OGLCompiler
            deqp_glslang_OSDependent
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_glslang ####################

#################### deqp_glslang_OGLCompiler ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/OGLCompilersDLL/InitializeDll.cpp
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
    add_library(deqp_glslang_OGLCompiler SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_OGLCompiler
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/OGLCompilersDLL
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_OGLCompiler PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_OGLCompiler
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_OGLCompiler ####################

#################### deqp_glslang_OSDependent ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/OSDependent/Unix/ossource.cpp
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
    add_library(deqp_glslang_OSDependent SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_OSDependent
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/OSDependent
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_OSDependent PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_OSDependent
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_OSDependent ####################

#################### deqp_glslang_SPVRemapper ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/SPVRemapper.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV/doc.cpp
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
    add_library(deqp_glslang_SPVRemapper SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_SPVRemapper
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/SPIRV
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_SPVRemapper PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_SPVRemapper
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_SPVRemapper ####################

#################### deqp_glslang_GenericCodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/GenericCodeGen/CodeGen.cpp
        ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/GenericCodeGen/Link.cpp
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
    add_library(deqp_glslang_GenericCodeGen SHARED ${SOURCE_FILES})

    target_include_directories(deqp_glslang_GenericCodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/glslang
            ${ANDROID_ROOT}/external/deqp-deps/glslang/glslang/GenericCodeGen
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/glslang/deqp_glslang_gen_build_info_h/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_glslang_GenericCodeGen PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_glslang_GenericCodeGen
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_glslang_GenericCodeGen ####################
