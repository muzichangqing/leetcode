package top100

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxWidth := 1
	q := []pair{{root, 1}}
	for len(q) > 0 {
		maxWidth = max(maxWidth, q[len(q)-1].index-q[0].index+1)
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, 2 * p.index})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, 2*p.index + 1})
			}
		}
	}
	return maxWidth
}
