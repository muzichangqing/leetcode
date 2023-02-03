package bilibili

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[int]*TreeNode)
	st := []*TreeNode{root}
	for len(st) > 0 {
		n := st[len(st)-1]
		st = st[:len(st)-1]
		if n.Left != nil {
			parents[n.Left.Val] = n
			st = append(st, n.Left)
		}
		if n.Right != nil {
			parents[n.Right.Val] = n
			st = append(st, n.Right)
		}

	}
	ans := make([]int, 0)
	var travel func(*TreeNode, int, int)
	travel = func(tn *TreeNode, from, k int) {
		if k == 0 {
			ans = append(ans, tn.Val)
			return
		}
		if tn.Left != nil && tn.Left.Val != from {
			travel(tn.Left, tn.Val, k-1)
		}
		if tn.Right != nil && tn.Right.Val != from {
			travel(tn.Right, tn.Val, k-1)
		}
		if parent, ok := parents[tn.Val]; ok && parent.Val != from {
			travel(parent, tn.Val, k-1)
		}
	}
	travel(target, target.Val, k)
	return ans
}
