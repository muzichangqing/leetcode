package top100

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, sum int) bool {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum == targetSum
		}
		if node.Left != nil {
			if exist := dfs(node.Left, sum); exist {
				return exist
			}
		}
		if node.Right != nil {
			if exist := dfs(node.Right, sum); exist {
				return true
			}
		}
		return false
	}
	return dfs(root, 0)
}
