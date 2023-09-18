package everyday

// 337. 打家劫舍 III
func rob(root *TreeNode) int {
	var rob func(*TreeNode, bool) int
	c1, c2 := make(map[*TreeNode]int), make(map[*TreeNode]int)

	rob = func(tn *TreeNode, b bool) int {
		var c map[*TreeNode]int
		if b {
			c = c1
		} else {
			c = c2
		}
		if res, ok := c[tn]; ok {
			return res
		}
		var res int
		if tn == nil {
			res = 0
		} else if b {
			res1 := tn.Val + rob(tn.Left, false) + rob(tn.Right, false)
			res2 := rob(tn.Left, true) + rob(tn.Right, true)
			if res1 > res2 {
				res = res1
			} else {
				res = res2
			}
		} else {
			res = rob(tn.Left, true) + rob(tn.Right, true)
		}
		c[tn] = res
		return res
	}
	res1, res2 := rob(root, true), rob(root, false)
	if res1 > res2 {
		return res1
	}
	return res2
}
