package top100

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		maxGain := node.Val
		maxGain = max(maxGain, maxGain+left)
		maxGain = max(maxGain, maxGain+right)
		ans = max(ans, maxGain)
		childMax := max(max(left, right), 0)
		return childMax + node.Val
	}
	dfs(root)
	return ans
}
