.intel_syntax noprefix

inc rax
inc eax
inc ax
inc al

add rbx, rax
add rax, rbx

mov eax, [ecx]

int 3
int 100

rep movsd
