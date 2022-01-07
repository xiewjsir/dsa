package main

import (
	"fmt"
	"unsafe"
)

func main() {
	Foo()
}

func Foo() {
	var local [1]struct{
		a bool
		b int16
		c []byte
	}
	var SP = &local[0]
	
	fmt.Println(&local[0].a,fmt.Sprintf("%p",SP),fmt.Sprintf("%x",uintptr(unsafe.Pointer(&SP))),unsafe.Pointer(uintptr(unsafe.Pointer(&SP))+unsafe.Offsetof(local[0].a)))
	
	//_ = -(unsafe.Sizeof(local)-unsafe.Offsetof(local[0].a)) + uintptr(unsafe.Pointer(&SP)) // a
	//_ = -(unsafe.Sizeof(local)-unsafe.Offsetof(local[0].b)) + uintptr(unsafe.Pointer(&SP)) // b
	//_ = -(unsafe.Sizeof(local)-unsafe.Offsetof(local[0].c)) + uintptr(unsafe.Pointer(&SP)) // c
}
