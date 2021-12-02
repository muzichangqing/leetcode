package codetop

import "testing"

func TestCodeTop(t *testing.T) {
	isSymmetric(nil)
	isSymmetricT(nil)
	ans := generateParenthesis(3)
	t.Log(ans)
	majorityElement([]int{3, 2, 3})
	t.Log(multiply("123", "456"))
	invertTree(nil)
}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	if ans := minPathSum(grid); ans != 7 {
		t.Errorf("expected: %d  got: %d", 7, ans)
	}
}

func TestCompareVersion(t *testing.T) {
	if ans := compareVersion("1.0", "1.0.0"); ans != 0 {
		t.Errorf("expected: %d  got: %d", 0, ans)
	}
	if ans := compareVersion("1.0", "1.0.1"); ans != -1 {
		t.Errorf("expected: %d  got: %d", -1, ans)
	}
	if ans := compareVersion("1.1", "1.0.1"); ans != 1 {
		t.Errorf("expected: %d  got: %d", 1, ans)
	}
}

func TestSearchRange(t *testing.T) {
	if ans := searchRange([]int{5, 7, 7, 8, 8, 10}, 8); ans[0] != 3 || ans[1] != 4 {
		t.Errorf("expected: %v  got: %v", []int{3, 4}, ans)
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	head := createLinkedList([]int{1, 1, 2, 3, 3})
	if ans := deleteDuplicates2(head); !listNodeEqual(ans, createLinkedList([]int{1, 2, 3})) {
		t.Error("error")
	}
}

func TestCombinationSum(t *testing.T) {
	t.Log(combinationSum([]int{2, 6, 3, 7}, 7))
}
