package parsers

var (
	// from asm-generic/fcntl.h
	// NOT sequential values
	// gap
	O_DIRECT    = SystemFunctionArgument{rawValue:  040000, stringValue: "O_DIRECT"}
	O_LARGEFILE = SystemFunctionArgument{rawValue: 0100000, stringValue: "O_LARGEFILE"}
	O_DIRECTORY = SystemFunctionArgument{rawValue: 0200000, stringValue: "O_DIRECTORY"}
	O_NOFOLLOW  = SystemFunctionArgument{rawValue: 0400000, stringValue: "O_NOFOLLOW"}
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
	PTRACE_SINGLEBLOCK =  SystemFunctionArgument{rawValue: 12, stringValue: "PTRACE_SINGLEBLOCK"}
	// gap
	PTRACE_OLDSETOPTIONS =  SystemFunctionArgument{rawValue: 21, stringValue: "PTRACE_OLDSETOPTIONS"}
	// gap
	PTRACE_PEEKUSR_AREA = SystemFunctionArgument{rawValue: 0x5000, stringValue: "PTRACE_PEEKUSR_AREA"}
	PTRACE_POKEUSR_AREA = SystemFunctionArgument{rawValue: 0x5001, stringValue: "PTRACE_POKEUSR_AREA"}
	PTRACE_PEEKTEXT_AREA = SystemFunctionArgument{rawValue: 0x5002, stringValue: "PTRACE_PEEKTEXT_AREA"}
	PTRACE_PEEKDATA_AREA = SystemFunctionArgument{rawValue: 0x5003, stringValue: "PTRACE_PEEKDATA_AREA"}
	PTRACE_POKETEXT_AREA = SystemFunctionArgument{rawValue: 0x5004, stringValue: "PTRACE_POKETEXT_AREA"}
	PTRACE_POKEDATA_AREA = SystemFunctionArgument{rawValue: 0x5005, stringValue: "PTRACE_POKEDATA_AREA"}
	PTRACE_GET_LAST_BREAK = SystemFunctionArgument{rawValue: 0x5006, stringValue: "PTRACE_GET_LAST_BREAK"}
	PTRACE_PEEK_SYSTEM_CALL = SystemFunctionArgument{rawValue: 0x5007, stringValue: "PTRACE_PEEK_SYSTEM_CALL"}
	PTRACE_POKE_SYSTEM_CALL = SystemFunctionArgument{rawValue: 0x5008, stringValue: "PTRACE_POKE_SYSTEM_CALL"}
	PTRACE_ENABLE_TE = SystemFunctionArgument{rawValue: 0x5009, stringValue: "PTRACE_ENABLE_TE"}
	PTRACE_DISABLE_TE = SystemFunctionArgument{rawValue: 0x5010, stringValue: "PTRACE_DISABLE_TE"}
	PTRACE_TE_ABORT_RAND = SystemFunctionArgument{rawValue: 0x5011, stringValue: "PTRACE_TE_ABORT_RAND"}
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
	PTRACE_SINGLEBLOCK,
	PTRACE_ATTACH,
	PTRACE_DETACH,
	PTRACE_OLDSETOPTIONS,
	PTRACE_SYSCALL,
	PTRACE_SYSEMU,
	PTRACE_SYSEMU_SINGLESTEP,
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
	// S/390 specific non posix ptrace requests
	PTRACE_PEEKUSR_AREA,
	PTRACE_POKEUSR_AREA,
	PTRACE_PEEKTEXT_AREA,
	PTRACE_PEEKDATA_AREA,
	PTRACE_POKETEXT_AREA,
	PTRACE_POKEDATA_AREA,
	PTRACE_GET_LAST_BREAK,
	PTRACE_PEEK_SYSTEM_CALL,
	PTRACE_POKE_SYSTEM_CALL,
	PTRACE_ENABLE_TE,
	PTRACE_DISABLE_TE,
	PTRACE_TE_ABORT_RAND,
}
