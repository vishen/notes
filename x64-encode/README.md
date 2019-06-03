
## OS and Arch

Executables do depend on both the OS and the CPU:

Instruction Set: The binary instructions in the executable are decoded by the CPU according to some instruction set. Most consumer CPUs support the x86 (“32bit”) and/or AMD64 (“64bit”) instruction sets. A program can be compiled for either of these instruction sets, but not both. There are extensions to these instruction sets; support for these can be queried at runtime. Such extensions offer SIMD support, for example. Optimizing compilers might try to take advantage of these extensions if they are present, but usually also offer a code path that works without any extensions.

Binary Format: The executable has to conform to a certain binary format, which allows the operating system to correctly load, initialize, and start the program. Windows mainly uses the Portable Executable format, while Linux uses ELF.

System APIs: The program may be using libraries, which have to be present on the executing system. If a program uses functions from Windows APIs, it can't be run on Linux. In the Unix world, the central operating system APIs have been standardized to POSIX: a program using only the POSIX functions will be able to run on any conformant Unix system, such as Mac OS X and Solaris.

## Bin vs ELF

### Bin

A binary file is a pure binary file wit no memory fix-ups or
relocations, more that likely it has explicit instructions to be 
loaded at a specific memory address.

### ELF

Executable Linkable Format which consists of symbol look-ups and relocatable table, that is,
it can be loaded at any memory address by the kernel. All symbols used are
adjusted to the offset from that memory address where is was loaded into. 
Usually ELF files have a number of sections, such as; "data", "text", "bss".
It is within those sections where the run-time can calculate where to adjust the symbols
memory reference dynamically at run-time.

## x86-64 encoding

An x86-64 instruction may be at most 15 bytes in length. It consists of the following
components in the given order, where the prefixes are at the least-significant (lowest)
address in memory:

- Legacy prefixes (1-4 bytes, optional)
- Opcode with prefixes (1-4 bytes, required)
- ModR/M (1 byte, if required)
- SIB (1 byte, if required)
- Displacement(1, 2, 4 or 8 bytes, if required)
- Immediate (1, 2, 4 or 8 bytes, if required)

### Registers

X.Reg -> registers 8-bit, 16-bit, 32-bit, 64-bit
- 0.000 (0)  -> AL,   AX,   EAX,  RAX  | Accumulator
- 0.001 (1)  -> CL,   CX,   ECX,  RCX  | Counter
- 0.010 (2)  -> DL,   DX,   EDX,  RDX  | Data
- 0.011 (3)  -> BL,   BX,   EBX,  RBX  | Base
- 0.100 (4)  -> AH,   SP,   ESP,  RSP  | Stack Pointer
- 0.101 (5)  -> CH,   BP,   EBP,  RBP  | Stack Base Pointer
- 0.110 (6)  -> DH,   SI,   ESI,  RSI  | Source
- 0.111 (7)  -> BH,   DI,   EDI,  RDI  | Destination
- 1.000 (8)  -> R8L,  R8W,  R8D,  R8
- 1.001 (9)  -> R9L,  R9W,  R9D,  R9
- 1.010 (10) -> R10L, R10W, R10D, R10,
- 1.011 (11) -> R11L, R11W, R11D, R11
- 1.100 (12) -> R12L, R12W, R12D, R12
- 1.101 (13) -> R13L, R13W, R13D, R13
- 1.110 (14) -> R14L, R14W, R14D, R14
- 1.111 (15) -> R15L, R15W, R15D, R15

### REX

REX prefixes are instruction-prefix bytes used in 64-bit mode. A REX prefix
is required only if an instruction references one of the extended registers
or uses 64-bit operand. 

If an REX prefix is used when it has no meaning, it is ignored.

The REX prefix is only available in long mode.

### Legacy Prefixes

Each instructions can have up to four prefixes. Order does not matter.
When there are two or more prefixes from a single group, the behaviour
is undefined.

#### Prefix group 1

- 0xF0: LOCK prefix
- 0xF2: REPNE/REPNZ prefix
- 0xF3: REP or REPE/REPZ prefix

##### Lock Prefix

With the LOCK prefix, certain read-modify-write instructions are executed atomically.

- ADC
- ADD
- AND
- BTC
- BTR
- BTS
- CMPXCHG
- CMPXCHG8B
- CMPXCHG16B
- DEC
- INC
- NEG
- NOT
- OR
- SBB
- SUB
- XADD
- XCHG
- XOR

##### REPNE/REPNZ, REP and REPE/REPZ prefixes

The repeated prefixes causes string handling instructions to be repeated.

#### Prefix group 2

- 0x2E: CS segment override
- 0x36: SS segment override
- 0x3E: DS segment override
- 0x26: ES segment override
- 0x64: FS segment override
- 0x65: GS segment override
- 0x2E: Branch not taken
- 0x3E: Branch taken

##### CS, SS, DS, ES, FS and GS segment override prefixes

Segment overrides are used with instructions that reference non-stack memory.

Each thread will have a GS that corresponds to an area in memory.

Used to be essential before 32-bit addresssing.

Nowadays, on x64, they are mostly all 0.

##### Branch taken/not taken prefixes

Branch hints may be used to lessen the impact of branch misprediction somewhat. The
`branch taken` hint is a string hint, while the `branch not taken` is a weak hint. Possibly
only useful for intel.

#### Prefix group 3

- 0x66: Operand-size override prefix

#### Prefix group 4

- 0x67: Address-size override prefix

##### Operand-size and address-size override prefix

THe default operand- and address-size can be overwritten with these prefix.

## Exploratory

```
rep movsd ; f3 a5
```

```
inc rax ; 48 ff c0
inc eax ; ff c0
inc ax  ; 66 ff c0
inc al  ; fe c0

add rax, rbx ; 48 01 d8
```

```
mov eax, [ecx] ; 67 8b 01
```

```
int 3   ; cc
int 100 ; cc 64
```

- REX/64-bit mode: 0x48
- REP: 0xf3
- int 3: 0xcc
- operand size override: 0x66
- address size override: 0x67

## Resources

- https://www.youtube.com/watch?v=N9B2KeAWXgE&t=1416s
- https://www-user.tu-chemnitz.de/~heha/viewchm.php/hs/x86.chm/x64.htm
- https://www.systutorials.com/72643/beginners-guide-x86-64-instruction-encoding/
- https://stackoverflow.com/questions/2427011/what-is-the-difference-between-elf-files-and-bin-files
- https://softwareengineering.stackexchange.com/questions/251250/why-do-executables-depend-on-the-os-but-not-on-the-cpu
- https://wiki.osdev.org/X86-64_Instruction_Encoding
