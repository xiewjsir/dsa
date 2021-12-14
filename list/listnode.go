package list


type listNode struct {
	next, prev *listNode
	value      string
}

// GetValue 获取节点值
func (node *listNode) GetValue() string {
	return node.value
}
// GetPrev 获取节点前驱节点
func (node *listNode) GetPrev() *listNode {
	return node.prev
}
// GetNext 获取节点后驱节点
func (node *listNode) GetNext() *listNode {
	return node.next
}
// HashNext 是否存在后驱节点
func (node *listNode) HashNext() bool {
	return node.prev != nil
}
// HashPre 是否存在前驱节点
func (node *listNode) HashPre() bool {
	return node.next != nil
}
// IsNil 是否为空节点
func (node *listNode) IsNil() bool {
	return node == nil
}
