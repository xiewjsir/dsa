#include "textflag.h"

GLOBL ·boolValue(SB),NOPTR,$1   // 未初始化

GLOBL ·trueValue(SB),NOPTR,$1   // var trueValue = true
DATA ·trueValue(SB)/1,$1  // 非 0 均为 true

GLOBL ·falseValue(SB),NOPTR,$1  // var falseValue = true
DATA ·falseValue(SB)/1,$0
