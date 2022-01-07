#include "textflag.h"

TEXT ·If(SB), NOSPLIT, $0-32
    MOVQ ok+8*0(FP), CX // ok
    MOVQ a+8*1(FP), AX  // a
    MOVQ b+8*2(FP), BX  // b

    CMPQ CX, $0         // test ok
    JZ   L              // if ok == 0, goto L
    MOVQ AX, ret+24(FP) // return a
    RET

L:
    MOVQ BX, ret+24(FP) // return b
    RET
