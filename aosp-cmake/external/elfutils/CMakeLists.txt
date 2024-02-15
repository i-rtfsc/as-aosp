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

project(libelf)

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

#################### libelf ####################
set(FILES
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_checksum.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_fsize.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_getchdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_getehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_getphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_getshdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_newehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_newphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_offscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_updatefile.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_updatenull.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_xlatetof.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf32_xlatetom.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_checksum.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_fsize.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_getchdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_getehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_getphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_getshdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_newehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_newphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_offscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_updatefile.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_updatenull.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_xlatetof.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf64_xlatetom.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_begin.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_clone.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_cntl.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_compress.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_compress_gnu.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_end.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_error.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_fill.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagdata.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagelf.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_flagshdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getarhdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getaroff.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getarsym.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getbase.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getdata.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getdata_rawchunk.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getident.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getphdrnum.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getshdrnum.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_getshdrstrndx.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_gnu_hash.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_hash.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_kind.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_memory.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_ndxscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_newdata.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_newscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_next.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_nextscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_rand.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_rawdata.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_rawfile.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_readall.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_scnshndx.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_strptr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_update.c
        ${ANDROID_ROOT}/external/elfutils/libelf/elf_version.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_checksum.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_fsize.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getauxv.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getchdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getclass.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getdyn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getlib.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getmove.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getnote.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getrel.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getrela.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getshdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getsym.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getsyminfo.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getsymshndx.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getverdaux.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getverdef.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getvernaux.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getverneed.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_getversym.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_newehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_newphdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_offscn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_auxv.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_dyn.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_ehdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_lib.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_move.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_phdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_rel.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_rela.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_shdr.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_sym.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_syminfo.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_symshndx.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_verdaux.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_verdef.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_vernaux.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_verneed.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_update_versym.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_xlate.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_xlatetof.c
        ${ANDROID_ROOT}/external/elfutils/libelf/gelf_xlatetom.c
        ${ANDROID_ROOT}/external/elfutils/libelf/libelf_crc32.c
        ${ANDROID_ROOT}/external/elfutils/libelf/libelf_next_prime.c
        ${ANDROID_ROOT}/external/elfutils/libelf/nlist.c
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
    add_library(libelf SHARED ${SOURCE_FILES})

    target_include_directories(libelf
            PUBLIC
            ${ANDROID_ROOT}/external/elfutils/libelf
            ${ANDROID_ROOT}/external/elfutils
            ${ANDROID_ROOT}/external/elfutils/bionic-fixup
            ${ANDROID_ROOT}/external/elfutils/include
            ${ANDROID_ROOT}/external/elfutils/lib
    )

    # Add c/c++ flags
    target_compile_options(libelf PRIVATE
            -Werror
    )

    target_link_libraries(libelf
            elfutils_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### libelf ####################
