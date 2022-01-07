package main

import (
	"fmt"
	"unsafe"
)

func main() {
	
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

