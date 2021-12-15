package queue

import (
	"fmt"
	"testing"
)

func TestarrayQueue(t *testing.T) {
	queue := new(arrayQueue)
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