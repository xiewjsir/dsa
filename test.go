package main

import (
	"fmt"
)

const PtrSize = 4 << (^uintptr(0) >> 63)

func main() {
	fmt.Println(PtrSize, ^PtrSize)

	fmt.Println(fmt.Sprintf("%b", ^uintptr(0)))
}
