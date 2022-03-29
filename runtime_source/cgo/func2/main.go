package main

//int sum(int a, int b);
import "C"

//export sum
func sum(a, b C.int) C.int {
	return a + b
}

//go build -buildmode=c-archive -o sum.a main.go
func main() {}