package list

import (
	"fmt"
	"testing"
)

func TestDoubleList(t *testing.T) {
	list := new(doubleList)
	// 在列表头部插入新元素
	list.AddNodeFromHead(0, "I")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")

	list.AddNodeFromTail(list.Len()-1, "begin second")
	list.AddNodeFromHead(list.Len()-1, "end second")
	// 正常遍历，比较慢，因为内部会遍历拿到值返回
	for i := 0; i < list.Len(); i++ {
		// 从头部开始索引
		node := list.IndexFromHead(i)
		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsNil() {
			fmt.Println(node.GetValue())
		}
	}
	fmt.Println("----------")
	// 正常遍历，特别快，因为直接拿到的链表节点
	// 先取出第一个元素
	first := list.First()
	for !first.IsNil() {
		// 如果非空就一直遍历
		fmt.Println(first.GetValue())
		// 接着下一个节点
		first = first.GetNext()
	}
	fmt.Println("----------")
	// 元素一个个 POP 出来
	for {
		node := list.PopFromHead(0)
		if node.IsNil() {
			// 没有元素了，直接返回
			break
		}
		fmt.Println(node.GetValue())
	}
	fmt.Println("----------")
	fmt.Println("len", list.Len())
}

