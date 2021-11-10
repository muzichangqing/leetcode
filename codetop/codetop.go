package codetop

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// 146. LRU(最近最少使用)缓存机制
type LRUCache struct {
	capacity int                  // 容量
	length   int                  // 当前长度
	data     map[int]*dLinkedNode // key和双向链表节点的映射
	head     *dLinkedNode         // 双向链表的头
	tail     *dLinkedNode         // 双向链表尾
}

// dLinkedNode 双向链表
type dLinkedNode struct {
	key  int
	val  int
	prev *dLinkedNode
	next *dLinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		length:   0,
		data:     map[int]*dLinkedNode{},
		head:     &dLinkedNode{},
		tail:     &dLinkedNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// 从data中获取key的节点，如果没有，则返回-1
// 如果存在，则把该节点移动到链表的头部，并返回该节点的值
func (cache *LRUCache) Get(key int) int {
	node, ok := cache.data[key]
	if !ok {
		return -1
	}
	cache.moveToHead(node)
	return node.val
}

// 如果该节点存在，则更新节点值，并移动到头部
// 如果该节点不存在，则添加节点到头部，并判断length是否大于capacity，如果是，则删除尾部节点
func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.data[key]
	if ok {
		node.val = value
		cache.moveToHead(node)
		return
	}
	node = &dLinkedNode{key: key, val: value}
	cache.addToHead(node)
	cache.data[key] = node
	cache.length++
	if cache.length > cache.capacity {
		tail := cache.removeTail()
		delete(cache.data, tail.key)
		cache.length--
	}
}

// 把节点从链表中去除
func (cache *LRUCache) removeNode(node *dLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 把节点添加到链表首部
func (cache *LRUCache) addToHead(node *dLinkedNode) {
	node.next = cache.head.next
	node.prev = cache.head
	cache.head.next.prev = node
	cache.head.next = node
}

// 把节点移动到首部
func (cache *LRUCache) moveToHead(node *dLinkedNode) {
	cache.removeNode(node)
	cache.addToHead(node)
}

// 删除尾部节点
func (cache *LRUCache) removeTail() *dLinkedNode {
	removed := cache.tail.prev
	cache.removeNode(removed)
	return removed
}

// 3. 无重复字符的最长字串
func lengthOfLongestSubstring(s string) int {
	var i, j, max, cLen int
	for ; j < len(s); j++ {
		// 判断s[j]在 i 到 j-1之间是否存在
		for k := i; k < j; k++ {
			if s[k] == s[j] {
				// 发现重复的字符，把i移到此为止
				i = k + 1
				break
			}
		}
		// 计算当前字串长度
		cLen = j - i + 1
		if cLen > max {
			max = cLen
		}
	}
	return max
}

// 25. K个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead, newTail, currentGroupHead, nextGroupHead *ListNode
	currentGroupHead = head
	i := 1
	for head != nil {
		if i == k {
			// 记录下一组的起点
			nextGroupHead = head.Next
			head.Next = nil
			// 翻转当前组
			newGroupHead := reverseList(currentGroupHead)
			if newHead == nil {
				newHead = newGroupHead
			} else {
				newTail.Next = newGroupHead
			}
			// 尾节点更新
			newTail = currentGroupHead
			// 开始下一组
			head = nextGroupHead
			currentGroupHead = nextGroupHead
			i = 1
		} else {
			head = head.Next
			i++
		}

	}
	// 剩下的部分
	if currentGroupHead != nil {
		newTail.Next = currentGroupHead
	}
	return newHead
}

// 215. 数组中第K个最大元素
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// 快速排序，找出倒数第index个元素时（即index左侧的元素全部小于index)，即为数组中倒数第index个最大元素
func quickSelect(nums []int, l, r, index int) int {
	if l == r {
		return nums[l]
	}
	// 随机选择一个元素作为基准值，把该元素放到最右侧，遍历 l 到 r - 1，如果该元素大于选择的基准值，则移动到右侧
	randIndex := rand.Intn(r-l) + l
	nums[randIndex], nums[r] = nums[r], nums[randIndex]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 0 - i 都是小于等于基准值的, i就是基准值的位置
	nums[i+1], nums[r] = nums[r], nums[i+1]
	if i+1 == index {
		return nums[i+1]
	} else if i+1 < index {
		return quickSelect(nums, i+2, r, index)
	} else {
		return quickSelect(nums, l, i, index)
	}
}

// 94.二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalCore(root, []int{})
}

func inorderTraversalCore(root *TreeNode, vals []int) []int {
	if root == nil {
		return vals
	}
	vals = inorderTraversalCore(root.Left, vals)
	vals = append(vals, root.Val)
	vals = inorderTraversalCore(root.Right, vals)
	return vals
}

// 300.最长递增子序列
// 动态规划
func lengthOfLIS(nums []int) int {
	numsL := len(nums)
	dp := make([]int, numsL)
	dp[0] = 1
	max := 1
	for i := 1; i < numsL; i++ {
		dpj := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > dpj {
				dpj = dp[j]
			}
		}
		dp[i] = dpj + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 贪心+二分查找
func lengthOfLIS2(nums []int) int {
	numsL := len(nums)
	dp := []int{0, nums[0]}
	len := 1
	for i := 1; i < numsL; i++ {
		if nums[i] > dp[len] {
			len = len + 1
			dp = append(dp, nums[i])
		} else {
			l, r, pos := 1, len, 0
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]
		}
	}
	return len
}

// 143. 重排链表
func reorderList(head *ListNode) {
	// 找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 断开链表，反转后半部分
	next := slow.Next
	slow.Next = nil
	slow = next
	var head2 *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}
	// 合并
	c1, c2 := head, head2
	for c2 != nil {
		c2Next := c2.Next
		c2.Next = c1.Next
		c1.Next = c2
		c1 = c2.Next
		c2 = c2Next
	}
}

//199.二叉树的右视图
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		nextQueue := []*TreeNode{}
		ans = append(ans, queue[0].Val)
		for _, node := range queue {
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
		}
		queue = nextQueue
	}
	return ans
}

// 70.爬楼梯
func climbStairs(n int) int {
	id1, id2, ans := 0, 0, 1
	for i := 1; i <= n; i++ {
		id1 = id2
		id2 = ans
		ans = id1 + id2
	}
	return ans
}

// 剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

//124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		nodeGain := node.Val + max(leftGain, rightGain)
		maxSum = max(maxSum, node.Val+leftGain+rightGain)
		return nodeGain
	}
	maxGain(root)
	return maxSum
}

//69.Sqrt(x)
// 二分查找
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right, ans := 0, x, -1
	for left <= right {
		mid := (left + right) >> 1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// 牛顿迭代
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

//56. 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	size := len(intervals)
	i := 0
	for i < size {
		current := intervals[i]
		j := i + 1
		for j < size {
			cursor := intervals[j]
			if cursor[0] > current[1] {
				break
			}
			if cursor[1] > current[1] {
				current[1] = cursor[1]
			}
			j++
		}
		ans = append(ans, current)
		i = j
	}
	return ans
}

//19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
