package queue

import (
	"fmt"
	"testing"
)

func TestLinkQueue(t *testing.T) {
	linkQueue := new(LinkQueue)
	linkQueue.Add("cat")
	linkQueue.Add("dog")
	linkQueue.Add("hen")
	fmt.Println("size:", linkQueue.Size())
	fmt.Println("pop:", linkQueue.Remove())
	fmt.Println("pop:", linkQueue.Remove())
	fmt.Println("size:", linkQueue.Size())
	linkQueue.Add("drag")
	fmt.Println("pop:", linkQueue.Remove())
}