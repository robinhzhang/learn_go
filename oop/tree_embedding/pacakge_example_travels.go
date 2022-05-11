package tree //遍历树

func (node *Node) TraverseP() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
