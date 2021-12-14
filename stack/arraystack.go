package stack

import "sync"

type stack interface {
	Pop() string
	Push(string2 string) *ArrayStack
	Peek() string
	Size() int
}

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func NewStack() *ArrayStack {
	return new(ArrayStack)
}

func (s *ArrayStack) Push(item string) *ArrayStack {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	s.array = append(s.array, item)
	s.size++
	return s
}

func (s *ArrayStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	if s.size == 0 {
		panic("stack is empty")
	}
	
	v := s.array[s.size-1]
	array := make([]string, s.size-1, s.size-1)
	for i := 0; i < s.size-1; i++ {
		array[i] = s.array[i]
	}
	
	s.array = array
	s.size--
	
	return v
}

func (s *ArrayStack) Peek() string {
	if s.size == 0{
		panic("stack is empty")
	}
	
	return s.array[s.size-1]
}

func (s *ArrayStack) Size() int {
	return s.size
}
