package everyday

// 337. 打家劫舍 III
func rob(root *TreeNode) int {
	m1, m2 := make(map[*TreeNode]int), make(map[*TreeNode]int)
	var dfs func(*TreeNode)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		dfs(tn.Right)
		m1[tn] = tn.Val + m2[tn.Left] + m2[tn.Right]
		m2[tn] = max(m1[tn.Left], m2[tn.Left]) + max(m1[tn.Right], m2[tn.Right])
	}
	dfs(root)
	return max(m1[root], m2[root])
}
