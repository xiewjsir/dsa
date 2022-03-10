package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
		8: 21,
	}

	for key, val := range tests {
		if n := Fibonacci(key); n != val {
			t.Errorf(fmt.Sprintf("fibonacci %d expected be %d, but %d got", key, val, n))
		}
	}
}

func TestFibonacci2(t *testing.T) {
	tests := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
		8: 21,
	}

	for key, val := range tests {
		if n := Fibonacci2(key, 1, 1); n != val {
			t.Errorf(fmt.Sprintf("fibonacci2 %d expected be %d, but %d got", key, val, n))
		}
	}
}

func TestFibonacci3(t *testing.T) {
	tests := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
		8: 21,
	}

	for key, val := range tests {
		if n := Fibonacci3(key); n != val {
			t.Errorf(fmt.Sprintf("fibonacci3 %d expected be %d, but %d got", key, val, n))
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(15)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci2(15, 1, 1)
	}
}

func BenchmarkFibonacci3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci3(15)
	}
}
