package leetcode

import "testing"

func Test_remove_duplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	newSize := removeDuplicates(nums[:])
	if newSize != 5 {
		t.Fatalf("error size return")
	}
	for i := 1; i < newSize; i++ {
		if nums[i] <= nums[i-1] {
			t.Fatalf("error sort or duplicate values")
		}
	}
}

func Test_merge_two_sorted_lists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	result := mergeTwoLists(list1, list2)
	if result == nil {
		t.Fatalf("mergeTwoList shouldn't return nil")
	}
	cursor := result
	count := 1
	for cursor.Next != nil {
		if cursor.Val > cursor.Next.Val {
			t.Fatalf("the lists has not sorted")
		}
		cursor = cursor.Next
		count++
	}
	if count != 6 {
		t.Fatalf("error number of nodes")
	}

	_ = mergeTwoLists(nil, nil)
}
