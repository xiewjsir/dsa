package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tr := &TreeNode{Data: "A"}
	tr.Left = &TreeNode{Data: "B"}
	tr.Right = &TreeNode{Data: "C"}
	tr.Left.Left = &TreeNode{Data: "D"}
	tr.Left.Right = &TreeNode{Data: "E"}
	tr.Right.Left = &TreeNode{Data: "F"}
	fmt.Println("先序排序：")
	PreOrder(tr)
	fmt.Println("\n中序排序：")
	MidOrder(tr)
	fmt.Println("\n后序排序")
	PostOrder(tr)
	fmt.Println("\n层序遍历")
	LayerOrder(tr)
	fmt.Println("")
}
