//go:build (ppc64le || ppc64)

package parsers

var (
	// from asm-generic/fcntl.h
	// NOT sequential values
	// gap
	O_DIRECTORY = SystemFunctionArgument{rawValue:  040000, stringValue: "O_DIRECTORY"}
	O_NOFOLLOW  = SystemFunctionArgument{rawValue: 0100000, stringValue: "O_NOFOLLOW"}
	O_LARGEFILE = SystemFunctionArgument{rawValue: 0200000, stringValue: "O_LARGEFILE"}
	O_DIRECT    = SystemFunctionArgument{rawValue: 0400000, stringValue: "O_DIRECT"}
)

var openFlagsValues = []SystemFunctionArgument{
	// O_ACCMODE, // macro for access mode, so not included

	// special cases checked before the loop in ParseOpenFlagArgument
	// O_RDONLY,
	// O_WRONLY,
	// O_RDWR,
	O_CREAT,
	O_EXCL,
	O_NOCTTY,
	O_TRUNC,
	O_APPEND,
	O_NONBLOCK,
	O_DSYNC,
	O_SYNC,
	FASYNC,
	O_DIRECT,
	O_LARGEFILE,
	O_DIRECTORY,
	O_NOFOLLOW,
	O_NOATIME,
	O_CLOEXEC,
	O_PATH,
	O_TMPFILE,
}

var (
	// from linux/ptrace.h and sys/ptrace.h
	// NOT sequential values
	// gap
	PTRACE_GETREGS   = SystemFunctionArgument{rawValue: 12, stringValue: "PTRACE_GETREGS"}
	PTRACE_SETREGS   = SystemFunctionArgument{rawValue: 13, stringValue: "PTRACE_SETREGS"}
	PTRACE_GETFPREGS = SystemFunctionArgument{rawValue: 14, stringValue: "PTRACE_GETFPREGS"}
	PTRACE_SETFPREGS = SystemFunctionArgument{rawValue: 15, stringValue: "PTRACE_SETFPREGS"}
	// gap
	PTRACE_GETVRREGS = SystemFunctionArgument{rawValue: 18, stringValue: "PTRACE_GETVRREGS"}
	PTRACE_SETVRREGS = SystemFunctionArgument{rawValue: 19, stringValue: "PTRACE_SETVRREGS"}
	PTRACE_GETEVRREGS = SystemFunctionArgument{rawValue: 20, stringValue: "PTRACE_GETEVRREGS"}
	PTRACE_SETEVRREGS = SystemFunctionArgument{rawValue: 21, stringValue: "PTRACE_SETEVRREGS"}
	PTRACE_GETREGS64   = SystemFunctionArgument{rawValue: 22, stringValue: "PTRACE_GETREGS64"}
	PTRACE_SETREGS64   = SystemFunctionArgument{rawValue: 23, stringValue: "PTRACE_SETREGS64"}
	// gap
	PTRACE_PEEKMTETAGS = SystemFunctionArgument{rawValue: 33, stringValue: "PTRACE_PEEKMTETAGS"}
	PTRACE_POKEMTETAGS = SystemFunctionArgument{rawValue: 34, stringValue: "PTRACE_POKEMTETAGS"}
	// gap
	PTRACE_SINGLEBLOCK = SystemFunctionArgument{rawValue: 256, stringValue: "PTRACE_SINGLEBLOCK"}
)

var ptraceRequestValues = []SystemFunctionArgument{
	PTRACE_TRACEME,
	PTRACE_PEEKTEXT,
	PTRACE_PEEKDATA,
	PTRACE_PEEKUSR,
	PTRACE_POKETEXT,
	PTRACE_POKEDATA,
	PTRACE_POKEUSR,
	PTRACE_CONT,
	PTRACE_KILL,
	PTRACE_SINGLESTEP,
	PTRACE_GETREGS,
	PTRACE_SETREGS,
	PTRACE_GETFPREGS,
	PTRACE_SETFPREGS,
	PTRACE_ATTACH,
	PTRACE_DETACH,
	PTRACE_SYSCALL,
	PTRACE_GETVRREGS,
	PTRACE_SETVRREGS,
	PTRACE_GETEVRREGS,
	PTRACE_SETEVRREGS,
	PTRACE_GETREGS64,
	PTRACE_SETREGS64,
	PTRACE_SYSEMU,
	PTRACE_SYSEMU_SINGLESTEP,
	PTRACE_SINGLEBLOCK,
	PTRACE_PEEKMTETAGS,
	PTRACE_POKEMTETAGS,
	PTRACE_SETOPTIONS,
	PTRACE_GETEVENTMSG,
	PTRACE_GETSIGINFO,
	PTRACE_SETSIGINFO,
	PTRACE_GETREGSET,
	PTRACE_SETREGSET,
	PTRACE_SEIZE,
	PTRACE_INTERRUPT,
	PTRACE_LISTEN,
	PTRACE_PEEKSIGINFO,
	PTRACE_GETSIGMASK,
	PTRACE_SETSIGMASK,
	PTRACE_SECCOMP_GET_FILTER,
	PTRACE_SECCOMP_GET_METADATA,
	PTRACE_GET_SYSCALL_INFO,
	PTRACE_GET_RSEQ_CONFIGURATION,
	PTRACE_SET_SYSCALL_USER_DISPATCH_CONFIG,
	PTRACE_GET_SYSCALL_USER_DISPATCH_CONFIG,
}
