package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := new(ArrayQueue)
	queue.Add("cat")
	queue.Add("dog")
	queue.Add("hen")
	fmt.Println("size:", queue.Size())
	fmt.Println("remove:", queue.Remove())
	fmt.Println("remove:", queue.Remove())
	fmt.Println("size:", queue.Size())
	queue.Add("drag")
	fmt.Println("remove:", queue.Remove())
}