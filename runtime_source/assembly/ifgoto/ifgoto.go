package main

import "fmt"

func If(ok bool, a, b int) int

func main() {
	fmt.Println(If(true,1,2))
}
/*
func If(ok bool, a, b int) int {
	if ok { return a } else { return b }
}

func If(ok int, a, b int) int {
	if ok == 0 { goto L }
	return a
L:
	return b
}
*/
