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

project(libicui18n)

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

#################### libicui18n ####################
set(FILES
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/alphaindex.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/anytrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/astro.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/basictz.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/bocsu.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/brktrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/buddhcal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/calendar.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/casetrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/cecal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/chnsecal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/choicfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/coleitr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/coll.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collation.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationbuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationcompare.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationdata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationdatabuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationdatareader.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationdatawriter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationfastlatin.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationfastlatinbuilder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationfcd.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationkeys.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationroot.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationrootelements.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationruleparser.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationsets.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationsettings.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationtailoring.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/collationweights.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/compactdecimalformat.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/coptccal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/cpdtrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csdetect.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csmatch.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csr2022.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csrecog.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csrmbcs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csrsbcs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csrucode.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/csrutf8.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/curramt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/currfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/currpinf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/currunit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dangical.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/datefmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dayperiodrules.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dcfmtsym.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/decContext.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/decNumber.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/decimfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/displayoptions.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-bignum-dtoa.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-bignum.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-cached-powers.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-double-to-string.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-fast-dtoa.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-string-to-double.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/double-conversion-strtod.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dtfmtsym.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dtitvfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dtitvinf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dtptngen.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/dtrule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/erarules.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/esctrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ethpccal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/fmtable.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/fmtable_cnv.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/format.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/formatted_string_builder.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/formattedval_iterimpl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/formattedval_sbimpl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/formattedvalue.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/fphdlimp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/fpositer.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/funcrepl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/gender.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/gregocal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/gregoimp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/hebrwcal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/indiancal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/inputext.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/islamcal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/japancal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/listformatter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/measfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/measunit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/measunit_extra.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/measure.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/msgfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/name2uni.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/nfrs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/nfrule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/nfsubs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/nortrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/nultrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_affixutils.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_asformat.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_capi.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_compact.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_currencysymbols.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_decimalquantity.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_decimfmtprops.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_fluent.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_formatimpl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_grouping.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_integerwidth.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_longnames.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_mapper.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_modifiers.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_multiplier.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_notation.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_output.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_padding.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_patternmodifier.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_patternstring.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_rounding.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_scientific.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_skeletons.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_symbolswrapper.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_usageprefs.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/number_utils.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_affixes.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_compositions.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_currency.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_decimal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_parsednumber.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_scientific.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_symbols.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numparse_validators.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numrange_capi.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numrange_fluent.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numrange_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/numsys.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/olsontz.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/persncal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/pluralranges.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/plurfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/plurrule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/quant.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/quantityformatter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbnf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbt_data.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbt_pars.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbt_rule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbt_set.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rbtz.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/regexcmp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/regeximp.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/regexst.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/regextxt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/region.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/reldatefmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/reldtfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rematch.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/remtrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/repattrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/rulebasedcollator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/scientificnumberformatter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/scriptset.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/search.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/selfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/sharedbreakiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/simpletz.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/smpdtfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/smpdtfst.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/sortkey.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/standardplural.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/string_segment.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/strmatch.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/strrepl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/stsearch.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/taiwncal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/timezone.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/titletrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tmunit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tmutamt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tmutfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tolowtrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/toupptrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/translit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/transreg.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tridpars.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tzfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tzgnames.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tznames.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tznames_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tzrule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/tztrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucal.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucln_in.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucol.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucol_res.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucol_sit.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucoleitr.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ucsdet.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/udat.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/udateintervalformat.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/udatpg.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ufieldpositer.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uitercollationiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ulistformatter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ulocdata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/umsg.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/unesctrn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uni2name.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/units_complexconverter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/units_converter.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/units_data.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/units_router.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/unum.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/unumsys.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/upluralrules.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uregex.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uregexc.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uregion.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/usearch.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uspoof.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uspoof_build.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uspoof_conf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/uspoof_impl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/utf16collationiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/utf8collationiterator.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/utmscale.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/utrans.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/vtzone.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/vzone.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/windtfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/winnmfmt.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/wintzimpl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/zonemeta.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/zrule.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/i18n/ztrans.cpp
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
    add_library(libicui18n SHARED ${SOURCE_FILES})

    target_include_directories(libicui18n
            PUBLIC
            ${ANDROID_ROOT}/external/icu/icu4c/source/i18n
    )

    # Add c/c++ flags
    target_compile_options(libicui18n PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-const-variable
    )

    target_link_libraries(libicui18n
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n_headers
            libicuuc
            libm
            libwinpthread
    )
endif ()
#################### libicui18n ####################
