package llrbtree

import "fmt"

func IsRed(node *LLRBTreeNode) bool {
	if node == nil {
		return false
	}

	return node.Color == RED
}

// RotateLeft 左旋
func RotateLeft(h *LLRBTreeNode) *LLRBTreeNode {
	if h == nil {
		return nil
	}

	x := h.Right
	h.Right = x.Left
	x.Left = h

	x.Color = h.Color
	h.Color = RED
	return x
}

// RotateRight 右旋
func RotateRight(h *LLRBTreeNode) *LLRBTreeNode {
	if h == nil {
		return nil
	}

	x := h.Left
	h.Left = x.Right
	x.Right = h

	x.Color = h.Color
	h.Color = RED
	return x
}

func ColorChange(node *LLRBTreeNode) *LLRBTreeNode {
	if node == nil {
		return nil
	}

	node.Color = !node.Color
	node.Left.Color = !node.Left.Color
	node.Right.Color = !node.Right.Color
	return node
}

// MoveRedLeft 红色左移, 节点 h 是红节点，其左儿子和左儿子的左儿子都为黑节点，左移后使得其左儿子或左儿子的左儿子有一个是红色节点
func MoveRedLeft(h *LLRBTreeNode) *LLRBTreeNode {
	ColorChange(h)

	// 右儿子的左儿子有红链接
	if IsRed(h.Right.Left) {
		h.Right = RotateRight(h.Right)
		h = RotateLeft(h)
		ColorChange(h)
	}
	return h
}

// MoveRedRight 红色右移 节点 h 是红节点，其右儿子和右儿子的左儿子都为黑节点，右移后使得其右儿子或右儿子的右儿子有一个是红色节点
func MoveRedRight(h *LLRBTreeNode) *LLRBTreeNode {
	// 应该确保 isRed(h) && !isRed(h.right) && !isRed(h.right.left);
	ColorChange(h)

	// 左儿子有左红链接
	if IsRed(h.Left.Left) {
		// 右旋
		h = RotateRight(h)
		// 变色
		ColorChange(h)
	}

	return h
}

type LLRBTreeNode struct {
	Value       int64
	Times       int64
	Left, Right *LLRBTreeNode
	Color       bool
}

// Add 往节点添加元素
func (node *LLRBTreeNode) Add(value int64) *LLRBTreeNode {
	// 插入的节点为空，将其链接颜色设置为红色，并返回
	if node == nil {
		return &LLRBTreeNode{
			Value: value,
			Color: RED,
		}
	}

	// 插入的元素重复
	if value == node.Value {
		node.Times = node.Times + 1
	} else if value > node.Value {
		// 插入的元素比节点值大，往右子树插入
		node.Right = node.Right.Add(value)
	} else {
		// 插入的元素比节点值小，往左子树插入
		node.Left = node.Left.Add(value)
	}

	// 辅助变量
	nowNode := node

	// 右链接为红色，那么进行左旋，确保树是左倾的
	// 这里做完操作后就可以结束了，因为插入操作，新插入的右红链接左旋后，nowNode节点不会出现连续两个红左链接，因为它只有一个左红链接
	if IsRed(nowNode.Right) && !IsRed(nowNode.Left) {
		nowNode = RotateLeft(nowNode)
	} else {
		// 连续两个左链接为红色，那么进行右旋
		if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
			nowNode = RotateRight(nowNode)
		}

		// 旋转后，可能左右链接都为红色，需要变色
		if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
			ColorChange(nowNode)
		}
	}

	return nowNode
}

// Delete 对该节点所在的子树删除元素
func (node *LLRBTreeNode) Delete(value int64) *LLRBTreeNode {
	// 辅助变量
	nowNode := node
	// 删除的元素比子树根节点小，需要从左子树删除
	if value < nowNode.Value {
		// 因为从左子树删除，所以要判断是否需要红色左移
		if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
			// 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
			nowNode = MoveRedLeft(nowNode)
		}

		// 现在可以从左子树中删除了
		nowNode.Left = nowNode.Left.Delete(value)
	} else {
		// 删除的元素等于或大于树根节点

		// 左节点为红色，那么需要右旋，方便后面可以红色右移
		if IsRed(nowNode.Left) {
			nowNode = RotateRight(nowNode)
		}

		// 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		// 为什么呢，反证，它没有右儿子，但有左儿子，因为左倾红黑树的特征，那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
		if value == nowNode.Value && nowNode.Right == nil {
			return nil
		}

		// 因为从右子树删除，所以要判断是否需要红色右移
		if !IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left) {
			// 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
			nowNode = MoveRedRight(nowNode)
		}

		// 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
		if value == nowNode.Value {
			minNode := nowNode.Right.FindMinValue()
			nowNode.Value = minNode.Value
			nowNode.Times = minNode.Times

			// 删除其最小后驱节点
			nowNode.Right = nowNode.Right.DeleteMin()
		} else {
			// 删除的元素比子树根节点大，需要从右子树删除
			nowNode.Right = nowNode.Right.Delete(value)
		}
	}

	// 最后，删除叶子节点后，需要恢复左倾红黑树特征
	return nowNode.FixUp()
}

func (node *LLRBTreeNode) FindMinValue() *LLRBTreeNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.FindMinValue()
}

func (node *LLRBTreeNode) FindMaxValue() *LLRBTreeNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

func (node *LLRBTreeNode) Find(value int64) *LLRBTreeNode {
	if value == node.Value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Right.Find(value)
	}
}

func (node *LLRBTreeNode) MidOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 打印右子树
	node.Right.MidOrder()
}

// DeleteMin 对该节点所在的子树删除最小元素
func (node *LLRBTreeNode) DeleteMin() *LLRBTreeNode {
	// 辅助变量
	nowNode := node

	// 没有左子树，那么删除它自己
	if nowNode.Left == nil {
		return nil
	}

	// 判断是否需要红色左移，因为最小元素在左子树中
	if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
		nowNode = MoveRedLeft(nowNode)
	}

	// 递归从左子树删除
	nowNode.Left = nowNode.Left.DeleteMin()

	// 修复左倾红黑树特征
	return nowNode.FixUp()
}

// FixUp 修复左倾红黑树特征
func (node *LLRBTreeNode) FixUp() *LLRBTreeNode {
	// 辅助变量
	nowNode := node

	// 红链接在右边，左旋恢复，让红链接只出现在左边
	if IsRed(nowNode.Right) {
		nowNode = RotateLeft(nowNode)
	}

	// 连续两个左链接为红色，那么进行右旋
	if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
		nowNode = RotateRight(nowNode)
	}

	// 旋转后，可能左右链接都为红色，需要变色
	if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
		ColorChange(nowNode)
	}

	return nowNode
}
