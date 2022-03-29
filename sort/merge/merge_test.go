package merge

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{5}
	MergeSort(list, 0, len(list))
	fmt.Println(list)
	list1 := []int{5, 9}
	MergeSort(list1, 0, len(list1))
	fmt.Println(list1)
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}

func TestMergeSort2(t *testing.T) {
	list := []int{5}
	MergeSort2(list, 0, len(list))
	fmt.Println(list)
	list1 := []int{5, 9}
	MergeSort2(list1, 0, len(list1))
	fmt.Println(list1)
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort2(list2, 0, len(list2))
	fmt.Println(list2)
}