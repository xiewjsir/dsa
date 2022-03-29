package improvemerge

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{5}
	MergeSort(list, len(list))
	fmt.Println(list)
	list1 := []int{5, 9}
	MergeSort(list1, len(list1))
	fmt.Println(list1)
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, len(list2))
	fmt.Println(list2)
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 45, 67, 2, 5, 24, 56, 34, 24, 56, 2, 2, 21, 4, 1, 4, 7, 9}
	MergeSort(list3, len(list3))
	fmt.Println(list3)
}