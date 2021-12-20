package avl


func RightRotation(root *AVLTreeNode) *AVLTreeNode {
	pivot := root.Left
	b := pivot.Right
	pivot.Right = root
	root.Left = b
	
	pivot.UpdateHeight()
	root.UpdateHeight()
	return pivot
}

func LeftRotation(root *AVLTreeNode) *AVLTreeNode {
	pivot := root.Right
	b := pivot.Left
	pivot.Left = root
	root.Right = b
	
	pivot.UpdateHeight()
	root.UpdateHeight()
	return pivot
}

func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = LeftRotation(node.Right)
	return RightRotation(node)
}
