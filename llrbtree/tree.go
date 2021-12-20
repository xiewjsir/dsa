package llrbtree

// 左倾红黑树实现
// Left-leaning red-black tree

const (
	RED = true
	BLACK = false
)

// LLRBTree 左倾红黑树
type LLRBTree struct {
	Root *LLRBTreeNode
}

// Add 左倾红黑树添加元素
func (tree *LLRBTree) Add(value int64) {
	// 跟节点开始添加元素，因为可能调整，所以需要将返回的节点赋值回根节点
	tree.Root = tree.Root.Add(value)
	// 根节点的链接永远都是黑色的
	tree.Root.Color = BLACK
}

// FindMinValue 找出最小值的节点
func (tree *LLRBTree) FindMinValue() *LLRBTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	
	return tree.Root.FindMinValue()
}

// FindMaxValue 找出最大值的节点
func (tree *LLRBTree) FindMaxValue() *LLRBTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	
	return tree.Root.FindMaxValue()
}

// Find 查找指定节点
func (tree *LLRBTree) Find(value int64) *LLRBTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	
	return tree.Root.Find(value)
}

// MidOrder 中序遍历
func (tree *LLRBTree) MidOrder() {
	tree.Root.MidOrder()
}

// Delete 左倾红黑树删除元素
func (tree *LLRBTree) Delete(value int64) {
	// 当找不到值时直接返回
	if tree.Find(value) == nil {
		return
	}
	
	if !IsRed(tree.Root.Left) && !IsRed(tree.Root.Right) {
		// 左右子树都是黑节点，那么先将根节点变为红节点，方便后面的红色左移或右移
		tree.Root.Color = RED
	}
	
	tree.Root = tree.Root.Delete(value)
	
	// 最后，如果根节点非空，永远都要为黑节点，赋值黑色
	if tree.Root != nil {
		tree.Root.Color = BLACK
	}
}
