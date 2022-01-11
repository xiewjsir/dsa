package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	
	//test1()
	test2()

}

func test1() {
	// uintptr 指针指向地址的整型数字表示
	// *uintptr 指针指向的地址
	// **uintptr 指针的地址
	
	var nu int32 = 10
	var data *int
	*(**uintptr)(unsafe.Pointer(&data)) = (*uintptr)(unsafe.Pointer(&nu))
	
	fmt.Println(&data,data,&nu,**&data)
	fmt.Println((uintptr)(unsafe.Pointer(&data)),fmt.Sprintf("%0X",(uintptr)(unsafe.Pointer(&data))))
	fmt.Println(*(*uintptr)(unsafe.Pointer(&data)),fmt.Sprintf("%0X",*(*uintptr)(unsafe.Pointer(&data))))
	fmt.Println((**uintptr)(unsafe.Pointer(&data)),*(**uintptr)(unsafe.Pointer(&data)),**(**uintptr)(unsafe.Pointer(&data)))
}

func test2() {
	var x struct {
		a bool  // bool 占一个字节,b占两个字节,需对齐,空一个字节
		b int16 // 两个字节
		c []int // 3*8个字节,c.data,c.len,c.cap各占8个字节
	}

	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	
	//m := x{}
	x.a = true
	x.b = 132
	x.c = []int{1,2,2,3,4}
	log.Println("Sizeof:")
	log.Println(unsafe.Sizeof(x.a))
	log.Println(unsafe.Sizeof(x.b))
	log.Println(unsafe.Sizeof(x.c))
	log.Println(unsafe.Sizeof(x))
	log.Println("Offsetof:")
	log.Println(unsafe.Offsetof(x.a))
	log.Println(unsafe.Offsetof(x.b))
	log.Println(unsafe.Offsetof(x.c))
	log.Println("ttt:")
	type SizeOfE struct {
		A byte  // 1
		C byte  // 1 //调换一下B C的顺序，择Sizeof 整个结构体的大小由24变为了16
		B int64 // 8
	}
	var tt SizeOfE
	log.Println(unsafe.Sizeof(SizeOfE{}))
	log.Println(unsafe.Sizeof(tt.A))
	log.Println(unsafe.Sizeof(tt.B))
	log.Println(unsafe.Sizeof(tt.C))
	log.Println(unsafe.Sizeof(tt))
	log.Println("AlignOf:")
	log.Println(unsafe.Alignof(tt.A))
	log.Println(unsafe.Alignof(tt.B))
	log.Println(unsafe.Alignof(tt.C))
	log.Println(unsafe.Alignof(tt))
}