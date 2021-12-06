package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)
	fmt.Println(list)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
		BubbleSort(list)
	}
}
