package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	averages := []float64{}
	myQueue := NewMyQueue()
	myQueue.PushBack(root)
	currentLevelCount := 1
	nextLevelCount := 0
	sum := 0
	count := 0
	for !myQueue.IsEmpty() {
		node := myQueue.Front()
		sum += node.Val
		count++
		if node.Left != nil {
			myQueue.PushBack(node.Left)
			nextLevelCount++
		}
		if node.Right != nil {
			myQueue.PushBack(node.Right)
			nextLevelCount++
		}
		if count == currentLevelCount {
			average := float64(sum) / float64(count)
			averages = append(averages, average)
			sum = 0
			count = 0
			currentLevelCount = nextLevelCount
			nextLevelCount = 0
		}
	}

	return averages
}

// MyQueue int类型的栈
type MyQueue struct {
	data []*TreeNode
}

// NewMyQueue Create a queue
func NewMyQueue() *MyQueue {
	data := make([]*TreeNode, 0)
	return &MyQueue{data}
}

// IsEmpty if myQueue is empty return true
func (myQueue *MyQueue) IsEmpty() bool {
	return len(myQueue.data) == 0
}

// PushBack an element to queue
func (myQueue *MyQueue) PushBack(node *TreeNode) {
	myQueue.data = append(myQueue.data, node)
}

// Front an element if exists, else return nil
func (myQueue *MyQueue) Front() *TreeNode {
	if len(myQueue.data) == 0 {
		return nil
	}
	node := myQueue.data[0]
	myQueue.data = myQueue.data[1:]
	return node
}
