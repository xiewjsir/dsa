package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T)  {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)
}

func TestSelectGoodsSort(t *testing.T)  {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectGoodSort(list)
	fmt.Println(list)
}