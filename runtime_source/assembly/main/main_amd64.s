TEXT ·main(SB),$16-0
    MOVQ ·helloWorld+0(SB),AX;MOVQ AX,0(SP)
    MOVQ ·helloWorld+8(SB),BX;MOVQ BX,8(SP)
    CALL ·print(SB)
    RET
