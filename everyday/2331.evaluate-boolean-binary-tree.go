package everyday

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	l, r := evaluateTree(root.Left), evaluateTree(root.Right)
	if root.Val == 2 {
		return l || r
	}
	return l && r
}
