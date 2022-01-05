package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// GenerateNatural2 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// PrimeFilter2 管道过滤器: 删除能被素数整除的数
func PrimeFilter2(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural2() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter2(ch, prime) // 基于新素数构造的过滤器
	}
	
	http.ListenAndServe(":60000",nil)
}
