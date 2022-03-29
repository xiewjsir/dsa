package main

import (
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f,_ := os.OpenFile("trace.out",os.O_APPEND|os.O_CREATE|os.O_RDWR,os.ModePerm)
	trace.Start(f)
	defer trace.Stop()
	
	var total int
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				total += readNumber()
			}
			wg.Done()
		}()
	}
	
	wg.Wait()
}

//go:noinline
func readNumber() int {
	return rand.Intn(10)
}
