package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GC()
	fmt.Println("hello world")
}