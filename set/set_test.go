package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	//other()
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)
	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())
	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Has(2) {
		fmt.Println("2 does exist")
	}
	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
