package interviewtop

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) != 2 || result[0] != 0 || result[1] != 1 {
		t.Fatalf("expect [0, 1], got %v", result)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	rs := addTwoNumbers(l1, l2)
	correct := []int{7, 0, 8}
	result := []int{}
	for ; rs != nil; rs = rs.Next {
		result = append(result, rs.Val)

	}
	if len(correct) != len(result) {
		t.Fatalf("expect %v, got %v", correct, result)
	}
	for i, v := range result {
		if v != correct[i] {
			t.Fatalf("expect %v, got %v", correct, result)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	correct := 3
	rs := lengthOfLongestSubstring(s)
	if rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
}
