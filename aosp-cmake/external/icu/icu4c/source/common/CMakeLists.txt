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

project(libicuuc)

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

#################### libicuuc ####################
set(FILES
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/appendable.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bmpset.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/brkeng.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/brkiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bytesinkutil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bytestream.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bytestrie.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bytestriebuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/bytestrieiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/caniter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/characterproperties.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/chariter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/charstr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/cmemory.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/cstr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/cstring.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/cwchar.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/dictbe.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/dictionarydata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/dtintrv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/edits.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/emojiprops.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/errorcode.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/filteredbrk.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/filterednormalizer2.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/icudataver.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/icuplug.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/loadednormalizer2impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/localebuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/localematcher.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/localeprioritylist.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locavailable.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locbased.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locdispnames.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locdistance.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locdspnm.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locid.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/loclikely.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/loclikelysubtags.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locmap.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locresdata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/locutil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/lsr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/lstmbe.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/messagepattern.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/mlbe.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/normalizer2.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/normalizer2impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/normlzr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/parsepos.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/patternprops.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/pluralmap.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/propname.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/propsvec.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/punycode.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/putil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbi.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbi_cache.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbidata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbinode.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbirb.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbiscan.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbisetb.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbistbl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/rbbitblb.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/resbund.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/resbund_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/resource.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/restrace.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ruleiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/schriter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/serv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servlk.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servlkf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servls.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servnotf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servrbf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/servslkf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/sharedobject.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/simpleformatter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/static_unicode_sets.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/stringpiece.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/stringtriebuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uarrsort.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubidi.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubidi_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubidiln.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubiditransform.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubidiwrt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ubrk.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucase.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucasemap.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucasemap_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucat.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uchar.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucharstrie.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucharstriebuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucharstrieiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uchriter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucln_cmn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucmndata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv2022.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_bld.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_cb.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_ct.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_err.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_ext.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_io.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_lmb.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_set.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_u16.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_u32.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_u7.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnv_u8.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvbocu.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvdisp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvhz.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvisci.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvlat1.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvmbcs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvscsu.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucnvsel.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucol_swp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucptrie.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ucurr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/udata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/udatamem.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/udataswp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uenum.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uhash.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uhash_us.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uidna.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uinit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uinvchar.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ulist.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uloc.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uloc_keytype.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uloc_tag.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/umapfile.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/umath.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/umutablecptrie.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/umutex.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unames.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unifiedcache.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unifilt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unifunct.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uniset.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uniset_closure.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uniset_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unisetspan.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr_case.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr_case_locale.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unistr_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unorm.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/unormcmp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uobject.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uprops.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ures_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uresbund.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uresdata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/usc_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uscript.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uscript_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uset.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uset_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/usetiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ushape.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/usprep.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustack.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustr_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustr_titlecase_brkiter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustr_wcs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustrcase.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustrcase_locale.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustrenum.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustrfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustring.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/ustrtrns.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utext.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utf_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/util.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/util_props.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utrace.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utrie.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utrie2.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utrie2_builder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utrie_swap.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uts46.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/utypes.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uvector.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uvectr32.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/uvectr64.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/common/wintz.cpp
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
    add_library(libicuuc SHARED ${SOURCE_FILES})

    target_include_directories(libicuuc
            PUBLIC
            ${ANDROID_ROOT}/external/icu/icu4c/source/common
    )

    # Add c/c++ flags
    target_compile_options(libicuuc PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-const-variable
    )

    target_link_libraries(libicuuc
            libandroidicuinit
            libc++
            libc++_static
            libclang_rt.builtins
            libicuuc_headers
            libicuuc_stubdata
            libwinpthread
    )
endif ()
#################### libicuuc ####################
