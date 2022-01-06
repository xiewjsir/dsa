#include "textflag.h"
TEXT ·neg(SB), NOSPLIT,$0
    MOVQ     x+0(FP), AX
    NEGQ     AX
    MOVQ     AX, ret+8(FP)
    RET

TEXT ·negBak(SB),NOPTR,$0-16
    MOVQ +8(SP),AX
    NEGQ AX
    MOVQ AX,ret+16(SP)
    RET
