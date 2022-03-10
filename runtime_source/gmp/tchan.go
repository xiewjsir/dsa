package main

import "fmt"

func main() {
	fmt.Println("broker point 1")
	ch := make(chan int)
	wait := make(chan int)
	go func() {
		for i := range ch {
			fmt.Println("broker point 2", i)
		}
	}()

	ch <- 2

	fmt.Println("broker point 3")
	<-wait
}
