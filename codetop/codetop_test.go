package codetop

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	lists := [][]int{
		{},
		{1, 2, 3, 4, 5},
		{1, 2},
	}

	results := [][]int{
		{},
		{5, 4, 3, 2, 1},
		{2, 1},
	}

	for index, values := range lists {
		list := sliceToSinglyList(values)
		reverseNode := reverseList(list)
		result := singlyListToSlice(reverseNode)
		correctResult := results[index]

		if len(result) != len(correctResult) {
			t.Errorf("reversing fail: %v => %v", values, result)
			continue
		}

		for i, val := range result {
			if val != correctResult[i] {
				t.Errorf("reversing fail: %v => %v", values, result)
				break
			}
		}
	}
}

func sliceToSinglyList(values []int) *ListNode {
	list := ListNode{0, nil}
	cursor := &list

	for _, val := range values {
		cursor.Next = &ListNode{val, nil}
		cursor = cursor.Next
	}

	return list.Next
}

func singlyListToSlice(list *ListNode) []int {
	var result []int

	cursor := list
	for cursor != nil {
		result = append(result, cursor.Val)
		cursor = cursor.Next
	}

	return result
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if val := cache.Get(1); val != 1 {
		t.Errorf("key 1 => value 1, not %d", val)
	}
	cache.Put(3, 3)
	if val := cache.Get(2); val != -1 {
		t.Errorf("key 2 => value -1, not %d", val)
	}
	cache.Put(4, 4)
	if val := cache.Get(1); val != -1 {
		t.Errorf("key 1 => value -1, not %d", val)
	}
	if val := cache.Get(3); val != 3 {
		t.Errorf("key 3 => value 3, not %d", val)
	}
	if val := cache.Get(4); val != 4 {
		t.Errorf("key 4 => value 4, not %d", val)
	}

	cache = Constructor(1)
	cache.Put(2, 1)
	if val := cache.Get(2); val != 1 {
		t.Errorf("key 2 => value 1, not %d", val)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	strs := []struct {
		str string
		len int
	}{
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, item := range strs {
		if len := lengthOfLongestSubstring(item.str); len != item.len {
			t.Errorf("error len of '%s', should be %d, got %d", item.str, item.len, len)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	lists := [][]int{
		{1, 1},
		{2, 1, 2, 3, 4, 5},
		{3, 1, 2, 3, 4, 5},
	}

	results := [][]int{
		{1},
		{2, 1, 4, 3, 5},
		{3, 2, 1, 4, 5},
	}

	for index, values := range lists {
		list := sliceToSinglyList(values[1:])
		reverseNode := reverseKGroup(list, values[0])
		result := singlyListToSlice(reverseNode)
		correctResult := results[index]

		if len(result) != len(correctResult) {
			t.Errorf("reversing K %d group fail: %v => %v", values[0], values[1:], result)
			continue
		}

		for i, val := range result {
			if val != correctResult[i] {
				t.Errorf("reversing K %d group fail: %v => %v", values[0], values[1:], result)
				break
			}
		}
	}
}

func TestFindKthLargest(t *testing.T) {
	numsList := [][]int{
		{2, 3, 2, 1, 5, 6, 4},
		{4, 3, 2, 3, 1, 2, 4, 5, 5, 6},
	}
	results := []int{5, 4}
	for index, nums := range numsList {
		if result := findKthLargest(nums[1:], nums[0]); result != results[index] {
			t.Errorf("%dth largest of %v is %d, got %d", nums[0], nums[1:], results[index], result)
		}
	}
}
func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	ans := inorderTraversal(root)
	correct := []int{1, 3, 2}
	for i, v := range correct {
		if v != ans[i] {
			t.Fatalf("中序遍历错误：%v √  %v ×", correct, ans)
		}
	}
}

func TestLengthOfLTS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	correctAns := 4
	if ans := lengthOfLIS(nums); ans != correctAns {
		t.Errorf("最长递增子序列（动态规划）错误 %d √  %d ×", correctAns, ans)
	}
	if ans := lengthOfLIS2(nums); ans != correctAns {
		t.Errorf("最长递增子序列（贪心+二分查找）错误 %d √  %d ×", correctAns, ans)
	}
}

func TestReorderList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	reorderList(head)
	correct := []int{1, 4, 2, 3}
	for _, val := range correct {
		if head == nil || val != head.Val {
			t.Fatal("error")
		}
		head = head.Next
	}
}

func TestRightSideView(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	correct := []int{1, 3, 4}
	ans := rightSideView(root)
	if len(ans) != len(correct) {
		t.Fatalf("%v √  %v ×", correct, ans)
	}
	for i, v := range correct {
		if v != ans[i] {
			t.Fatalf("%v √  %v ×", correct, ans)
		}
	}
}

func TestClimStairs(t *testing.T) {
	if ans := climbStairs(3); ans != 3 {
		t.Fatalf("n = %d, %d √  %d ×", 3, 3, ans)
	}
}

func TestGetKthFromEnd(t *testing.T) {
	head := createLinkedList([]int{1, 2, 3, 4, 5, 6})
	if node := getKthFromEnd(head, 2); node.Val != 5 {
		t.Fatal("error")
	}
}

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	if max := maxPathSum(root); max != 42 {
		t.Fatalf("%d √  %d ×", 42, max)
	}
}

func TestMySqrt(t *testing.T) {
	if ans := mySqrt(8); ans != 2 {
		t.Fatalf("%d √  %d ×", 2, ans)
	}
	if ans := mySqrt2(8); ans != 2 {
		t.Fatalf("%d √  %d ×", 2, ans)
	}
}

func TestMergeIntervals(t *testing.T) {
	ans := mergeIntervals([][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}})
	correct := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if len(ans) != len(correct) {
		t.Fatalf("%v √  %v ×", correct, ans)
	}
	for i, interval := range correct {
		ansInterval := ans[i]
		if ansInterval[0] != interval[0] || ansInterval[1] != interval[1] {
			t.Fatalf("%v √  %v ×", correct, ans)
		}
	}
}

func TestRemoveKthFromEnd(t *testing.T) {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 2)
	correct := []int{1, 2, 3, 5}
	for i := 0; head != nil; i++ {
		if head.Val != correct[i] {
			t.Fatal("error")
		}
		head = head.Next
	}
}
