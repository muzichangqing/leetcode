package top100

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode, []int, int)
	dfs = func(node *TreeNode, path []int, sum int) {
		sum += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				p := make([]int, len(path))
				copy(p, path)
				ans = append(ans, p)
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, path, sum)
		}
		if node.Right != nil {
			dfs(node.Right, path, sum)
		}
		path = path[:len(path)-1]
	}
	dfs(root, []int{}, 0)
	return ans
}
