package leetcode

import "testing"

func Test_combine_test(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	combines := combinationSum(candidates, target)
	t.Logf("rs: %v", combines)
}

func Test_combine(t *testing.T) {
	n := 1
	k := 1
	combines := combine(n, k)
	if len(combines) != 1 {
		t.Fatalf("Error combines count return %v", len(combines))
	}
	t.Log(combines)
}

func Test_top_k_frequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	topFrequentNums := topKFrequent(nums, k)
	if len(topFrequentNums) != 2 {
		t.Fatalf("Error element count return")
	}
	if !(topFrequentNums[0] == 1 && topFrequentNums[1] == 2 || topFrequentNums[0] == 2 && topFrequentNums[1] == 1) {
		t.Fatalf("Error element return")
	}
}

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
