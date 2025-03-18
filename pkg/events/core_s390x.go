//go:build s390x

package events

// s390x 64bit syscall numbers (used as event IDs for the Syscall Events)
// https://github.com/torvalds/linux/blob/master/arch/s390/kernel/syscalls/syscall.tbl

const (
	Exit                     	ID = 1
	Fork                     	ID = 2
	Read                     	ID = 3
	Write                    	ID = 4
	Open                     	ID = 5
	Close                    	ID = 6
	RestartSyscall           	ID = 7
	Creat                    	ID = 8
	Link                     	ID = 9
	Unlink                   	ID = 10
	Execve                   	ID = 11
	Chdir                    	ID = 12
	Mknod                    	ID = 14
	Chmod                    	ID = 15
	Lseek                    	ID = 19
	Getpid                   	ID = 20
	Mount                    	ID = 21
	Umount                   	ID = 22
	Ptrace                   	ID = 26
	Alarm                    	ID = 27
	Pause                    	ID = 29
	Utime                    	ID = 30
	Access                   	ID = 33
	Nice                     	ID = 34
	Sync                     	ID = 36
	Kill                     	ID = 37
	Rename                   	ID = 38
	Mkdir                    	ID = 39
	Rmdir                    	ID = 40
	Dup                      	ID = 41
	Pipe                     	ID = 42
	Times                    	ID = 43
	Brk                      	ID = 45
	Signal                   	ID = 48
	Acct                     	ID = 51
	Umount2                  	ID = 52
	Ioctl                    	ID = 54
	Fcntl                    	ID = 55
	Setpgid                  	ID = 57
	Umask                    	ID = 60
	Chroot                   	ID = 61
	Ustat                    	ID = 62
	Dup2                     	ID = 63
	Getppid                  	ID = 64
	Getpgrp                  	ID = 65
	Setsid                   	ID = 66
	Sigaction                	ID = 67
	Sigsuspend               	ID = 72
	Sigpending               	ID = 73
	Sethostname              	ID = 74
	Setrlimit                	ID = 75
	Getrusage                	ID = 77
	Gettimeofday             	ID = 78
	Settimeofday             	ID = 79
	Symlink                  	ID = 83
	Readlink                 	ID = 85
	Uselib                   	ID = 86
	Swapon                   	ID = 87
	Reboot                   	ID = 88
	Readdir                  	ID = 89
	Mmap                     	ID = 90
	Munmap                   	ID = 91
	Truncate                 	ID = 92
	Ftruncate                	ID = 93
	Fchmod                   	ID = 94
	Getpriority              	ID = 96
	Setpriority              	ID = 97
	Statfs                   	ID = 99
	Fstatfs                  	ID = 100
	Socketcall               	ID = 102
	Syslog                   	ID = 103
	Setitimer                	ID = 104
	Getitimer                	ID = 105
	Stat                     	ID = 106
	Lstat                    	ID = 107
	Fstat                    	ID = 108
	LookupDcookie            	ID = 110
	Vhangup                  	ID = 111
	Idle                     	ID = 112
	Wait4                    	ID = 114
	Swapoff                  	ID = 115
	Sysinfo                  	ID = 116
	Ipc                      	ID = 117
	Fsync                    	ID = 118
	Sigreturn                	ID = 119
	Clone                    	ID = 120
	Setdomainname            	ID = 121
	Uname                    	ID = 122
	Adjtimex                 	ID = 124
	Mprotect                 	ID = 125
	Sigprocmask              	ID = 126
	CreateModule             	ID = 127
	InitModule               	ID = 128
	DeleteModule             	ID = 129
	GetKernelSyms            	ID = 130
	Quotactl                 	ID = 131
	Getpgid                  	ID = 132
	Fchdir                   	ID = 133
	Bdflush                  	ID = 134
	Sysfs                    	ID = 135
	Personality              	ID = 136
	AfsSyscall               	ID = 137
	Getdents                 	ID = 141
	Select                   	ID = 142
	Flock                    	ID = 143
	Msync                    	ID = 144
	Readv                    	ID = 145
	Writev                   	ID = 146
	Getsid                   	ID = 147
	Fdatasync                	ID = 148
	Sysctl                   	ID = 149
	Mlock                    	ID = 150
	Munlock                  	ID = 151
	Mlockall                 	ID = 152
	Munlockall               	ID = 153
	SchedSetparam            	ID = 154
	SchedGetparam            	ID = 155
	SchedSetscheduler        	ID = 156
	SchedGetscheduler        	ID = 157
	SchedYield               	ID = 158
	SchedGetPriorityMax      	ID = 159
	SchedGetPriorityMin      	ID = 160
	SchedRrGetInterval       	ID = 161
	Nanosleep                	ID = 162
	Mremap                   	ID = 163
	QueryModule              	ID = 167
	Poll                     	ID = 168
	Nfsservctl               	ID = 169
	Prctl                    	ID = 172
	RtSigreturn              	ID = 173
	RtSigaction              	ID = 174
	RtSigprocmask            	ID = 175
	RtSigpending             	ID = 176
	RtSigtimedwait           	ID = 177
	RtSigqueueinfo           	ID = 178
	RtSigsuspend             	ID = 179
	Pread64                  	ID = 180
	Pwrite64                 	ID = 181
	Getcwd                   	ID = 183
	Capget                   	ID = 184
	Capset                   	ID = 185
	Sigaltstack              	ID = 186
	Sendfile                 	ID = 187
	Getpmsg                  	ID = 188
	Putpmsg                  	ID = 189
	Vfork                    	ID = 190
	Getrlimit                	ID = 191
	Lchown                   	ID = 198
	Getuid                   	ID = 199
	Getgid                   	ID = 200
	Geteuid                  	ID = 201
	Getegid                  	ID = 202
	Setreuid                 	ID = 203
	Setregid                 	ID = 204
	Getgroups                	ID = 205
	Setgroups                	ID = 206
	Fchown                   	ID = 207
	Setresuid                	ID = 208
	Getresuid                	ID = 209
	Setresgid                	ID = 210
	Getresgid                	ID = 211
	Chown                    	ID = 212
	Setuid                   	ID = 213
	Setgid                   	ID = 214
	Setfsuid                 	ID = 215
	Setfsgid                 	ID = 216
	PivotRoot                	ID = 217
	Mincore                  	ID = 218
	Madvise                  	ID = 219
	Getdents64               	ID = 220
	Readahead                	ID = 222
	Setxattr                 	ID = 224
	Lsetxattr                	ID = 225
	Fsetxattr                	ID = 226
	Getxattr                 	ID = 227
	Lgetxattr                	ID = 228
	Fgetxattr                	ID = 229
	Listxattr                	ID = 230
	Llistxattr               	ID = 231
	Flistxattr               	ID = 232
	Removexattr              	ID = 233
	Lremovexattr             	ID = 234
	Fremovexattr             	ID = 235
	Gettid                   	ID = 236
	Tkill                    	ID = 237
	Futex                    	ID = 238
	SchedSetaffinity         	ID = 239
	SchedGetaffinity         	ID = 240
	Tgkill                   	ID = 241
	IoSetup                  	ID = 243
	IoDestroy                	ID = 244
	IoGetevents              	ID = 245
	IoSubmit                 	ID = 246
	IoCancel                 	ID = 247
	ExitGroup                	ID = 248
	EpollCreate              	ID = 249
	EpollCtl                 	ID = 250
	EpollWait                	ID = 251
	SetTidAddress            	ID = 252
	Fadvise64                	ID = 253
	TimerCreate              	ID = 254
	TimerSettime             	ID = 255
	TimerGettime             	ID = 256
	TimerGetoverrun          	ID = 257
	TimerDelete              	ID = 258
	ClockSettime             	ID = 259
	ClockGettime             	ID = 260
	ClockGetres              	ID = 261
	ClockNanosleep           	ID = 262
	Statfs64                 	ID = 265
	Fstatfs64                	ID = 266
	RemapFilePages           	ID = 267
	Mbind                    	ID = 268
	GetMempolicy             	ID = 269
	SetMempolicy             	ID = 270
	MqOpen                   	ID = 271
	MqUnlink                 	ID = 272
	MqTimedsend              	ID = 273
	MqTimedreceive           	ID = 274
	MqNotify                 	ID = 275
	MqGetsetattr             	ID = 276
	KexecLoad                	ID = 277
	AddKey                   	ID = 278
	RequestKey               	ID = 279
	Keyctl                   	ID = 280
	Waitid                   	ID = 281
	IoprioSet                	ID = 282
	IoprioGet                	ID = 283
	InotifyInit              	ID = 284
	InotifyAddWatch          	ID = 285
	InotifyRmWatch           	ID = 286
	MigratePages             	ID = 287
	Openat                   	ID = 288
	Mkdirat                  	ID = 289
	Mknodat                  	ID = 290
	Fchownat                 	ID = 291
	Futimesat                	ID = 292
	Newfstatat               	ID = 293
	Unlinkat                 	ID = 294
	Renameat                 	ID = 295
	Linkat                   	ID = 296
	Symlinkat                	ID = 297
	Readlinkat               	ID = 298
	Fchmodat                 	ID = 299
	Faccessat                	ID = 300
	Pselect6                 	ID = 301
	Ppoll                    	ID = 302
	Unshare                  	ID = 303
	SetRobustList            	ID = 304
	GetRobustList            	ID = 305
	Splice                   	ID = 306
	SyncFileRange            	ID = 307
	Tee                      	ID = 308
	Vmsplice                 	ID = 309
	MovePages                	ID = 310
	Getcpu                   	ID = 311
	EpollPwait               	ID = 312
	Utimes                   	ID = 313
	Fallocate                	ID = 314
	Utimensat                	ID = 315
	Signalfd                 	ID = 316
	Timerfd                  	ID = 317
	Eventfd                  	ID = 318
	TimerfdCreate            	ID = 319
	TimerfdSettime           	ID = 320
	TimerfdGettime           	ID = 321
	Signalfd4                	ID = 322
	Eventfd2                 	ID = 323
	InotifyInit1             	ID = 324
	Pipe2                    	ID = 325
	Dup3                     	ID = 326
	EpollCreate1             	ID = 327
	Preadv                   	ID = 328
	Pwritev                  	ID = 329
	RtTgsigqueueinfo         	ID = 330
	PerfEventOpen            	ID = 331
	FanotifyInit             	ID = 332
	FanotifyMark             	ID = 333
	Prlimit64                	ID = 334
	NameToHandleAt           	ID = 335
	OpenByHandleAt           	ID = 336
	ClockAdjtime             	ID = 337
	Syncfs                   	ID = 338
	Setns                    	ID = 339
	ProcessVmReadv           	ID = 340
	ProcessVmWritev          	ID = 341
	S390RuntimeInstr         	ID = 342
	Kcmp                     	ID = 343
	FinitModule              	ID = 344
	SchedSetattr             	ID = 345
	SchedGetattr             	ID = 346
	Renameat2                	ID = 347
	Seccomp                  	ID = 348
	Getrandom                	ID = 349
	MemfdCreate              	ID = 350
	Bpf                      	ID = 351
	S390PciMmioWrite         	ID = 352
	S390PciMmioRead          	ID = 353
	Execveat                 	ID = 354
	Userfaultfd              	ID = 355
	Membarrier               	ID = 356
	Recvmmsg                 	ID = 357
	Sendmmsg                 	ID = 358
	Socket                   	ID = 359
	Socketpair               	ID = 360
	Bind                     	ID = 361
	Connect                  	ID = 362
	Listen                   	ID = 363
	Accept4                  	ID = 364
	Getsockopt               	ID = 365
	Setsockopt               	ID = 366
	Getsockname              	ID = 367
	Getpeername              	ID = 368
	Sendto                   	ID = 369
	Sendmsg                  	ID = 370
	Recvfrom                 	ID = 371
	Recvmsg                  	ID = 372
	Shutdown                 	ID = 373
	Mlock2                   	ID = 374
	CopyFileRange            	ID = 375
	Preadv2                  	ID = 376
	Pwritev2                 	ID = 377
	S390GuardedStorage       	ID = 378
	Statx                    	ID = 379
	S390Sthyi                	ID = 380
	KexecFileLoad            	ID = 381
	IoPgetevents             	ID = 382
	Rseq                     	ID = 383
	PkeyMprotect             	ID = 384
	PkeyAlloc                	ID = 385
	PkeyFree                 	ID = 386
	Semtimedop               	ID = 392
	Semget                   	ID = 393
	Semctl                   	ID = 394
	Shmget                   	ID = 395
	Shmctl                   	ID = 396
	Shmat                    	ID = 397
	Shmdt                    	ID = 398
	Msgget                   	ID = 399
	Msgsnd                   	ID = 400
	Msgrcv                   	ID = 401
	Msgctl                   	ID = 402
	PidfdSendSignal          	ID = 424
	IoUringSetup             	ID = 425
	IoUringEnter             	ID = 426
	IoUringRegister          	ID = 427
	OpenTree                 	ID = 428
	MoveMount                	ID = 429
	Fsopen                   	ID = 430
	Fsconfig                 	ID = 431
	Fsmount                  	ID = 432
	Fspick                   	ID = 433
	PidfdOpen                	ID = 434
	Clone3                   	ID = 435
	CloseRange               	ID = 436
	Openat2                  	ID = 437
	PidfdGetfd               	ID = 438
	Faccessat2               	ID = 439
	ProcessMadvise           	ID = 440
	EpollPwait2              	ID = 441
	MountSetattr             	ID = 442
	QuotactlFd               	ID = 443
	LandlockCreateRuleset    	ID = 444
	LandlockAddRule          	ID = 445
	LandlockRestrictSelf     	ID = 446
	MemfdSecret              	ID = 447
	ProcessMrelease          	ID = 448
	FutexWaitv               	ID = 449
	SetMempolicyHomeNode     	ID = 450
	Cachestat                	ID = 451
	Fchmodat2                	ID = 452
	MapShadowStack           	ID = 453
	FutexWake                	ID = 454
	FutexWait                	ID = 455
	FutexRequeue             	ID = 456
	Statmount                	ID = 457
	Listmount                	ID = 458
	LsmGetSelfAttr           	ID = 459
	LsmSetSelfAttr           	ID = 460
	LsmListModules           	ID = 461
	Mseal                    	ID = 462
	Setxattrat               	ID = 463
	Getxattrat               	ID = 464
	Listxattrat              	ID = 465
	Removexattrat            	ID = 466
	MaxSyscallID            	ID = iota
)

// following syscalls are undefined on s390x
const (
	Accept                   	ID = iota + Unsupported
	Afs
	Afs_syscall
	ArchPrctl
	Break
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
	Ftime
	Ftruncate64
	FutexTime32
	GetThreadArea
	Getegid16
	Geteuid16
	Getgid16
	Getgroups16
	Getresgid16
	Getresuid16
	Getuid16
	Gtty
	IoPgeteventsTime32
	Ioperm
	Iopl
	Lchown16
	Llseek
	Lock
	Lstat64
	Mmap2
	ModifyLdt
	Mpx
	MqTimedreceiveTime32
	MqTimedsendTime32
	OldGetrlimit
	OldSelect
	Oldfstat
	Oldlstat
	Oldolduname
	Oldstat
	Olduname
	PpollTime32
	Prof
	Profil
	Pselect6Time32
	RecvmmsgTime32
	RtSigtimedwaitTime32
	SchedRrGetInterval32
	Security
	Semop
	Sendfile32
	SetThreadArea
	Setfsgid16
	Setfsuid16
	Setgid16
	Setgroups16
	Setregid16
	Setresgid16
	Setresuid16
	Setreuid16
	Setuid16
	Sgetmask
	Ssetmask
	Stat64
	Stime
	Stty
	Time
	TimerGettime32
	TimerSettime32
	TimerfdGettime32
	TimerfdSettime32
	Truncate64
	Tuxcall
	Ulimit
	UtimensatTime32
	Vm86
	Vm86old
	Vserver
	Waitpid
)

// s390x 32-bit syscall numbers (used as event IDs for the Syscall Events)
// https://github.com/torvalds/linux/blob/master/arch/s390/kernel/syscalls/syscall.tbl

const (
	Sys32exit                	ID = 1
	Sys32fork                	ID = 2
	Sys32read                	ID = 3
	Sys32write               	ID = 4
	Sys32open                	ID = 5
	Sys32close               	ID = 6
	Sys32restart_syscall     	ID = 7
	Sys32creat               	ID = 8
	Sys32link                	ID = 9
	Sys32unlink              	ID = 10
	Sys32execve              	ID = 11
	Sys32chdir               	ID = 12
	Sys32time                	ID = 13
	Sys32mknod               	ID = 14
	Sys32chmod               	ID = 15
	Sys32lchown              	ID = 16
	Sys32lseek               	ID = 19
	Sys32getpid              	ID = 20
	Sys32mount               	ID = 21
	Sys32umount              	ID = 22
	Sys32setuid              	ID = 23
	Sys32getuid              	ID = 24
	Sys32stime               	ID = 25
	Sys32ptrace              	ID = 26
	Sys32alarm               	ID = 27
	Sys32pause               	ID = 29
	Sys32utime               	ID = 30
	Sys32access              	ID = 33
	Sys32nice                	ID = 34
	Sys32sync                	ID = 36
	Sys32kill                	ID = 37
	Sys32rename              	ID = 38
	Sys32mkdir               	ID = 39
	Sys32rmdir               	ID = 40
	Sys32dup                 	ID = 41
	Sys32pipe                	ID = 42
	Sys32times               	ID = 43
	Sys32brk                 	ID = 45
	Sys32setgid              	ID = 46
	Sys32getgid              	ID = 47
	Sys32signal              	ID = 48
	Sys32geteuid             	ID = 49
	Sys32getegid             	ID = 50
	Sys32acct                	ID = 51
	Sys32umount2             	ID = 52
	Sys32ioctl               	ID = 54
	Sys32fcntl               	ID = 55
	Sys32setpgid             	ID = 57
	Sys32umask               	ID = 60
	Sys32chroot              	ID = 61
	Sys32ustat               	ID = 62
	Sys32dup2                	ID = 63
	Sys32getppid             	ID = 64
	Sys32getpgrp             	ID = 65
	Sys32setsid              	ID = 66
	Sys32sigaction           	ID = 67
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
	Sys32symlink             	ID = 83
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
	Sys32lookup_dcookie      	ID = 110
	Sys32vhangup             	ID = 111
	Sys32idle                	ID = 112
	Sys32wait4               	ID = 114
	Sys32swapoff             	ID = 115
	Sys32sysinfo             	ID = 116
	Sys32ipc                 	ID = 117
	Sys32fsync               	ID = 118
	Sys32sigreturn           	ID = 119
	Sys32clone               	ID = 120
	Sys32setdomainname       	ID = 121
	Sys32uname               	ID = 122
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
	Sys32query_module        	ID = 167
	Sys32poll                	ID = 168
	Sys32nfsservctl          	ID = 169
	Sys32setresgid           	ID = 170
	Sys32getresgid           	ID = 171
	Sys32prctl               	ID = 172
	Sys32rt_sigreturn        	ID = 173
	Sys32rt_sigaction        	ID = 174
	Sys32rt_sigprocmask      	ID = 175
	Sys32rt_sigpending       	ID = 176
	Sys32rt_sigtimedwait     	ID = 177
	Sys32rt_sigqueueinfo     	ID = 178
	Sys32rt_sigsuspend       	ID = 179
	Sys32pread64             	ID = 180
	Sys32pwrite64            	ID = 181
	Sys32chown               	ID = 182
	Sys32getcwd              	ID = 183
	Sys32capget              	ID = 184
	Sys32capset              	ID = 185
	Sys32sigaltstack         	ID = 186
	Sys32sendfile            	ID = 187
	Sys32getpmsg             	ID = 188
	Sys32putpmsg             	ID = 189
	Sys32vfork               	ID = 190
	Sys32ugetrlimit          	ID = 191
	Sys32mmap2               	ID = 192
	Sys32truncate64          	ID = 193
	Sys32ftruncate64         	ID = 194
	Sys32stat64              	ID = 195
	Sys32lstat64             	ID = 196
	Sys32fstat64             	ID = 197
	Sys32lchown32            	ID = 198
	Sys32getuid32            	ID = 199
	Sys32getgid32            	ID = 200
	Sys32geteuid32           	ID = 201
	Sys32getegid32           	ID = 202
	Sys32setreuid32          	ID = 203
	Sys32setregid32          	ID = 204
	Sys32getgroups32         	ID = 205
	Sys32setgroups32         	ID = 206
	Sys32fchown32            	ID = 207
	Sys32setresuid32         	ID = 208
	Sys32getresuid32         	ID = 209
	Sys32setresgid32         	ID = 210
	Sys32getresgid32         	ID = 211
	Sys32chown32             	ID = 212
	Sys32setuid32            	ID = 213
	Sys32setgid32            	ID = 214
	Sys32setfsuid32          	ID = 215
	Sys32setfsgid32          	ID = 216
	Sys32pivot_root          	ID = 217
	Sys32mincore             	ID = 218
	Sys32madvise             	ID = 219
	Sys32getdents64          	ID = 220
	Sys32fcntl64             	ID = 221
	Sys32readahead           	ID = 222
	Sys32sendfile64          	ID = 223
	Sys32setxattr            	ID = 224
	Sys32lsetxattr           	ID = 225
	Sys32fsetxattr           	ID = 226
	Sys32getxattr            	ID = 227
	Sys32lgetxattr           	ID = 228
	Sys32fgetxattr           	ID = 229
	Sys32listxattr           	ID = 230
	Sys32llistxattr          	ID = 231
	Sys32flistxattr          	ID = 232
	Sys32removexattr         	ID = 233
	Sys32lremovexattr        	ID = 234
	Sys32fremovexattr        	ID = 235
	Sys32gettid              	ID = 236
	Sys32tkill               	ID = 237
	Sys32futex               	ID = 238
	Sys32sched_setaffinity   	ID = 239
	Sys32sched_getaffinity   	ID = 240
	Sys32tgkill              	ID = 241
	Sys32io_setup            	ID = 243
	Sys32io_destroy          	ID = 244
	Sys32io_getevents        	ID = 245
	Sys32io_submit           	ID = 246
	Sys32io_cancel           	ID = 247
	Sys32exit_group          	ID = 248
	Sys32epoll_create        	ID = 249
	Sys32epoll_ctl           	ID = 250
	Sys32epoll_wait          	ID = 251
	Sys32set_tid_address     	ID = 252
	Sys32fadvise64           	ID = 253
	Sys32timer_create        	ID = 254
	Sys32timer_settime       	ID = 255
	Sys32timer_gettime       	ID = 256
	Sys32timer_getoverrun    	ID = 257
	Sys32timer_delete        	ID = 258
	Sys32clock_settime       	ID = 259
	Sys32clock_gettime       	ID = 260
	Sys32clock_getres        	ID = 261
	Sys32clock_nanosleep     	ID = 262
	Sys32fadvise64_64        	ID = 264
	Sys32statfs64            	ID = 265
	Sys32fstatfs64           	ID = 266
	Sys32remap_file_pages    	ID = 267
	Sys32mbind               	ID = 268
	Sys32get_mempolicy       	ID = 269
	Sys32set_mempolicy       	ID = 270
	Sys32mq_open             	ID = 271
	Sys32mq_unlink           	ID = 272
	Sys32mq_timedsend        	ID = 273
	Sys32mq_timedreceive     	ID = 274
	Sys32mq_notify           	ID = 275
	Sys32mq_getsetattr       	ID = 276
	Sys32kexec_load          	ID = 277
	Sys32add_key             	ID = 278
	Sys32request_key         	ID = 279
	Sys32keyctl              	ID = 280
	Sys32waitid              	ID = 281
	Sys32ioprio_set          	ID = 282
	Sys32ioprio_get          	ID = 283
	Sys32inotify_init        	ID = 284
	Sys32inotify_add_watch   	ID = 285
	Sys32inotify_rm_watch    	ID = 286
	Sys32migrate_pages       	ID = 287
	Sys32openat              	ID = 288
	Sys32mkdirat             	ID = 289
	Sys32mknodat             	ID = 290
	Sys32fchownat            	ID = 291
	Sys32futimesat           	ID = 292
	Sys32fstatat64           	ID = 293
	Sys32unlinkat            	ID = 294
	Sys32renameat            	ID = 295
	Sys32linkat              	ID = 296
	Sys32symlinkat           	ID = 297
	Sys32readlinkat          	ID = 298
	Sys32fchmodat            	ID = 299
	Sys32faccessat           	ID = 300
	Sys32pselect6            	ID = 301
	Sys32ppoll               	ID = 302
	Sys32unshare             	ID = 303
	Sys32set_robust_list     	ID = 304
	Sys32get_robust_list     	ID = 305
	Sys32splice              	ID = 306
	Sys32sync_file_range     	ID = 307
	Sys32tee                 	ID = 308
	Sys32vmsplice            	ID = 309
	Sys32move_pages          	ID = 310
	Sys32getcpu              	ID = 311
	Sys32epoll_pwait         	ID = 312
	Sys32utimes              	ID = 313
	Sys32fallocate           	ID = 314
	Sys32utimensat           	ID = 315
	Sys32signalfd            	ID = 316
	Sys32timerfd             	ID = 317
	Sys32eventfd             	ID = 318
	Sys32timerfd_create      	ID = 319
	Sys32timerfd_settime     	ID = 320
	Sys32timerfd_gettime     	ID = 321
	Sys32signalfd4           	ID = 322
	Sys32eventfd2            	ID = 323
	Sys32inotify_init1       	ID = 324
	Sys32pipe2               	ID = 325
	Sys32dup3                	ID = 326
	Sys32epoll_create1       	ID = 327
	Sys32preadv              	ID = 328
	Sys32pwritev             	ID = 329
	Sys32rt_tgsigqueueinfo   	ID = 330
	Sys32perf_event_open     	ID = 331
	Sys32fanotify_init       	ID = 332
	Sys32fanotify_mark       	ID = 333
	Sys32prlimit64           	ID = 334
	Sys32name_to_handle_at   	ID = 335
	Sys32open_by_handle_at   	ID = 336
	Sys32clock_adjtime       	ID = 337
	Sys32syncfs              	ID = 338
	Sys32setns               	ID = 339
	Sys32process_vm_readv    	ID = 340
	Sys32process_vm_writev   	ID = 341
	Sys32s390_runtime_instr  	ID = 342
	Sys32kcmp                	ID = 343
	Sys32finit_module        	ID = 344
	Sys32sched_setattr       	ID = 345
	Sys32sched_getattr       	ID = 346
	Sys32renameat2           	ID = 347
	Sys32seccomp             	ID = 348
	Sys32getrandom           	ID = 349
	Sys32memfd_create        	ID = 350
	Sys32bpf                 	ID = 351
	Sys32s390_pci_mmio_write 	ID = 352
	Sys32s390_pci_mmio_read  	ID = 353
	Sys32execveat            	ID = 354
	Sys32userfaultfd         	ID = 355
	Sys32membarrier          	ID = 356
	Sys32recvmmsg            	ID = 357
	Sys32sendmmsg            	ID = 358
	Sys32socket              	ID = 359
	Sys32socketpair          	ID = 360
	Sys32bind                	ID = 361
	Sys32connect             	ID = 362
	Sys32listen              	ID = 363
	Sys32accept4             	ID = 364
	Sys32getsockopt          	ID = 365
	Sys32setsockopt          	ID = 366
	Sys32getsockname         	ID = 367
	Sys32getpeername         	ID = 368
	Sys32sendto              	ID = 369
	Sys32sendmsg             	ID = 370
	Sys32recvfrom            	ID = 371
	Sys32recvmsg             	ID = 372
	Sys32shutdown            	ID = 373
	Sys32mlock2              	ID = 374
	Sys32copy_file_range     	ID = 375
	Sys32preadv2             	ID = 376
	Sys32pwritev2            	ID = 377
	Sys32s390_guarded_storage	ID = 378
	Sys32statx               	ID = 379
	Sys32s390_sthyi          	ID = 380
	Sys32kexec_file_load     	ID = 381
	Sys32io_pgetevents       	ID = 382
	Sys32rseq                	ID = 383
	Sys32pkey_mprotect       	ID = 384
	Sys32pkey_alloc          	ID = 385
	Sys32pkey_free           	ID = 386
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
	Sys32memfd_secret        	ID = 447
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

// following syscalls are undefined on s390x in 32-bit mode
const (
	Sys32arch_prctl          	ID = iota + Unsupported
	Sys32break
	Sys32ftime
	Sys32get_thread_area
	Sys32gtty
	Sys32iopl
	Sys32lock
	Sys32modify_ldt
	Sys32mpx
	Sys32oldfstat
	Sys32oldlstat
	Sys32oldolduname
	Sys32oldstat
	Sys32olduname
	Sys32prof
	Sys32profil
	Sys32select
	Sys32set_thread_area
	Sys32sgetmask
	Sys32ssetmask
	Sys32stty
	Sys32ulimit
	Sys32vm86
	Sys32vm86old
	Sys32vserver
	Sys32waitpid
)

const SyscallPrefix = "__s390_sys_"
const SyscallNotImplemented = "NOT_IMPLEMENTED"

type KernelRestrictions struct {
	Below string
	Above string
	Name  string
}

var SyscallSymbolNames = map[ID][]KernelRestrictions{
	0:   {{Name: SyscallNotImplemented}},
	1:   {{Name: "exit"}},
	2:   {{Name: "fork"}},
	3:   {{Name: "read"}},
	4:   {{Name: "write"}},
	5:   {{Name: "open"}},
	6:   {{Name: "close"}},
	7:   {{Name: "restart_syscall"}},
	8:   {{Name: "creat"}},
	9:   {{Name: "link"}},
	10:  {{Name: "unlink"}},
	11:  {{Name: "execve"}},
	12:  {{Name: "chdir"}},
	13:  {{Name: SyscallNotImplemented}},
	14:  {{Name: "mknod"}},
	15:  {{Name: "chmod"}},
	16:  {{Name: SyscallNotImplemented}},
	17:  {{Name: SyscallNotImplemented}},
	18:  {{Name: SyscallNotImplemented}},
	19:  {{Name: "lseek"}},
	20:  {{Name: "getpid"}},
	21:  {{Name: "mount"}},
	22:  {{Name: "umount"}},
	23:  {{Name: SyscallNotImplemented}},
	24:  {{Name: SyscallNotImplemented}},
	25:  {{Name: SyscallNotImplemented}},
	26:  {{Name: "ptrace"}},
	27:  {{Name: "alarm"}},
	28:  {{Name: SyscallNotImplemented}},
	29:  {{Name: "pause"}},
	30:  {{Name: "utime"}},
	31:  {{Name: SyscallNotImplemented}},
	32:  {{Name: SyscallNotImplemented}},
	33:  {{Name: "access"}},
	34:  {{Name: "nice"}},
	35:  {{Name: SyscallNotImplemented}},
	36:  {{Name: "sync"}},
	37:  {{Name: "kill"}},
	38:  {{Name: "rename"}},
	39:  {{Name: "mkdir"}},
	40:  {{Name: "rmdir"}},
	41:  {{Name: "dup"}},
	42:  {{Name: "pipe"}},
	43:  {{Name: "times"}},
	44:  {{Name: SyscallNotImplemented}},
	45:  {{Name: "brk"}},
	46:  {{Name: SyscallNotImplemented}},
	47:  {{Name: SyscallNotImplemented}},
	48:  {{Name: "signal"}},
	49:  {{Name: SyscallNotImplemented}},
	50:  {{Name: SyscallNotImplemented}},
	51:  {{Name: "acct"}},
	52:  {{Name: "umount2"}},
	53:  {{Name: SyscallNotImplemented}},
	54:  {{Name: "ioctl"}},
	55:  {{Name: "fcntl"}},
	56:  {{Name: SyscallNotImplemented}},
	57:  {{Name: "setpgid"}},
	58:  {{Name: SyscallNotImplemented}},
	59:  {{Name: SyscallNotImplemented}},
	60:  {{Name: "umask"}},
	61:  {{Name: "chroot"}},
	62:  {{Name: "ustat"}},
	63:  {{Name: "dup2"}},
	64:  {{Name: "getppid"}},
	65:  {{Name: "getpgrp"}},
	66:  {{Name: "setsid"}},
	67:  {{Name: "sigaction"}},
	68:  {{Name: SyscallNotImplemented}},
	69:  {{Name: SyscallNotImplemented}},
	70:  {{Name: SyscallNotImplemented}},
	71:  {{Name: SyscallNotImplemented}},
	72:  {{Name: "sigsuspend"}},
	73:  {{Name: "sigpending"}},
	74:  {{Name: "sethostname"}},
	75:  {{Name: "setrlimit"}},
	76:  {{Name: SyscallNotImplemented}},
	77:  {{Name: "getrusage"}},
	78:  {{Name: "gettimeofday"}},
	79:  {{Name: "settimeofday"}},
	80:  {{Name: SyscallNotImplemented}},
	81:  {{Name: SyscallNotImplemented}},
	82:  {{Name: SyscallNotImplemented}},
	83:  {{Name: "symlink"}},
	84:  {{Name: SyscallNotImplemented}},
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
	95:  {{Name: SyscallNotImplemented}},
	96:  {{Name: "getpriority"}},
	97:  {{Name: "setpriority"}},
	98:  {{Name: SyscallNotImplemented}},
	99:  {{Name: "statfs"}},
	100: {{Name: "fstatfs"}},
	101: {{Name: SyscallNotImplemented}},
	102: {{Name: "socketcall"}},
	103: {{Name: "syslog"}},
	104: {{Name: "setitimer"}},
	105: {{Name: "getitimer"}},
	106: {{Name: "stat"}},
	107: {{Name: "lstat"}},
	108: {{Name: "fstat"}},
	109: {{Name: SyscallNotImplemented}},
	110: {{Name: SyscallNotImplemented + "lookup_dcookie"}},
	111: {{Name: "vhangup"}},
	112: {{Name: SyscallNotImplemented + "idle"}},
	113: {{Name: SyscallNotImplemented}},
	114: {{Name: "wait4"}},
	115: {{Name: "swapoff"}},
	116: {{Name: "sysinfo"}},
	117: {{Name: "ipc"}},
	118: {{Name: "fsync"}},
	119: {{Name: "sigreturn"}},
	120: {{Name: "clone"}},
	121: {{Name: "setdomainname"}},
	122: {{Name: "uname"}},
	123: {{Name: SyscallNotImplemented}},
	124: {{Name: "adjtimex"}},
	125: {{Name: "mprotect"}},
	126: {{Name: "sigprocmask"}},
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
	138: {{Name: SyscallNotImplemented}},
	139: {{Name: SyscallNotImplemented}},
	140: {{Name: SyscallNotImplemented}},
	141: {{Name: "getdents"}},
	142: {{Name: "select"}},
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
	164: {{Name: SyscallNotImplemented}},
	165: {{Name: SyscallNotImplemented}},
	166: {{Name: SyscallNotImplemented}},
	167: {{Name: SyscallNotImplemented + "query_module"}},
	168: {{Name: "poll"}},
	169: {{Name: SyscallNotImplemented + "nfsservctl"}},
	170: {{Name: SyscallNotImplemented}},
	171: {{Name: SyscallNotImplemented}},
	172: {{Name: "prctl"}},
	173: {{Name: "rt_sigreturn"}},
	174: {{Name: "rt_sigaction"}},
	175: {{Name: "rt_sigprocmask"}},
	176: {{Name: "rt_sigpending"}},
	177: {{Name: "rt_sigtimedwait"}},
	178: {{Name: "rt_sigqueueinfo"}},
	179: {{Name: "rt_sigsuspend"}},
	180: {{Name: "pread64"}},
	181: {{Name: "pwrite64"}},
	182: {{Name: SyscallNotImplemented}},
	183: {{Name: "getcwd"}},
	184: {{Name: "capget"}},
	185: {{Name: "capset"}},
	186: {{Name: "sigaltstack"}},
	187: {{Name: "sendfile"}},
	188: {{Name: SyscallNotImplemented + "getpmsg"}},
	189: {{Name: SyscallNotImplemented + "putpmsg"}},
	190: {{Name: "vfork"}},
	191: {{Name: "getrlimit"}},
	192: {{Name: SyscallNotImplemented}},
	193: {{Name: SyscallNotImplemented}},
	194: {{Name: SyscallNotImplemented}},
	195: {{Name: SyscallNotImplemented}},
	196: {{Name: SyscallNotImplemented}},
	197: {{Name: SyscallNotImplemented}},
	198: {{Name: "lchown"}},
	199: {{Name: "getuid"}},
	200: {{Name: "getgid"}},
	201: {{Name: "geteuid"}},
	202: {{Name: "getegid"}},
	203: {{Name: "setreuid"}},
	204: {{Name: "setregid"}},
	205: {{Name: "getgroups"}},
	206: {{Name: "setgroups"}},
	207: {{Name: "fchown"}},
	208: {{Name: "setresuid"}},
	209: {{Name: "getresuid"}},
	210: {{Name: "setresgid"}},
	211: {{Name: "getresgid"}},
	212: {{Name: "chown"}},
	213: {{Name: "setuid"}},
	214: {{Name: "setgid"}},
	215: {{Name: "setfsuid"}},
	216: {{Name: "setfsgid"}},
	217: {{Name: "pivot_root"}},
	218: {{Name: "mincore"}},
	219: {{Name: "madvise"}},
	220: {{Name: "getdents64"}},
	221: {{Name: SyscallNotImplemented}},
	222: {{Name: "readahead"}},
	223: {{Name: SyscallNotImplemented}},
	224: {{Name: "setxattr"}},
	225: {{Name: "lsetxattr"}},
	226: {{Name: "fsetxattr"}},
	227: {{Name: "getxattr"}},
	228: {{Name: "lgetxattr"}},
	229: {{Name: "fgetxattr"}},
	230: {{Name: "listxattr"}},
	231: {{Name: "llistxattr"}},
	232: {{Name: "flistxattr"}},
	233: {{Name: "removexattr"}},
	234: {{Name: "lremovexattr"}},
	235: {{Name: "fremovexattr"}},
	236: {{Name: "gettid"}},
	237: {{Name: "tkill"}},
	238: {{Name: "futex"}},
	239: {{Name: "sched_setaffinity"}},
	240: {{Name: "sched_getaffinity"}},
	241: {{Name: "tgkill"}},
	242: {{Name: SyscallNotImplemented}},
	243: {{Name: "io_setup"}},
	244: {{Name: "io_destroy"}},
	245: {{Name: "io_getevents"}},
	246: {{Name: "io_submit"}},
	247: {{Name: "io_cancel"}},
	248: {{Name: "exit_group"}},
	249: {{Name: "epoll_create"}},
	250: {{Name: "epoll_ctl"}},
	251: {{Name: "epoll_wait"}},
	252: {{Name: "set_tid_address"}},
	253: {{Name: "fadvise64"}},
	254: {{Name: "timer_create"}},
	255: {{Name: "timer_settime"}},
	256: {{Name: "timer_gettime"}},
	257: {{Name: "timer_getoverrun"}},
	258: {{Name: "timer_delete"}},
	259: {{Name: "clock_settime"}},
	260: {{Name: "clock_gettime"}},
	261: {{Name: "clock_getres"}},
	262: {{Name: "clock_nanosleep"}},
	263: {{Name: SyscallNotImplemented}},
	264: {{Name: SyscallNotImplemented}},
	265: {{Name: "statfs64"}},
	266: {{Name: "fstatfs64"}},
	267: {{Name: "remap_file_pages"}},
	268: {{Name: "mbind"}},
	269: {{Name: "get_mempolicy"}},
	270: {{Name: "set_mempolicy"}},
	271: {{Name: "mq_open"}},
	272: {{Name: "mq_unlink"}},
	273: {{Name: "mq_timedsend"}},
	274: {{Name: "mq_timedreceive"}},
	275: {{Name: "mq_notify"}},
	276: {{Name: "mq_getsetattr"}},
	277: {{Name: "kexec_load"}},
	278: {{Name: "add_key"}},
	279: {{Name: "request_key"}},
	280: {{Name: "keyctl"}},
	281: {{Name: "waitid"}},
	282: {{Name: "ioprio_set"}},
	283: {{Name: "ioprio_get"}},
	284: {{Name: "inotify_init"}},
	285: {{Name: "inotify_add_watch"}},
	286: {{Name: "inotify_rm_watch"}},
	287: {{Name: "migrate_pages"}},
	288: {{Name: "openat"}},
	289: {{Name: "mkdirat"}},
	290: {{Name: "mknodat"}},
	291: {{Name: "fchownat"}},
	292: {{Name: "futimesat"}},
	293: {{Name: "newfstatat"}},
	294: {{Name: "unlinkat"}},
	295: {{Name: "renameat"}},
	296: {{Name: "linkat"}},
	297: {{Name: "symlinkat"}},
	298: {{Name: "readlinkat"}},
	299: {{Name: "fchmodat"}},
	300: {{Name: "faccessat"}},
	301: {{Name: "pselect6"}},
	302: {{Name: "ppoll"}},
	303: {{Name: "unshare"}},
	304: {{Name: "set_robust_list"}},
	305: {{Name: "get_robust_list"}},
	306: {{Name: "splice"}},
	307: {{Name: "sync_file_range"}},
	308: {{Name: "tee"}},
	309: {{Name: "vmsplice"}},
	310: {{Name: "move_pages"}},
	311: {{Name: "getcpu"}},
	312: {{Name: "epoll_pwait"}},
	313: {{Name: "utimes"}},
	314: {{Name: "fallocate"}},
	315: {{Name: "utimensat"}},
	316: {{Name: "signalfd"}},
	317: {{Name: SyscallNotImplemented + "timerfd"}},
	318: {{Name: "eventfd"}},
	319: {{Name: "timerfd_create"}},
	320: {{Name: "timerfd_settime"}},
	321: {{Name: "timerfd_gettime"}},
	322: {{Name: "signalfd4"}},
	323: {{Name: "eventfd2"}},
	324: {{Name: "inotify_init1"}},
	325: {{Name: "pipe2"}},
	326: {{Name: "dup3"}},
	327: {{Name: "epoll_create1"}},
	328: {{Name: "preadv"}},
	329: {{Name: "pwritev"}},
	330: {{Name: "rt_tgsigqueueinfo"}},
	331: {{Name: "perf_event_open"}},
	332: {{Name: "fanotify_init"}},
	333: {{Name: "fanotify_mark"}},
	334: {{Name: "prlimit64"}},
	335: {{Name: "name_to_handle_at"}},
	336: {{Name: "open_by_handle_at"}},
	337: {{Name: "clock_adjtime"}},
	338: {{Name: "syncfs"}},
	339: {{Name: "setns"}},
	340: {{Name: "process_vm_readv"}},
	341: {{Name: "process_vm_writev"}},
	342: {{Name: "s390_runtime_instr"}},
	343: {{Name: "kcmp"}},
	344: {{Name: "finit_module"}},
	345: {{Name: "sched_setattr"}},
	346: {{Name: "sched_getattr"}},
	347: {{Name: "renameat2"}},
	348: {{Name: "seccomp"}},
	349: {{Name: "getrandom"}},
	350: {{Name: "memfd_create"}},
	351: {{Name: "bpf"}},
	352: {{Name: "s390_pci_mmio_write"}},
	353: {{Name: "s390_pci_mmio_read"}},
	354: {{Name: "execveat"}},
	355: {{Name: "userfaultfd"}},
	356: {{Name: "membarrier"}},
	357: {{Name: "recvmmsg"}},
	358: {{Name: "sendmmsg"}},
	359: {{Name: "socket"}},
	360: {{Name: "socketpair"}},
	361: {{Name: "bind"}},
	362: {{Name: "connect"}},
	363: {{Name: "listen"}},
	364: {{Name: "accept4"}},
	365: {{Name: "getsockopt"}},
	366: {{Name: "setsockopt"}},
	367: {{Name: "getsockname"}},
	368: {{Name: "getpeername"}},
	369: {{Name: "sendto"}},
	370: {{Name: "sendmsg"}},
	371: {{Name: "recvfrom"}},
	372: {{Name: "recvmsg"}},
	373: {{Name: "shutdown"}},
	374: {{Name: "mlock2"}},
	375: {{Name: "copy_file_range"}},
	376: {{Name: "preadv2"}},
	377: {{Name: "pwritev2"}},
	378: {{Name: "s390_guarded_storage"}},
	379: {{Name: "statx"}},
	380: {{Name: "s390_sthyi"}},
	381: {{Name: "kexec_file_load"}},
	382: {{Name: "io_pgetevents"}},
	383: {{Name: "rseq"}},
	384: {{Name: "pkey_mprotect"}},
	385: {{Name: "pkey_alloc"}},
	386: {{Name: "pkey_free"}},
	387: {{Name: SyscallNotImplemented}},
	388: {{Name: SyscallNotImplemented}},
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
	447: {{Name: "memfd_secret"}},
	448: {{Name: "process_mrelease"}},
	449: {{Name: "futex_waitv"}},
	450: {{Name: "set_mempolicy_home_node"}},
	451: {{Name: "cachestat"}},
	452: {{Name: "fchmodat2"}},
	453: {{Name: "map_shadow_stack"}},
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
