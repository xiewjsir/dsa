TEXT ·main(SB), $16-0
    // var temp int

    // 将新的值写入a对应内存
    MOVQ $10, AX        // AX = 10
    MOVQ AX, temp-8(SP) // temp = AX

    // 以a为参数调用函数
    runtime·print(SB)

    // 函数调用后, AX 可能被污染, 需要重新加载
    MOVQ temp-8*1(SP), AX // AX = temp

    // 计算b值, 不需要写入内存
    MOVQ AX, BX        // BX = AX  // b = a
    ADDQ BX, BX        // BX += BX // b += a
    IMULQ AX, BX       // BX *= AX // b *= a

    runtime·print(SB)

