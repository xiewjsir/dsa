package main

import "unsafe"

/*
package main

func NewTwiceFunClosure(x int) func() int {
    return func() int {
        x *= 2
        return x
    }
}

func main() {
    fnTwice := NewTwiceFunClosure(1)

    println(fnTwice()) // 1*2 => 2
    println(fnTwice()) // 2*2 => 4
    println(fnTwice()) // 4*2 => 8
}
 */

type FunTwiceClosure struct {
	F uintptr
	X int
}

func ptrToFunc(p unsafe.Pointer) func() int

func asmFunTwiceClosureAddr() uintptr
func asmFunTwiceClosureBody() int

func NewTwiceFunClosure(x int) func() int {
	var p = &FunTwiceClosure{
		F: asmFunTwiceClosureAddr(),
		X: x,
	}
	return ptrToFunc(unsafe.Pointer(p))
}

func main() {
	fnTwice := NewTwiceFunClosure(1)
	
	println(fnTwice()) // 1*2 => 2
	println(fnTwice()) // 2*2 => 4
	println(fnTwice()) // 4*2 => 8
}