package tree

import (
	"fmt"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Print(tree.Data, "")

	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	fmt.Print(tree.Data, "")
	MidOrder(tree.Right)
}

func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Print(tree.Data, "")
}

func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	q := new(LinkQueue)
	q.Add(tree)
	for q.size > 0 {
		element := q.Remove()

		fmt.Print(element.Data, "")

		if element.Left != nil {
			q.Add(element.Left)
		}

		if element.Right != nil {
			q.Add(element.Right)
		}
	}
}
