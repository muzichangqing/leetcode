package codetop

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	size := 2
	for size > 0 {
		for i := 0; i < len(queue)/2; i++ {
			l := queue[i]
			r := queue[size-1-i]
			if l == nil && r == nil {
				continue
			}
			if l == nil || r == nil || l.Val != r.Val {
				return false
			}
		}
		temp := make([]*TreeNode, 0, size*2)
		for _, tn := range queue {
			if tn != nil {
				temp = append(temp, tn.Left, tn.Right)
			}
		}
		queue = temp
		size = len(temp)
	}
	return true
}

// 101. 对称二叉树 递归
func isSymmetricT(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(tn1, tn2 *TreeNode) bool {
		if tn1 == nil && tn2 == nil {
			return true
		}
		if tn1 == nil || tn2 == nil {
			return false
		}
		return tn1.Val == tn2.Val && check(tn1.Left, tn2.Right) && check(tn1.Right, tn2.Left)
	}
	return check(root, root)
}

// 22. 括号生成
func generateParenthesis(n int) []string {
	ans := []string{}
	size := 2 * n
	bts := make([]byte, size)
	var backtrace func(btsIndex, leftCount int)
	backtrace = func(btsIndex, leftCount int) {
		if btsIndex == size && leftCount == 0 {
			ans = append(ans, string(bts))
			return
		}
		if leftCount > 0 {
			bts[btsIndex] = ')'
			backtrace(btsIndex+1, leftCount-1)
		}
		if leftCount < size-1-btsIndex {
			bts[btsIndex] = '('
			backtrace(btsIndex+1, leftCount+1)
		}
	}
	backtrace(0, 0)
	return ans
}

// 169. 多数元素
func majorityElement(nums []int) int {
	ans := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			ans = num
		}
		if ans == num {
			count++
		} else {
			count--
		}
	}
	return ans
}

// 43. 字符串相乘
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m > n {
		m, n = n, m
		num1, num2 = num2, num1
	}
	var add func([]byte, []byte) []byte
	add = func(num1, num2 []byte) []byte {
		if len(num1) > len(num2) {
			return add(num2, num1)
		}
		res := []byte{}
		c := byte(0)
		i := 0
		for ; i < len(num1); i++ {
			sum := num1[i] + num2[i] + c
			res = append(res, sum%10)
			c = sum / 10
		}
		for ; i < len(num2); i++ {
			sum := num2[i] + c
			res = append(res, sum%10)
			c = sum / 10
		}
		if c != 0 {
			res = append(res, c)
		}
		return res
	}
	ans := []byte{0}
	for i := 0; i < m; i++ {
		multi := []byte{}
		c := byte(0)
		digit1 := num1[m-1-i] - '0'
		for j := 0; j < n; j++ {
			digit2 := num2[n-1-j] - '0'
			res := digit1*digit2 + c
			multi = append(multi, res%10)
			c = res / 10
		}
		if c != 0 {
			multi = append(multi, c)
		}
		multi = append(make([]byte, i), multi...)
		ans = add(multi, ans)
	}
	l, r := 0, len(ans)-1
	for l <= r {
		ans[l], ans[r] = ans[r]+'0', ans[l]+'0'
		l++
		r--
	}
	if ans[0] == '0' {
		return "0"
	}
	return string(ans)
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 165. 比较版本号
func compareVersion(version1 string, version2 string) int {
	version1S := strings.Split(version1, ".")
	version2S := strings.Split(version2, ".")
	m, n := len(version1S), len(version2S)
	i := 0
	for ; i < m && i < n; i++ {
		num1, _ := strconv.Atoi(version1S[i])
		num2, _ := strconv.Atoi(version2S[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	for ; i < m; i++ {
		num1, _ := strconv.Atoi(version1S[i])
		if num1 > 0 {
			return 1
		}
	}
	for ; i < n; i++ {
		num2, _ := strconv.Atoi(version2S[i])
		if num2 > 0 {
			return -1
		}
	}
	return 0
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)
	return root
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	binarySearch := func(lower bool) int {
		l, r := 0, len(nums)-1
		ans := len(nums)
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] > target || lower && nums[mid] == target {
				ans = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return ans
	}
	left := binarySearch(true)
	right := binarySearch(false) - 1
	if left <= right && right <= len(nums) {
		return []int{left, right}
	}
	return []int{-1, -1}
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(candidates)
	combination := []int{}
	var backtrace func(int, int)
	backtrace = func(index, left int) {
		if left == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			ans = append(ans, temp)
			return
		}
		if index == len(candidates) {
			return
		}
		candidate := candidates[index]
		if candidate > left {
			return
		}
		i := 0
		for left >= 0 {
			backtrace(index+1, left)
			i++
			combination = append(combination, candidate)
			left -= candidate
		}
		combination = combination[:len(combination)-i]
	}
	backtrace(0, target)
	return ans
}

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}

// 912. 排序数组
func sortArray(nums []int) []int {
	var maxHeapify func(index, length int)
	maxHeapify = func(index, length int) {
		left := 2*index + 1
		right := left + 1
		if left > length {
			return
		}
		maxIndex := left
		if right <= length && nums[right] > nums[left] {
			maxIndex = right
		}
		if nums[maxIndex] > nums[index] {
			nums[index], nums[maxIndex] = nums[maxIndex], nums[index]
			maxHeapify(maxIndex, length)
		}
	}

	length := len(nums) - 1
	for i := (length - 1) >> 1; i >= 0; i-- {
		maxHeapify(i, length)
	}

	for i := length; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(0, i-1)
	}

	return nums
}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	numsSet := make(map[int]bool)
	for _, num := range nums {
		numsSet[num] = true
	}
	longest := 0
	for num := range numsSet {
		currentNum := num
		currentLength := 1
		if numsSet[currentNum-1] {
			continue
		}
		for numsSet[currentNum+1] {
			currentLength++
			currentNum++
		}
		if currentLength > longest {
			longest = currentLength
		}
	}
	return longest
}

// 62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		row := make([]int, n)
		dp[i] = row
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 468. 验证IP地址
func validIPAddress(queryIP string) string {
	var (
		neither = "Neither"
		ipv4    = "IPv4"
		ipv6    = "IPv6"
	)
	if strings.Contains(queryIP, ".") {
		segs := strings.Split(queryIP, ".")
		if len(segs) != 4 {
			return neither
		}
		for i, seg := range segs {
			num, err := strconv.Atoi(seg)
			if err != nil {
				return neither
			}
			if num > 255 || num < 0 || i == 0 && num == 0 {
				return neither
			}
			if num != 0 && seg[0] == '0' || num == 0 && len(seg) > 1 {
				return neither
			}
		}
		return ipv4
	} else if strings.Contains(queryIP, ":") {
		segs := strings.Split(queryIP, ":")
		if len(segs) != 8 {
			return neither
		}
		for _, seg := range segs {
			if len(seg) > 4 || len(seg) < 1 {
				return neither
			}
			trimSeg := strings.TrimLeft(seg, "0")
			if trimSeg == "" {
				continue
			}
			num, err := strconv.ParseInt(trimSeg, 16, 0)
			if err != nil || num < 0 {
				return neither
			}
		}
		return ipv6
	}
	return neither
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// 240. 搜索二维矩阵 II
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// 221. 最大正方形
func maximalSquare(matrix [][]byte) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	maxSide := 0
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}

		}
	}
	return maxSide * maxSide
}

// 162. 寻找峰值
func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	l, r := 0, n-1
	for {
		mid := (l + r) >> 1
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid-1) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	longestCommonPrefixOfTwo := func(str1, str2 string) string {
		m, n := len(str1), len(str2)
		i := 0
		for i < m && i < n {
			if str1[i] == str2[i] {
				i++
			} else {
				break
			}
		}
		if i == 0 {
			return ""
		}
		return str1[:i]
	}
	prefix := strs[0]
	n := len(strs)
	for i := 1; i < n; i++ {
		prefix = longestCommonPrefixOfTwo(prefix, strs[i])
		if prefix == "" {
			return prefix
		}
	}
	return prefix
}

// 695. 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return dfs(i+1, j) + dfs(i, j+1) + dfs(i-1, j) + dfs(i, j-1) + 1
	}
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area := dfs(i, j); area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// 179. 最大数
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		return num1+num2 > num2+num1
	})
	if nums[0] == 0 {
		return "0"
	}
	number := ""
	for _, num := range nums {
		number += strconv.Itoa(num)
	}
	return number
}

// 227. 基本计算器 II
func calculate(s string) int {
	numStack := []int{}
	numStackTop := -1
	operatorStack := []byte{}
	operatorStackTop := -1

	num := -1
	operator := byte(0)
	bs := append([]byte(s), ' ')
	for _, b := range bs {
		if b >= '0' && b <= '9' {
			if num == -1 {
				num = int(b - '0')
			} else {
				num = num*10 + int(b-'0')
			}
		} else {
			if num != -1 {
				if operator == '*' {
					numStack[numStackTop] *= num
					operator = byte(0)
				} else if operator == '/' {
					numStack[numStackTop] /= num
					operator = byte(0)
				} else {
					numStackTop++
					numStack = append(numStack, num)
				}
				num = -1
			}
			if b == ' ' {
				continue
			} else if b == '+' || b == '-' {
				operatorStackTop++
				operatorStack = append(operatorStack, b)
			} else {
				operator = b
			}
		}
	}
	i := 0
	for _, operator := range operatorStack {
		if operator == '+' {
			numStack[i+1] = numStack[i] + numStack[i+1]
		} else {
			numStack[i+1] = numStack[i] - numStack[i+1]
		}
		i++
	}
	return numStack[i]
}

// 958. 二叉树的完全性检验
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		newQueue := []*TreeNode{}
		for i, node := range queue {
			if node == nil {
				for _, node := range newQueue {
					if node != nil {
						return false
					}
				}
				for j := i + 1; j < len(queue); j++ {
					if queue[j] != nil {
						return false
					}
				}
				return true
			} else {
				newQueue = append(newQueue, node.Left, node.Right)
			}
		}
		queue = newQueue
	}
	return true
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		ans, isAns := dfs(node.Right)
		if isAns {
			return ans, isAns
		}
		k--
		if k == 0 {
			return node.Val, true
		}
		ans, isAns = dfs(node.Left)
		if isAns {
			return ans, isAns
		}
		return 0, false
	}
	ans, _ := dfs(root)
	return ans
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	nodePointerMap := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		newNode := &Node{Val: cur.Val}
		nodePointerMap[cur] = newNode

	}

	dummyHead := &Node{}
	pre, cur := dummyHead, head
	for cur != nil {
		newNode := nodePointerMap[cur]
		newNode.Random = nodePointerMap[cur.Random]
		pre.Next = newNode
		pre, cur = newNode, cur.Next
	}

	return dummyHead.Next
}

// 394. 字符串解码
func decodeString(s string) string {
	stack := []byte{}
	for _, b := range []byte(s) {
		if b == ']' {
			// 找到[]中的字符串
			top := len(stack) - 1
			bs := []byte{}
			for {
				tb := stack[top]
				top--
				if tb == '[' {
					break
				}
				bs = append(bs, tb)
			}
			bsLen := len(bs)
			for i := 0; i < bsLen/2; i++ {
				bs[i], bs[bsLen-1-i] = bs[bsLen-1-i], bs[i]
			}
			// 找到[]前的数字
			counter, i := 0, 1
			for top >= 0 {
				tb := stack[top]
				if tb >= '0' && tb <= '9' {
					counter += int(tb-'0') * i
					i *= 10
					top--
				} else {
					break
				}
			}
			// 弹出
			stack = stack[:top+1]
			// 压入
			for counter > 0 {
				counter--
				stack = append(stack, bs...)
			}
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		prev := cur.Next
		after := cur.Next.Next
		cur.Next = after
		prev.Next = after.Next
		after.Next = prev
		cur = prev
	}
	return dummy.Next
}
