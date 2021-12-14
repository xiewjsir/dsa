package stack

import (
	"sync"
)

type linkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Value    string
	NextNode *LinkNode
}

func (l *linkStack) Push(v string) *linkStack {
	l.lock.Lock()
	defer l.lock.Unlock()
	
	node := new(LinkNode)
	node.Value = v
	
	if l.root != nil {
		prevNode := l.root
		node.NextNode = prevNode
	}
	
	l.root = node
	
	l.size++
	return l
}

func (l *linkStack) Pop() string {
	l.lock.Lock()
	defer l.lock.Unlock()
	
	if l.size == 0 {
		panic("stack is empty")
	}
	
	node := l.root
	l.root = node.NextNode
	l.size--
	
	return node.Value
}

func (l *linkStack) Peek() string {
	if l.size == 0 {
		panic("stack is empty")
	}
	
	return l.root.Value
}

func (l *linkStack) Size() int {
	return l.size
}
