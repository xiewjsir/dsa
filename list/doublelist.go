package list

import (
	"sync"
)

type doubleList struct {
	head *listNode
	tail *listNode
	len  int
	lock sync.Mutex
}

// Len 返回列表长度
func (list *doubleList) Len() int {
	return list.len
}

// AddNodeFromHead 从头部开始，添加节点到第N+1个元素之前，
// N=0表示添加到第一个元素之前，表示新节点成为新的头部，
// N=1表示添加到第二个元素之前，以此类推
func (list *doubleList) AddNodeFromHead(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()
	
	if n != 0 && n >= list.len {
		panic("index out")
	}
	
	newNode := new(listNode)
	newNode.value = v
	
	node := list.head
	for ; n > 0; n-- {
		node = node.next
	}
	
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		prev := node.prev
		if prev.IsNil() {
			node.prev = newNode
			newNode.next = node
			list.head = newNode
		} else {
			prev.next = newNode
			newNode.prev = prev
			newNode.next = node
			node.prev = newNode
		}
	}
	
	list.len++
	
	return
}

func (list *doubleList) AddNodeFromTail(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()
	
	if n != 0 && n >= list.len {
		panic("index out")
	}
	
	newNode := new(listNode)
	newNode.value = v

	node := list.tail
	for ; n > 0; n-- {
		node = node.prev
	}
	
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		next := node.next
		if next.IsNil() {
			newNode.prev = node
			node.next = newNode
			list.tail = newNode
		} else {
			next.prev = newNode
			newNode.next = next
			
			node.next = newNode
			newNode.prev = node
		}
	}
	
	list.len++
	return
}

// IndexFromHead 从头部开始往后找，获取第N+1个位置的节点，索引从0开始。
func (list *doubleList) IndexFromHead(n int) *listNode {
	if n > list.len {
		panic("index out")
	}
	
	node := list.head
	for ; n > 0; n-- {
		node = node.next
	}
	
	return node
}

// IndexFromTail 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
func (list *doubleList) IndexFromTail(n int) *listNode {
	if n > list.len {
		panic("index out")
	}
	
	node := list.tail
	for ; n > 0; n-- {
		node = node.prev
	}
	
	return node
}

// PopFromHead 从头部开始往后找, 获取第N+1个位置的节点, 并移除返回
func (list *doubleList) PopFromHead(n int) *listNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()
	
	if n > list.len {
		panic("index out")
	}
	
	node := list.head
	for ; n > 0; n-- {
		node = node.next
	}
	
	prev := node.prev
	next := node.next
	if prev.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if prev.IsNil() {
		list.head = next
		next.prev = nil
	} else if next.IsNil() {
		list.tail = prev
		prev.next = nil
	} else {
		prev.next = next
		next.prev = prev
	}
	
	list.len--
	return node
}


// PopFromTail 从尾部开始往前找, 获取第N+1个位置的节点, 并移除返回
func (list *doubleList) PopFromTail(n int) *listNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()
	
	if n > list.len {
		panic("index out")
	}
	
	node := list.head
	for ; n > 0; n-- {
		node = node.prev
	}
	
	prev := node.prev
	next := node.next
	if prev.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if prev.IsNil() {
		list.head = next
		next.prev = nil
	} else if next.IsNil() {
		list.tail = prev
		prev.next = nil
	} else {
		prev.next = next
		next.prev = prev
	}
	
	list.len--
	return node
}

// First 返回列表链表头结点
func (list *doubleList) First() *listNode {
	return list.head
}
// Last 返回列表链表尾结点
func (list *doubleList) Last() *listNode {
	return list.tail
}