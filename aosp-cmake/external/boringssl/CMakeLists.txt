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

project(acvp_modulewrapper)

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

#################### bssl ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/tool/args.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/ciphers.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/client.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/const.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/digest.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/fd.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/file.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/generate_ech.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/generate_ed25519.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/genrsa.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/pkcs12.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/rand.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/server.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/sign.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/speed.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/tool.cc
        ${ANDROID_ROOT}/external/boringssl/src/tool/transport_common.cc
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
    add_executable(bssl ${SOURCE_FILES})

    target_include_directories(bssl
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl
            ${ANDROID_ROOT}/external/boringssl/src/include
    )

    # Add c/c++ flags
    target_compile_options(bssl PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(bssl
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libssl
    )
endif ()
#################### bssl ####################

#################### libacvp_modulewrapper ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/util/fipstools/acvp/modulewrapper/modulewrapper.cc
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
    add_library(libacvp_modulewrapper SHARED ${SOURCE_FILES})

    target_include_directories(libacvp_modulewrapper
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/util/fipstools/acvp/modulewrapper
            ${ANDROID_ROOT}/external/boringssl
            ${ANDROID_ROOT}/external/boringssl/src/include
    )

    # Add c/c++ flags
    target_compile_options(libacvp_modulewrapper PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libacvp_modulewrapper
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
    )
endif ()
#################### libacvp_modulewrapper ####################

#################### libssl_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/ssl/bio_ssl.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_pkt.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_srtp.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/dtls_method.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/dtls_record.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/encrypted_client_hello.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/extensions.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handoff.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake_client.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake_server.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_pkt.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_aead_ctx.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_asn1.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_buffer.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_cert.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_cipher.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_file.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_key_share.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_privkey.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_session.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_stat.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_transcript.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_versions.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_x509.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/t1_enc.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_client.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_enc.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_server.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls_method.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls_record.cc
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
    add_library(libssl_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libssl_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/include
            ${ANDROID_ROOT}/external/boringssl
    )

    # Add c/c++ flags
    target_compile_options(libssl_baremetal PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libssl_baremetal
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_baremetal
            libdl
            libm
    )
endif ()
#################### libssl_baremetal ####################

#################### libcrypto_fuzz_unsafe ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/err_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bitstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_d2i_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_dup.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_gentm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_i2d_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_mbstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_object.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_octet.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strex.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strnid.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_type.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_utctm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_par.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn_pack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_string.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/posix_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_dec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_enc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_fre.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_new.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_typ.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/base64/base64.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio_mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/connect.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/fd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/hexdump.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/pair.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/printf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket_helper.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/blake2/blake2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/bn_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/convert.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/buf/buf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/asn1_compat.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/ber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbb.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/unicode.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/chacha/chacha.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/cipher_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/derive_key.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesctrhmac.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesgcmsiv.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_chacha20poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_null.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_tls.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/tls_cbc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/conf/conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_apple.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_intel.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/crypto.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/curve25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/spake25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/des/des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/dh_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/params.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/digest_extra/digest_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_derive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/hash_to_curve.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdh_extra/ecdh_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdsa_extra/ecdsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/engine/engine.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/err/err.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_ctx.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_hkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/pbkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/scrypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ex_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/fips_shared_support.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hpke/hpke.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hrss/hrss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/keccak.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/kyber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/lhash/lhash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj_xref.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_oth.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pk8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_xaux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/p5_pbev2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_vec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pool/pool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/deterministic.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/forkunsafe.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/passive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/rand_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/windows.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rc4/rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_c11.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_lock.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_crypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/siphash/siphash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/stack/stack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_none.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_pthread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/pmbtoken.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/trust_token.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/voprf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_digest.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_verify.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/algorithm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/asn1_gen.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_dir.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/i2d_pr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/name_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/policy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/rsa_pss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_att.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_cmp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_d2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_def.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_ext.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_lu.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_set.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_trs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_txt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_v3.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vfy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vpm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509cset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509rset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_algor.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_attrib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_exten.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pubkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_sig.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_val.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akeya.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_alt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bitst.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_cpols.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_crld.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_enum.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_extku.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_genn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ia5.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ncons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ocsp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pmaps.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_prn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_purp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_skey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/bcm.c
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/chacha/chacha-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/cipher_extra/chacha20_poly1305_armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/test/trampoline-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-gcm-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/armv8-mont-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/bn-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghash-neon-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghashv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256_beeu-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha1-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha256-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha512-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/vpaes-armv8-linux.S
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
    add_library(libcrypto_fuzz_unsafe SHARED ${SOURCE_FILES})

    target_include_directories(libcrypto_fuzz_unsafe
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/include
            ${ANDROID_ROOT}/external/boringssl/src/crypto
            ${ANDROID_ROOT}/external/boringssl
    )

    # Add c/c++ flags
    target_compile_options(libcrypto_fuzz_unsafe PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libcrypto_fuzz_unsafe
            libc
            libclang_rt.builtins
            libdl
            libgtest
            libgtest_main
            libm
            libwinpthread
    )
endif ()
#################### libcrypto_fuzz_unsafe ####################

#################### libcrypto_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/err_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bitstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_d2i_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_dup.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_gentm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_i2d_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_mbstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_object.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_octet.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strex.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strnid.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_type.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_utctm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_par.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn_pack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_string.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/posix_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_dec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_enc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_fre.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_new.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_typ.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/base64/base64.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio_mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/connect.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/fd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/hexdump.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/pair.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/printf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket_helper.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/blake2/blake2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/bn_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/convert.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/buf/buf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/asn1_compat.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/ber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbb.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/unicode.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/chacha/chacha.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/cipher_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/derive_key.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesctrhmac.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesgcmsiv.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_chacha20poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_null.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_tls.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/tls_cbc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/conf/conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_apple.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_intel.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/crypto.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/curve25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/spake25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/des/des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/dh_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/params.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/digest_extra/digest_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_derive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/hash_to_curve.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdh_extra/ecdh_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdsa_extra/ecdsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/engine/engine.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/err/err.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_ctx.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_hkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/pbkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/scrypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ex_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/fips_shared_support.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hpke/hpke.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hrss/hrss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/keccak.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/kyber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/lhash/lhash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj_xref.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_oth.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pk8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_xaux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/p5_pbev2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_vec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pool/pool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/deterministic.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/forkunsafe.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/passive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/rand_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/windows.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rc4/rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_c11.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_lock.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_crypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/siphash/siphash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/stack/stack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_none.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_pthread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/pmbtoken.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/trust_token.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/voprf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_digest.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_verify.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/algorithm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/asn1_gen.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_dir.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/i2d_pr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/name_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/policy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/rsa_pss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_att.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_cmp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_d2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_def.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_ext.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_lu.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_set.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_trs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_txt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_v3.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vfy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vpm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509cset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509rset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_algor.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_attrib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_exten.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pubkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_sig.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_val.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akeya.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_alt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bitst.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_cpols.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_crld.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_enum.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_extku.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_genn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ia5.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ncons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ocsp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pmaps.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_prn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_purp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_skey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/bcm.c
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/chacha/chacha-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/cipher_extra/chacha20_poly1305_armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/test/trampoline-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-gcm-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/armv8-mont-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/bn-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghash-neon-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghashv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256_beeu-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha1-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha256-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha512-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/vpaes-armv8-linux.S
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
    add_library(libcrypto_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libcrypto_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/include
            ${ANDROID_ROOT}/external/boringssl/src/crypto
            ${ANDROID_ROOT}/external/boringssl
    )

    # Add c/c++ flags
    target_compile_options(libcrypto_baremetal PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libcrypto_baremetal
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libcrypto_baremetal ####################

#################### libbssl_rust_support ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/rust/bssl-sys/rust_wrapper.c
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
    add_library(libbssl_rust_support SHARED ${SOURCE_FILES})

    target_include_directories(libbssl_rust_support
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl
            ${ANDROID_ROOT}/external/boringssl/src/include
    )

    # Add c/c++ flags
    target_compile_options(libbssl_rust_support PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbssl_rust_support
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libssl
    )
endif ()
#################### libbssl_rust_support ####################

#################### acvp_modulewrapper ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/util/fipstools/acvp/modulewrapper/main.cc
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
    add_executable(acvp_modulewrapper ${SOURCE_FILES})

    target_include_directories(acvp_modulewrapper
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl
            ${ANDROID_ROOT}/external/boringssl/src/util/fipstools/acvp/modulewrapper
            ${ANDROID_ROOT}/external/boringssl/src/include
    )

    # Add c/c++ flags
    target_compile_options(acvp_modulewrapper PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(acvp_modulewrapper
            libacvp_modulewrapper
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
    )
endif ()
#################### acvp_modulewrapper ####################

#################### libbssl_rust_support_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/rust/bssl-sys/rust_wrapper.c
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
    add_library(libbssl_rust_support_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(libbssl_rust_support_baremetal
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl
            ${ANDROID_ROOT}/external/boringssl/src/include
    )

    # Add c/c++ flags
    target_compile_options(libbssl_rust_support_baremetal PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbssl_rust_support_baremetal
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_baremetal
            libdl
            libm
            libssl_baremetal
    )
endif ()
#################### libbssl_rust_support_baremetal ####################

#################### libcrypto_static ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/err_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bitstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_bool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_d2i_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_dup.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_gentm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_i2d_fp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_mbstr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_object.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_octet.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strex.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_strnid.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_type.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/a_utctm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn1_par.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/asn_pack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/f_string.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/posix_time.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_dec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_enc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_fre.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_new.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_typ.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/asn1/tasn_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/base64/base64.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/bio_mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/connect.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/fd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/hexdump.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/pair.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/printf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bio/socket_helper.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/blake2/blake2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/bn_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bn_extra/convert.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/buf/buf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/asn1_compat.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/ber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbb.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/cbs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/bytestring/unicode.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/chacha/chacha.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/cipher_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/derive_key.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesctrhmac.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_aesgcmsiv.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_chacha20poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_null.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/e_tls.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cipher_extra/tls_cbc.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/conf/conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_apple.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_aarch64_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_freebsd.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_arm_linux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/cpu_intel.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/crypto.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/curve25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/curve25519/spake25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/des/des.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/dh_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dh_extra/params.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/digest_extra/digest_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/dsa/dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/ec_derive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ec_extra/hash_to_curve.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdh_extra/ecdh_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ecdsa_extra/ecdsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/engine/engine.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/err/err.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/evp_ctx.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_dsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ec_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_ed25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_hkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/p_x25519_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/pbkdf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/scrypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/evp/sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/ex_data.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/fips_shared_support.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hpke/hpke.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/hrss/hrss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/keccak.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/kyber/kyber.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/lhash/lhash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/mem.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/obj/obj_xref.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_oth.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pk8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pem/pem_xaux.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs7/pkcs7_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/p5_pbev2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pkcs8/pkcs8_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_arm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/poly1305/poly1305_vec.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/pool/pool.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/deterministic.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/forkunsafe.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/fuchsia.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/passive.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/rand_extra.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rand_extra/windows.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rc4/rc4.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_c11.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/refcount_lock.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_asn1.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_crypt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/rsa_extra/rsa_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/siphash/siphash.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/stack/stack.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_none.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_pthread.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/thread_win.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/pmbtoken.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/trust_token.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/trust_token/voprf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_digest.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_sign.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/a_verify.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/algorithm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/asn1_gen.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_dir.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/by_file.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/i2d_pr.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/name_print.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/policy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/rsa_pss.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/t_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_att.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_cmp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_d2.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_def.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_ext.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_lu.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_obj.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_set.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_trs.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_txt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_v3.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vfy.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509_vpm.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509cset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509rset.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x509spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_algor.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_all.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_attrib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_crl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_exten.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_name.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_pubkey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_req.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_sig.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_spki.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_val.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509/x_x509a.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_akeya.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_alt.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_bitst.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_conf.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_cpols.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_crld.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_enum.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_extku.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_genn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ia5.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_info.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_int.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_lib.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ncons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_ocsp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pcons.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_pmaps.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_prn.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_purp.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_skey.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/x509v3/v3_utl.c
        ${ANDROID_ROOT}/external/boringssl/src/crypto/fipsmodule/bcm.c
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/chacha/chacha-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/cipher_extra/chacha20_poly1305_armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/test/trampoline-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/aesv8-gcm-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/armv8-mont-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/bn-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghash-neon-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/ghashv8-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/p256_beeu-armv8-asm-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha1-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha256-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/sha512-armv8-linux.S
        ${ANDROID_ROOT}/external/boringssl/linux-aarch64/crypto/fipsmodule/vpaes-armv8-linux.S
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
    add_library(libcrypto_static SHARED ${SOURCE_FILES})

    target_include_directories(libcrypto_static
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/include
            ${ANDROID_ROOT}/external/boringssl/src/crypto
            ${ANDROID_ROOT}/external/boringssl
    )

    # Add c/c++ flags
    target_compile_options(libcrypto_static PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libcrypto_static
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libcrypto_static ####################

#################### libssl_fuzz_unsafe ####################
set(FILES
        ${ANDROID_ROOT}/external/boringssl/src/ssl/bio_ssl.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_pkt.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/d1_srtp.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/dtls_method.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/dtls_record.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/encrypted_client_hello.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/extensions.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handoff.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake_client.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/handshake_server.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/s3_pkt.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_aead_ctx.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_asn1.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_buffer.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_cert.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_cipher.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_file.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_key_share.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_lib.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_privkey.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_session.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_stat.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_transcript.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_versions.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/ssl_x509.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/t1_enc.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_both.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_client.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_enc.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls13_server.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls_method.cc
        ${ANDROID_ROOT}/external/boringssl/src/ssl/tls_record.cc
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
    add_library(libssl_fuzz_unsafe SHARED ${SOURCE_FILES})

    target_include_directories(libssl_fuzz_unsafe
            PUBLIC
            ${ANDROID_ROOT}/external/boringssl/src/include
            ${ANDROID_ROOT}/external/boringssl
    )

    # Add c/c++ flags
    target_compile_options(libssl_fuzz_unsafe PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libssl_fuzz_unsafe
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_fuzz_unsafe
            libdl
            libgtest
            libgtest_main
            libm
    )
endif ()
#################### libssl_fuzz_unsafe ####################
