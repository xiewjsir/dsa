package set

import "sync"

type set struct {
	m map[int]struct{}
	len int
	sync.RWMutex
}

func NewSet(cap int64) *set{
	m := make(map[int]struct{},cap)
	return &set{
		m:m,
	}
}

func (s *set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.m[item]= struct{}{}
	s.len = len(s.m)
}

func (s *set) delete(item int) {
	s.Lock()
	defer s.Unlock()

	if s.len == 0{
		return
	}

	delete(s.m,item)
	s.len = len(s.m)
}

func (s *set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()

	_,ok := s.m[item]
	return ok
}

func (s *set) Size() int {
	return s.len
}

func (s *set) IsEmpty() bool {
	if s.len == 0{
		return true
	}

	return false
}

func (s *set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = make(map[int]struct{},0)
	s.len = 0
}

func (s *set)List() []int {
	s.RLock()
	defer s.RUnlock()

	slice := make([]int,0,s.len)
	for item := range s.m{
		slice = append(slice,item)
	}

	return slice
}