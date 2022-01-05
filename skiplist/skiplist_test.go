package skiplist

import "testing"

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Add(12, "twelve")
	sl.Add(11, "eleven")
	t.Logf("now skip list: %d", sl.Len())
	t.Logf("get %d from skip : %+v", 11, sl.Search(11))
	sl.Remove(11)
	t.Logf("now skip list: %d", sl.Len())
	t.Logf("get %d from skip : %+v", 11, sl.Search(11))
}

//func BenchmarkNewSkipList(b *testing.B) {
//	sl := NewSkipList()
//	for i := 0; i < b.N; i++ {
//		sl.Add(11, "eleven")
//		sl.Search(11)
//		sl.Remove(11)
//		sl.Search(11)
//	}
//}