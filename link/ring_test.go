package link

import (
	"fmt"
	"testing"
)
func TestRingCreate(t *testing.T) {
	// 第一个节点
	r := NewRing(5)
	
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		
		// 移到下一个节点
		node = node.Prev()
		
		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func TestRingLink(t *testing.T) {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 0})
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Prev()

		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}


func TestRingUnLink(t *testing.T) {
	// 第一个节点
	r := &Ring{Value: -1}
	
	// 链接新的五个节点
	r.Link(&Ring{Value: 0})
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		
		// 移到下一个节点
		node = node.Prev()
		
		//  如果节点回到了起点，结束
		if node == r {
			break
		}
	}
	
	fmt.Println("len=",r.Len())
	
	r.Unlink(3)
	for {
		// 打印节点值
		fmt.Println(node.Value)
		
		// 移到下一个节点
		node = node.Prev()
		
		//  如果节点回到了起点，结束
		if node == r {
			break
		}
	}
	fmt.Println("len=",r.Len())
}