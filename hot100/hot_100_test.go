package hot100

import "testing"

func Test_1_twoSum(t *testing.T) {
	nums_list := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}
	target_list := []int{9, 6, 6}
	result_list := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}
	for index, nums := range nums_list {
		target := target_list[index]
		result := twoSum(nums, target)
		t.Logf("nums: %v, target: %d", nums, target)
		if result == nil || len(result) != 2 {
			t.Error("error size of result")
			return
		}
		expected_result := result_list[index]
		if !(result[0] == expected_result[0] &&
			result[1] == expected_result[1] ||
			result[0] == expected_result[1] &&
				result[1] == expected_result[0]) {
			t.Errorf("correct result is [0, 1] or [1, 0], return is %v", result)
		}
	}
}

func Test_2_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	target := 807
	lr := addTwoNumbers(l1, l2)
	sum := 0
	for rate := 1; lr != nil; lr = lr.Next {
		sum += lr.Val * rate
		rate *= 10
	}
	if sum != target {
		t.Errorf("expected %d, got %d", target, sum)
	}
}

func Test_3_lengthOfLongestSubstring(t *testing.T) {
	str_list := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		" ",
	}
	result_list := []int{3, 1, 3, 0, 1}

	for index, str := range str_list {
		result := lengthOfLongestSubstring(str)
		if result != result_list[index] {
			t.Errorf("%s length of lonhest sub-string is %d, got %d", str, result_list[index], result)
		}
	}
}
