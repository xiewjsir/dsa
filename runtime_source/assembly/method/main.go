package main

import "fmt"

type MyInt int
func (v MyInt) Twice() int
func MyIntTwice(v MyInt) int

func main() {
	var a MyInt = 2
	fmt.Println(MyIntTwice(a))
	fmt.Println(a.Twice())
}

//func (v MyInt) Twice() int {
//	return int(v)*2
//}
//
//func MyIntTwice(v MyInt) int {
//	return int(v)*2
//}


