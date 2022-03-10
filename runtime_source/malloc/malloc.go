package main

import (
	"fmt"
	"unsafe"
)

type smallobj struct {
	arr [1 << 10]byte
}

type largeobj struct {
	arr [1 << 26]byte
}

func tiny() {
	y := 100000
	fmt.Println(y) // call $runtime.convT64
}

func large() {
	large := largeobj{}
	println(&large)
	println(unsafe.Sizeof(large) / 1024 / 1024)
}

func small() {
	small := smallobj{}
	print(&small)
}

func main() {
	//var a int64 = 255 //runtime.convT64
	// var s = "" //runtime.convTstring
	tiny()
	//small()
	//large()
}
