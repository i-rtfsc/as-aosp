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

project(tpm2-tss2-sys)

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

#################### tpm2-tss2-sys ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/sysapi_util.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_AC_GetCapability.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_AC_Send.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ActivateCredential.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Certify.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_CertifyCreation.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ChangeEPS.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ChangePPS.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Clear.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ClearControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ClockRateAdjust.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ClockSet.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Commit.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ContextLoad.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ContextSave.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Create.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_CreateLoaded.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_CreatePrimary.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_DictionaryAttackLockReset.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_DictionaryAttackParameters.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Duplicate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ECC_Parameters.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ECDH_KeyGen.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ECDH_ZGen.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_EC_Ephemeral.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_EncryptDecrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_EncryptDecrypt2.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_EventSequenceComplete.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_EvictControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Execute.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_FieldUpgradeData.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_FieldUpgradeStart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Finalize.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_FirmwareRead.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_FlushContext.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetCapability.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetCommandAuditDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetCommandCode.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetContextSize.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetCpBuffer.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetDecryptParam.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetEncryptParam.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetRandom.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetRpBuffer.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetRspAuths.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetSessionAuditDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetTctiContext.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetTestResult.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_GetTime.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_HMAC.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_HMAC_Start.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Hash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_HashSequenceStart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_HierarchyChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_HierarchyControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Import.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_IncrementalSelfTest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Initialize.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Load.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_LoadExternal.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_MakeCredential.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_Certify.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_ChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_DefineSpace.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_Extend.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_GlobalWriteLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_Increment.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_Read.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_ReadLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_ReadPublic.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_SetBits.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_UndefineSpace.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_UndefineSpaceSpecial.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_Write.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_NV_WriteLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ObjectChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_Allocate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_Event.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_Extend.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_Read.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_Reset.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_SetAuthPolicy.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PCR_SetAuthValue.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PP_Commands.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyAuthValue.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyAuthorize.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyAuthorizeNV.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyCommandCode.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyCounterTimer.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyCpHash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyDuplicationSelect.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyGetDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyLocality.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyNV.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyNameHash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyNvWritten.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyOR.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyPCR.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyPassword.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyPhysicalPresence.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyRestart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicySecret.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicySigned.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyTemplate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_PolicyTicket.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Policy_AC_SendSelect.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Quote.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_RSA_Decrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_RSA_Encrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ReadClock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ReadPublic.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Rewrap.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SelfTest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SequenceComplete.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SequenceUpdate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetAlgorithmSet.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetCmdAuths.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetCommandCodeAuditStatus.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetDecryptParam.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetEncryptParam.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_SetPrimaryPolicy.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Shutdown.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Sign.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_StartAuthSession.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Startup.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_StirRandom.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_TestParms.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Unseal.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_Vendor_TCG_Test.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_VerifySignature.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api/Tss2_Sys_ZGen_2Phase.c
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
    add_library(tpm2-tss2-sys SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-sys
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-sys/api
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-sys PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-sys
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            tpm2-tss2-mu
            tpm2-tss2-util
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-sys ####################

#################### tpm2-tss2-tcti ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-tcti/tcti-common.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-tcti/tctildr.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-tcti/tctildr-nodl.c
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
    add_library(tpm2-tss2-tcti SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-tcti
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-tcti
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-tcti PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-tcti
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            tpm2-tss2-mu
            tpm2-tss2-util
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-tcti ####################

#################### tpm2-tss2-rc ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-rc/tss2_rc.c
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
    add_library(tpm2-tss2-rc SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-rc
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-rc
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-rc PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-rc
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            tpm2-tss2-sys
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-rc ####################

#################### tpm2-tss2-esys ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ActivateCredential.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Certify.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_CertifyCreation.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ChangeEPS.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ChangePPS.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Clear.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ClearControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ClockRateAdjust.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ClockSet.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Commit.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ContextLoad.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ContextSave.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Create.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_CreateLoaded.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_CreatePrimary.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_DictionaryAttackLockReset.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_DictionaryAttackParameters.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Duplicate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ECC_Parameters.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ECDH_KeyGen.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ECDH_ZGen.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_EC_Ephemeral.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_EncryptDecrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_EncryptDecrypt2.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_EventSequenceComplete.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_EvictControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_FieldUpgradeData.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_FieldUpgradeStart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_FirmwareRead.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_FlushContext.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetCapability.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetCommandAuditDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetRandom.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetSessionAuditDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetTestResult.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_GetTime.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_HMAC.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_HMAC_Start.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Hash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_HashSequenceStart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_HierarchyChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_HierarchyControl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Import.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_IncrementalSelfTest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Load.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_LoadExternal.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_MakeCredential.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_Certify.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_ChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_DefineSpace.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_Extend.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_GlobalWriteLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_Increment.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_Read.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_ReadLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_ReadPublic.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_SetBits.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_UndefineSpace.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_UndefineSpaceSpecial.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_Write.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_NV_WriteLock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ObjectChangeAuth.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_Allocate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_Event.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_Extend.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_Read.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_Reset.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_SetAuthPolicy.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PCR_SetAuthValue.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PP_Commands.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyAuthValue.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyAuthorize.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyAuthorizeNV.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyCommandCode.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyCounterTimer.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyCpHash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyDuplicationSelect.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyGetDigest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyLocality.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyNV.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyNameHash.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyNvWritten.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyOR.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyPCR.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyPassword.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyPhysicalPresence.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyRestart.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicySecret.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicySigned.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyTemplate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_PolicyTicket.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Quote.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_RSA_Decrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_RSA_Encrypt.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ReadClock.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ReadPublic.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Rewrap.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SelfTest.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SequenceComplete.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SequenceUpdate.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SetAlgorithmSet.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SetCommandCodeAuditStatus.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_SetPrimaryPolicy.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Shutdown.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Sign.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_StartAuthSession.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Startup.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_StirRandom.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_TestParms.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Unseal.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_Vendor_TCG_Test.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_VerifySignature.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/api/Esys_ZGen_2Phase.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_context.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_crypto.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_crypto_bssl.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_free.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_iutil.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_mu.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys/esys_tr.c
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
    add_library(tpm2-tss2-esys SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-esys
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-esys
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-esys PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-esys
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libcrypto
            libcrypto
            libdl
            libm
            libwinpthread
            tpm2-tss2-mu
            tpm2-tss2-sys
            tpm2-tss2-tcti
            tpm2-tss2-util
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-esys ####################

#################### tpm2-tss2-util ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/util/io.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/util/key-value-parse.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/util/log.cpp
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
    add_library(tpm2-tss2-util SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-util
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-util PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-util
            libbase
            libbase
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-util ####################

#################### tpm2-tss2-mu ####################
set(FILES
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/base-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpm2b-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpma-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpml-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpms-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpmt-types.c
        ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu/tpmu-types.c
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
    add_library(tpm2-tss2-mu SHARED ${SOURCE_FILES})

    target_include_directories(tpm2-tss2-mu
            PUBLIC
            ${ANDROID_ROOT}/external/tpm2-tss/include
            ${ANDROID_ROOT}/external/tpm2-tss/include/tss2
            ${ANDROID_ROOT}/external/tpm2-tss/src
            ${ANDROID_ROOT}/external/tpm2-tss/src/tss2-mu
            ${ANDROID_ROOT}/external/tpm2-tss
    )

    # Add c/c++ flags
    target_compile_options(tpm2-tss2-mu PRIVATE
            -Werror
    )

    target_link_libraries(tpm2-tss2-mu
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            tpm2-tss2-util
            libc
            libdl
            libm
    )
endif ()
#################### tpm2-tss2-mu ####################
