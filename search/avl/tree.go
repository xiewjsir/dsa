package avl

type AVLTree struct {
	Root *AVLTreeNode
}

// NewAVLTree 初始化一个AVL树
func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

func (tree *AVLTree) Add(value int64) {
	tree.Root = tree.Root.Add(value)
}