package top100

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)

	q := []*TreeNode{root}
	for len(q) > 0 {
		nq := make([]*TreeNode, 0, len(q)*2)
		ans = append(ans, q[len(q)-1].Val)
		for _, node := range q {
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		q = nq
	}
	return ans
}
