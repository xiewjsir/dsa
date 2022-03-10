package avl

type AVLTreeNode struct {
	Value  int64
	Times  int64
	Height int64
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 哪个子树高算哪棵的
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	// 高度加上自己那一层
	node.Height = maxHeight + 1
}

func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	return leftHeight - rightHeight
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	if node == nil {
		return &AVLTreeNode{Value: value}
	}

	if node.Value == value {
		node.Times++
		return node
	}

	var newTreeNode *AVLTreeNode

	if value < node.Value {
		node.Left = node.Left.Add(value)
		factor := node.BalanceFactor()
		if factor == 2 {
			if value < node.Left.Value {
				newTreeNode = RightRotation(node)
			} else {
				newTreeNode = LeftRightRotation(node)
			}
		}
	} else {
		node.Right = node.Right.Add(value)
		factor := node.BalanceFactor()
		if factor == -2 {
			if value > node.Right.Value {
				newTreeNode = LeftRotation(node)
			} else {
				newTreeNode = RightLeftRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newTreeNode.UpdateHeight()
		return newTreeNode
	}

	return node
}
