package everyday

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	queue := []*TreeNode{root}
	needSwap := true

	for {
		if queue[0].Left == nil {
			break
		}
		newQueue := make([]*TreeNode, len(queue)*2)
		for i, node := range queue {
			newQueue[2*i], newQueue[2*i+1] = node.Left, node.Right
		}
		if needSwap {
			for i := 0; i < len(newQueue)/2; i++ {
				swapNodeChildren(newQueue[i], newQueue[len(newQueue)-1-i])
				newQueue[i], newQueue[len(newQueue)-1-i] =
					newQueue[len(newQueue)-1-i], newQueue[i]
			}
			for i, node := range queue {
				node.Left, node.Right = newQueue[2*i], newQueue[2*i+1]
			}
		}
		needSwap = !needSwap
		queue = newQueue
	}

	return root
}

func swapNodeChildren(n1, n2 *TreeNode) {
	if n1 == nil || n2 == nil {
		panic("TreeNode shouldn't be nil")
	}
	n1.Left, n1.Right, n2.Left, n2.Right =
		n2.Left, n2.Right, n1.Left, n1.Right
}
