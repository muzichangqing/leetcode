package codetop

import (
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
