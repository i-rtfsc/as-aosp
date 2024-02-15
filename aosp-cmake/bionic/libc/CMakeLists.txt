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

project(libc_jemalloc_wrapper)

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

#################### libc_jemalloc_wrapper ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/jemalloc_wrapper.cpp
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
    add_library(libc_jemalloc_wrapper SHARED ${SOURCE_FILES})

    target_include_directories(libc_jemalloc_wrapper
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_jemalloc_wrapper PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_jemalloc_wrapper
            libc_headers
            libclang_rt.builtins
            libjemalloc5
            liblog_headers
    )
endif ()
#################### libc_jemalloc_wrapper ####################

#################### libc_netbsd ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/common/lib/libc/stdlib/random.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/gen/nice.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/gen/psignal.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/gen/utime.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/inet/nsap_addr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/regex/regcomp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/regex/regerror.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/regex/regexec.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/regex/regfree.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/bsearch.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/drand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/erand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/jrand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/lcong48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/lrand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/mrand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/nrand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/_rand48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/rand_r.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/reallocarr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/seed48.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/stdlib/srand48.c
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
    add_library(libc_netbsd SHARED ${SOURCE_FILES})

    target_include_directories(libc_netbsd
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_netbsd PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_netbsd
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_netbsd ####################

#################### libc_freebsd ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/gen/ldexp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/getopt_long.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/hcreate.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/hcreate_r.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/hdestroy_r.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/hsearch_r.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/qsort.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/stdlib/quick_exit.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcpcpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcpncpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcscasecmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcscat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcschr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcscmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcscpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcscspn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsdup.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcslcat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcslen.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsncasecmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsncat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsncmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsncpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsnlen.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcspbrk.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsrchr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsspn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcsstr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wcstok.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wmemchr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wmemcmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wmemcpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wmemmove.c
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/string/wmemset.c
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
    add_library(libc_freebsd SHARED ${SOURCE_FILES})

    target_include_directories(libc_freebsd
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/android/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_freebsd PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_freebsd
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_freebsd ####################

#################### libc_dynamic_dispatch ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/dynamic_function_dispatch.cpp
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
    add_library(libc_dynamic_dispatch SHARED ${SOURCE_FILES})

    target_include_directories(libc_dynamic_dispatch
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_dynamic_dispatch PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_dynamic_dispatch
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_dynamic_dispatch ####################

#################### libc_ndk ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/ether_aton.c
        ${ANDROID_ROOT}/bionic/libc/bionic/ether_ntoa.c
        ${ANDROID_ROOT}/bionic/libc/bionic/exit.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/initgroups.c
        ${ANDROID_ROOT}/bionic/libc/bionic/isatty.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_cpualloc.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_cpucount.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sysprop_helpers.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/fmemopen.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/parsefloat.c
        ${ANDROID_ROOT}/bionic/libc/stdio/refill.c
        ${ANDROID_ROOT}/bionic/libc/stdio/stdio.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/stdio_ext.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/vfscanf.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/vfwscanf.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/gwp_asan_wrappers.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/heap_tagging.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/malloc_common.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/malloc_limit.cpp
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
    add_library(libc_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libc_ndk
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libc_ndk PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_ndk
            gwp_asan
            gwp_asan_crash_handler
            gwp_asan_headers
            libarm-optimized-routines-string
            libasync_safe
            libc_bionic_ndk
            libc_bootstrap
            libc_fortify
            libc_freebsd
            libc_freebsd_large_stack
            libc_gdtoa
            libc_headers
            libc_netbsd
            libc_openbsd_large_stack
            libc_openbsd_ndk
            libc_syscalls
            libc_tzcode
            libclang_rt.builtins
            liblog_headers
            libm
            libscudo
            libstdc++
    )
endif ()
#################### libc_ndk ####################

#################### libseccomp_policy ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/seccomp_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_sources_arm/android_${OUT_ARCH_CPU}/gen/arm_app_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_sources_arm/android_${OUT_ARCH_CPU}/gen/arm64_app_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_zygote_sources_arm/android_${OUT_ARCH_CPU}/gen/arm_app_zygote_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_zygote_sources_arm/android_${OUT_ARCH_CPU}/gen/arm64_app_zygote_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_system_sources_arm/android_${OUT_ARCH_CPU}/gen/arm_system_policy.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_system_sources_arm/android_${OUT_ARCH_CPU}/gen/arm64_system_policy.cpp
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
    add_library(libseccomp_policy SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_policy
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/seccomp/include
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
            ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_sources_arm/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_app_zygote_sources_arm/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/bionic/libc/libseccomp_policy_system_sources_arm/android_${OUT_ARCH_CPU}/gen
            ${OUT_INTERMEDIATES_ROOT}/bionic/libc/func_to_syscall_nrs/android_${OUT_ARCH_CPU}/gen
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_policy PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_policy
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libseccomp_policy ####################

#################### libc_syscalls ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/__set_errno.cpp
        ${OUT_INTERMEDIATES_ROOT}/bionic/libc/syscalls-arm64/gen/syscalls-arm64.S
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
    add_library(libc_syscalls SHARED ${SOURCE_FILES})

    target_include_directories(libc_syscalls
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_syscalls PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_syscalls
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_syscalls ####################

#################### libseccomp_gen_syscall_nrs_riscv64 ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/gen_syscall_nrs.cpp
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
    add_library(libseccomp_gen_syscall_nrs_riscv64 SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_gen_syscall_nrs_riscv64
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-riscv
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_gen_syscall_nrs_riscv64 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_gen_syscall_nrs_riscv64
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libseccomp_gen_syscall_nrs_riscv64 ####################

#################### crtbegin_static ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/arch-common/bionic/crtbegin.c
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
    add_library(crtbegin_static SHARED ${SOURCE_FILES})

    target_include_directories(crtbegin_static
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/include
            ${ANDROID_ROOT}/bionic/libc/bionic
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(crtbegin_static PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(crtbegin_static
    )
endif ()
#################### crtbegin_static ####################

#################### libseccomp_gen_syscall_nrs_x86 ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/gen_syscall_nrs_x86.cpp
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
    add_library(libseccomp_gen_syscall_nrs_x86 SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_gen_syscall_nrs_x86
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-x86
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_gen_syscall_nrs_x86 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_gen_syscall_nrs_x86
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libseccomp_gen_syscall_nrs_x86 ####################

#################### crtbegin_so ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/arch-common/bionic/crtbegin_so.c
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
    add_library(crtbegin_so SHARED ${SOURCE_FILES})

    target_include_directories(crtbegin_so
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/include
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(crtbegin_so PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(crtbegin_so
    )
endif ()
#################### crtbegin_so ####################

#################### libc_bootstrap ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/__libc_init_main_thread.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__stack_chk_fail.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_call_ifunc_resolver.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getauxval.cpp
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/__set_tls.c
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
    add_library(libc_bootstrap SHARED ${SOURCE_FILES})

    target_include_directories(libc_bootstrap
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_bootstrap PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_bootstrap
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_bootstrap ####################

#################### libc_bionic_systrace ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_systrace.cpp
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
    add_library(libc_bionic_systrace SHARED ${SOURCE_FILES})

    target_include_directories(libc_bionic_systrace
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_bionic_systrace PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_bionic_systrace
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_bionic_systrace ####################

#################### libc_bionic ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/sysconf.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/vdso.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/setjmp_cookie.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/android_set_abort_message.cpp
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/kryo785/bionic/__memcpy.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/kryo785/bionic/memcpy.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/kryo785/bionic/memmove.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/__bionic_clone.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/_exit_with_stack_teardown.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/setjmp.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/syscall.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/bionic/vfork.S
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
    add_library(libc_bionic SHARED ${SOURCE_FILES})

    target_include_directories(libc_bionic
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_bionic PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_bionic
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_bionic ####################

#################### libc_tzcode ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/tzcode/asctime.c
        ${ANDROID_ROOT}/bionic/libc/tzcode/difftime.c
        ${ANDROID_ROOT}/bionic/libc/tzcode/localtime.c
        ${ANDROID_ROOT}/bionic/libc/tzcode/strftime.c
        ${ANDROID_ROOT}/bionic/libc/tzcode/strptime.c
        ${ANDROID_ROOT}/bionic/libc/tzcode/bionic.cpp
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/_def_time.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/time/wcsftime.c
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
    add_library(libc_tzcode SHARED ${SOURCE_FILES})

    target_include_directories(libc_tzcode
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/tzcode
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_tzcode PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_tzcode
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_tzcode ####################

#################### libc_openbsd ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/crypt/arc4random.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/crypt/arc4random_uniform.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/stpncpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strcat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strlcat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strlcpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strncat.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strncpy.c
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
    add_library(libc_openbsd SHARED ${SOURCE_FILES})

    target_include_directories(libc_openbsd
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/private
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_openbsd PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_openbsd
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_openbsd ####################

#################### libseccomp_gen_syscall_nrs_arm ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/gen_syscall_nrs.cpp
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
    add_library(libseccomp_gen_syscall_nrs_arm SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_gen_syscall_nrs_arm
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_gen_syscall_nrs_arm PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_gen_syscall_nrs_arm
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libseccomp_gen_syscall_nrs_arm ####################

#################### libseccomp_gen_syscall_nrs_x86_64 ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/gen_syscall_nrs_x86_64.cpp
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
    add_library(libseccomp_gen_syscall_nrs_x86_64 SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_gen_syscall_nrs_x86_64
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-x86
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_gen_syscall_nrs_x86_64 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_gen_syscall_nrs_x86_64
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libseccomp_gen_syscall_nrs_x86_64 ####################

#################### libc_freebsd_large_stack ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/lib/libc/gen/glob.c
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
    add_library(libc_freebsd_large_stack SHARED ${SOURCE_FILES})

    target_include_directories(libc_freebsd_large_stack
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/upstream-freebsd/android/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_freebsd_large_stack PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_freebsd_large_stack
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_freebsd_large_stack ####################

#################### libc_init_dynamic ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/libc_init_dynamic.cpp
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
    add_library(libc_init_dynamic SHARED ${SOURCE_FILES})

    target_include_directories(libc_init_dynamic
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_init_dynamic PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_init_dynamic
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_init_dynamic ####################

#################### libstdc++ ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/__cxa_guard.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__cxa_pure_virtual.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/new.cpp
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
    add_library(libstdc++ SHARED ${SOURCE_FILES})

    target_include_directories(libstdc++
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libstdc++ PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libstdc++
            libasync_safe
            libc
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libstdc++ ####################

#################### libseccomp_gen_syscall_nrs_arm64 ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/seccomp/gen_syscall_nrs.cpp
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
    add_library(libseccomp_gen_syscall_nrs_arm64 SHARED ${SOURCE_FILES})

    target_include_directories(libseccomp_gen_syscall_nrs_arm64
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64
            ${ANDROID_ROOT}/bionic/libc/kernel/uapi
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libseccomp_gen_syscall_nrs_arm64 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libseccomp_gen_syscall_nrs_arm64
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### libseccomp_gen_syscall_nrs_arm64 ####################

#################### libc_bionic_ndk ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/NetdClientDispatch.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__bionic_get_shell_path.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__cmsg_nxthdr.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__errno.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__gnu_basename.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__libc_current_sigrtmax.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__libc_current_sigrtmin.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/abort.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/accept.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/access.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/arpa_inet.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/assert.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/atof.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_allocator.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_arc4random.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_futex.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_netlink.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_time_conversions.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/brk.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/c16rtomb.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/c32rtomb.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/chmod.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/chown.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/clearenv.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/clock.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/clock_getcpuclockid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/clock_nanosleep.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/clone.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ctype.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/dirent.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/dup.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/environ.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/error.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/eventfd.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/exec.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/execinfo.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/faccessat.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fchmod.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fchmodat.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fcntl.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fdsan.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fdtrack.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ffs.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fgetxattr.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/flistxattr.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fpclassify.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fsetxattr.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ftruncate.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ftw.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/futimens.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getcwd.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getdomainname.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getentropy.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/gethostname.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getloadavg.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getpagesize.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getpgrp.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getpid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/getpriority.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/gettid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/get_device_api_level.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/grp_pwd.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/grp_pwd_file.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/heap_zero_init.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/iconv.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/icu_wrappers.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ifaddrs.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/inotify_init.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ioctl.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/killpg.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/langinfo.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/lchown.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/lfs64_support.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/libc_init_common.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/libgen.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/link.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/locale.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/lockf.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/lstat.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mblen.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mbrtoc16.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mbrtoc32.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mempcpy.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/memset_explicit.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mkdir.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mkfifo.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mknod.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mntent.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/mremap.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/net_if.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/netdb.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/netinet_in.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/nl_types.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/open.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pathconf.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pause.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pidfd.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pipe.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/poll.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/posix_fadvise.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/posix_fallocate.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/posix_madvise.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/posix_timers.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/preadv_pwritev.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/ptrace.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pty.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/raise.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/rand.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/readlink.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/realpath.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/reboot.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/recv.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/recvmsg.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/rename.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/rmdir.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/scandir.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_getaffinity.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_getcpu.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/semaphore.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/send.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/setegid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/seteuid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/setpgrp.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sigaction.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/signal.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sigprocmask.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sleep.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/socketpair.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/spawn.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/stat.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/stdlib_l.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strerror.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/string_l.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strings_l.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strsignal.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strtol.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strtold.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/swab.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/symlink.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sync_file_range.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_epoll.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_msg.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_sem.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_shm.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_signalfd.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_statfs.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_statvfs.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_time.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sysinfo.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/syslog.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/system.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/system_property_api.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/system_property_set.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/tdestroy.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/termios.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/thread_private.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/threads.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/time_l.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/timespec_get.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/tmpfile.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/umount.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/unlink.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/usleep.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/utmp.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wait.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wchar.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wchar_l.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wcstod.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wctype.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wcwidth.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/wmempcpy.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/icu_static.cpp
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
    add_library(libc_bionic_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libc_bionic_ndk
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/stdio
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
            ${OUT_INTERMEDIATES_ROOT}/bionic/libc/generated_android_ids/gen
    )

    # Add c/c++ flags
    target_compile_options(libc_bionic_ndk PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_bionic_ndk
            libc_bionic_systrace
            libc_headers
            libclang_rt.builtins
            liblog_headers
            libsystemproperties
    )
endif ()
#################### libc_bionic_ndk ####################

#################### libc_pthread ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/bionic_elf_tls.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_atfork.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_attr.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_barrier.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_cond.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_create.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_detach.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_equal.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_exit.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_getcpuclockid.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_getschedparam.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_gettid_np.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_internal.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_join.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_key.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_kill.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_mutex.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_once.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_rwlock.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_sigqueue.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_self.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_setname_np.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_setschedparam.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/pthread_spinlock.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/sys_thread_properties.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/__cxa_thread_atexit_impl.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/android_unsafe_frame_pointer_chase.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/atexit.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/fork.cpp
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
    add_library(libc_pthread SHARED ${SOURCE_FILES})

    target_include_directories(libc_pthread
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_pthread PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_pthread
            bionic_libc_platform_headers
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_pthread ####################

#################### libc_openbsd_ndk ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/alarm.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/ctype_.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/daemon.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/err.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/errx.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/fnmatch.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/ftok.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/getprogname.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/setprogname.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/verr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/verrx.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/vwarn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/vwarnx.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/warn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gen/warnx.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/btowc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/mbrlen.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/mbstowcs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/mbtowc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/wcscoll.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/wcstombs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/wcsxfrm.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/wctob.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/locale/wctomb.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/base64.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/htonl.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/htons.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_lnaof.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_makeaddr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_netof.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_ntoa.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_ntop.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/inet_pton.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/ntohl.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/ntohs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/net/res_random.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fgetln.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fgetwc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fgetws.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/flags.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fpurge.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fputwc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fputws.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fvwrite.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/fwide.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/getdelim.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/gets.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/makebuf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/mktemp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/open_memstream.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/open_wmemstream.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/rget.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/setvbuf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/ungetc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/ungetwc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/vasprintf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/vdprintf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/vsscanf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/vswprintf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/vswscanf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/wbuf.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio/wsetup.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/abs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/div.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/getenv.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/getsubopt.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/insque.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/imaxabs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/imaxdiv.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/labs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/ldiv.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/llabs.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/lldiv.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/lsearch.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/recallocarray.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/remque.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/setenv.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/tfind.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdlib/tsearch.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/memccpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strcasecmp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strcasestr.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strcoll.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strcspn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strdup.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strndup.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strpbrk.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strsep.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strspn.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strtok.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strxfrm.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/wcslcpy.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/wcswidth.c
        ${ANDROID_ROOT}/bionic/libc/bionic/fts.c
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
    add_library(libc_openbsd_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libc_openbsd_ndk
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/private
            ${ANDROID_ROOT}/bionic/libc/stdio
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/include
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_openbsd_ndk PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_openbsd_ndk
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_openbsd_ndk ####################

#################### libc_openbsd_large_stack ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/stdio/vfprintf.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/vfwprintf.cpp
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/memmem.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/string/strstr.c
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
    add_library(libc_openbsd_large_stack SHARED ${SOURCE_FILES})

    target_include_directories(libc_openbsd_large_stack
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/private
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/include
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/stdio
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_openbsd_large_stack PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_openbsd_large_stack
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_openbsd_large_stack ####################

#################### libc_nopthread ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/ether_aton.c
        ${ANDROID_ROOT}/bionic/libc/bionic/ether_ntoa.c
        ${ANDROID_ROOT}/bionic/libc/bionic/exit.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/initgroups.c
        ${ANDROID_ROOT}/bionic/libc/bionic/isatty.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_cpualloc.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sched_cpucount.c
        ${ANDROID_ROOT}/bionic/libc/bionic/sysprop_helpers.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/fmemopen.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/parsefloat.c
        ${ANDROID_ROOT}/bionic/libc/stdio/refill.c
        ${ANDROID_ROOT}/bionic/libc/stdio/stdio.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/stdio_ext.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/vfscanf.cpp
        ${ANDROID_ROOT}/bionic/libc/stdio/vfwscanf.cpp
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
    add_library(libc_nopthread SHARED ${SOURCE_FILES})

    target_include_directories(libc_nopthread
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
    )

    # Add c/c++ flags
    target_compile_options(libc_nopthread PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_nopthread
            libarm-optimized-routines-string
            libasync_safe
            libc_bionic
            libc_bionic_ndk
            libc_bootstrap
            libc_dns
            libc_fortify
            libc_freebsd
            libc_freebsd_large_stack
            libc_gdtoa
            libc_headers
            libc_netbsd
            libc_openbsd
            libc_openbsd_large_stack
            libc_openbsd_ndk
            libc_syscalls
            libc_tzcode
            libclang_rt.builtins
            liblog_headers
            libstdc++
    )
endif ()
#################### libc_nopthread ####################

#################### librust_baremetal ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/fortify.cpp
        ${ANDROID_ROOT}/bionic/libc/bionic/strtol.cpp
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/string/__memcpy_chk.S
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
    add_library(librust_baremetal SHARED ${SOURCE_FILES})

    target_include_directories(librust_baremetal
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(librust_baremetal PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librust_baremetal
            libarm-optimized-routines-mem
            libc_headers
            libc_netbsd
            libclang_rt.builtins
    )
endif ()
#################### librust_baremetal ####################

#################### libc_gdtoa ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/android/gdtoa_support.cpp
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/dmisc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/dtoa.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/gdtoa.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/gethex.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/gmisc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/hd_init.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/hdtoa.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/hexnan.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/ldtoa.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/misc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/smisc.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/strtod.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/strtodg.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/strtof.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/strtord.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/sum.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/ulp.c
        ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/gdtoa/strtorQ.c
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
    add_library(libc_gdtoa SHARED ${SOURCE_FILES})

    target_include_directories(libc_gdtoa
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/private
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/upstream-openbsd/lib/libc/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_gdtoa PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_gdtoa
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_gdtoa ####################

#################### libc_dns ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_name.c
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_netint.c
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_parse.c
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_print.c
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_samedomain.c
        ${ANDROID_ROOT}/bionic/libc/dns/nameser/ns_ttl.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/getaddrinfo.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/gethnamaddr.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/getnameinfo.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/getservent.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/nsdispatch.c
        ${ANDROID_ROOT}/bionic/libc/dns/net/sethostent.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/herror.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_cache.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_comp.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_data.cpp
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_debug.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_init.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_mkquery.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_query.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_send.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_state.c
        ${ANDROID_ROOT}/bionic/libc/dns/resolv/res_stats.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/isc/ev_streams.c
        ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/isc/ev_timers.c
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
    add_library(libc_dns SHARED ${SOURCE_FILES})

    target_include_directories(libc_dns
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/dns/include
            ${ANDROID_ROOT}/bionic/libc/private
            ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/lib/libc/include
            ${ANDROID_ROOT}/bionic/libc/upstream-netbsd/android/include
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_dns PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libc_dns
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_dns ####################

#################### libc_init_static ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/libc_init_static.cpp
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
    add_library(libc_init_static SHARED ${SOURCE_FILES})

    target_include_directories(libc_init_static
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_init_static PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_init_static
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_init_static ####################

#################### libc_unwind_static ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/dl_iterate_phdr_static.cpp
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
    add_library(libc_unwind_static SHARED ${SOURCE_FILES})

    target_include_directories(libc_unwind_static
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_unwind_static PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_unwind_static
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_unwind_static ####################

#################### libc_fortify ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/bionic/fortify.cpp
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/string/__memcpy_chk.S
        ${ANDROID_ROOT}/bionic/libc/arch-arm64/string/__memset_chk.S
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
    add_library(libc_fortify SHARED ${SOURCE_FILES})

    target_include_directories(libc_fortify
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/async_safe/include
            ${ANDROID_ROOT}/bionic/libc/platform
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(libc_fortify PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libc_fortify
            libc_headers
            libclang_rt.builtins
            liblog_headers
    )
endif ()
#################### libc_fortify ####################

#################### crtbegin_dynamic ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/arch-common/bionic/crtbegin.c
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
    add_library(crtbegin_dynamic SHARED ${SOURCE_FILES})

    target_include_directories(crtbegin_dynamic
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/include
            ${ANDROID_ROOT}/bionic/libc/bionic
            ${ANDROID_ROOT}/bionic/libc
    )

    # Add c/c++ flags
    target_compile_options(crtbegin_dynamic PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(crtbegin_dynamic
    )
endif ()
#################### crtbegin_dynamic ####################
