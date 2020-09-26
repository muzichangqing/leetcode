package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	paths := [][]int{}
	if root == nil {
		return paths
	}
	path := []int{}

	var backtrace func(node *TreeNode, currentSum int)
	backtrace = func(node *TreeNode, currentSum int) {
		currentSum += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if currentSum != sum {
				return
			}
			newResult := append([]int{}, path...)
			paths = append(paths, newResult)
			return
		}
		pathLen := len(path)
		if node.Left != nil {
			backtrace(node.Left, currentSum)
		}
		path = path[:pathLen]
		if node.Right != nil {
			backtrace(node.Right, currentSum)
		}
	}
	backtrace(root, 0)

	return paths
}
