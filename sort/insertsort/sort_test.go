package insertsort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	//list := []int{5}
	//InsertSort(list)
	//fmt.Println(list)
	list1 := []int{9, 5}
	InsertSort(list1)
	fmt.Println(list1)
	//list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//InsertSort(list2)
	//fmt.Println(list2)
}