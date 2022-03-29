package main

import (
	"fmt"
)

func g2(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)
	
	go g2(100, ch)
	
	fmt.Println(<-ch)
}