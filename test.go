package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int8
	num2 int32
	num3 int16
}

func main() {
	fmt.Println(unsafe.Alignof(Args{}),unsafe.Sizeof(Args{}))
	fmt.Println(unsafe.Alignof(Flag{}),unsafe.Sizeof(Flag{}))
}