package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string,64*6.5+1)
	m[1] = "ffff"
	fmt.Println(m[1])
}