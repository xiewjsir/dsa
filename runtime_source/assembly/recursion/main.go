package main

import "fmt"

func sum(n int) (result int)

/*
func sum2(n int) int {
	if n > 0 {
		return n + sum(n-1)
	} else {
		return 0
	}
}

func sum(n int) (result int) {
	var AX = n
	var BX int

	if n > 0 {
		goto L_STEP_TO_END
	}
	goto L_END

L_STEP_TO_END:
	AX -= 1
	BX = sum(AX)

	AX = n // 调用函数后, AX重新恢复为n
	BX += AX

	return BX

L_END:
	return 0
}
*/

func main() {
	fmt.Println(sum(50))
}
