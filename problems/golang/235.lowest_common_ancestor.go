package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var min, max, current, ancestor *TreeNode
	if q.Val > p.Val {
		max = q
		min = p
	} else {
		max = p
		min = q
	}
	current = root
	for {
		if current.Val > max.Val {
			current = current.Left
		} else if current.Val < min.Val {
			current = current.Right
		} else {
			ancestor = current
			break
		}
	}

	return ancestor
}
