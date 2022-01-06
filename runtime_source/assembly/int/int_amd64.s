#include "textflag.h"

GLOBL ·int32Value(SB),NOPTR,$4
DATA ·int32Value+0(SB)/1,$0x01  // 第0字节
DATA ·int32Value+1(SB)/1,$0x02  // 第1字节
DATA ·int32Value+2(SB)/2,$0x03  // 第3-4字节

GLOBL ·uint32Value(SB),NOPTR,$4
DATA ·uint32Value(SB)/4,$0x01020304 // 第1-4字节

