package main

import (
	"fmt"
	"log"
	"time"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural3() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			log.Println("aaaa", i)
			ch <- i
			log.Println("bbbb", i)
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter3(in <-chan int, prime int) chan int {
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
	ch := GenerateNatural3() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter3(ch, prime) // 基于新素数构造的过滤器
	}

	time.Sleep(10 * time.Second)
}
