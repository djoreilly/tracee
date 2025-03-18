//go:build (ppc64 || ppc64le)

package probes

const SyscallPrefix = "sys_"
const SyscallPrefixCompat = "compat_sys_"
const SyscallPrefixCompat2 = "NOT_SUPPORTED"
