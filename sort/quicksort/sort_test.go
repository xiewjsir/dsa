package quicksort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	//list := []int{5}
	//QuickSort(list, 0, len(list)-1)
	//fmt.Println(list)
	//list1 := []int{5, 9}
	//QuickSort(list1, 0, len(list1)-1)
	//fmt.Println(list1)
	//list2 := []int{5, 9, 1}
	//QuickSort(list2, 0, len(list2)-1)
	//fmt.Println(list2)
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}