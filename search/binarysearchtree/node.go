package binarysearchtree

import "fmt"

type TreeNode struct {
	Value int64
	Times int64
	Left  *TreeNode
	Right *TreeNode
}

func (treeNode *TreeNode) Add(value int64) {
	if treeNode.Value == value {
		treeNode.Times++
		return
	}
	
	if value < treeNode.Value {
		if treeNode.Left == nil {
			treeNode.Left = &TreeNode{Value: value}
		} else {
			treeNode.Left.Add(value)
		}
	} else {
		if treeNode.Right == nil {
			treeNode.Right = &TreeNode{Value: value}
		} else {
			treeNode.Right.Add(value)
		}
	}
}

func (treeNode *TreeNode) FindMaxValue() *TreeNode {
	for treeNode.Right == nil {
		return treeNode
	}
	
	return treeNode.Right.FindMaxValue()
}

func (treeNode *TreeNode) FindMinValue() *TreeNode {
	for treeNode.Left == nil {
		return treeNode
	}
	
	return treeNode.Left.FindMinValue()
}

func (treeNode *TreeNode) Find(value int64) *TreeNode {
	if treeNode.Value == value {
		return treeNode
	}
	
	if value < treeNode.Value {
		if treeNode.Left == nil {
			return nil
		}
		
		return treeNode.Left.Find(value)
	} else {
		if treeNode.Right == nil {
			return nil
		}
		
		return treeNode.Right.Find(value)
	}
}

func (treeNode *TreeNode) FindParent(value int64) *TreeNode {
	if value < treeNode.Value {
		leftTree := treeNode.Left
		if leftTree == nil {
			return nil
		}
		
		if leftTree.Value == value {
			return treeNode
		}
		
		return leftTree.FindParent(value)
	}
	
	rightTree := treeNode.Right
	if rightTree == nil {
		return nil
	}
	
	if rightTree.Value == value {
		return treeNode
	}
	
	return rightTree.FindParent(value)
}

func (treeNode *TreeNode) MidOrder() {
	if treeNode == nil{
		return
	}
	
	treeNode.Left.MidOrder()
	times := treeNode.Times
	for ;times>=0;times--{
		fmt.Print(treeNode.Value, " ")
	}
	treeNode.Right.MidOrder()
	
}
