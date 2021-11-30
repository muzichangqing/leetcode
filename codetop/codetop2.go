package codetop

import (
	"container/heap"
	"math"
	"strconv"
	"strings"
)

//8. 字符串转换整数(atoi)
func myAtoi(s string) int {
	const (
		START = iota
		SIGN
		IN_NUMBER
		END
	)
	states := [][]int{
		{START, SIGN, IN_NUMBER, END},
		{END, END, IN_NUMBER, END},
		{END, END, IN_NUMBER, END},
		{END, END, END, END},
	}

	state := START
	sign := 1
	ans := 0
	getCol := func(c rune) int {
		if c == ' ' {
			return 0
		}
		if c == '+' || c == '-' {
			return 1
		}
		if c >= '0' && c <= '9' {
			return 2
		}
		return 3
	}
	for _, c := range s {
		state = states[state][getCol(c)]
		if state == END {
			break
		}
		if state == SIGN && c == '-' {
			sign = -1
		} else if state == IN_NUMBER {
			ans = ans*10 + int(c-'0')
			if ans > math.MaxInt32 {
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		}
	}
	return ans * sign
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	c1, c2, c3 := l1, l2, head
	carry := 0
	for c1 != nil && c2 != nil {
		val := c1.Val + c2.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c1 = c1.Next
		c2 = c2.Next
		c3 = c3.Next
	}
	for c1 != nil {
		val := c1.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c1 = c1.Next
		c3 = c3.Next
	}
	for c2 != nil {
		val := c2.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c2 = c2.Next
		c3 = c3.Next
	}
	if carry == 1 {
		c3.Next = &ListNode{Val: 1}
	}
	return head.Next
}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, corrent := dummy, head
	for corrent != nil {
		next := corrent.Next
		for next != nil && next.Val == corrent.Val {
			next = next.Next
		}
		if next != corrent.Next {
			prev.Next = next
			corrent = next
		} else {
			prev = corrent
			corrent = corrent.Next
		}
	}
	return dummy.Next
}

// 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return ans
}

// 148.排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	merge := func(h1, h2 *ListNode) *ListNode {
		h3 := &ListNode{}
		c1, c2, c3 := h1, h2, h3
		for c1 != nil && c2 != nil {
			if c1.Val > c2.Val {
				c3.Next = c2
				c2 = c2.Next
			} else {
				c3.Next = c1
				c1 = c1.Next
			}
			c3 = c3.Next
		}
		if c1 != nil {
			c3.Next = c1
		} else if c2 != nil {
			c3.Next = c2
		}

		return h3.Next
	}

	var mergeSort func(*ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}
		slow, fast := head, head.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		h1 := mergeSort(slow.Next)
		slow.Next = nil
		h2 := mergeSort(head)
		return merge(h1, h2)
	}
	return mergeSort(head)
}

// 31. 下一个排列
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 {
		if nums[i+1] > nums[i] {
			break
		}
		i--
	}
	if i >= 0 {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				break
			}
		}
	}

	i++
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

//76. 最小覆盖子串
func minWindow(s string, t string) string {
	cur := make(map[byte]int)
	ori := make(map[byte]int)
	for _, c := range []byte(t) {
		ori[c] += 1
	}
	check := func() bool {
		for c, n := range ori {
			if cur[c] < n {
				return false
			}
		}
		return true
	}
	var lo, hi int
	ansL, ansR := -1, -1
	length := math.MaxInt32
	sLen := len(s)
	for hi < sLen {
		if ori[s[hi]] > 0 {
			cur[s[hi]]++
		}
		for check() && lo <= hi {
			if hi-lo+1 < length {
				length = hi - lo + 1
				ansL, ansR = lo, hi
			}
			if cur[s[lo]] > 0 {
				cur[s[lo]]--
			}
			lo++
		}
		hi++
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL : ansR+1]
}

// 151. 翻转字符串里的单词
func reverseWords(s string) string {
	words := []string{}
	l := -1
	for i, c := range s {
		if c == ' ' && l != -1 {
			words = append(words, s[l:i])
			l = -1
		} else if c != ' ' && l == -1 {
			l = i
		}
	}
	if l != -1 {
		words = append(words, s[l:])
	}
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// 72. 编辑距离
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	min := func(x, y, z int) int {
		r := x
		if y < r {
			r = y
		}
		if r < z {
			z = r
		}
		return z
	}
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]-1) + 1
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	k := (m + n + 1) / 2

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	left, right := 0, m
	for left < right {
		i := (left + right + 1) / 2
		j := k - i
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}

	i, j := left, k-left
	nums1Left := math.MinInt32
	if i != 0 {
		nums1Left = nums1[i-1]
	}
	nums2Left := math.MinInt32
	if j != 0 {
		nums2Left = nums2[j-1]
	}
	if (m+n)%2 == 1 {
		return float64(max(nums1Left, nums2Left))
	}
	nums1Right := math.MaxInt32
	if i != m {
		nums1Right = nums1[i]
	}
	nums2Right := math.MaxInt32
	if j != n {
		nums2Right = nums2[j]
	}
	return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2.0
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var height func(*TreeNode) (int, bool)
	height = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		lHeight, blance := height(root.Left)
		if !blance {
			return 0, false
		}
		rHeight, blance := height(root.Right)
		if !blance {
			return 0, false
		}
		if math.Abs(float64(lHeight-rHeight)) > 1 {
			return 0, false
		}
		return max(lHeight, rHeight) + 1, true
	}
	_, blanced := height(root)
	return blanced
}

// 239. 滑动窗口最大值
type hp struct {
	data [][2]int
}

func (h *hp) Len() int {
	return len(h.data)
}

func (h *hp) Less(i, j int) bool {
	return h.data[i][1] > h.data[j][1]
}

func (h *hp) Push(x interface{}) {
	h.data = append(h.data, x.([2]int))
}

func (h *hp) Pop() interface{} {
	x := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return x
}

func (h *hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func maxSlidingWindow(nums []int, k int) []int {
	numsLength := len(nums)
	ans := make([]int, 1, numsLength-k+1)

	h := &hp{data: make([][2]int, k)}
	for i := 0; i < k; i++ {
		h.data[i] = [2]int{i, nums[i]}
	}
	heap.Init(h)
	ans[0] = h.data[0][1]

	for i := k; i < numsLength; i++ {
		heap.Push(h, [2]int{i, nums[i]})
		for h.data[0][0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, h.data[0][1])
	}

	return ans
}

// 93. 复原 IP 地址
func restoreIpAddresses(s string) []string {
	ans := []string{}
	const SEG_COUNT = 4
	segs := [SEG_COUNT]int{}
	sLen := len(s)

	var dfs func(int, int)
	dfs = func(segId, sIndex int) {
		if segId == SEG_COUNT {
			if sIndex == sLen {
				addr := ""
				for _, seg := range segs {
					addr += strconv.Itoa(seg)
					addr += "."
				}
				ans = append(ans, strings.TrimRight(addr, "."))
			}
			return
		}
		if sIndex == sLen {
			return
		}
		if s[sIndex] == '0' {
			segs[segId] = 0
			dfs(segId+1, sIndex+1)
			return
		}
		seg := 0
		for i := sIndex; i < sLen; i++ {
			seg = seg*10 + int(s[i]-'0')
			if seg > 0 && seg <= 0xFF {
				segs[segId] = seg
				dfs(segId+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return ans
}
