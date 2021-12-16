package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	values := []int64{3, 6, 8, 20, 9, 2, 6, 8, 9, 3, 5, 40, 7, 9, 13, 6, 8}
	// 初始化二叉查找树并添加元素
	tree := NewTree()
	for _, v := range values {
		tree.Add(v)
	}
	// 找到最大值或最小值的节点
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())
	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}
	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}
	// 删除存在的9后，再查找9
	tree.Delete(9)
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}
	// 中序遍历，实现排序
	tree.MidOrder()
	fmt.Println("")
}
