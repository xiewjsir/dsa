package queue

import "sync"

type queue interface {
	Add(v string) *ArrayQueue
	Remove() string
	Size() int
}

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (a *ArrayQueue) Add(v string) *ArrayQueue {
	a.lock.Lock()
	defer a.lock.Unlock()
	
	a.array = append(a.array, v)
	a.size++
	
	return a
}

func (a *ArrayQueue) Remove() string {
	a.lock.Lock()
	defer a.lock.Unlock()
	
	if a.size == 0 {
		panic("queue is empty")
	}
	
	v := a.array[0]
	a.size--
	/*    直接原位移动，但缩容后继的空间不会被释放
	      for i := 1; i < a.size; i++ {
	          // 从第一位开始进行数据移动
	          a.array[i-1] = a.array[i]
	      }
	      // 原数组缩容
	      a.array = a.array[0 : a.size-1]
	*/
	
	// 创建新的数组，移动次数过多
	array := make([]string, a.size, a.size)
	for i := 0; i < a.size; i++ {
		array[i] = a.array[i+1]
	}
	a.array = array
	
	return v
}

func (a *ArrayQueue) Size() int {
	return a.size
}
