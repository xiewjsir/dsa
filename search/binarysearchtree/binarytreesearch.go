package binarysearchtree

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree{
	return new(BinarySearchTree)
}

func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil{
		tree.Root = &BinarySearchTreeNode{Value:value}
	}

	tree.Root.Add(value)
}