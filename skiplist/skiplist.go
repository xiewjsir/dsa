package skiplist

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

const defaultLevel = 12 // default level

type skipListNode struct {
	key  int
	data interface{}
	next []*skipListNode
}

type skipList struct {
	head, tail *skipListNode
	length     int
	level      int
	m          sync.RWMutex
	rand       *rand.Rand
}

func NewSkipList(opt ...int) *skipList {
	var level = defaultLevel
	if opt != nil {
		level = opt[0]
	}
	var sl = &skipList{
		head:  &skipListNode{next: make([]*skipListNode, level, level)},
		tail:  &skipListNode{key: math.MinInt32},
		level: level,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for l := range sl.head.next {
		sl.head.next[l] = sl.tail
	}

	return sl
}

func (sl *skipList) Add(key int, data interface{}) {
	sl.m.Lock()
	defer sl.m.Unlock()

	// 1. 随机一层用来开始添加该节点
	l := sl.randomLevel()

	// 2. 从该层还是寻找第一个比key大的节点
	addPos := make([]*skipListNode, l)
	node := sl.head
	for index := l - 1; index >= 0; index-- {
		for {
			temp := node.next[index]
			if temp == sl.tail || temp.key > key {
				addPos[index] = node // 找到插入节点
				break
			} else if temp.key == key {
				temp.data = data // 如果已经存在节点，就直接更新节点数据
				return
			} else {
				node = temp
			}
		}
	}

	nowNode := &skipListNode{key: key, data: data, next: make([]*skipListNode, l)}
	for level, pos := range addPos {
		nowNode.next[level] = pos.next[level]
		pos.next[level] = nowNode
	}

	sl.length++
}

func (sl *skipList) randomLevel() int {
	level := 1
	for ; level < sl.level && sl.rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}

func (sl *skipList) Remove(key int) bool {
	sl.m.Lock()
	defer sl.m.Unlock()

	node := sl.head
	remPos := make([]*skipListNode, sl.level)
	var target *skipListNode
	for level := sl.level - 1; level >= 0; level-- {
		for {
			temp := node.next[level]
			if temp == sl.tail || temp.key > key {
				break
			} else if temp.key == key {
				remPos[level] = node
				target = temp
				break
			} else {
				node = temp
			}
		}
	}

	if target != nil {
		for l, pos := range remPos {
			if pos != nil {
				pos.next[l] = target.next[l]
			}
		}
		target = nil
		sl.level--
	}

	return true
}

func (sl *skipList) Search(key int) interface{} {
	sl.m.RLock()
	defer sl.m.RUnlock()

	node := sl.head
	for level := sl.level - 1; level >= 0; level-- {
		for {
			temp := node.next[level]
			if temp == sl.tail || temp.key > key {
				break
			} else if temp.key == key {
				return temp.data
			} else {
				node = temp
			}
		}
	}

	return false
}

func (sl *skipList) Len() int {
	sl.m.Lock()
	defer sl.m.Unlock()
	return sl.length
}
