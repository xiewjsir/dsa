#include "textflag.h"
// func (v MyInt) Twice() int
TEXT ·MyInt·Twice(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX   // v
    ADDQ AX, AX        // AX *= 2
    MOVQ AX, ret+8(FP) // return v
    RET

TEXT ·MyIntTwice(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX   // v
    ADDQ AX, AX        // AX *= 2
    MOVQ AX, ret+8(FP) // return v
    RET
