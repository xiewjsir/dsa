package main

import (
	"fmt"
	"sync"
)


type Person struct {
	mux sync.Mutex
	rwmux sync.RWMutex
}

func Reduce(p1 *Person) {
	fmt.Println("step...", )
	p1.mux.Lock()
	fmt.Println(p1)
	defer p1.mux.Unlock()
	fmt.Println("over...")
}

func main() {
	var p *Person = &Person{}
	p.mux.Lock()
	go Reduce(p)
	p.mux.Unlock()
	fmt.Println(111)
	
	p.rwmux.RLock()
	
	for {}
	
}
