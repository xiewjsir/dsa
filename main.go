package main

import (
	"fmt"
)

func counting(c chan<- int){
	for i := 0; i < 10; i++{
		c <- i
	}
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus{
		fmt.Println("count : ", count)
	}
}