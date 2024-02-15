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

project(curl)

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

#################### libcurl ####################
set(FILES
        ${ANDROID_ROOT}/external/curl/lib/altsvc.c
        ${ANDROID_ROOT}/external/curl/lib/amigaos.c
        ${ANDROID_ROOT}/external/curl/lib/asyn-ares.c
        ${ANDROID_ROOT}/external/curl/lib/asyn-thread.c
        ${ANDROID_ROOT}/external/curl/lib/base64.c
        ${ANDROID_ROOT}/external/curl/lib/bufref.c
        ${ANDROID_ROOT}/external/curl/lib/c-hyper.c
        ${ANDROID_ROOT}/external/curl/lib/cf-https-connect.c
        ${ANDROID_ROOT}/external/curl/lib/cf-socket.c
        ${ANDROID_ROOT}/external/curl/lib/cfilters.c
        ${ANDROID_ROOT}/external/curl/lib/conncache.c
        ${ANDROID_ROOT}/external/curl/lib/connect.c
        ${ANDROID_ROOT}/external/curl/lib/content_encoding.c
        ${ANDROID_ROOT}/external/curl/lib/cookie.c
        ${ANDROID_ROOT}/external/curl/lib/curl_addrinfo.c
        ${ANDROID_ROOT}/external/curl/lib/curl_des.c
        ${ANDROID_ROOT}/external/curl/lib/curl_endian.c
        ${ANDROID_ROOT}/external/curl/lib/curl_fnmatch.c
        ${ANDROID_ROOT}/external/curl/lib/curl_get_line.c
        ${ANDROID_ROOT}/external/curl/lib/curl_gethostname.c
        ${ANDROID_ROOT}/external/curl/lib/curl_gssapi.c
        ${ANDROID_ROOT}/external/curl/lib/curl_log.c
        ${ANDROID_ROOT}/external/curl/lib/curl_memrchr.c
        ${ANDROID_ROOT}/external/curl/lib/curl_multibyte.c
        ${ANDROID_ROOT}/external/curl/lib/curl_ntlm_core.c
        ${ANDROID_ROOT}/external/curl/lib/curl_ntlm_wb.c
        ${ANDROID_ROOT}/external/curl/lib/curl_path.c
        ${ANDROID_ROOT}/external/curl/lib/curl_range.c
        ${ANDROID_ROOT}/external/curl/lib/curl_rtmp.c
        ${ANDROID_ROOT}/external/curl/lib/curl_sasl.c
        ${ANDROID_ROOT}/external/curl/lib/curl_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/curl_threads.c
        ${ANDROID_ROOT}/external/curl/lib/dict.c
        ${ANDROID_ROOT}/external/curl/lib/doh.c
        ${ANDROID_ROOT}/external/curl/lib/dynbuf.c
        ${ANDROID_ROOT}/external/curl/lib/easy.c
        ${ANDROID_ROOT}/external/curl/lib/easygetopt.c
        ${ANDROID_ROOT}/external/curl/lib/easyoptions.c
        ${ANDROID_ROOT}/external/curl/lib/escape.c
        ${ANDROID_ROOT}/external/curl/lib/file.c
        ${ANDROID_ROOT}/external/curl/lib/fileinfo.c
        ${ANDROID_ROOT}/external/curl/lib/fopen.c
        ${ANDROID_ROOT}/external/curl/lib/formdata.c
        ${ANDROID_ROOT}/external/curl/lib/ftp.c
        ${ANDROID_ROOT}/external/curl/lib/ftplistparser.c
        ${ANDROID_ROOT}/external/curl/lib/getenv.c
        ${ANDROID_ROOT}/external/curl/lib/getinfo.c
        ${ANDROID_ROOT}/external/curl/lib/gopher.c
        ${ANDROID_ROOT}/external/curl/lib/h2h3.c
        ${ANDROID_ROOT}/external/curl/lib/hash.c
        ${ANDROID_ROOT}/external/curl/lib/headers.c
        ${ANDROID_ROOT}/external/curl/lib/hmac.c
        ${ANDROID_ROOT}/external/curl/lib/hostasyn.c
        ${ANDROID_ROOT}/external/curl/lib/hostip.c
        ${ANDROID_ROOT}/external/curl/lib/hostip4.c
        ${ANDROID_ROOT}/external/curl/lib/hostip6.c
        ${ANDROID_ROOT}/external/curl/lib/hostsyn.c
        ${ANDROID_ROOT}/external/curl/lib/hsts.c
        ${ANDROID_ROOT}/external/curl/lib/http.c
        ${ANDROID_ROOT}/external/curl/lib/http2.c
        ${ANDROID_ROOT}/external/curl/lib/http_aws_sigv4.c
        ${ANDROID_ROOT}/external/curl/lib/http_chunks.c
        ${ANDROID_ROOT}/external/curl/lib/http_digest.c
        ${ANDROID_ROOT}/external/curl/lib/http_negotiate.c
        ${ANDROID_ROOT}/external/curl/lib/http_ntlm.c
        ${ANDROID_ROOT}/external/curl/lib/http_proxy.c
        ${ANDROID_ROOT}/external/curl/lib/idn.c
        ${ANDROID_ROOT}/external/curl/lib/if2ip.c
        ${ANDROID_ROOT}/external/curl/lib/imap.c
        ${ANDROID_ROOT}/external/curl/lib/inet_ntop.c
        ${ANDROID_ROOT}/external/curl/lib/inet_pton.c
        ${ANDROID_ROOT}/external/curl/lib/krb5.c
        ${ANDROID_ROOT}/external/curl/lib/ldap.c
        ${ANDROID_ROOT}/external/curl/lib/llist.c
        ${ANDROID_ROOT}/external/curl/lib/md4.c
        ${ANDROID_ROOT}/external/curl/lib/md5.c
        ${ANDROID_ROOT}/external/curl/lib/memdebug.c
        ${ANDROID_ROOT}/external/curl/lib/mime.c
        ${ANDROID_ROOT}/external/curl/lib/mprintf.c
        ${ANDROID_ROOT}/external/curl/lib/mqtt.c
        ${ANDROID_ROOT}/external/curl/lib/multi.c
        ${ANDROID_ROOT}/external/curl/lib/netrc.c
        ${ANDROID_ROOT}/external/curl/lib/nonblock.c
        ${ANDROID_ROOT}/external/curl/lib/noproxy.c
        ${ANDROID_ROOT}/external/curl/lib/openldap.c
        ${ANDROID_ROOT}/external/curl/lib/parsedate.c
        ${ANDROID_ROOT}/external/curl/lib/pingpong.c
        ${ANDROID_ROOT}/external/curl/lib/pop3.c
        ${ANDROID_ROOT}/external/curl/lib/progress.c
        ${ANDROID_ROOT}/external/curl/lib/psl.c
        ${ANDROID_ROOT}/external/curl/lib/rand.c
        ${ANDROID_ROOT}/external/curl/lib/rename.c
        ${ANDROID_ROOT}/external/curl/lib/rtsp.c
        ${ANDROID_ROOT}/external/curl/lib/select.c
        ${ANDROID_ROOT}/external/curl/lib/sendf.c
        ${ANDROID_ROOT}/external/curl/lib/setopt.c
        ${ANDROID_ROOT}/external/curl/lib/sha256.c
        ${ANDROID_ROOT}/external/curl/lib/share.c
        ${ANDROID_ROOT}/external/curl/lib/slist.c
        ${ANDROID_ROOT}/external/curl/lib/smb.c
        ${ANDROID_ROOT}/external/curl/lib/smtp.c
        ${ANDROID_ROOT}/external/curl/lib/socketpair.c
        ${ANDROID_ROOT}/external/curl/lib/socks.c
        ${ANDROID_ROOT}/external/curl/lib/socks_gssapi.c
        ${ANDROID_ROOT}/external/curl/lib/socks_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/speedcheck.c
        ${ANDROID_ROOT}/external/curl/lib/splay.c
        ${ANDROID_ROOT}/external/curl/lib/strcase.c
        ${ANDROID_ROOT}/external/curl/lib/strdup.c
        ${ANDROID_ROOT}/external/curl/lib/strerror.c
        ${ANDROID_ROOT}/external/curl/lib/strtok.c
        ${ANDROID_ROOT}/external/curl/lib/strtoofft.c
        ${ANDROID_ROOT}/external/curl/lib/system_win32.c
        ${ANDROID_ROOT}/external/curl/lib/telnet.c
        ${ANDROID_ROOT}/external/curl/lib/tftp.c
        ${ANDROID_ROOT}/external/curl/lib/timediff.c
        ${ANDROID_ROOT}/external/curl/lib/timeval.c
        ${ANDROID_ROOT}/external/curl/lib/transfer.c
        ${ANDROID_ROOT}/external/curl/lib/url.c
        ${ANDROID_ROOT}/external/curl/lib/urlapi.c
        ${ANDROID_ROOT}/external/curl/lib/version.c
        ${ANDROID_ROOT}/external/curl/lib/version_win32.c
        ${ANDROID_ROOT}/external/curl/lib/warnless.c
        ${ANDROID_ROOT}/external/curl/lib/ws.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/cleartext.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/cram.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/digest.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/digest_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/gsasl.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/krb5_gssapi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/krb5_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/ntlm.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/ntlm_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/oauth2.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/spnego_gssapi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/spnego_sspi.c
        ${ANDROID_ROOT}/external/curl/lib/vauth/vauth.c
        ${ANDROID_ROOT}/external/curl/lib/vquic/curl_msh3.c
        ${ANDROID_ROOT}/external/curl/lib/vquic/curl_ngtcp2.c
        ${ANDROID_ROOT}/external/curl/lib/vquic/curl_quiche.c
        ${ANDROID_ROOT}/external/curl/lib/vquic/vquic.c
        ${ANDROID_ROOT}/external/curl/lib/vssh/libssh.c
        ${ANDROID_ROOT}/external/curl/lib/vssh/libssh2.c
        ${ANDROID_ROOT}/external/curl/lib/vssh/wolfssh.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/bearssl.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/gskit.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/gtls.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/hostcheck.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/keylog.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/mbedtls.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/mbedtls_threadlock.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/nss.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/openssl.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/rustls.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/schannel.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/schannel_verify.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/sectransp.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/vtls.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/wolfssl.c
        ${ANDROID_ROOT}/external/curl/lib/vtls/x509asn1.c
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
    add_library(libcurl SHARED ${SOURCE_FILES})

    target_include_directories(libcurl
            PUBLIC
            ${ANDROID_ROOT}/external/curl/include
            ${ANDROID_ROOT}/external/curl/lib
            ${ANDROID_ROOT}/external/curl
    )

    # Add c/c++ flags
    target_compile_options(libcurl PRIVATE
            -Werror
    )

    target_link_libraries(libcurl
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libssl
            libz
    )
endif ()
#################### libcurl ####################

#################### curl ####################
set(FILES
        ${ANDROID_ROOT}/external/curl/src/slist_wc.c
        ${ANDROID_ROOT}/external/curl/src/tool_binmode.c
        ${ANDROID_ROOT}/external/curl/src/tool_bname.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_dbg.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_hdr.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_prg.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_rea.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_see.c
        ${ANDROID_ROOT}/external/curl/src/tool_cb_wrt.c
        ${ANDROID_ROOT}/external/curl/src/tool_cfgable.c
        ${ANDROID_ROOT}/external/curl/src/tool_dirhie.c
        ${ANDROID_ROOT}/external/curl/src/tool_doswin.c
        ${ANDROID_ROOT}/external/curl/src/tool_easysrc.c
        ${ANDROID_ROOT}/external/curl/src/tool_filetime.c
        ${ANDROID_ROOT}/external/curl/src/tool_findfile.c
        ${ANDROID_ROOT}/external/curl/src/tool_formparse.c
        ${ANDROID_ROOT}/external/curl/src/tool_getparam.c
        ${ANDROID_ROOT}/external/curl/src/tool_getpass.c
        ${ANDROID_ROOT}/external/curl/src/tool_help.c
        ${ANDROID_ROOT}/external/curl/src/tool_helpers.c
        ${ANDROID_ROOT}/external/curl/src/tool_hugehelp.c
        ${ANDROID_ROOT}/external/curl/src/tool_libinfo.c
        ${ANDROID_ROOT}/external/curl/src/tool_listhelp.c
        ${ANDROID_ROOT}/external/curl/src/tool_main.c
        ${ANDROID_ROOT}/external/curl/src/tool_msgs.c
        ${ANDROID_ROOT}/external/curl/src/tool_operate.c
        ${ANDROID_ROOT}/external/curl/src/tool_operhlp.c
        ${ANDROID_ROOT}/external/curl/src/tool_paramhlp.c
        ${ANDROID_ROOT}/external/curl/src/tool_parsecfg.c
        ${ANDROID_ROOT}/external/curl/src/tool_progress.c
        ${ANDROID_ROOT}/external/curl/src/tool_setopt.c
        ${ANDROID_ROOT}/external/curl/src/tool_sleep.c
        ${ANDROID_ROOT}/external/curl/src/tool_stderr.c
        ${ANDROID_ROOT}/external/curl/src/tool_strdup.c
        ${ANDROID_ROOT}/external/curl/src/tool_urlglob.c
        ${ANDROID_ROOT}/external/curl/src/tool_util.c
        ${ANDROID_ROOT}/external/curl/src/tool_vms.c
        ${ANDROID_ROOT}/external/curl/src/tool_writeout.c
        ${ANDROID_ROOT}/external/curl/src/tool_writeout_json.c
        ${ANDROID_ROOT}/external/curl/src/tool_xattr.c
        ${ANDROID_ROOT}/external/curl/lib/dynbuf.c
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
    add_executable(curl ${SOURCE_FILES})

    target_include_directories(curl
            PUBLIC
            ${ANDROID_ROOT}/external/curl/include
            ${ANDROID_ROOT}/external/curl/lib
            ${ANDROID_ROOT}/external/curl
    )

    # Add c/c++ flags
    target_compile_options(curl PRIVATE
            -Werror
    )

    target_link_libraries(curl
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcurl
            libdl
            libm
            libssl
            libz
    )
endif ()
#################### curl ####################
