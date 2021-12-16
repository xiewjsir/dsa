package binarysearchtree

type Tree struct {
	Root *TreeNode
}

func NewTree() *Tree {
	return new(Tree)
}

func (tree *Tree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &TreeNode{Value: value}
		return
	}
	
	tree.Root.Add(value)
}

func (tree *Tree) FindMaxValue() *TreeNode {
	if tree.Root == nil {
		return nil
	}
	
	return tree.Root.FindMaxValue()
}

func (tree *Tree) FindMinValue() *TreeNode {
	if tree.Root == nil {
		return nil
	}
	
	return tree.Root.FindMinValue()
}

func (tree *Tree) Find(value int64) *TreeNode {
	if tree.Root == nil {
		return nil
	}
	
	return tree.Root.Find(value)
}

func (tree *Tree) FindParent(value int64) *TreeNode {
	if tree.Root == nil {
		return nil
	}
	
	if tree.Root.Value == value {
		return nil
	}
	
	return tree.Root.FindParent(value)
}

func (tree *Tree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	
	node := tree.Root.Find(value)
	if node == nil{
		return
	}
	
	parent := tree.Root.FindParent(value)
	if parent == nil && node.Left == nil && node.Right == nil{
		tree.Root = nil
		return
	}
	
	if parent != nil && node.Left == nil && node.Right == nil{
		if parent.Left != nil && parent.Left.Value == node.Value{
			parent.Left = nil
		}else {
			parent.Right = nil
		}
		
		return
	}
	
	if node.Left != nil && node.Right != nil{
		minNode := node.Right
		if minNode.Left != nil{
			minNode = minNode.Left
		}
		
		tree.Delete(minNode.Value)
		
		node.Value = minNode.Value
		node.Times = minNode.Times
		
		return
	}
	
	if parent == nil{
		if node.Left != nil{
			tree.Root = node.Left
		}else{
			tree.Root = node.Right
		}
		
		return
	}
	
	if node.Left != nil{
		if parent.Left != nil && value == parent.Left.Value{
			parent.Left = node.Left
		}else{
			parent.Right = node.Left
		}
		return
	}
	
	if node.Right != nil{
		if parent.Right != nil && value == parent.Right.Value{
			parent.Right = node.Right
		}else{
			parent.Left = node.Right
		}
		return
	}
}

func (tree *Tree) MidOrder() {
	if tree.Root == nil{
		return
	}
	
	tree.Root.MidOrder()
}