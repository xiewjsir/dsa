package queue

import (
	"sync"
)

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Value    string
	NextNode *LinkNode
}

func (l *LinkQueue) Add(v string) *LinkQueue {
	l.lock.Lock()
	defer l.lock.Unlock()

	nNode := new(LinkNode)
	nNode.Value = v

	if l.root == nil {
		l.root = nNode
	} else {
		nowNode := l.root
		for nowNode.NextNode != nil {
			nowNode = nowNode.NextNode
		}

		nowNode.NextNode = nNode
	}

	l.size++
	return l
}

func (l *LinkQueue) Remove() string {
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

func (l *LinkQueue) Size() int {
	return l.size
}
