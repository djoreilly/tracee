//go:build (ppc64 || ppc64le)

package events

// ppc64 64bit syscall numbers (used as event IDs for the Syscall Events)
// https://github.com/torvalds/linux/blob/master/arch/powerpc/kernel/syscalls/syscall.tbl

const (
	RestartSyscall  	ID = 0
	Exit            	ID = 1
	Fork            	ID = 2
	Read            	ID = 3
	Write           	ID = 4
	Open            	ID = 5
	Close           	ID = 6
	Waitpid         	ID = 7
	Creat           	ID = 8
	Link            	ID = 9
	Unlink          	ID = 10
	Execve          	ID = 11
	Chdir           	ID = 12
	Time            	ID = 13
	Mknod           	ID = 14
	Chmod           	ID = 15
	Lchown          	ID = 16
	Break           	ID = 17
	Oldstat         	ID = 18
	Lseek           	ID = 19
	Getpid          	ID = 20
	Mount           	ID = 21
	Umount          	ID = 22
	Setuid          	ID = 23
	Getuid          	ID = 24
	Stime           	ID = 25
	Ptrace          	ID = 26
	Alarm           	ID = 27
	Oldfstat        	ID = 28
	Pause           	ID = 29
	Utime           	ID = 30
	Stty            	ID = 31
	Gtty            	ID = 32
	Access          	ID = 33
	Nice            	ID = 34
	Ftime           	ID = 35
	Sync            	ID = 36
	Kill            	ID = 37
	Rename          	ID = 38
	Mkdir           	ID = 39
	Rmdir           	ID = 40
	Dup             	ID = 41
	Pipe            	ID = 42
	Times           	ID = 43
	Prof            	ID = 44
	Brk             	ID = 45
	Setgid          	ID = 46
	Getgid          	ID = 47
	Signal          	ID = 48
	Geteuid         	ID = 49
	Getegid         	ID = 50
	Acct            	ID = 51
	Umount2         	ID = 52
	Lock            	ID = 53
	Ioctl           	ID = 54
	Fcntl           	ID = 55
	Mpx             	ID = 56
	Setpgid         	ID = 57
	Ulimit          	ID = 58
	Oldolduname     	ID = 59
	Umask           	ID = 60
	Chroot          	ID = 61
	Ustat           	ID = 62
	Dup2            	ID = 63
	Getppid         	ID = 64
	Getpgrp         	ID = 65
	Setsid          	ID = 66
	Sigaction       	ID = 67
	Sgetmask        	ID = 68
	Ssetmask        	ID = 69
	Setreuid        	ID = 70
	Setregid        	ID = 71
	Sigsuspend      	ID = 72
	Sigpending      	ID = 73
	Sethostname     	ID = 74
	Setrlimit       	ID = 75
	Getrlimit       	ID = 76
	Getrusage       	ID = 77
	Gettimeofday    	ID = 78
	Settimeofday    	ID = 79
	Getgroups       	ID = 80
	Setgroups       	ID = 81
	Select          	ID = 82
	Symlink         	ID = 83
	Oldlstat        	ID = 84
	Readlink        	ID = 85
	Uselib          	ID = 86
	Swapon          	ID = 87
	Reboot          	ID = 88
	Readdir         	ID = 89
	Mmap            	ID = 90
	Munmap          	ID = 91
	Truncate        	ID = 92
	Ftruncate       	ID = 93
	Fchmod          	ID = 94
	Fchown          	ID = 95
	Getpriority     	ID = 96
	Setpriority     	ID = 97
	Profil          	ID = 98
	Statfs          	ID = 99
	Fstatfs         	ID = 100
	Ioperm          	ID = 101
	Socketcall      	ID = 102
	Syslog          	ID = 103
	Setitimer       	ID = 104
	Getitimer       	ID = 105
	Stat            	ID = 106
	Lstat           	ID = 107
	Fstat           	ID = 108
	Olduname        	ID = 109
	Iopl            	ID = 110
	Vhangup         	ID = 111
	Idle            	ID = 112
	Vm86            	ID = 113
	Wait4           	ID = 114
	Swapoff         	ID = 115
	Sysinfo         	ID = 116
	Ipc             	ID = 117
	Fsync           	ID = 118
	Sigreturn       	ID = 119
	Clone           	ID = 120
	Setdomainname   	ID = 121
	Uname           	ID = 122
	ModifyLdt       	ID = 123
	Adjtimex        	ID = 124
	Mprotect        	ID = 125
	Sigprocmask     	ID = 126
	CreateModule    	ID = 127
	InitModule      	ID = 128
	DeleteModule    	ID = 129
	GetKernelSyms   	ID = 130
	Quotactl        	ID = 131
	Getpgid         	ID = 132
	Fchdir          	ID = 133
	Bdflush         	ID = 134
	Sysfs           	ID = 135
	Personality     	ID = 136
	AfsSyscall      	ID = 137
	Setfsuid        	ID = 138
	Setfsgid        	ID = 139
	Llseek          	ID = 140
	Getdents        	ID = 141
	Newselect       	ID = 142
	Flock           	ID = 143
	Msync           	ID = 144
	Readv           	ID = 145
	Writev          	ID = 146
	Getsid          	ID = 147
	Fdatasync       	ID = 148
	Sysctl          	ID = 149
	Mlock           	ID = 150
	Munlock         	ID = 151
	Mlockall        	ID = 152
	Munlockall      	ID = 153
	SchedSetparam   	ID = 154
	SchedGetparam   	ID = 155
	SchedSetscheduler	ID = 156
	SchedGetscheduler	ID = 157
	SchedYield      	ID = 158
	SchedGetPriorityMax	ID = 159
	SchedGetPriorityMin	ID = 160
	SchedRrGetInterval	ID = 161
	Nanosleep       	ID = 162
	Mremap          	ID = 163
	Setresuid       	ID = 164
	Getresuid       	ID = 165
	QueryModule     	ID = 166
	Poll            	ID = 167
	Nfsservctl      	ID = 168
	Setresgid       	ID = 169
	Getresgid       	ID = 170
	Prctl           	ID = 171
	RtSigreturn     	ID = 172
	RtSigaction     	ID = 173
	RtSigprocmask   	ID = 174
	RtSigpending    	ID = 175
	RtSigtimedwait  	ID = 176
	RtSigqueueinfo  	ID = 177
	RtSigsuspend    	ID = 178
	Pread64         	ID = 179
	Pwrite64        	ID = 180
	Chown           	ID = 181
	Getcwd          	ID = 182
	Capget          	ID = 183
	Capset          	ID = 184
	Sigaltstack     	ID = 185
	Sendfile        	ID = 186
	Getpmsg         	ID = 187
	Putpmsg         	ID = 188
	Vfork           	ID = 189
	Ugetrlimit      	ID = 190
	Readahead       	ID = 191
	PciconfigRead   	ID = 198
	PciconfigWrite  	ID = 199
	PciconfigIobase 	ID = 200
	Multiplexer     	ID = 201
	Getdents64      	ID = 202
	PivotRoot       	ID = 203
	Madvise         	ID = 205
	Mincore         	ID = 206
	Gettid          	ID = 207
	Tkill           	ID = 208
	Setxattr        	ID = 209
	Lsetxattr       	ID = 210
	Fsetxattr       	ID = 211
	Getxattr        	ID = 212
	Lgetxattr       	ID = 213
	Fgetxattr       	ID = 214
	Listxattr       	ID = 215
	Llistxattr      	ID = 216
	Flistxattr      	ID = 217
	Removexattr     	ID = 218
	Lremovexattr    	ID = 219
	Fremovexattr    	ID = 220
	Futex           	ID = 221
	SchedSetaffinity	ID = 222
	SchedGetaffinity	ID = 223
	Tuxcall         	ID = 225
	IoSetup         	ID = 227
	IoDestroy       	ID = 228
	IoGetevents     	ID = 229
	IoSubmit        	ID = 230
	IoCancel        	ID = 231
	SetTidAddress   	ID = 232
	Fadvise64       	ID = 233
	ExitGroup       	ID = 234
	LookupDcookie   	ID = 235
	EpollCreate     	ID = 236
	EpollCtl        	ID = 237
	EpollWait       	ID = 238
	RemapFilePages  	ID = 239
	TimerCreate     	ID = 240
	TimerSettime    	ID = 241
	TimerGettime    	ID = 242
	TimerGetoverrun 	ID = 243
	TimerDelete     	ID = 244
	ClockSettime    	ID = 245
	ClockGettime    	ID = 246
	ClockGetres     	ID = 247
	ClockNanosleep  	ID = 248
	Swapcontext     	ID = 249
	Tgkill          	ID = 250
	Utimes          	ID = 251
	Statfs64        	ID = 252
	Fstatfs64       	ID = 253
	Rtas            	ID = 255
	SysDebugSetcontext	ID = 256
	MigratePages    	ID = 258
	Mbind           	ID = 259
	GetMempolicy    	ID = 260
	SetMempolicy    	ID = 261
	MqOpen          	ID = 262
	MqUnlink        	ID = 263
	MqTimedsend     	ID = 264
	MqTimedreceive  	ID = 265
	MqNotify        	ID = 266
	MqGetsetattr    	ID = 267
	KexecLoad       	ID = 268
	AddKey          	ID = 269
	RequestKey      	ID = 270
	Keyctl          	ID = 271
	Waitid          	ID = 272
	IoprioSet       	ID = 273
	IoprioGet       	ID = 274
	InotifyInit     	ID = 275
	InotifyAddWatch 	ID = 276
	InotifyRmWatch  	ID = 277
	SpuRun          	ID = 278
	SpuCreate       	ID = 279
	Pselect6        	ID = 280
	Ppoll           	ID = 281
	Unshare         	ID = 282
	Splice          	ID = 283
	Tee             	ID = 284
	Vmsplice        	ID = 285
	Openat          	ID = 286
	Mkdirat         	ID = 287
	Mknodat         	ID = 288
	Fchownat        	ID = 289
	Futimesat       	ID = 290
	Newfstatat      	ID = 291
	Unlinkat        	ID = 292
	Renameat        	ID = 293
	Linkat          	ID = 294
	Symlinkat       	ID = 295
	Readlinkat      	ID = 296
	Fchmodat        	ID = 297
	Faccessat       	ID = 298
	GetRobustList   	ID = 299
	SetRobustList   	ID = 300
	MovePages       	ID = 301
	Getcpu          	ID = 302
	EpollPwait      	ID = 303
	Utimensat       	ID = 304
	Signalfd        	ID = 305
	TimerfdCreate   	ID = 306
	Eventfd         	ID = 307
	SyncFileRange2  	ID = 308
	Fallocate       	ID = 309
	SubpageProt     	ID = 310
	TimerfdSettime  	ID = 311
	TimerfdGettime  	ID = 312
	Signalfd4       	ID = 313
	Eventfd2        	ID = 314
	EpollCreate1    	ID = 315
	Dup3            	ID = 316
	Pipe2           	ID = 317
	InotifyInit1    	ID = 318
	PerfEventOpen   	ID = 319
	Preadv          	ID = 320
	Pwritev         	ID = 321
	RtTgsigqueueinfo	ID = 322
	FanotifyInit    	ID = 323
	FanotifyMark    	ID = 324
	Prlimit64       	ID = 325
	Socket          	ID = 326
	Bind            	ID = 327
	Connect         	ID = 328
	Listen          	ID = 329
	Accept          	ID = 330
	Getsockname     	ID = 331
	Getpeername     	ID = 332
	Socketpair      	ID = 333
	Send            	ID = 334
	Sendto          	ID = 335
	Recv            	ID = 336
	Recvfrom        	ID = 337
	Shutdown        	ID = 338
	Setsockopt      	ID = 339
	Getsockopt      	ID = 340
	Sendmsg         	ID = 341
	Recvmsg         	ID = 342
	Recvmmsg        	ID = 343
	Accept4         	ID = 344
	NameToHandleAt  	ID = 345
	OpenByHandleAt  	ID = 346
	ClockAdjtime    	ID = 347
	Syncfs          	ID = 348
	Sendmmsg        	ID = 349
	Setns           	ID = 350
	ProcessVmReadv  	ID = 351
	ProcessVmWritev 	ID = 352
	FinitModule     	ID = 353
	Kcmp            	ID = 354
	SchedSetattr    	ID = 355
	SchedGetattr    	ID = 356
	Renameat2       	ID = 357
	Seccomp         	ID = 358
	Getrandom       	ID = 359
	MemfdCreate     	ID = 360
	Bpf             	ID = 361
	Execveat        	ID = 362
	SwitchEndian    	ID = 363
	Userfaultfd     	ID = 364
	Membarrier      	ID = 365
	Mlock2          	ID = 378
	CopyFileRange   	ID = 379
	Preadv2         	ID = 380
	Pwritev2        	ID = 381
	KexecFileLoad   	ID = 382
	Statx           	ID = 383
	PkeyAlloc       	ID = 384
	PkeyFree        	ID = 385
	PkeyMprotect    	ID = 386
	Rseq            	ID = 387
	IoPgetevents    	ID = 388
	Semtimedop      	ID = 392
	Semget          	ID = 393
	Semctl          	ID = 394
	Shmget          	ID = 395
	Shmctl          	ID = 396
	Shmat           	ID = 397
	Shmdt           	ID = 398
	Msgget          	ID = 399
	Msgsnd          	ID = 400
	Msgrcv          	ID = 401
	Msgctl          	ID = 402
	PidfdSendSignal 	ID = 424
	IoUringSetup    	ID = 425
	IoUringEnter    	ID = 426
	IoUringRegister 	ID = 427
	OpenTree        	ID = 428
	MoveMount       	ID = 429
	Fsopen          	ID = 430
	Fsconfig        	ID = 431
	Fsmount         	ID = 432
	Fspick          	ID = 433
	PidfdOpen       	ID = 434
	Clone3          	ID = 435
	CloseRange      	ID = 436
	Openat2         	ID = 437
	PidfdGetfd      	ID = 438
	Faccessat2      	ID = 439
	ProcessMadvise  	ID = 440
	EpollPwait2     	ID = 441
	MountSetattr    	ID = 442
	QuotactlFd      	ID = 443
	LandlockCreateRuleset	ID = 444
	LandlockAddRule 	ID = 445
	LandlockRestrictSelf	ID = 446
	ProcessMrelease 	ID = 448
	FutexWaitv      	ID = 449
	SetMempolicyHomeNode	ID = 450
	Cachestat       	ID = 451
	Fchmodat2       	ID = 452
	MapShadowStack  	ID = 453
	FutexWake       	ID = 454
	FutexWait       	ID = 455
	FutexRequeue    	ID = 456
	Statmount       	ID = 457
	Listmount       	ID = 458
	LsmGetSelfAttr  	ID = 459
	LsmSetSelfAttr  	ID = 460
	LsmListModules  	ID = 461
	Mseal           	ID = 462
	Setxattrat      	ID = 463
	Getxattrat      	ID = 464
	Listxattrat     	ID = 465
	Removexattrat   	ID = 466
	MaxSyscallID		ID = 449
)

// following syscalls are undefined on powerpc64
const (
	ArchPrctl ID = iota + Unsupported
	Afs
	Afs_syscall
	Chown16
	ClockAdjtime64
	ClockGetresTime32
	ClockGettime32
	ClockNanosleepTime32
	ClockSettime32
	EpollCtlOld
	EpollWaitOld
	Fadvise64_64
	Fchown16
	Fcntl64
	Fstat64
	Ftruncate64
	FutexTime32
	Getegid16
	Geteuid16
	Getgid16
	Getgroups16
	Getresgid16
	Getresuid16
	GetThreadArea
	Getuid16
	IoPgeteventsTime32
	Lchown16
	Lstat64
	MemfdSecret
	Mmap2
	MqTimedreceiveTime32
	MqTimedsendTime32
	OldGetrlimit
	OldSelect
	PpollTime32
	Pselect6Time32
	RecvmmsgTime32
	RtSigtimedwaitTime32
	SchedRrGetInterval32
	Security
	Semop
	Sendfile32
	Setfsgid16
	Setfsuid16
	Setgid16
	Setgroups16
	Setregid16
	Setresgid16
	Setresuid16
	Setreuid16
	SetThreadArea
	Setuid16
	Stat64
	SyncFileRange
	TimerfdGettime32
	TimerfdSettime32
	TimerGettime32
	TimerSettime32
	Truncate64
	UtimensatTime32
	Vm86old
	Vserver
)


// ppc64 32bit syscall numbers (used as event IDs for the Syscall Events)
// https://github.com/torvalds/linux/blob/master/arch/powerpc/kernel/syscalls/syscall.tbl


const (
	Sys32restart_syscall     	ID = 0
	Sys32exit                	ID = 1
	Sys32fork                	ID = 2
	Sys32read                	ID = 3
	Sys32write               	ID = 4
	Sys32open                	ID = 5
	Sys32close               	ID = 6
	Sys32waitpid             	ID = 7
	Sys32creat               	ID = 8
	Sys32link                	ID = 9
	Sys32unlink              	ID = 10
	Sys32execve              	ID = 11
	Sys32chdir               	ID = 12
	Sys32time                	ID = 13
	Sys32mknod               	ID = 14
	Sys32chmod               	ID = 15
	Sys32lchown              	ID = 16
	Sys32break               	ID = 17
	Sys32oldstat             	ID = 18
	Sys32lseek               	ID = 19
	Sys32getpid              	ID = 20
	Sys32mount               	ID = 21
	Sys32umount              	ID = 22
	Sys32setuid              	ID = 23
	Sys32getuid              	ID = 24
	Sys32stime               	ID = 25
	Sys32ptrace              	ID = 26
	Sys32alarm               	ID = 27
	Sys32oldfstat            	ID = 28
	Sys32pause               	ID = 29
	Sys32utime               	ID = 30
	Sys32stty                	ID = 31
	Sys32gtty                	ID = 32
	Sys32access              	ID = 33
	Sys32nice                	ID = 34
	Sys32ftime               	ID = 35
	Sys32sync                	ID = 36
	Sys32kill                	ID = 37
	Sys32rename              	ID = 38
	Sys32mkdir               	ID = 39
	Sys32rmdir               	ID = 40
	Sys32dup                 	ID = 41
	Sys32pipe                	ID = 42
	Sys32times               	ID = 43
	Sys32prof                	ID = 44
	Sys32brk                 	ID = 45
	Sys32setgid              	ID = 46
	Sys32getgid              	ID = 47
	Sys32signal              	ID = 48
	Sys32geteuid             	ID = 49
	Sys32getegid             	ID = 50
	Sys32acct                	ID = 51
	Sys32umount2             	ID = 52
	Sys32lock                	ID = 53
	Sys32ioctl               	ID = 54
	Sys32fcntl               	ID = 55
	Sys32mpx                 	ID = 56
	Sys32setpgid             	ID = 57
	Sys32ulimit              	ID = 58
	Sys32oldolduname         	ID = 59
	Sys32umask               	ID = 60
	Sys32chroot              	ID = 61
	Sys32ustat               	ID = 62
	Sys32dup2                	ID = 63
	Sys32getppid             	ID = 64
	Sys32getpgrp             	ID = 65
	Sys32setsid              	ID = 66
	Sys32sigaction           	ID = 67
	Sys32sgetmask            	ID = 68
	Sys32ssetmask            	ID = 69
	Sys32setreuid            	ID = 70
	Sys32setregid            	ID = 71
	Sys32sigsuspend          	ID = 72
	Sys32sigpending          	ID = 73
	Sys32sethostname         	ID = 74
	Sys32setrlimit           	ID = 75
	Sys32getrlimit           	ID = 76
	Sys32getrusage           	ID = 77
	Sys32gettimeofday        	ID = 78
	Sys32settimeofday        	ID = 79
	Sys32getgroups           	ID = 80
	Sys32setgroups           	ID = 81
	Sys32select              	ID = 82
	Sys32symlink             	ID = 83
	Sys32oldlstat            	ID = 84
	Sys32readlink            	ID = 85
	Sys32uselib              	ID = 86
	Sys32swapon              	ID = 87
	Sys32reboot              	ID = 88
	Sys32readdir             	ID = 89
	Sys32mmap                	ID = 90
	Sys32munmap              	ID = 91
	Sys32truncate            	ID = 92
	Sys32ftruncate           	ID = 93
	Sys32fchmod              	ID = 94
	Sys32fchown              	ID = 95
	Sys32getpriority         	ID = 96
	Sys32setpriority         	ID = 97
	Sys32profil              	ID = 98
	Sys32statfs              	ID = 99
	Sys32fstatfs             	ID = 100
	Sys32ioperm              	ID = 101
	Sys32socketcall          	ID = 102
	Sys32syslog              	ID = 103
	Sys32setitimer           	ID = 104
	Sys32getitimer           	ID = 105
	Sys32stat                	ID = 106
	Sys32lstat               	ID = 107
	Sys32fstat               	ID = 108
	Sys32olduname            	ID = 109
	Sys32iopl                	ID = 110
	Sys32vhangup             	ID = 111
	Sys32idle                	ID = 112
	Sys32vm86                	ID = 113
	Sys32wait4               	ID = 114
	Sys32swapoff             	ID = 115
	Sys32sysinfo             	ID = 116
	Sys32ipc                 	ID = 117
	Sys32fsync               	ID = 118
	Sys32sigreturn           	ID = 119
	Sys32clone               	ID = 120
	Sys32setdomainname       	ID = 121
	Sys32uname               	ID = 122
	Sys32modify_ldt          	ID = 123
	Sys32adjtimex            	ID = 124
	Sys32mprotect            	ID = 125
	Sys32sigprocmask         	ID = 126
	Sys32create_module       	ID = 127
	Sys32init_module         	ID = 128
	Sys32delete_module       	ID = 129
	Sys32get_kernel_syms     	ID = 130
	Sys32quotactl            	ID = 131
	Sys32getpgid             	ID = 132
	Sys32fchdir              	ID = 133
	Sys32bdflush             	ID = 134
	Sys32sysfs               	ID = 135
	Sys32personality         	ID = 136
	Sys32afs_syscall         	ID = 137
	Sys32setfsuid            	ID = 138
	Sys32setfsgid            	ID = 139
	Sys32_llseek             	ID = 140
	Sys32getdents            	ID = 141
	Sys32_newselect          	ID = 142
	Sys32flock               	ID = 143
	Sys32msync               	ID = 144
	Sys32readv               	ID = 145
	Sys32writev              	ID = 146
	Sys32getsid              	ID = 147
	Sys32fdatasync           	ID = 148
	Sys32_sysctl             	ID = 149
	Sys32mlock               	ID = 150
	Sys32munlock             	ID = 151
	Sys32mlockall            	ID = 152
	Sys32munlockall          	ID = 153
	Sys32sched_setparam      	ID = 154
	Sys32sched_getparam      	ID = 155
	Sys32sched_setscheduler  	ID = 156
	Sys32sched_getscheduler  	ID = 157
	Sys32sched_yield         	ID = 158
	Sys32sched_get_priority_max	ID = 159
	Sys32sched_get_priority_min	ID = 160
	Sys32sched_rr_get_interval	ID = 161
	Sys32nanosleep           	ID = 162
	Sys32mremap              	ID = 163
	Sys32setresuid           	ID = 164
	Sys32getresuid           	ID = 165
	Sys32query_module        	ID = 166
	Sys32poll                	ID = 167
	Sys32nfsservctl          	ID = 168
	Sys32setresgid           	ID = 169
	Sys32getresgid           	ID = 170
	Sys32prctl               	ID = 171
	Sys32rt_sigreturn        	ID = 172
	Sys32rt_sigaction        	ID = 173
	Sys32rt_sigprocmask      	ID = 174
	Sys32rt_sigpending       	ID = 175
	Sys32rt_sigtimedwait     	ID = 176
	Sys32rt_sigqueueinfo     	ID = 177
	Sys32rt_sigsuspend       	ID = 178
	Sys32pread64             	ID = 179
	Sys32pwrite64            	ID = 180
	Sys32chown               	ID = 181
	Sys32getcwd              	ID = 182
	Sys32capget              	ID = 183
	Sys32capset              	ID = 184
	Sys32sigaltstack         	ID = 185
	Sys32sendfile            	ID = 186
	Sys32getpmsg             	ID = 187
	Sys32putpmsg             	ID = 188
	Sys32vfork               	ID = 189
	Sys32ugetrlimit          	ID = 190
	Sys32readahead           	ID = 191
	Sys32mmap2               	ID = 192
	Sys32truncate64          	ID = 193
	Sys32ftruncate64         	ID = 194
	Sys32stat64              	ID = 195
	Sys32lstat64             	ID = 196
	Sys32fstat64             	ID = 197
	Sys32pciconfig_read      	ID = 198
	Sys32pciconfig_write     	ID = 199
	Sys32pciconfig_iobase    	ID = 200
	Sys32multiplexer         	ID = 201
	Sys32getdents64          	ID = 202
	Sys32pivot_root          	ID = 203
	Sys32fcntl64             	ID = 204
	Sys32madvise             	ID = 205
	Sys32mincore             	ID = 206
	Sys32gettid              	ID = 207
	Sys32tkill               	ID = 208
	Sys32setxattr            	ID = 209
	Sys32lsetxattr           	ID = 210
	Sys32fsetxattr           	ID = 211
	Sys32getxattr            	ID = 212
	Sys32lgetxattr           	ID = 213
	Sys32fgetxattr           	ID = 214
	Sys32listxattr           	ID = 215
	Sys32llistxattr          	ID = 216
	Sys32flistxattr          	ID = 217
	Sys32removexattr         	ID = 218
	Sys32lremovexattr        	ID = 219
	Sys32fremovexattr        	ID = 220
	Sys32futex               	ID = 221
	Sys32sched_setaffinity   	ID = 222
	Sys32sched_getaffinity   	ID = 223
	Sys32tuxcall             	ID = 225
	Sys32sendfile64          	ID = 226
	Sys32io_setup            	ID = 227
	Sys32io_destroy          	ID = 228
	Sys32io_getevents        	ID = 229
	Sys32io_submit           	ID = 230
	Sys32io_cancel           	ID = 231
	Sys32set_tid_address     	ID = 232
	Sys32fadvise64           	ID = 233
	Sys32exit_group          	ID = 234
	Sys32lookup_dcookie      	ID = 235
	Sys32epoll_create        	ID = 236
	Sys32epoll_ctl           	ID = 237
	Sys32epoll_wait          	ID = 238
	Sys32remap_file_pages    	ID = 239
	Sys32timer_create        	ID = 240
	Sys32timer_settime       	ID = 241
	Sys32timer_gettime       	ID = 242
	Sys32timer_getoverrun    	ID = 243
	Sys32timer_delete        	ID = 244
	Sys32clock_settime       	ID = 245
	Sys32clock_gettime       	ID = 246
	Sys32clock_getres        	ID = 247
	Sys32clock_nanosleep     	ID = 248
	Sys32swapcontext         	ID = 249
	Sys32tgkill              	ID = 250
	Sys32utimes              	ID = 251
	Sys32statfs64            	ID = 252
	Sys32fstatfs64           	ID = 253
	Sys32fadvise64_64        	ID = 254
	Sys32rtas                	ID = 255
	Sys32sys_debug_setcontext	ID = 256
	Sys32migrate_pages       	ID = 258
	Sys32mbind               	ID = 259
	Sys32get_mempolicy       	ID = 260
	Sys32set_mempolicy       	ID = 261
	Sys32mq_open             	ID = 262
	Sys32mq_unlink           	ID = 263
	Sys32mq_timedsend        	ID = 264
	Sys32mq_timedreceive     	ID = 265
	Sys32mq_notify           	ID = 266
	Sys32mq_getsetattr       	ID = 267
	Sys32kexec_load          	ID = 268
	Sys32add_key             	ID = 269
	Sys32request_key         	ID = 270
	Sys32keyctl              	ID = 271
	Sys32waitid              	ID = 272
	Sys32ioprio_set          	ID = 273
	Sys32ioprio_get          	ID = 274
	Sys32inotify_init        	ID = 275
	Sys32inotify_add_watch   	ID = 276
	Sys32inotify_rm_watch    	ID = 277
	Sys32spu_run             	ID = 278
	Sys32spu_create          	ID = 279
	Sys32pselect6            	ID = 280
	Sys32ppoll               	ID = 281
	Sys32unshare             	ID = 282
	Sys32splice              	ID = 283
	Sys32tee                 	ID = 284
	Sys32vmsplice            	ID = 285
	Sys32openat              	ID = 286
	Sys32mkdirat             	ID = 287
	Sys32mknodat             	ID = 288
	Sys32fchownat            	ID = 289
	Sys32futimesat           	ID = 290
	Sys32fstatat64           	ID = 291
	Sys32unlinkat            	ID = 292
	Sys32renameat            	ID = 293
	Sys32linkat              	ID = 294
	Sys32symlinkat           	ID = 295
	Sys32readlinkat          	ID = 296
	Sys32fchmodat            	ID = 297
	Sys32faccessat           	ID = 298
	Sys32get_robust_list     	ID = 299
	Sys32set_robust_list     	ID = 300
	Sys32move_pages          	ID = 301
	Sys32getcpu              	ID = 302
	Sys32epoll_pwait         	ID = 303
	Sys32utimensat           	ID = 304
	Sys32signalfd            	ID = 305
	Sys32timerfd_create      	ID = 306
	Sys32eventfd             	ID = 307
	Sys32sync_file_range2    	ID = 308
	Sys32fallocate           	ID = 309
	Sys32subpage_prot        	ID = 310
	Sys32timerfd_settime     	ID = 311
	Sys32timerfd_gettime     	ID = 312
	Sys32signalfd4           	ID = 313
	Sys32eventfd2            	ID = 314
	Sys32epoll_create1       	ID = 315
	Sys32dup3                	ID = 316
	Sys32pipe2               	ID = 317
	Sys32inotify_init1       	ID = 318
	Sys32perf_event_open     	ID = 319
	Sys32preadv              	ID = 320
	Sys32pwritev             	ID = 321
	Sys32rt_tgsigqueueinfo   	ID = 322
	Sys32fanotify_init       	ID = 323
	Sys32fanotify_mark       	ID = 324
	Sys32prlimit64           	ID = 325
	Sys32socket              	ID = 326
	Sys32bind                	ID = 327
	Sys32connect             	ID = 328
	Sys32listen              	ID = 329
	Sys32accept              	ID = 330
	Sys32getsockname         	ID = 331
	Sys32getpeername         	ID = 332
	Sys32socketpair          	ID = 333
	Sys32send                	ID = 334
	Sys32sendto              	ID = 335
	Sys32recv                	ID = 336
	Sys32recvfrom            	ID = 337
	Sys32shutdown            	ID = 338
	Sys32setsockopt          	ID = 339
	Sys32getsockopt          	ID = 340
	Sys32sendmsg             	ID = 341
	Sys32recvmsg             	ID = 342
	Sys32recvmmsg            	ID = 343
	Sys32accept4             	ID = 344
	Sys32name_to_handle_at   	ID = 345
	Sys32open_by_handle_at   	ID = 346
	Sys32clock_adjtime       	ID = 347
	Sys32syncfs              	ID = 348
	Sys32sendmmsg            	ID = 349
	Sys32setns               	ID = 350
	Sys32process_vm_readv    	ID = 351
	Sys32process_vm_writev   	ID = 352
	Sys32finit_module        	ID = 353
	Sys32kcmp                	ID = 354
	Sys32sched_setattr       	ID = 355
	Sys32sched_getattr       	ID = 356
	Sys32renameat2           	ID = 357
	Sys32seccomp             	ID = 358
	Sys32getrandom           	ID = 359
	Sys32memfd_create        	ID = 360
	Sys32bpf                 	ID = 361
	Sys32execveat            	ID = 362
	Sys32switch_endian       	ID = 363
	Sys32userfaultfd         	ID = 364
	Sys32membarrier          	ID = 365
	Sys32mlock2              	ID = 378
	Sys32copy_file_range     	ID = 379
	Sys32preadv2             	ID = 380
	Sys32pwritev2            	ID = 381
	Sys32kexec_file_load     	ID = 382
	Sys32statx               	ID = 383
	Sys32pkey_alloc          	ID = 384
	Sys32pkey_free           	ID = 385
	Sys32pkey_mprotect       	ID = 386
	Sys32rseq                	ID = 387
	Sys32io_pgetevents       	ID = 388
	Sys32semget              	ID = 393
	Sys32semctl              	ID = 394
	Sys32shmget              	ID = 395
	Sys32shmctl              	ID = 396
	Sys32shmat               	ID = 397
	Sys32shmdt               	ID = 398
	Sys32msgget              	ID = 399
	Sys32msgsnd              	ID = 400
	Sys32msgrcv              	ID = 401
	Sys32msgctl              	ID = 402
	Sys32clock_gettime64     	ID = 403
	Sys32clock_settime64     	ID = 404
	Sys32clock_adjtime64     	ID = 405
	Sys32clock_getres_time64 	ID = 406
	Sys32clock_nanosleep_time64	ID = 407
	Sys32timer_gettime64     	ID = 408
	Sys32timer_settime64     	ID = 409
	Sys32timerfd_gettime64   	ID = 410
	Sys32timerfd_settime64   	ID = 411
	Sys32utimensat_time64    	ID = 412
	Sys32pselect6_time64     	ID = 413
	Sys32ppoll_time64        	ID = 414
	Sys32io_pgetevents_time64	ID = 416
	Sys32recvmmsg_time64     	ID = 417
	Sys32mq_timedsend_time64 	ID = 418
	Sys32mq_timedreceive_time64	ID = 419
	Sys32semtimedop_time64   	ID = 420
	Sys32rt_sigtimedwait_time64	ID = 421
	Sys32futex_time64        	ID = 422
	Sys32sched_rr_get_interval_time64	ID = 423
	Sys32pidfd_send_signal   	ID = 424
	Sys32io_uring_setup      	ID = 425
	Sys32io_uring_enter      	ID = 426
	Sys32io_uring_register   	ID = 427
	Sys32open_tree           	ID = 428
	Sys32move_mount          	ID = 429
	Sys32fsopen              	ID = 430
	Sys32fsconfig            	ID = 431
	Sys32fsmount             	ID = 432
	Sys32fspick              	ID = 433
	Sys32pidfd_open          	ID = 434
	Sys32clone3              	ID = 435
	Sys32close_range         	ID = 436
	Sys32openat2             	ID = 437
	Sys32pidfd_getfd         	ID = 438
	Sys32faccessat2          	ID = 439
	Sys32process_madvise     	ID = 440
	Sys32epoll_pwait2        	ID = 441
	Sys32mount_setattr       	ID = 442
	Sys32quotactl_fd         	ID = 443
	Sys32landlock_create_ruleset	ID = 444
	Sys32landlock_add_rule   	ID = 445
	Sys32landlock_restrict_self	ID = 446
	Sys32process_mrelease    	ID = 448
	Sys32futex_waitv         	ID = 449
	Sys32set_mempolicy_home_node	ID = 450
	Sys32cachestat           	ID = 451
	Sys32fchmodat2           	ID = 452
	Sys32map_shadow_stack    	ID = 453
	Sys32futex_wake          	ID = 454
	Sys32futex_wait          	ID = 455
	Sys32futex_requeue       	ID = 456
	Sys32statmount           	ID = 457
	Sys32listmount           	ID = 458
	Sys32lsm_get_self_attr   	ID = 459
	Sys32lsm_set_self_attr   	ID = 460
	Sys32lsm_list_modules    	ID = 461
	Sys32mseal               	ID = 462
	Sys32setxattrat          	ID = 463
	Sys32getxattrat          	ID = 464
	Sys32listxattrat         	ID = 465
	Sys32removexattrat       	ID = 466
)

const (
	Sys32semop = iota + Unsupported
	Sys32arch_prctl
	Sys32chown32
	Sys32fchown32
	Sys32get_thread_area
	Sys32getegid32
	Sys32geteuid32
	Sys32getgid32
	Sys32getgroups32
	Sys32getresgid32
	Sys32getresuid32
	Sys32getuid32
	Sys32lchown32
	Sys32memfd_secret
	Sys32set_thread_area
	Sys32setfsgid32
	Sys32setfsuid32
	Sys32setgid32
	Sys32setgroups32
	Sys32setregid32
	Sys32setresgid32
	Sys32setresuid32
	Sys32setreuid32
	Sys32setuid32
	Sys32sync_file_range
	Sys32vm86old
	Sys32vserver
)

const SyscallPrefix = "sys_"
const SyscallNotImplemented = "NOT_IMPLEMENTED"

type KernelRestrictions struct {
	Below string
	Above string
	Name  string
}

var SyscallSymbolNames = map[ID][]KernelRestrictions{
	0:   {{Name: "restart_syscall"}},
	1:   {{Name: "exit"}},
	2:   {{Name: "fork"}},
	3:   {{Name: "read"}},
	4:   {{Name: "write"}},
	5:   {{Name: "open"}},
	6:   {{Name: "close"}},
	7:   {{Name: "waitpid"}},
	8:   {{Name: "creat"}},
	9:   {{Name: "link"}},
	10:  {{Name: "unlink"}},
	11:  {{Name: "execve"}},
	12:  {{Name: "chdir"}},
	13:  {{Name: "time"}},
	14:  {{Name: "mknod"}},
	15:  {{Name: "chmod"}},
	16:  {{Name: "lchown"}},
	17:  {{Name: SyscallNotImplemented + "break"}},
	18:  {{Name: SyscallNotImplemented + "oldstat"}},
	19:  {{Name: "lseek"}},
	20:  {{Name: "getpid"}},
	21:  {{Name: "mount"}},
	22:  {{Name: SyscallNotImplemented + "umount"}},
	23:  {{Name: "setuid"}},
	24:  {{Name: "getuid"}},
	25:  {{Name: "stime"}},
	26:  {{Name: "ptrace"}},
	27:  {{Name: "alarm"}},
	28:  {{Name: SyscallNotImplemented + "oldfstat"}},
	29:  {{Name: "pause"}},
	30:  {{Name: "utime"}},
	31:  {{Name: SyscallNotImplemented + "stty"}},
	32:  {{Name: SyscallNotImplemented + "gtty"}},
	33:  {{Name: "access"}},
	34:  {{Name: "nice"}},
	35:  {{Name: SyscallNotImplemented + "ftime"}},
	36:  {{Name: "sync"}},
	37:  {{Name: "kill"}},
	38:  {{Name: "rename"}},
	39:  {{Name: "mkdir"}},
	40:  {{Name: "rmdir"}},
	41:  {{Name: "dup"}},
	42:  {{Name: "pipe"}},
	43:  {{Name: "times"}},
	44:  {{Name: SyscallNotImplemented + "prof"}},
	45:  {{Name: "brk"}},
	46:  {{Name: "setgid"}},
	47:  {{Name: "getgid"}},
	48:  {{Name: "signal"}},
	49:  {{Name: "geteuid"}},
	50:  {{Name: "getegid"}},
	51:  {{Name: "acct"}},
	52:  {{Name: "umount2"}},
	53:  {{Name: SyscallNotImplemented + "lock"}},
	54:  {{Name: "ioctl"}},
	55:  {{Name: "fcntl"}},
	56:  {{Name: SyscallNotImplemented + "mpx"}},
	57:  {{Name: "setpgid"}},
	58:  {{Name: SyscallNotImplemented + "ulimit"}},
	59:  {{Name: SyscallNotImplemented + "oldolduname"}},
	60:  {{Name: "umask"}},
	61:  {{Name: "chroot"}},
	62:  {{Name: "ustat"}},
	63:  {{Name: "dup2"}},
	64:  {{Name: "getppid"}},
	65:  {{Name: "getpgrp"}},
	66:  {{Name: "setsid"}},
	67:  {{Name: SyscallNotImplemented + "sigaction"}},
	68:  {{Name: "sgetmask"}},
	69:  {{Name: "ssetmask"}},
	70:  {{Name: "setreuid"}},
	71:  {{Name: "setregid"}},
	72:  {{Name: SyscallNotImplemented + "sigsuspend"}},
	73:  {{Name: SyscallNotImplemented + "sigpending"}},
	74:  {{Name: "sethostname"}},
	75:  {{Name: "setrlimit"}},
	76:  {{Name: SyscallNotImplemented + "getrlimit"}},
	77:  {{Name: "getrusage"}},
	78:  {{Name: "gettimeofday"}},
	79:  {{Name: "settimeofday"}},
	80:  {{Name: "getgroups"}},
	81:  {{Name: "setgroups"}},
	82:  {{Name: SyscallNotImplemented + "select"}},
	83:  {{Name: "symlink"}},
	84:  {{Name: SyscallNotImplemented + "oldlstat"}},
	85:  {{Name: "readlink"}},
	86:  {{Name: "uselib"}},
	87:  {{Name: "swapon"}},
	88:  {{Name: "reboot"}},
	89:  {{Name: SyscallNotImplemented + "readdir"}},
	90:  {{Name: "mmap"}},
	91:  {{Name: "munmap"}},
	92:  {{Name: "truncate"}},
	93:  {{Name: "ftruncate"}},
	94:  {{Name: "fchmod"}},
	95:  {{Name: "fchown"}},
	96:  {{Name: "getpriority"}},
	97:  {{Name: "setpriority"}},
	98:  {{Name: SyscallNotImplemented + "profil"}},
	99:  {{Name: "statfs"}},
	100: {{Name: "fstatfs"}},
	101: {{Name: SyscallNotImplemented + "ioperm"}},
	102: {{Name: "socketcall"}},
	103: {{Name: "syslog"}},
	104: {{Name: "setitimer"}},
	105: {{Name: "getitimer"}},
	106: {{Name: "stat"}},
	107: {{Name: "lstat"}},
	108: {{Name: "fstat"}},
	109: {{Name: SyscallNotImplemented + "olduname"}},
	110: {{Name: SyscallNotImplemented + "iopl"}},
	111: {{Name: "vhangup"}},
	112: {{Name: SyscallNotImplemented + "idle"}},
	113: {{Name: SyscallNotImplemented + "vm86"}},
	114: {{Name: "wait4"}},
	115: {{Name: "swapoff"}},
	116: {{Name: "sysinfo"}},
	117: {{Name: "ipc"}},
	118: {{Name: "fsync"}},
	119: {{Name: SyscallNotImplemented + "sigreturn"}},
	120: {{Name: "clone"}},
	121: {{Name: "setdomainname"}},
	122: {{Name: "uname"}},
	123: {{Name: SyscallNotImplemented + "modify_ldt"}},
	124: {{Name: "adjtimex"}},
	125: {{Name: "mprotect"}},
	126: {{Name: SyscallNotImplemented + "sigprocmask"}},
	127: {{Name: SyscallNotImplemented + "create_module"}},
	128: {{Name: "init_module"}},
	129: {{Name: "delete_module"}},
	130: {{Name: SyscallNotImplemented + "get_kernel_syms"}},
	131: {{Name: "quotactl"}},
	132: {{Name: "getpgid"}},
	133: {{Name: "fchdir"}},
	134: {{Name: SyscallNotImplemented + "bdflush"}},
	135: {{Name: "sysfs"}},
	136: {{Name: "personality"}},
	137: {{Name: SyscallNotImplemented + "afs_syscall"}},
	138: {{Name: "setfsuid"}},
	139: {{Name: "setfsgid"}},
	140: {{Name: "_llseek"}},
	141: {{Name: "getdents"}},
	142: {{Name: "_newselect"}},
	143: {{Name: "flock"}},
	144: {{Name: "msync"}},
	145: {{Name: "readv"}},
	146: {{Name: "writev"}},
	147: {{Name: "getsid"}},
	148: {{Name: "fdatasync"}},
	149: {{Name: SyscallNotImplemented + "_sysctl"}},
	150: {{Name: "mlock"}},
	151: {{Name: "munlock"}},
	152: {{Name: "mlockall"}},
	153: {{Name: "munlockall"}},
	154: {{Name: "sched_setparam"}},
	155: {{Name: "sched_getparam"}},
	156: {{Name: "sched_setscheduler"}},
	157: {{Name: "sched_getscheduler"}},
	158: {{Name: "sched_yield"}},
	159: {{Name: "sched_get_priority_max"}},
	160: {{Name: "sched_get_priority_min"}},
	161: {{Name: "sched_rr_get_interval"}},
	162: {{Name: "nanosleep"}},
	163: {{Name: "mremap"}},
	164: {{Name: "setresuid"}},
	165: {{Name: "getresuid"}},
	166: {{Name: SyscallNotImplemented + "query_module"}},
	167: {{Name: "poll"}},
	168: {{Name: SyscallNotImplemented + "nfsservctl"}},
	169: {{Name: "setresgid"}},
	170: {{Name: "getresgid"}},
	171: {{Name: "prctl"}},
	172: {{Name: "rt_sigreturn"}},
	173: {{Name: "rt_sigaction"}},
	174: {{Name: "rt_sigprocmask"}},
	175: {{Name: "rt_sigpending"}},
	176: {{Name: "rt_sigtimedwait"}},
	177: {{Name: "rt_sigqueueinfo"}},
	178: {{Name: "rt_sigsuspend"}},
	179: {{Name: "pread64"}},
	180: {{Name: "pwrite64"}},
	181: {{Name: "chown"}},
	182: {{Name: "getcwd"}},
	183: {{Name: "capget"}},
	184: {{Name: "capset"}},
	185: {{Name: "sigaltstack"}},
	186: {{Name: "sendfile"}},
	187: {{Name: SyscallNotImplemented + "getpmsg"}},
	188: {{Name: SyscallNotImplemented + "putpmsg"}},
	189: {{Name: "vfork"}},
	190: {{Name: "ugetrlimit"}},
	191: {{Name: "readahead"}},
	192: {{Name: SyscallNotImplemented}},
	193: {{Name: SyscallNotImplemented}},
	194: {{Name: SyscallNotImplemented}},
	195: {{Name: SyscallNotImplemented}},
	196: {{Name: SyscallNotImplemented}},
	197: {{Name: SyscallNotImplemented}},
	198: {{Name: "pciconfig_read"}},
	199: {{Name: "pciconfig_write"}},
	200: {{Name: "pciconfig_iobase"}},
	201: {{Name: SyscallNotImplemented + "multiplexer"}},
	202: {{Name: "getdents64"}},
	203: {{Name: "pivot_root"}},
	204: {{Name: SyscallNotImplemented}},
	205: {{Name: "madvise"}},
	206: {{Name: "mincore"}},
	207: {{Name: "gettid"}},
	208: {{Name: "tkill"}},
	209: {{Name: "setxattr"}},
	210: {{Name: "lsetxattr"}},
	211: {{Name: "fsetxattr"}},
	212: {{Name: "getxattr"}},
	213: {{Name: "lgetxattr"}},
	214: {{Name: "fgetxattr"}},
	215: {{Name: "listxattr"}},
	216: {{Name: "llistxattr"}},
	217: {{Name: "flistxattr"}},
	218: {{Name: "removexattr"}},
	219: {{Name: "lremovexattr"}},
	220: {{Name: "fremovexattr"}},
	221: {{Name: "futex"}},
	222: {{Name: "sched_setaffinity"}},
	223: {{Name: "sched_getaffinity"}},
	224: {{Name: SyscallNotImplemented}},
	225: {{Name: SyscallNotImplemented + "tuxcall"}},
	226: {{Name: SyscallNotImplemented}},
	227: {{Name: "io_setup"}},
	228: {{Name: "io_destroy"}},
	229: {{Name: "io_getevents"}},
	230: {{Name: "io_submit"}},
	231: {{Name: "io_cancel"}},
	232: {{Name: "set_tid_address"}},
	233: {{Name: "fadvise64"}},
	234: {{Name: "exit_group"}},
	235: {{Name: SyscallNotImplemented + "lookup_dcookie"}},
	236: {{Name: "epoll_create"}},
	237: {{Name: "epoll_ctl"}},
	238: {{Name: "epoll_wait"}},
	239: {{Name: "remap_file_pages"}},
	240: {{Name: "timer_create"}},
	241: {{Name: "timer_settime"}},
	242: {{Name: "timer_gettime"}},
	243: {{Name: "timer_getoverrun"}},
	244: {{Name: "timer_delete"}},
	245: {{Name: "clock_settime"}},
	246: {{Name: "clock_gettime"}},
	247: {{Name: "clock_getres"}},
	248: {{Name: "clock_nanosleep"}},
	249: {{Name: "swapcontext"}},
	250: {{Name: "tgkill"}},
	251: {{Name: "utimes"}},
	252: {{Name: "statfs64"}},
	253: {{Name: "fstatfs64"}},
	254: {{Name: SyscallNotImplemented}},
	255: {{Name: "rtas"}},
	256: {{Name: SyscallNotImplemented + "sys_debug_setcontext"}},
	257: {{Name: SyscallNotImplemented}},
	258: {{Name: "migrate_pages"}},
	259: {{Name: "mbind"}},
	260: {{Name: "get_mempolicy"}},
	261: {{Name: "set_mempolicy"}},
	262: {{Name: "mq_open"}},
	263: {{Name: "mq_unlink"}},
	264: {{Name: "mq_timedsend"}},
	265: {{Name: "mq_timedreceive"}},
	266: {{Name: "mq_notify"}},
	267: {{Name: "mq_getsetattr"}},
	268: {{Name: "kexec_load"}},
	269: {{Name: "add_key"}},
	270: {{Name: "request_key"}},
	271: {{Name: "keyctl"}},
	272: {{Name: "waitid"}},
	273: {{Name: "ioprio_set"}},
	274: {{Name: "ioprio_get"}},
	275: {{Name: "inotify_init"}},
	276: {{Name: "inotify_add_watch"}},
	277: {{Name: "inotify_rm_watch"}},
	278: {{Name: "spu_run"}},
	279: {{Name: "spu_create"}},
	280: {{Name: "pselect6"}},
	281: {{Name: "ppoll"}},
	282: {{Name: "unshare"}},
	283: {{Name: "splice"}},
	284: {{Name: "tee"}},
	285: {{Name: "vmsplice"}},
	286: {{Name: "openat"}},
	287: {{Name: "mkdirat"}},
	288: {{Name: "mknodat"}},
	289: {{Name: "fchownat"}},
	290: {{Name: "futimesat"}},
	291: {{Name: "newfstatat"}},
	292: {{Name: "unlinkat"}},
	293: {{Name: "renameat"}},
	294: {{Name: "linkat"}},
	295: {{Name: "symlinkat"}},
	296: {{Name: "readlinkat"}},
	297: {{Name: "fchmodat"}},
	298: {{Name: "faccessat"}},
	299: {{Name: "get_robust_list"}},
	300: {{Name: "set_robust_list"}},
	301: {{Name: "move_pages"}},
	302: {{Name: "getcpu"}},
	303: {{Name: "epoll_pwait"}},
	304: {{Name: "utimensat"}},
	305: {{Name: "signalfd"}},
	306: {{Name: "timerfd_create"}},
	307: {{Name: "eventfd"}},
	308: {{Name: "sync_file_range2"}},
	309: {{Name: "fallocate"}},
	310: {{Name: "subpage_prot"}},
	311: {{Name: "timerfd_settime"}},
	312: {{Name: "timerfd_gettime"}},
	313: {{Name: "signalfd4"}},
	314: {{Name: "eventfd2"}},
	315: {{Name: "epoll_create1"}},
	316: {{Name: "dup3"}},
	317: {{Name: "pipe2"}},
	318: {{Name: "inotify_init1"}},
	319: {{Name: "perf_event_open"}},
	320: {{Name: "preadv"}},
	321: {{Name: "pwritev"}},
	322: {{Name: "rt_tgsigqueueinfo"}},
	323: {{Name: "fanotify_init"}},
	324: {{Name: "fanotify_mark"}},
	325: {{Name: "prlimit64"}},
	326: {{Name: "socket"}},
	327: {{Name: "bind"}},
	328: {{Name: "connect"}},
	329: {{Name: "listen"}},
	330: {{Name: "accept"}},
	331: {{Name: "getsockname"}},
	332: {{Name: "getpeername"}},
	333: {{Name: "socketpair"}},
	334: {{Name: "send"}},
	335: {{Name: "sendto"}},
	336: {{Name: "recv"}},
	337: {{Name: "recvfrom"}},
	338: {{Name: "shutdown"}},
	339: {{Name: "setsockopt"}},
	340: {{Name: "getsockopt"}},
	341: {{Name: "sendmsg"}},
	342: {{Name: "recvmsg"}},
	343: {{Name: "recvmmsg"}},
	344: {{Name: "accept4"}},
	345: {{Name: "name_to_handle_at"}},
	346: {{Name: "open_by_handle_at"}},
	347: {{Name: "clock_adjtime"}},
	348: {{Name: "syncfs"}},
	349: {{Name: "sendmmsg"}},
	350: {{Name: "setns"}},
	351: {{Name: "process_vm_readv"}},
	352: {{Name: "process_vm_writev"}},
	353: {{Name: "finit_module"}},
	354: {{Name: "kcmp"}},
	355: {{Name: "sched_setattr"}},
	356: {{Name: "sched_getattr"}},
	357: {{Name: "renameat2"}},
	358: {{Name: "seccomp"}},
	359: {{Name: "getrandom"}},
	360: {{Name: "memfd_create"}},
	361: {{Name: "bpf"}},
	362: {{Name: "execveat"}},
	363: {{Name: "switch_endian"}},
	364: {{Name: "userfaultfd"}},
	365: {{Name: "membarrier"}},
	366: {{Name: SyscallNotImplemented}},
	367: {{Name: SyscallNotImplemented}},
	368: {{Name: SyscallNotImplemented}},
	369: {{Name: SyscallNotImplemented}},
	370: {{Name: SyscallNotImplemented}},
	371: {{Name: SyscallNotImplemented}},
	372: {{Name: SyscallNotImplemented}},
	373: {{Name: SyscallNotImplemented}},
	374: {{Name: SyscallNotImplemented}},
	375: {{Name: SyscallNotImplemented}},
	376: {{Name: SyscallNotImplemented}},
	377: {{Name: SyscallNotImplemented}},
	378: {{Name: "mlock2"}},
	379: {{Name: "copy_file_range"}},
	380: {{Name: "preadv2"}},
	381: {{Name: "pwritev2"}},
	382: {{Name: "kexec_file_load"}},
	383: {{Name: "statx"}},
	384: {{Name: "pkey_alloc"}},
	385: {{Name: "pkey_free"}},
	386: {{Name: "pkey_mprotect"}},
	387: {{Name: "rseq"}},
	388: {{Name: "io_pgetevents"}},
	389: {{Name: SyscallNotImplemented}},
	390: {{Name: SyscallNotImplemented}},
	391: {{Name: SyscallNotImplemented}},
	392: {{Name: "semtimedop"}},
	393: {{Name: "semget"}},
	394: {{Name: "semctl"}},
	395: {{Name: "shmget"}},
	396: {{Name: "shmctl"}},
	397: {{Name: "shmat"}},
	398: {{Name: "shmdt"}},
	399: {{Name: "msgget"}},
	400: {{Name: "msgsnd"}},
	401: {{Name: "msgrcv"}},
	402: {{Name: "msgctl"}},
	403: {{Name: SyscallNotImplemented}},
	404: {{Name: SyscallNotImplemented}},
	405: {{Name: SyscallNotImplemented}},
	406: {{Name: SyscallNotImplemented}},
	407: {{Name: SyscallNotImplemented}},
	408: {{Name: SyscallNotImplemented}},
	409: {{Name: SyscallNotImplemented}},
	410: {{Name: SyscallNotImplemented}},
	411: {{Name: SyscallNotImplemented}},
	412: {{Name: SyscallNotImplemented}},
	413: {{Name: SyscallNotImplemented}},
	414: {{Name: SyscallNotImplemented}},
	415: {{Name: SyscallNotImplemented}},
	416: {{Name: SyscallNotImplemented}},
	417: {{Name: SyscallNotImplemented}},
	418: {{Name: SyscallNotImplemented}},
	419: {{Name: SyscallNotImplemented}},
	420: {{Name: SyscallNotImplemented}},
	421: {{Name: SyscallNotImplemented}},
	422: {{Name: SyscallNotImplemented}},
	423: {{Name: SyscallNotImplemented}},
	424: {{Name: "pidfd_send_signal"}},
	425: {{Name: "io_uring_setup"}},
	426: {{Name: "io_uring_enter"}},
	427: {{Name: "io_uring_register"}},
	428: {{Name: "open_tree"}},
	429: {{Name: "move_mount"}},
	430: {{Name: "fsopen"}},
	431: {{Name: "fsconfig"}},
	432: {{Name: "fsmount"}},
	433: {{Name: "fspick"}},
	434: {{Name: "pidfd_open"}},
	435: {{Name: "clone3"}},
	436: {{Name: "close_range"}},
	437: {{Name: "openat2"}},
	438: {{Name: "pidfd_getfd"}},
	439: {{Name: "faccessat2"}},
	440: {{Name: "process_madvise"}},
	441: {{Name: "epoll_pwait2"}},
	442: {{Name: "mount_setattr"}},
	443: {{Name: "quotactl_fd"}},
	444: {{Name: "landlock_create_ruleset"}},
	445: {{Name: "landlock_add_rule"}},
	446: {{Name: "landlock_restrict_self"}},
	447: {{Name: SyscallNotImplemented}},
	448: {{Name: "process_mrelease"}},
	449: {{Name: "futex_waitv"}},
	450: {{Name: "set_mempolicy_home_node"}},
	451: {{Name: "cachestat"}},
	452: {{Name: "fchmodat2"}},
	453: {{Name: SyscallNotImplemented + "map_shadow_stack"}},
	454: {{Name: "futex_wake"}},
	455: {{Name: "futex_wait"}},
	456: {{Name: "futex_requeue"}},
	457: {{Name: "statmount"}},
	458: {{Name: "listmount"}},
	459: {{Name: "lsm_get_self_attr"}},
	460: {{Name: "lsm_set_self_attr"}},
	461: {{Name: "lsm_list_modules"}},
	462: {{Name: "mseal"}},
	463: {{Name: "setxattrat"}},
	464: {{Name: "getxattrat"}},
	465: {{Name: "listxattrat"}},
	466: {{Name: "removexattrat"}},
}
