package heapsort

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	list := []int{5, 9, 6, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	
	// 构建最大堆
	h := NewHeap(list)
	for _, v := range list {
		h.Push(v)
	}
	
	// 将堆元素移除
	for range list {
		h.Pop()
	}
	
	// 打印排序后的值
	fmt.Println(list)
}

func TestHeapSort(t *testing.T) {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	
	HeapSort(list)
	
	// 打印排序后的值
	fmt.Println(list)
}
