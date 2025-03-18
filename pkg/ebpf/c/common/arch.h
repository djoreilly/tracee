#ifndef __COMMON_ARCH_H__
#define __COMMON_ARCH_H__

#include <vmlinux.h>

#include <bpf/bpf_tracing.h>

#include <common/common.h>

// PROTOTYPES

statfunc bool is_x86_compat(struct task_struct *);
statfunc bool is_arm64_compat(struct task_struct *);
statfunc bool is_ppc64_compat(struct task_struct *);
statfunc bool is_s390x_compat(struct task_struct *);
statfunc bool is_compat(struct task_struct *);
statfunc int get_syscall_id_from_regs(struct pt_regs *);
statfunc struct pt_regs *get_current_task_pt_regs(void);
statfunc bool has_syscall_fd_arg(uint);
statfunc uint get_syscall_fd_num_from_arg(uint syscall_id, args_t *);

// FUNCTIONS

statfunc bool is_x86_compat(struct task_struct *task)
{
#if defined(bpf_target_x86)
    return BPF_CORE_READ(task, thread_info.status) & TS_COMPAT;
#else
    return false;
#endif
}

statfunc bool is_arm64_compat(struct task_struct *task)
{
#if defined(bpf_target_arm64)
    return BPF_CORE_READ(task, thread_info.flags) & _TIF_32BIT;
#else
    return false;
#endif
}

statfunc bool is_ppc64_compat(struct task_struct *task)
{
#if defined(bpf_target_powerpc)
    return BPF_CORE_READ(task, thread_info.flags) & _TIF_32BIT;
#else
    return false;
#endif
}

statfunc bool is_s390x_compat(struct task_struct *task)
{
#if defined(bpf_target_s390)
    return BPF_CORE_READ(task, thread_info.flags) & _TIF_31BIT;
#else
    return false;
#endif
}

statfunc bool is_compat(struct task_struct *task)
{
#if defined(bpf_target_x86)
    return is_x86_compat(task);
#elif defined(bpf_target_arm64)
    return is_arm64_compat(task);
#elif defined(bpf_target_powerpc)
    return is_ppc64_compat(task);
#elif defined(bpf_target_s390)
    return is_s390x_compat(task);
#else
    return false;
#endif
}

statfunc int get_syscall_id_from_regs(struct pt_regs *regs)
{
#if defined(bpf_target_x86)
    int id = BPF_CORE_READ(regs, orig_ax);
#elif defined(bpf_target_arm64)
    int id = BPF_CORE_READ(regs, syscallno);
#elif defined(bpf_target_powerpc)
    int id = BPF_CORE_READ(regs, gpr[0]);
#elif defined(bpf_target_s390)
    int id = BPF_CORE_READ(regs, gprs[2]);
#endif
    return id;
}

statfunc struct pt_regs *get_current_task_pt_regs(void)
{
    struct task_struct *task;

    // Use the bpf_task_pt_regs helper if possible
    if (bpf_core_enum_value_exists(enum bpf_func_id, BPF_FUNC_get_current_task_btf) &&
        bpf_core_enum_value_exists(enum bpf_func_id, BPF_FUNC_task_pt_regs)) {
        task = bpf_get_current_task_btf();
        return (struct pt_regs *) bpf_task_pt_regs(task);
    }

    // Helper not available, extract registers manually
    task = (struct task_struct *) bpf_get_current_task();

// THREAD_SIZE here is statistically defined and assumed to work for 4k page sizes.
#if defined(bpf_target_x86)
    void *__ptr = BPF_CORE_READ(task, stack) + THREAD_SIZE - TOP_OF_KERNEL_STACK_PADDING;
    return ((struct pt_regs *) __ptr) - 1;
#elif defined(bpf_target_arm64)
    return ((struct pt_regs *) (THREAD_SIZE + BPF_CORE_READ(task, stack)) - 1);
#elif defined(bpf_target_powerpc)
    return ((struct pt_regs *) (THREAD_SIZE + BPF_CORE_READ(task, stack)) - 1);
#elif defined(bpf_target_s390)
    return ((struct pt_regs *) (THREAD_SIZE + BPF_CORE_READ(task, stack)) - 1);
#endif
}

#define UNDEFINED_SYSCALL 1000
#define NO_SYSCALL        -1

#if defined(bpf_target_x86)
    #define SYSCALL_READ                   0
    #define SYSCALL_WRITE                  1
    #define SYSCALL_OPEN                   2
    #define SYSCALL_CLOSE                  3
    #define SYSCALL_FSTAT                  5
    #define SYSCALL_LSEEK                  8
    #define SYSCALL_MMAP                   9
    #define SYSCALL_MPROTECT               10
    #define SYSCALL_RT_SIGRETURN           15
    #define SYSCALL_IOCTL                  16
    #define SYSCALL_PREAD64                17
    #define SYSCALL_PWRITE64               18
    #define SYSCALL_READV                  19
    #define SYSCALL_WRITEV                 20
    #define SYSCALL_DUP                    32
    #define SYSCALL_DUP2                   33
    #define SYSCALL_SOCKET                 41
    #define SYSCALL_CONNECT                42
    #define SYSCALL_ACCEPT                 43
    #define SYSCALL_SENDTO                 44
    #define SYSCALL_RECVFROM               45
    #define SYSCALL_SENDMSG                46
    #define SYSCALL_RECVMSG                47
    #define SYSCALL_SHUTDOWN               48
    #define SYSCALL_BIND                   49
    #define SYSCALL_LISTEN                 50
    #define SYSCALL_GETSOCKNAME            51
    #define SYSCALL_GETPEERNAME            52
    #define SYSCALL_SETSOCKOPT             54
    #define SYSCALL_GETSOCKOPT             55
    #define SYSCALL_EXECVE                 59
    #define SYSCALL_EXIT                   60
    #define SYSCALL_FCNTL                  72
    #define SYSCALL_FLOCK                  73
    #define SYSCALL_FSYNC                  74
    #define SYSCALL_FDATASYNC              75
    #define SYSCALL_FTRUNCATE              77
    #define SYSCALL_GETDENTS               78
    #define SYSCALL_CHDIR                  80
    #define SYSCALL_FCHDIR                 81
    #define SYSCALL_FCHMOD                 91
    #define SYSCALL_FCHOWN                 93
    #define SYSCALL_PTRACE                 101
    #define SYSCALL_FSTATFS                138
    #define SYSCALL_ARCH_PRCTL             158
    #define SYSCALL_READAHEAD              187
    #define SYSCALL_FSETXATTR              190
    #define SYSCALL_FGETXATTR              193
    #define SYSCALL_FLISTXATTR             196
    #define SYSCALL_FREMOVEXATTR           199
    #define SYSCALL_GETDENTS64             217
    #define SYSCALL_FADVISE64              221
    #define SYSCALL_EXIT_GROUP             231
    #define SYSCALL_EPOLL_WAIT             232
    #define SYSCALL_EPOLL_CTL              233
    #define SYSCALL_INOTIFY_ADD_WATCH      254
    #define SYSCALL_INOTIFY_RM_WATCH       255
    #define SYSCALL_OPENAT                 257
    #define SYSCALL_MKDIRAT                258
    #define SYSCALL_MKNODAT                259
    #define SYSCALL_FCHOWNAT               260
    #define SYSCALL_FUTIMESAT              261
    #define SYSCALL_NEWFSTATAT             262
    #define SYSCALL_UNLINKAT               263
    #define SYSCALL_SYMLINKAT              266
    #define SYSCALL_READLINKAT             267
    #define SYSCALL_FCHMODAT               268
    #define SYSCALL_FACCESSAT              269
    #define SYSCALL_SYNC_FILE_RANGE        277
    #define SYSCALL_VMSPLICE               278
    #define SYSCALL_UTIMENSAT              280
    #define SYSCALL_EPOLL_PWAIT            281
    #define SYSCALL_SIGNALFD               282
    #define SYSCALL_FALLOCATE              285
    #define SYSCALL_TIMERFD_SETTIME        286
    #define SYSCALL_TIMERFD_GETTIME        287
    #define SYSCALL_ACCEPT4                288
    #define SYSCALL_SIGNALFD4              289
    #define SYSCALL_DUP3                   292
    #define SYSCALL_PREADV                 295
    #define SYSCALL_PWRITEV                296
    #define SYSCALL_PERF_EVENT_OPEN        298
    #define SYSCALL_RECVMMSG               299
    #define SYSCALL_NAME_TO_HANDLE_AT      303
    #define SYSCALL_OPEN_BY_HANDLE_AT      304
    #define SYSCALL_SYNCFS                 306
    #define SYSCALL_SENDMMSG               307
    #define SYSCALL_SETNS                  308
    #define SYSCALL_PROCESS_VM_WRITEV      311
    #define SYSCALL_FINIT_MODULE           313
    #define SYSCALL_EXECVEAT               322
    #define SYSCALL_PREADV2                327
    #define SYSCALL_PWRITEV2               328
    #define SYSCALL_PKEY_MPROTECT          329
    #define SYSCALL_STATX                  332
    #define SYSCALL_PIDFD_SEND_SIGNAL      424
    #define SYSCALL_IO_URING_ENTER         426
    #define SYSCALL_IO_URING_REGISTER      427
    #define SYSCALL_OPEN_TREE              428
    #define SYSCALL_FSCONFIG               431
    #define SYSCALL_FSMOUNT                432
    #define SYSCALL_FSPICK                 433
    #define SYSCALL_OPENAT2                437
    #define SYSCALL_FACCESSAT2             439
    #define SYSCALL_PROCESS_MADVISE        440
    #define SYSCALL_EPOLL_PWAIT2           441
    #define SYSCALL_MOUNT_SETATTR          442
    #define SYSCALL_QUOTACTL_FD            443
    #define SYSCALL_LANDLOCK_ADD_RULE      445
    #define SYSCALL_LANDLOCK_RESTRICT_SELF 446
    #define SYSCALL_PROCESS_MRELEASE       448
    #define SYSCALL_SOCKETCALL             473

#elif defined(bpf_target_arm64)
    #define SYSCALL_FSETXATTR              7
    #define SYSCALL_FGETXATTR              10
    #define SYSCALL_FLISTXATTR             13
    #define SYSCALL_FREMOVEXATTR           16
    #define SYSCALL_EPOLL_CTL              21
    #define SYSCALL_EPOLL_PWAIT            22
    #define SYSCALL_DUP                    23
    #define SYSCALL_DUP3                   24
    #define SYSCALL_FCNTL                  25
    #define SYSCALL_INOTIFY_ADD_WATCH      27
    #define SYSCALL_INOTIFY_RM_WATCH       28
    #define SYSCALL_IOCTL                  29
    #define SYSCALL_FLOCK                  32
    #define SYSCALL_MKNODAT                33
    #define SYSCALL_MKDIRAT                34
    #define SYSCALL_UNLINKAT               35
    #define SYSCALL_SYMLINKAT              36
    #define SYSCALL_FSTATFS                44
    #define SYSCALL_FTRUNCATE              46
    #define SYSCALL_FALLOCATE              47
    #define SYSCALL_FACCESSAT              48
    #define SYSCALL_CHDIR                  49
    #define SYSCALL_FCHDIR                 50
    #define SYSCALL_FCHMOD                 52
    #define SYSCALL_FCHMODAT               53
    #define SYSCALL_FCHOWNAT               54
    #define SYSCALL_FCHOWN                 55
    #define SYSCALL_OPENAT                 56
    #define SYSCALL_CLOSE                  57
    #define SYSCALL_GETDENTS64             61
    #define SYSCALL_LSEEK                  62
    #define SYSCALL_READ                   63
    #define SYSCALL_WRITE                  64
    #define SYSCALL_READV                  65
    #define SYSCALL_WRITEV                 66
    #define SYSCALL_PREAD64                67
    #define SYSCALL_PWRITE64               68
    #define SYSCALL_PREADV                 69
    #define SYSCALL_PWRITEV                70
    #define SYSCALL_SIGNALFD4              74
    #define SYSCALL_VMSPLICE               75
    #define SYSCALL_READLINKAT             78
    #define SYSCALL_FSTAT                  80
    #define SYSCALL_FSYNC                  82
    #define SYSCALL_FDATASYNC              83
    #define SYSCALL_SYNC_FILE_RANGE        84
    #define SYSCALL_TIMERFD_SETTIME        86
    #define SYSCALL_TIMERFD_GETTIME        87
    #define SYSCALL_UTIMENSAT              88
    #define SYSCALL_EXIT                   93
    #define SYSCALL_EXIT_GROUP             94
    #define SYSCALL_PTRACE                 117
    #define SYSCALL_RT_SIGRETURN           139
    #define SYSCALL_SOCKET                 198
    #define SYSCALL_BIND                   200
    #define SYSCALL_LISTEN                 201
    #define SYSCALL_ACCEPT                 202
    #define SYSCALL_CONNECT                203
    #define SYSCALL_GETSOCKNAME            204
    #define SYSCALL_GETPEERNAME            205
    #define SYSCALL_SENDTO                 206
    #define SYSCALL_RECVFROM               207
    #define SYSCALL_SETSOCKOPT             208
    #define SYSCALL_GETSOCKOPT             209
    #define SYSCALL_SHUTDOWN               210
    #define SYSCALL_SENDMSG                211
    #define SYSCALL_RECVMSG                212
    #define SYSCALL_READAHEAD              213
    #define SYSCALL_EXECVE                 221
    #define SYSCALL_MMAP                   222
    #define SYSCALL_FADVISE64              223
    #define SYSCALL_MPROTECT               226
    #define SYSCALL_PERF_EVENT_OPEN        241
    #define SYSCALL_ACCEPT4                242
    #define SYSCALL_RECVMMSG               243
    #define SYSCALL_NAME_TO_HANDLE_AT      264
    #define SYSCALL_OPEN_BY_HANDLE_AT      265
    #define SYSCALL_SYNCFS                 267
    #define SYSCALL_SETNS                  268
    #define SYSCALL_SENDMMSG               269
    #define SYSCALL_PROCESS_VM_WRITEV      271
    #define SYSCALL_FINIT_MODULE           273
    #define SYSCALL_EXECVEAT               281
    #define SYSCALL_PREADV2                286
    #define SYSCALL_PWRITEV2               287
    #define SYSCALL_PKEY_MPROTECT          288
    #define SYSCALL_STATX                  291
    #define SYSCALL_PIDFD_SEND_SIGNAL      424
    #define SYSCALL_IO_URING_ENTER         426
    #define SYSCALL_IO_URING_REGISTER      427
    #define SYSCALL_OPEN_TREE              428
    #define SYSCALL_FSCONFIG               431
    #define SYSCALL_FSMOUNT                432
    #define SYSCALL_FSPICK                 433
    #define SYSCALL_OPENAT2                437
    #define SYSCALL_FACCESSAT2             439
    #define SYSCALL_PROCESS_MADVISE        440
    #define SYSCALL_EPOLL_PWAIT2           441
    #define SYSCALL_MOUNT_SETATTR          442
    #define SYSCALL_QUOTACTL_FD            443
    #define SYSCALL_LANDLOCK_ADD_RULE      445
    #define SYSCALL_LANDLOCK_RESTRICT_SELF 446
    #define SYSCALL_PROCESS_MRELEASE       448
    #define SYSCALL_SOCKETCALL             UNDEFINED_SYSCALL
    #define SYSCALL_OPEN                   UNDEFINED_SYSCALL
    #define SYSCALL_DUP2                   UNDEFINED_SYSCALL
    #define SYSCALL_GETDENTS               UNDEFINED_SYSCALL
    #define SYSCALL_FUTIMESAT              UNDEFINED_SYSCALL
    #define SYSCALL_NEWFSTATAT             UNDEFINED_SYSCALL
    #define SYSCALL_EPOLL_WAIT             UNDEFINED_SYSCALL
    #define SYSCALL_SIGNALFD               UNDEFINED_SYSCALL
    #define SYSCALL_ARCH_PRCTL             UNDEFINED_SYSCALL
#elif defined(bpf_target_powerpc)
    #define SYSCALL_RESTART_SYSCALL         0
    #define SYSCALL_EXIT                    1
    #define SYSCALL_FORK                    2
    #define SYSCALL_READ                    3
    #define SYSCALL_WRITE                   4
    #define SYSCALL_OPEN                    5
    #define SYSCALL_CLOSE                   6
    #define SYSCALL_WAITPID                 7
    #define SYSCALL_CREAT                   8
    #define SYSCALL_LINK                    9
    #define SYSCALL_UNLINK                  10
    #define SYSCALL_EXECVE                  11
    #define SYSCALL_CHDIR                   12
    #define SYSCALL_TIME                    13
    #define SYSCALL_MKNOD                   14
    #define SYSCALL_CHMOD                   15
    #define SYSCALL_LCHOWN                  16
    #define SYSCALL_LSEEK                   19
    #define SYSCALL_GETPID                  20
    #define SYSCALL_MOUNT                   21
    #define SYSCALL_SETUID                  23
    #define SYSCALL_GETUID                  24
    #define SYSCALL_STIME                   25
    #define SYSCALL_PTRACE                  26
    #define SYSCALL_ALARM                   27
    #define SYSCALL_PAUSE                   29
    #define SYSCALL_UTIME                   30
    #define SYSCALL_ACCESS                  33
    #define SYSCALL_NICE                    34
    #define SYSCALL_SYNC                    36
    #define SYSCALL_KILL                    37
    #define SYSCALL_RENAME                  38
    #define SYSCALL_MKDIR                   39
    #define SYSCALL_RMDIR                   40
    #define SYSCALL_DUP                     41
    #define SYSCALL_PIPE                    42
    #define SYSCALL_TIMES                   43
    #define SYSCALL_BRK                     45
    #define SYSCALL_SETGID                  46
    #define SYSCALL_GETGID                  47
    #define SYSCALL_SIGNAL                  48
    #define SYSCALL_GETEUID                 49
    #define SYSCALL_GETEGID                 50
    #define SYSCALL_ACCT                    51
    #define SYSCALL_UMOUNT2                 52
    #define SYSCALL_IOCTL                   54
    #define SYSCALL_FCNTL                   55
    #define SYSCALL_SETPGID                 57
    #define SYSCALL_UMASK                   60
    #define SYSCALL_CHROOT                  61
    #define SYSCALL_USTAT                   62
    #define SYSCALL_DUP2                    63
    #define SYSCALL_GETPPID                 64
    #define SYSCALL_GETPGRP                 65
    #define SYSCALL_SETSID                  66
    #define SYSCALL_SGETMASK                68
    #define SYSCALL_SSETMASK                69
    #define SYSCALL_SETREUID                70
    #define SYSCALL_SETREGID                71
    #define SYSCALL_SETHOSTNAME             74
    #define SYSCALL_SETRLIMIT               75
    #define SYSCALL_GETRUSAGE               77
    #define SYSCALL_GETTIMEOFDAY            78
    #define SYSCALL_SETTIMEOFDAY            79
    #define SYSCALL_GETGROUPS               80
    #define SYSCALL_SETGROUPS               81
    #define SYSCALL_SYMLINK                 83
    #define SYSCALL_READLINK                85
    #define SYSCALL_USELIB                  86
    #define SYSCALL_SWAPON                  87
    #define SYSCALL_REBOOT                  88
    #define SYSCALL_MMAP                    90
    #define SYSCALL_MUNMAP                  91
    #define SYSCALL_TRUNCATE                92
    #define SYSCALL_FTRUNCATE               93
    #define SYSCALL_FCHMOD                  94
    #define SYSCALL_FCHOWN                  95
    #define SYSCALL_GETPRIORITY             96
    #define SYSCALL_SETPRIORITY             97
    #define SYSCALL_STATFS                  99
    #define SYSCALL_FSTATFS                 100
    #define SYSCALL_SOCKETCALL              102
    #define SYSCALL_SYSLOG                  103
    #define SYSCALL_SETITIMER               104
    #define SYSCALL_GETITIMER               105
    #define SYSCALL_STAT                    106
    #define SYSCALL_LSTAT                   107
    #define SYSCALL_FSTAT                   108
    #define SYSCALL_VHANGUP                 111
    #define SYSCALL_WAIT4                   114
    #define SYSCALL_SWAPOFF                 115
    #define SYSCALL_SYSINFO                 116
    #define SYSCALL_IPC                     117
    #define SYSCALL_FSYNC                   118
    #define SYSCALL_CLONE                   120
    #define SYSCALL_SETDOMAINNAME           121
    #define SYSCALL_UNAME                   122
    #define SYSCALL_ADJTIMEX                124
    #define SYSCALL_MPROTECT                125
    #define SYSCALL_INIT_MODULE             128
    #define SYSCALL_DELETE_MODULE           129
    #define SYSCALL_QUOTACTL                131
    #define SYSCALL_GETPGID                 132
    #define SYSCALL_FCHDIR                  133
    #define SYSCALL_SYSFS                   135
    #define SYSCALL_PERSONALITY             136
    #define SYSCALL_SETFSUID                138
    #define SYSCALL_SETFSGID                139
    #define SYSCALL__LLSEEK                 140
    #define SYSCALL_GETDENTS                141
    #define SYSCALL__NEWSELECT              142
    #define SYSCALL_FLOCK                   143
    #define SYSCALL_MSYNC                   144
    #define SYSCALL_READV                   145
    #define SYSCALL_WRITEV                  146
    #define SYSCALL_GETSID                  147
    #define SYSCALL_FDATASYNC               148
    #define SYSCALL_MLOCK                   150
    #define SYSCALL_MUNLOCK                 151
    #define SYSCALL_MLOCKALL                152
    #define SYSCALL_MUNLOCKALL              153
    #define SYSCALL_SCHED_SETPARAM          154
    #define SYSCALL_SCHED_GETPARAM          155
    #define SYSCALL_SCHED_SETSCHEDULER      156
    #define SYSCALL_SCHED_GETSCHEDULER      157
    #define SYSCALL_SCHED_YIELD             158
    #define SYSCALL_SCHED_GET_PRIORITY_MAX  159
    #define SYSCALL_SCHED_GET_PRIORITY_MIN  160
    #define SYSCALL_SCHED_RR_GET_INTERVAL   161
    #define SYSCALL_NANOSLEEP               162
    #define SYSCALL_MREMAP                  163
    #define SYSCALL_SETRESUID               164
    #define SYSCALL_GETRESUID               165
    #define SYSCALL_POLL                    167
    #define SYSCALL_SETRESGID               169
    #define SYSCALL_GETRESGID               170
    #define SYSCALL_PRCTL                   171
    #define SYSCALL_RT_SIGRETURN            172
    #define SYSCALL_RT_SIGACTION            173
    #define SYSCALL_RT_SIGPROCMASK          174
    #define SYSCALL_RT_SIGPENDING           175
    #define SYSCALL_RT_SIGTIMEDWAIT         176
    #define SYSCALL_RT_SIGQUEUEINFO         177
    #define SYSCALL_RT_SIGSUSPEND           178
    #define SYSCALL_PREAD64                 179
    #define SYSCALL_PWRITE64                180
    #define SYSCALL_CHOWN                   181
    #define SYSCALL_GETCWD                  182
    #define SYSCALL_CAPGET                  183
    #define SYSCALL_CAPSET                  184
    #define SYSCALL_SIGALTSTACK             185
    #define SYSCALL_SENDFILE                186
    #define SYSCALL_VFORK                   189
    #define SYSCALL_UGETRLIMIT              190
    #define SYSCALL_READAHEAD               191
    #define SYSCALL_PCICONFIG_READ          198
    #define SYSCALL_PCICONFIG_WRITE         199
    #define SYSCALL_PCICONFIG_IOBASE        200
    #define SYSCALL_GETDENTS64              202
    #define SYSCALL_PIVOT_ROOT              203
    #define SYSCALL_MADVISE                 205
    #define SYSCALL_MINCORE                 206
    #define SYSCALL_GETTID                  207
    #define SYSCALL_TKILL                   208
    #define SYSCALL_SETXATTR                209
    #define SYSCALL_LSETXATTR               210
    #define SYSCALL_FSETXATTR               211
    #define SYSCALL_GETXATTR                212
    #define SYSCALL_LGETXATTR               213
    #define SYSCALL_FGETXATTR               214
    #define SYSCALL_LISTXATTR               215
    #define SYSCALL_LLISTXATTR              216
    #define SYSCALL_FLISTXATTR              217
    #define SYSCALL_REMOVEXATTR             218
    #define SYSCALL_LREMOVEXATTR            219
    #define SYSCALL_FREMOVEXATTR            220
    #define SYSCALL_FUTEX                   221
    #define SYSCALL_SCHED_SETAFFINITY       222
    #define SYSCALL_SCHED_GETAFFINITY       223
    #define SYSCALL_IO_SETUP                227
    #define SYSCALL_IO_DESTROY              228
    #define SYSCALL_IO_GETEVENTS            229
    #define SYSCALL_IO_SUBMIT               230
    #define SYSCALL_IO_CANCEL               231
    #define SYSCALL_SET_TID_ADDRESS         232
    #define SYSCALL_FADVISE64               233
    #define SYSCALL_EXIT_GROUP              234
    #define SYSCALL_EPOLL_CREATE            236
    #define SYSCALL_EPOLL_CTL               237
    #define SYSCALL_EPOLL_WAIT              238
    #define SYSCALL_REMAP_FILE_PAGES        239
    #define SYSCALL_TIMER_CREATE            240
    #define SYSCALL_TIMER_SETTIME           241
    #define SYSCALL_TIMER_GETTIME           242
    #define SYSCALL_TIMER_GETOVERRUN        243
    #define SYSCALL_TIMER_DELETE            244
    #define SYSCALL_CLOCK_SETTIME           245
    #define SYSCALL_CLOCK_GETTIME           246
    #define SYSCALL_CLOCK_GETRES            247
    #define SYSCALL_CLOCK_NANOSLEEP         248
    #define SYSCALL_SWAPCONTEXT             249
    #define SYSCALL_TGKILL                  250
    #define SYSCALL_UTIMES                  251
    #define SYSCALL_STATFS64                252
    #define SYSCALL_FSTATFS64               253
    #define SYSCALL_RTAS                    255
    #define SYSCALL_MIGRATE_PAGES           258
    #define SYSCALL_MBIND                   259
    #define SYSCALL_GET_MEMPOLICY           260
    #define SYSCALL_SET_MEMPOLICY           261
    #define SYSCALL_MQ_OPEN                 262
    #define SYSCALL_MQ_UNLINK               263
    #define SYSCALL_MQ_TIMEDSEND            264
    #define SYSCALL_MQ_TIMEDRECEIVE         265
    #define SYSCALL_MQ_NOTIFY               266
    #define SYSCALL_MQ_GETSETATTR           267
    #define SYSCALL_KEXEC_LOAD              268
    #define SYSCALL_ADD_KEY                 269
    #define SYSCALL_REQUEST_KEY             270
    #define SYSCALL_KEYCTL                  271
    #define SYSCALL_WAITID                  272
    #define SYSCALL_IOPRIO_SET              273
    #define SYSCALL_IOPRIO_GET              274
    #define SYSCALL_INOTIFY_INIT            275
    #define SYSCALL_INOTIFY_ADD_WATCH       276
    #define SYSCALL_INOTIFY_RM_WATCH        277
    #define SYSCALL_SPU_RUN                 278
    #define SYSCALL_SPU_CREATE              279
    #define SYSCALL_PSELECT6                280
    #define SYSCALL_PPOLL                   281
    #define SYSCALL_UNSHARE                 282
    #define SYSCALL_SPLICE                  283
    #define SYSCALL_TEE                     284
    #define SYSCALL_VMSPLICE                285
    #define SYSCALL_OPENAT                  286
    #define SYSCALL_MKDIRAT                 287
    #define SYSCALL_MKNODAT                 288
    #define SYSCALL_FCHOWNAT                289
    #define SYSCALL_FUTIMESAT               290
    #define SYSCALL_NEWFSTATAT              291
    #define SYSCALL_UNLINKAT                292
    #define SYSCALL_RENAMEAT                293
    #define SYSCALL_LINKAT                  294
    #define SYSCALL_SYMLINKAT               295
    #define SYSCALL_READLINKAT              296
    #define SYSCALL_FCHMODAT                297
    #define SYSCALL_FACCESSAT               298
    #define SYSCALL_GET_ROBUST_LIST         299
    #define SYSCALL_SET_ROBUST_LIST         300
    #define SYSCALL_MOVE_PAGES              301
    #define SYSCALL_GETCPU                  302
    #define SYSCALL_EPOLL_PWAIT             303
    #define SYSCALL_UTIMENSAT               304
    #define SYSCALL_SIGNALFD                305
    #define SYSCALL_TIMERFD_CREATE          306
    #define SYSCALL_EVENTFD                 307
    #define SYSCALL_SYNC_FILE_RANGE         308
    #define SYSCALL_FALLOCATE               309
    #define SYSCALL_SUBPAGE_PROT            310
    #define SYSCALL_TIMERFD_SETTIME         311
    #define SYSCALL_TIMERFD_GETTIME         312
    #define SYSCALL_SIGNALFD4               313
    #define SYSCALL_EVENTFD2                314
    #define SYSCALL_EPOLL_CREATE1           315
    #define SYSCALL_DUP3                    316
    #define SYSCALL_PIPE2                   317
    #define SYSCALL_INOTIFY_INIT1           318
    #define SYSCALL_PERF_EVENT_OPEN         319
    #define SYSCALL_PREADV                  320
    #define SYSCALL_PWRITEV                 321
    #define SYSCALL_RT_TGSIGQUEUEINFO       322
    #define SYSCALL_FANOTIFY_INIT           323
    #define SYSCALL_FANOTIFY_MARK           324
    #define SYSCALL_PRLIMIT64               325
    #define SYSCALL_SOCKET                  326
    #define SYSCALL_BIND                    327
    #define SYSCALL_CONNECT                 328
    #define SYSCALL_LISTEN                  329
    #define SYSCALL_ACCEPT                  330
    #define SYSCALL_GETSOCKNAME             331
    #define SYSCALL_GETPEERNAME             332
    #define SYSCALL_SOCKETPAIR              333
    #define SYSCALL_SEND                    334
    #define SYSCALL_SENDTO                  335
    #define SYSCALL_RECV                    336
    #define SYSCALL_RECVFROM                337
    #define SYSCALL_SHUTDOWN                338
    #define SYSCALL_SETSOCKOPT              339
    #define SYSCALL_GETSOCKOPT              340
    #define SYSCALL_SENDMSG                 341
    #define SYSCALL_RECVMSG                 342
    #define SYSCALL_RECVMMSG                343
    #define SYSCALL_ACCEPT4                 344
    #define SYSCALL_NAME_TO_HANDLE_AT       345
    #define SYSCALL_OPEN_BY_HANDLE_AT       346
    #define SYSCALL_CLOCK_ADJTIME           347
    #define SYSCALL_SYNCFS                  348
    #define SYSCALL_SENDMMSG                349
    #define SYSCALL_SETNS                   350
    #define SYSCALL_PROCESS_VM_READV        351
    #define SYSCALL_PROCESS_VM_WRITEV       352
    #define SYSCALL_FINIT_MODULE            353
    #define SYSCALL_KCMP                    354
    #define SYSCALL_SCHED_SETATTR           355
    #define SYSCALL_SCHED_GETATTR           356
    #define SYSCALL_RENAMEAT2               357
    #define SYSCALL_SECCOMP                 358
    #define SYSCALL_GETRANDOM               359
    #define SYSCALL_MEMFD_CREATE            360
    #define SYSCALL_BPF                     361
    #define SYSCALL_EXECVEAT                362
    #define SYSCALL_SWITCH_ENDIAN           363
    #define SYSCALL_USERFAULTFD             364
    #define SYSCALL_MEMBARRIER              365
    #define SYSCALL_MLOCK2                  378
    #define SYSCALL_COPY_FILE_RANGE         379
    #define SYSCALL_PREADV2                 380
    #define SYSCALL_PWRITEV2                381
    #define SYSCALL_KEXEC_FILE_LOAD         382
    #define SYSCALL_STATX                   383
    #define SYSCALL_PKEY_ALLOC              384
    #define SYSCALL_PKEY_FREE               385
    #define SYSCALL_PKEY_MPROTECT           386
    #define SYSCALL_RSEQ                    387
    #define SYSCALL_IO_PGETEVENTS           388
    #define SYSCALL_SEMTIMEDOP              392
    #define SYSCALL_SEMGET                  393
    #define SYSCALL_SEMCTL                  394
    #define SYSCALL_SHMGET                  395
    #define SYSCALL_SHMCTL                  396
    #define SYSCALL_SHMAT                   397
    #define SYSCALL_SHMDT                   398
    #define SYSCALL_MSGGET                  399
    #define SYSCALL_MSGSND                  400
    #define SYSCALL_MSGRCV                  401
    #define SYSCALL_MSGCTL                  402
    #define SYSCALL_PIDFD_SEND_SIGNAL       424
    #define SYSCALL_IO_URING_SETUP          425
    #define SYSCALL_IO_URING_ENTER          426
    #define SYSCALL_IO_URING_REGISTER       427
    #define SYSCALL_OPEN_TREE               428
    #define SYSCALL_MOVE_MOUNT              429
    #define SYSCALL_FSOPEN                  430
    #define SYSCALL_FSCONFIG                431
    #define SYSCALL_FSMOUNT                 432
    #define SYSCALL_FSPICK                  433
    #define SYSCALL_PIDFD_OPEN              434
    #define SYSCALL_CLONE3                  435
    #define SYSCALL_CLOSE_RANGE             436
    #define SYSCALL_OPENAT2                 437
    #define SYSCALL_PIDFD_GETFD             438
    #define SYSCALL_FACCESSAT2              439
    #define SYSCALL_PROCESS_MADVISE         440
    #define SYSCALL_EPOLL_PWAIT2            441
    #define SYSCALL_MOUNT_SETATTR           442
    #define SYSCALL_QUOTACTL_FD             443
    #define SYSCALL_LANDLOCK_CREATE_RULESET 444
    #define SYSCALL_LANDLOCK_ADD_RULE       445
    #define SYSCALL_LANDLOCK_RESTRICT_SELF  446
    #define SYSCALL_PROCESS_MRELEASE        448
    #define SYSCALL_FUTEX_WAITV             449
    #define SYSCALL_SET_MEMPOLICY_HOME_NODE 450
    #define SYSCALL_CACHESTAT               451
    #define SYSCALL_FCHMODAT2               452
    #define SYSCALL_FUTEX_WAKE              454
    #define SYSCALL_FUTEX_WAIT              455
    #define SYSCALL_FUTEX_REQUEUE           456
    #define SYSCALL_STATMOUNT               457
    #define SYSCALL_LISTMOUNT               458
    #define SYSCALL_LSM_GET_SELF_ATTR       459
    #define SYSCALL_LSM_SET_SELF_ATTR       460
    #define SYSCALL_LSM_LIST_MODULES        461
    #define SYSCALL_MSEAL                   462
    #define SYSCALL_SETXATTRAT              463
    #define SYSCALL_GETXATTRAT              464
    #define SYSCALL_LISTXATTRAT             465
    #define SYSCALL_REMOVEXATTRAT           466
    #define SYSCALL_ARCH_PRCTL              UNDEFINED_SYSCALL
    #define SYSCALL_BREAK                   UNDEFINED_SYSCALL
    #define SYSCALL_OLDSTAT                 UNDEFINED_SYSCALL
    #define SYSCALL_UMOUNT                  UNDEFINED_SYSCALL
    #define SYSCALL_OLDFSTAT                UNDEFINED_SYSCALL
    #define SYSCALL_STTY                    UNDEFINED_SYSCALL
    #define SYSCALL_GTTY                    UNDEFINED_SYSCALL
    #define SYSCALL_FTIME                   UNDEFINED_SYSCALL
    #define SYSCALL_PROF                    UNDEFINED_SYSCALL
    #define SYSCALL_LOCK                    UNDEFINED_SYSCALL
    #define SYSCALL_MPX                     UNDEFINED_SYSCALL
    #define SYSCALL_ULIMIT                  UNDEFINED_SYSCALL
    #define SYSCALL_OLDOLDUNAME             UNDEFINED_SYSCALL
    #define SYSCALL_SIGACTION               UNDEFINED_SYSCALL
    #define SYSCALL_SIGSUSPEND              UNDEFINED_SYSCALL
    #define SYSCALL_SIGPENDING              UNDEFINED_SYSCALL
    #define SYSCALL_GETRLIMIT               UNDEFINED_SYSCALL
    #define SYSCALL_SELECT                  UNDEFINED_SYSCALL
    #define SYSCALL_OLDLSTAT                UNDEFINED_SYSCALL
    #define SYSCALL_READDIR                 UNDEFINED_SYSCALL
    #define SYSCALL_PROFIL                  UNDEFINED_SYSCALL
    #define SYSCALL_IOPERM                  UNDEFINED_SYSCALL
    #define SYSCALL_OLDUNAME                UNDEFINED_SYSCALL
    #define SYSCALL_IOPL                    UNDEFINED_SYSCALL
    #define SYSCALL_IDLE                    UNDEFINED_SYSCALL
    #define SYSCALL_VM86                    UNDEFINED_SYSCALL
    #define SYSCALL_SIGRETURN               UNDEFINED_SYSCALL
    #define SYSCALL_MODIFY_LDT              UNDEFINED_SYSCALL
    #define SYSCALL_SIGPROCMASK             UNDEFINED_SYSCALL
    #define SYSCALL_CREATE_MODULE           UNDEFINED_SYSCALL
    #define SYSCALL_GET_KERNEL_SYMS         UNDEFINED_SYSCALL
    #define SYSCALL_BDFLUSH                 UNDEFINED_SYSCALL
    #define SYSCALL_AFS_SYSCALL             UNDEFINED_SYSCALL
    #define SYSCALL__SYSCTL                 UNDEFINED_SYSCALL
    #define SYSCALL_QUERY_MODULE            UNDEFINED_SYSCALL
    #define SYSCALL_NFSSERVCTL              UNDEFINED_SYSCALL
    #define SYSCALL_GETPMSG                 UNDEFINED_SYSCALL
    #define SYSCALL_PUTPMSG                 UNDEFINED_SYSCALL
    #define SYSCALL_MULTIPLEXER             UNDEFINED_SYSCALL
    #define SYSCALL_TUXCALL                 UNDEFINED_SYSCALL
    #define SYSCALL_LOOKUP_DCOOKIE          UNDEFINED_SYSCALL
    #define SYSCALL_SYS_DEBUG_SETCONTEXT    UNDEFINED_SYSCALL
    #define SYSCALL_MAP_SHADOW_STACK        UNDEFINED_SYSCALL
#elif defined(bpf_target_s390)
    #define SYSCALL_EXIT                    1
    #define SYSCALL_FORK                    2
    #define SYSCALL_READ                    3
    #define SYSCALL_WRITE                   4
    #define SYSCALL_OPEN                    5
    #define SYSCALL_CLOSE                   6
    #define SYSCALL_RESTART_SYSCALL         7
    #define SYSCALL_CREAT                   8
    #define SYSCALL_LINK                    9
    #define SYSCALL_UNLINK                  10
    #define SYSCALL_EXECVE                  11
    #define SYSCALL_CHDIR                   12
    #define SYSCALL_MKNOD                   14
    #define SYSCALL_CHMOD                   15
    #define SYSCALL_LSEEK                   19
    #define SYSCALL_GETPID                  20
    #define SYSCALL_MOUNT                   21
    #define SYSCALL_UMOUNT                  22
    #define SYSCALL_PTRACE                  26
    #define SYSCALL_ALARM                   27
    #define SYSCALL_PAUSE                   29
    #define SYSCALL_UTIME                   30
    #define SYSCALL_ACCESS                  33
    #define SYSCALL_NICE                    34
    #define SYSCALL_SYNC                    36
    #define SYSCALL_KILL                    37
    #define SYSCALL_RENAME                  38
    #define SYSCALL_MKDIR                   39
    #define SYSCALL_RMDIR                   40
    #define SYSCALL_DUP                     41
    #define SYSCALL_PIPE                    42
    #define SYSCALL_TIMES                   43
    #define SYSCALL_BRK                     45
    #define SYSCALL_SIGNAL                  48
    #define SYSCALL_ACCT                    51
    #define SYSCALL_UMOUNT2                 52
    #define SYSCALL_IOCTL                   54
    #define SYSCALL_FCNTL                   55
    #define SYSCALL_SETPGID                 57
    #define SYSCALL_UMASK                   60
    #define SYSCALL_CHROOT                  61
    #define SYSCALL_USTAT                   62
    #define SYSCALL_DUP2                    63
    #define SYSCALL_GETPPID                 64
    #define SYSCALL_GETPGRP                 65
    #define SYSCALL_SETSID                  66
    #define SYSCALL_SIGACTION               67
    #define SYSCALL_SIGSUSPEND              72
    #define SYSCALL_SIGPENDING              73
    #define SYSCALL_SETHOSTNAME             74
    #define SYSCALL_SETRLIMIT               75
    #define SYSCALL_GETRUSAGE               77
    #define SYSCALL_GETTIMEOFDAY            78
    #define SYSCALL_SETTIMEOFDAY            79
    #define SYSCALL_SYMLINK                 83
    #define SYSCALL_READLINK                85
    #define SYSCALL_USELIB                  86
    #define SYSCALL_SWAPON                  87
    #define SYSCALL_REBOOT                  88
    #define SYSCALL_MMAP                    90
    #define SYSCALL_MUNMAP                  91
    #define SYSCALL_TRUNCATE                92
    #define SYSCALL_FTRUNCATE               93
    #define SYSCALL_FCHMOD                  94
    #define SYSCALL_GETPRIORITY             96
    #define SYSCALL_SETPRIORITY             97
    #define SYSCALL_STATFS                  99
    #define SYSCALL_FSTATFS                 100
    #define SYSCALL_SOCKETCALL              102
    #define SYSCALL_SYSLOG                  103
    #define SYSCALL_SETITIMER               104
    #define SYSCALL_GETITIMER               105
    #define SYSCALL_STAT                    106
    #define SYSCALL_LSTAT                   107
    #define SYSCALL_FSTAT                   108
    #define SYSCALL_VHANGUP                 111
    #define SYSCALL_WAIT4                   114
    #define SYSCALL_SWAPOFF                 115
    #define SYSCALL_SYSINFO                 116
    #define SYSCALL_IPC                     117
    #define SYSCALL_FSYNC                   118
    #define SYSCALL_SIGRETURN               119
    #define SYSCALL_CLONE                   120
    #define SYSCALL_SETDOMAINNAME           121
    #define SYSCALL_UNAME                   122
    #define SYSCALL_ADJTIMEX                124
    #define SYSCALL_MPROTECT                125
    #define SYSCALL_SIGPROCMASK             126
    #define SYSCALL_INIT_MODULE             128
    #define SYSCALL_DELETE_MODULE           129
    #define SYSCALL_QUOTACTL                131
    #define SYSCALL_GETPGID                 132
    #define SYSCALL_FCHDIR                  133
    #define SYSCALL_SYSFS                   135
    #define SYSCALL_PERSONALITY             136
    #define SYSCALL_GETDENTS                141
    #define SYSCALL_SELECT                  142
    #define SYSCALL_FLOCK                   143
    #define SYSCALL_MSYNC                   144
    #define SYSCALL_READV                   145
    #define SYSCALL_WRITEV                  146
    #define SYSCALL_GETSID                  147
    #define SYSCALL_FDATASYNC               148
    #define SYSCALL_MLOCK                   150
    #define SYSCALL_MUNLOCK                 151
    #define SYSCALL_MLOCKALL                152
    #define SYSCALL_MUNLOCKALL              153
    #define SYSCALL_SCHED_SETPARAM          154
    #define SYSCALL_SCHED_GETPARAM          155
    #define SYSCALL_SCHED_SETSCHEDULER      156
    #define SYSCALL_SCHED_GETSCHEDULER      157
    #define SYSCALL_SCHED_YIELD             158
    #define SYSCALL_SCHED_GET_PRIORITY_MAX  159
    #define SYSCALL_SCHED_GET_PRIORITY_MIN  160
    #define SYSCALL_SCHED_RR_GET_INTERVAL   161
    #define SYSCALL_NANOSLEEP               162
    #define SYSCALL_MREMAP                  163
    #define SYSCALL_POLL                    168
    #define SYSCALL_PRCTL                   172
    #define SYSCALL_RT_SIGRETURN            173
    #define SYSCALL_RT_SIGACTION            174
    #define SYSCALL_RT_SIGPROCMASK          175
    #define SYSCALL_RT_SIGPENDING           176
    #define SYSCALL_RT_SIGTIMEDWAIT         177
    #define SYSCALL_RT_SIGQUEUEINFO         178
    #define SYSCALL_RT_SIGSUSPEND           179
    #define SYSCALL_PREAD64                 180
    #define SYSCALL_PWRITE64                181
    #define SYSCALL_GETCWD                  183
    #define SYSCALL_CAPGET                  184
    #define SYSCALL_CAPSET                  185
    #define SYSCALL_SIGALTSTACK             186
    #define SYSCALL_SENDFILE                187
    #define SYSCALL_VFORK                   190
    #define SYSCALL_GETRLIMIT               191
    #define SYSCALL_LCHOWN                  198
    #define SYSCALL_GETUID                  199
    #define SYSCALL_GETGID                  200
    #define SYSCALL_GETEUID                 201
    #define SYSCALL_GETEGID                 202
    #define SYSCALL_SETREUID                203
    #define SYSCALL_SETREGID                204
    #define SYSCALL_GETGROUPS               205
    #define SYSCALL_SETGROUPS               206
    #define SYSCALL_FCHOWN                  207
    #define SYSCALL_SETRESUID               208
    #define SYSCALL_GETRESUID               209
    #define SYSCALL_SETRESGID               210
    #define SYSCALL_GETRESGID               211
    #define SYSCALL_CHOWN                   212
    #define SYSCALL_SETUID                  213
    #define SYSCALL_SETGID                  214
    #define SYSCALL_SETFSUID                215
    #define SYSCALL_SETFSGID                216
    #define SYSCALL_PIVOT_ROOT              217
    #define SYSCALL_MINCORE                 218
    #define SYSCALL_MADVISE                 219
    #define SYSCALL_GETDENTS64              220
    #define SYSCALL_READAHEAD               222
    #define SYSCALL_SETXATTR                224
    #define SYSCALL_LSETXATTR               225
    #define SYSCALL_FSETXATTR               226
    #define SYSCALL_GETXATTR                227
    #define SYSCALL_LGETXATTR               228
    #define SYSCALL_FGETXATTR               229
    #define SYSCALL_LISTXATTR               230
    #define SYSCALL_LLISTXATTR              231
    #define SYSCALL_FLISTXATTR              232
    #define SYSCALL_REMOVEXATTR             233
    #define SYSCALL_LREMOVEXATTR            234
    #define SYSCALL_FREMOVEXATTR            235
    #define SYSCALL_GETTID                  236
    #define SYSCALL_TKILL                   237
    #define SYSCALL_FUTEX                   238
    #define SYSCALL_SCHED_SETAFFINITY       239
    #define SYSCALL_SCHED_GETAFFINITY       240
    #define SYSCALL_TGKILL                  241
    #define SYSCALL_IO_SETUP                243
    #define SYSCALL_IO_DESTROY              244
    #define SYSCALL_IO_GETEVENTS            245
    #define SYSCALL_IO_SUBMIT               246
    #define SYSCALL_IO_CANCEL               247
    #define SYSCALL_EXIT_GROUP              248
    #define SYSCALL_EPOLL_CREATE            249
    #define SYSCALL_EPOLL_CTL               250
    #define SYSCALL_EPOLL_WAIT              251
    #define SYSCALL_SET_TID_ADDRESS         252
    #define SYSCALL_FADVISE64               253
    #define SYSCALL_TIMER_CREATE            254
    #define SYSCALL_TIMER_SETTIME           255
    #define SYSCALL_TIMER_GETTIME           256
    #define SYSCALL_TIMER_GETOVERRUN        257
    #define SYSCALL_TIMER_DELETE            258
    #define SYSCALL_CLOCK_SETTIME           259
    #define SYSCALL_CLOCK_GETTIME           260
    #define SYSCALL_CLOCK_GETRES            261
    #define SYSCALL_CLOCK_NANOSLEEP         262
    #define SYSCALL_STATFS64                265
    #define SYSCALL_FSTATFS64               266
    #define SYSCALL_REMAP_FILE_PAGES        267
    #define SYSCALL_MBIND                   268
    #define SYSCALL_GET_MEMPOLICY           269
    #define SYSCALL_SET_MEMPOLICY           270
    #define SYSCALL_MQ_OPEN                 271
    #define SYSCALL_MQ_UNLINK               272
    #define SYSCALL_MQ_TIMEDSEND            273
    #define SYSCALL_MQ_TIMEDRECEIVE         274
    #define SYSCALL_MQ_NOTIFY               275
    #define SYSCALL_MQ_GETSETATTR           276
    #define SYSCALL_KEXEC_LOAD              277
    #define SYSCALL_ADD_KEY                 278
    #define SYSCALL_REQUEST_KEY             279
    #define SYSCALL_KEYCTL                  280
    #define SYSCALL_WAITID                  281
    #define SYSCALL_IOPRIO_SET              282
    #define SYSCALL_IOPRIO_GET              283
    #define SYSCALL_INOTIFY_INIT            284
    #define SYSCALL_INOTIFY_ADD_WATCH       285
    #define SYSCALL_INOTIFY_RM_WATCH        286
    #define SYSCALL_MIGRATE_PAGES           287
    #define SYSCALL_OPENAT                  288
    #define SYSCALL_MKDIRAT                 289
    #define SYSCALL_MKNODAT                 290
    #define SYSCALL_FCHOWNAT                291
    #define SYSCALL_FUTIMESAT               292
    #define SYSCALL_NEWFSTATAT              293
    #define SYSCALL_UNLINKAT                294
    #define SYSCALL_RENAMEAT                295
    #define SYSCALL_LINKAT                  296
    #define SYSCALL_SYMLINKAT               297
    #define SYSCALL_READLINKAT              298
    #define SYSCALL_FCHMODAT                299
    #define SYSCALL_FACCESSAT               300
    #define SYSCALL_PSELECT6                301
    #define SYSCALL_PPOLL                   302
    #define SYSCALL_UNSHARE                 303
    #define SYSCALL_SET_ROBUST_LIST         304
    #define SYSCALL_GET_ROBUST_LIST         305
    #define SYSCALL_SPLICE                  306
    #define SYSCALL_SYNC_FILE_RANGE         307
    #define SYSCALL_TEE                     308
    #define SYSCALL_VMSPLICE                309
    #define SYSCALL_MOVE_PAGES              310
    #define SYSCALL_GETCPU                  311
    #define SYSCALL_EPOLL_PWAIT             312
    #define SYSCALL_UTIMES                  313
    #define SYSCALL_FALLOCATE               314
    #define SYSCALL_UTIMENSAT               315
    #define SYSCALL_SIGNALFD                316
    #define SYSCALL_EVENTFD                 318
    #define SYSCALL_TIMERFD_CREATE          319
    #define SYSCALL_TIMERFD_SETTIME         320
    #define SYSCALL_TIMERFD_GETTIME         321
    #define SYSCALL_SIGNALFD4               322
    #define SYSCALL_EVENTFD2                323
    #define SYSCALL_INOTIFY_INIT1           324
    #define SYSCALL_PIPE2                   325
    #define SYSCALL_DUP3                    326
    #define SYSCALL_EPOLL_CREATE1           327
    #define SYSCALL_PREADV                  328
    #define SYSCALL_PWRITEV                 329
    #define SYSCALL_RT_TGSIGQUEUEINFO       330
    #define SYSCALL_PERF_EVENT_OPEN         331
    #define SYSCALL_FANOTIFY_INIT           332
    #define SYSCALL_FANOTIFY_MARK           333
    #define SYSCALL_PRLIMIT64               334
    #define SYSCALL_NAME_TO_HANDLE_AT       335
    #define SYSCALL_OPEN_BY_HANDLE_AT       336
    #define SYSCALL_CLOCK_ADJTIME           337
    #define SYSCALL_SYNCFS                  338
    #define SYSCALL_SETNS                   339
    #define SYSCALL_PROCESS_VM_READV        340
    #define SYSCALL_PROCESS_VM_WRITEV       341
    #define SYSCALL_S390_RUNTIME_INSTR      342
    #define SYSCALL_KCMP                    343
    #define SYSCALL_FINIT_MODULE            344
    #define SYSCALL_SCHED_SETATTR           345
    #define SYSCALL_SCHED_GETATTR           346
    #define SYSCALL_RENAMEAT2               347
    #define SYSCALL_SECCOMP                 348
    #define SYSCALL_GETRANDOM               349
    #define SYSCALL_MEMFD_CREATE            350
    #define SYSCALL_BPF                     351
    #define SYSCALL_S390_PCI_MMIO_WRITE     352
    #define SYSCALL_S390_PCI_MMIO_READ      353
    #define SYSCALL_EXECVEAT                354
    #define SYSCALL_USERFAULTFD             355
    #define SYSCALL_MEMBARRIER              356
    #define SYSCALL_RECVMMSG                357
    #define SYSCALL_SENDMMSG                358
    #define SYSCALL_SOCKET                  359
    #define SYSCALL_SOCKETPAIR              360
    #define SYSCALL_BIND                    361
    #define SYSCALL_CONNECT                 362
    #define SYSCALL_LISTEN                  363
    #define SYSCALL_ACCEPT4                 364
    #define SYSCALL_GETSOCKOPT              365
    #define SYSCALL_SETSOCKOPT              366
    #define SYSCALL_GETSOCKNAME             367
    #define SYSCALL_GETPEERNAME             368
    #define SYSCALL_SENDTO                  369
    #define SYSCALL_SENDMSG                 370
    #define SYSCALL_RECVFROM                371
    #define SYSCALL_RECVMSG                 372
    #define SYSCALL_SHUTDOWN                373
    #define SYSCALL_MLOCK2                  374
    #define SYSCALL_COPY_FILE_RANGE         375
    #define SYSCALL_PREADV2                 376
    #define SYSCALL_PWRITEV2                377
    #define SYSCALL_S390_GUARDED_STORAGE    378
    #define SYSCALL_STATX                   379
    #define SYSCALL_S390_STHYI              380
    #define SYSCALL_KEXEC_FILE_LOAD         381
    #define SYSCALL_IO_PGETEVENTS           382
    #define SYSCALL_RSEQ                    383
    #define SYSCALL_PKEY_MPROTECT           384
    #define SYSCALL_PKEY_ALLOC              385
    #define SYSCALL_PKEY_FREE               386
    #define SYSCALL_SEMTIMEDOP              392
    #define SYSCALL_SEMGET                  393
    #define SYSCALL_SEMCTL                  394
    #define SYSCALL_SHMGET                  395
    #define SYSCALL_SHMCTL                  396
    #define SYSCALL_SHMAT                   397
    #define SYSCALL_SHMDT                   398
    #define SYSCALL_MSGGET                  399
    #define SYSCALL_MSGSND                  400
    #define SYSCALL_MSGRCV                  401
    #define SYSCALL_MSGCTL                  402
    #define SYSCALL_PIDFD_SEND_SIGNAL       424
    #define SYSCALL_IO_URING_SETUP          425
    #define SYSCALL_IO_URING_ENTER          426
    #define SYSCALL_IO_URING_REGISTER       427
    #define SYSCALL_OPEN_TREE               428
    #define SYSCALL_MOVE_MOUNT              429
    #define SYSCALL_FSOPEN                  430
    #define SYSCALL_FSCONFIG                431
    #define SYSCALL_FSMOUNT                 432
    #define SYSCALL_FSPICK                  433
    #define SYSCALL_PIDFD_OPEN              434
    #define SYSCALL_CLONE3                  435
    #define SYSCALL_CLOSE_RANGE             436
    #define SYSCALL_OPENAT2                 437
    #define SYSCALL_PIDFD_GETFD             438
    #define SYSCALL_FACCESSAT2              439
    #define SYSCALL_PROCESS_MADVISE         440
    #define SYSCALL_EPOLL_PWAIT2            441
    #define SYSCALL_MOUNT_SETATTR           442
    #define SYSCALL_QUOTACTL_FD             443
    #define SYSCALL_LANDLOCK_CREATE_RULESET 444
    #define SYSCALL_LANDLOCK_ADD_RULE       445
    #define SYSCALL_LANDLOCK_RESTRICT_SELF  446
    #define SYSCALL_MEMFD_SECRET            447
    #define SYSCALL_PROCESS_MRELEASE        448
    #define SYSCALL_FUTEX_WAITV             449
    #define SYSCALL_SET_MEMPOLICY_HOME_NODE 450
    #define SYSCALL_CACHESTAT               451
    #define SYSCALL_FCHMODAT2               452
    #define SYSCALL_MAP_SHADOW_STACK        453
    #define SYSCALL_FUTEX_WAKE              454
    #define SYSCALL_FUTEX_WAIT              455
    #define SYSCALL_FUTEX_REQUEUE           456
    #define SYSCALL_STATMOUNT               457
    #define SYSCALL_LISTMOUNT               458
    #define SYSCALL_LSM_GET_SELF_ATTR       459
    #define SYSCALL_LSM_SET_SELF_ATTR       460
    #define SYSCALL_LSM_LIST_MODULES        461
    #define SYSCALL_MSEAL                   462
    #define SYSCALL_SETXATTRAT              463
    #define SYSCALL_GETXATTRAT              464
    #define SYSCALL_LISTXATTRAT             465
    #define SYSCALL_REMOVEXATTRAT           466
    #define SYSCALL_ARCH_PRCTL              UNDEFINED_SYSCALL
    #define SYSCALL_ACCEPT                  UNDEFINED_SYSCALL
    #define SYSCALL_READDIR                 UNDEFINED_SYSCALL
    #define SYSCALL_LOOKUP_DCOOKIE          UNDEFINED_SYSCALL
    #define SYSCALL_IDLE                    UNDEFINED_SYSCALL
    #define SYSCALL_CREATE_MODULE           UNDEFINED_SYSCALL
    #define SYSCALL_GET_KERNEL_SYMS         UNDEFINED_SYSCALL
    #define SYSCALL_BDFLUSH                 UNDEFINED_SYSCALL
    #define SYSCALL_AFS_SYSCALL             UNDEFINED_SYSCALL
    #define SYSCALL__SYSCTL                 UNDEFINED_SYSCALL
    #define SYSCALL_QUERY_MODULE            UNDEFINED_SYSCALL
    #define SYSCALL_NFSSERVCTL              UNDEFINED_SYSCALL
    #define SYSCALL_GETPMSG                 UNDEFINED_SYSCALL
    #define SYSCALL_PUTPMSG                 UNDEFINED_SYSCALL
    #define SYSCALL_TIMERFD                 UNDEFINED_SYSCALL
#endif

#if SYSCALL_SOCKETCALL != UNDEFINED_SYSCALL
#define ARCH_HAS_SOCKETCALL 1
#endif

statfunc bool has_syscall_fd_arg(uint syscall_id)
{
    // Only syscalls with one fd argument so far
    switch (syscall_id) {
        case SYSCALL_READ:
        case SYSCALL_WRITE:
        case SYSCALL_CLOSE:
        case SYSCALL_FSTAT:
        case SYSCALL_LSEEK:
        case SYSCALL_MMAP:
        case SYSCALL_IOCTL:
        case SYSCALL_PREAD64:
        case SYSCALL_PWRITE64:
        case SYSCALL_READV:
        case SYSCALL_WRITEV:
        case SYSCALL_DUP:
        case SYSCALL_CONNECT:
        case SYSCALL_ACCEPT:
        case SYSCALL_SENDTO:
        case SYSCALL_RECVFROM:
        case SYSCALL_SENDMSG:
        case SYSCALL_RECVMSG:
        case SYSCALL_SHUTDOWN:
        case SYSCALL_BIND:
        case SYSCALL_LISTEN:
        case SYSCALL_GETSOCKNAME:
        case SYSCALL_GETPEERNAME:
        case SYSCALL_SETSOCKOPT:
        case SYSCALL_GETSOCKOPT:
        case SYSCALL_FCNTL:
        case SYSCALL_FLOCK:
        case SYSCALL_FSYNC:
        case SYSCALL_FDATASYNC:
        case SYSCALL_FTRUNCATE:
        case SYSCALL_FCHDIR:
        case SYSCALL_FCHMOD:
        case SYSCALL_FCHOWN:
        case SYSCALL_FSTATFS:
        case SYSCALL_READAHEAD:
        case SYSCALL_FSETXATTR:
        case SYSCALL_FGETXATTR:
        case SYSCALL_FLISTXATTR:
        case SYSCALL_FREMOVEXATTR:
        case SYSCALL_GETDENTS64:
        case SYSCALL_FADVISE64:
        case SYSCALL_INOTIFY_ADD_WATCH:
        case SYSCALL_INOTIFY_RM_WATCH:
        case SYSCALL_OPENAT:
        case SYSCALL_MKDIRAT:
        case SYSCALL_MKNODAT:
        case SYSCALL_FCHOWNAT:
        case SYSCALL_UNLINKAT:
        case SYSCALL_SYMLINKAT:
        case SYSCALL_READLINKAT:
        case SYSCALL_FCHMODAT:
        case SYSCALL_FACCESSAT:
        case SYSCALL_SYNC_FILE_RANGE:
        case SYSCALL_VMSPLICE:
        case SYSCALL_UTIMENSAT:
        case SYSCALL_FALLOCATE:
        case SYSCALL_TIMERFD_SETTIME:
        case SYSCALL_TIMERFD_GETTIME:
        case SYSCALL_ACCEPT4:
        case SYSCALL_SIGNALFD4:
        case SYSCALL_PREADV:
        case SYSCALL_PWRITEV:
        case SYSCALL_PERF_EVENT_OPEN:
        case SYSCALL_RECVMMSG:
        case SYSCALL_NAME_TO_HANDLE_AT:
        case SYSCALL_OPEN_BY_HANDLE_AT:
        case SYSCALL_SYNCFS:
        case SYSCALL_SENDMMSG:
        case SYSCALL_SETNS:
        case SYSCALL_FINIT_MODULE:
        case SYSCALL_EXECVEAT:
        case SYSCALL_PREADV2:
        case SYSCALL_PWRITEV2:
        case SYSCALL_STATX:
        case SYSCALL_PIDFD_SEND_SIGNAL:
        case SYSCALL_IO_URING_ENTER:
        case SYSCALL_IO_URING_REGISTER:
        case SYSCALL_OPEN_TREE:
        case SYSCALL_FSCONFIG:
        case SYSCALL_FSMOUNT:
        case SYSCALL_FSPICK:
        case SYSCALL_OPENAT2:
        case SYSCALL_FACCESSAT2:
        case SYSCALL_PROCESS_MADVISE:
        case SYSCALL_EPOLL_PWAIT2:
        case SYSCALL_MOUNT_SETATTR:
        case SYSCALL_QUOTACTL_FD:
        case SYSCALL_LANDLOCK_ADD_RULE:
        case SYSCALL_LANDLOCK_RESTRICT_SELF:
        case SYSCALL_PROCESS_MRELEASE:
#if !defined(bpf_target_arm64)
        case SYSCALL_GETDENTS:
        case SYSCALL_EPOLL_WAIT:
        case SYSCALL_FUTIMESAT:
        case SYSCALL_NEWFSTATAT:
        case SYSCALL_EPOLL_PWAIT:
        case SYSCALL_SIGNALFD:
#endif
            return true;
    }

    return false;
}

statfunc uint get_syscall_fd_num_from_arg(uint syscall_id, args_t *args)
{
    switch (syscall_id) {
        case SYSCALL_SYMLINKAT:
            return args->args[1];
        case SYSCALL_PERF_EVENT_OPEN:
            return args->args[3];
        case SYSCALL_MMAP:
            return args->args[4];
    }

    return args->args[0];
}

#endif
