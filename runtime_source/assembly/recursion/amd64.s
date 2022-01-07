// func sum(n int) (result int)
TEXT ·sum(SB), $16-16
    MOVQ n+0(FP), AX       // n
    MOVQ result+8(FP), BX  // result

    CMPQ AX, $0            // test n - 0
    JG   L_STEP_TO_END     // if > 0: goto L_STEP_TO_END
    JMP  L_END             // goto L_STEP_TO_END

L_STEP_TO_END:
    SUBQ $1, AX            // AX -= 1
    MOVQ AX, 0(SP)         // arg: n-1
    CALL ·sum(SB)          // call sum(n-1)
    MOVQ 8(SP), BX         // BX = sum(n-1)

    MOVQ n+0(FP), AX       // AX = n
    ADDQ AX, BX            // BX += AX
    MOVQ BX, result+8(FP)  // return BX
    RET

L_END:
    MOVQ $0, result+8(FP) // return 0
    RET
