package codetop

import (
	"math"
	"testing"
)

func TestDemo(t *testing.T) {
	var max int32
	max = math.MaxInt32
	t.Logf("%d, %T", max, max)
	max = max + 1
	t.Logf("%d, %T", max, max)
}

func TestMyAtio(t *testing.T) {
	items := map[string]int{
		"42":              42,
		"      -42":       -42,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
		"00000-42a1234":   0,
	}
	for key, val := range items {
		if ans := myAtoi(key); ans != val {
			t.Errorf("%s => %d ×", key, ans)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})
	correctAns := createLinkedList([]int{7, 0, 8})
	ans := addTwoNumbers(l1, l2)
	if !listNodeEqual(correctAns, ans) {
		t.Fatalf("AddTwoNumbers error")
	}
}

func TestDeleteDuplicates(t *testing.T) {
	head := createLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	correctAns := createLinkedList([]int{1, 2, 5})
	ans := deleteDuplicates(head)
	if !listNodeEqual(correctAns, ans) {
		t.Fatalf("DeleteDuplicates error")
	}
}

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	correctAns := []int{1, 2, 3}
	ans := preorderTraversal(root)
	if !sliceEqual(correctAns, ans) {
		t.Fatalf("%v √  %v ×", correctAns, ans)
	}
}

func TestSortList(t *testing.T) {
	head := createLinkedList([]int{4, 2, 1, 3})
	correctAns := createLinkedList([]int{1, 2, 3, 4})
	if ans := sortList(head); !listNodeEqual(correctAns, ans) {
		t.Fatalf("SortList err")
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	cAns := []int{1, 3, 2}
	nextPermutation(nums)
	if !sliceEqual(nums, cAns) {
		t.Fatalf("%v √  %v ×", cAns, nums)
	}
}

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	st := "ABC"
	cAns := "BANC"
	if ans := minWindow(s, st); ans != cAns {
		t.Fatalf("%s √  %s ×", cAns, ans)
	}
}

func TestReverseWords(t *testing.T) {
	s := "   Bob   Loves   Alice   "
	cAns := "Alice Loves Bob"
	if ans := reverseWords(s); ans != cAns {
		t.Fatalf("%s √  %s ×", cAns, ans)
	}
}

func TestMinDistance(t *testing.T) {
	if ans := minDistance("horse", "ros"); ans != 3 {
		t.Fatalf("%d √  %d ×", 3, ans)
	}
}

func TestBuildTree(t *testing.T) {
	buildTree([]int{3, 9, 20, 15, 7},
		[]int{3, 9, 20, 15, 7})
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := 4.0
	if rs := findMedianSortedArrays(nums1, nums2); math.Abs(rs-ans) > 1.0e-5 {
		t.Fatalf("%.2f √  %.2f ×", ans, rs)
	}
}

func TestMaxDepth(t *testing.T) {
	maxDepth(nil)
}
