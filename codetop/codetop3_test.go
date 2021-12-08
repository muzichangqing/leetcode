package codetop

import (
	"strings"
	"testing"
)

func TestCodeTop(t *testing.T) {
	isSymmetric(nil)
	isSymmetricT(nil)
	ans := generateParenthesis(3)
	t.Log(ans)
	majorityElement([]int{3, 2, 3})
	t.Log(multiply("123", "456"))
	invertTree(nil)
	searchMatrix([][]int{}, 0)
	maximalSquare([][]byte{})
	findPeakElement([]int{})
	longestCommonPrefix([]string{})
	maxAreaOfIsland([][]int{})
	largestNumber([]int{})
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

func TestFindMin(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := 0
	if ans := findMin(nums); ans != res {
		t.Errorf("expected %d, got %d", res, ans)
	}
}

func TestSortArray(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		sortArray(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestLongestConsecutive(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := 9
	if ans := longestConsecutive(nums); ans != res {
		t.Errorf("expected %d, got %d", res, ans)
	}
}

func TestUniquePaths(t *testing.T) {
	m, n := 3, 7
	res := 28
	if ans := uniquePaths(m, n); ans != res {
		t.Errorf("expected %d, got %d", res, ans)
	}
}

func TestValidIpAddress(t *testing.T) {
	ip := "2001:db8:85a3:0::8a2E:0370:7334"
	res := "Neither"
	if ans := validIPAddress(ip); ans != res {
		t.Errorf("expected %s, got %s", res, ans)
	}
}

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 2, 1, 2, 1}
	res := 4
	if ans := singleNumber(nums); ans != res {
		t.Errorf("expected %d, got %d", res, ans)
	}
}

func TestStringCompare(t *testing.T) {
	s1 := "3340"
	s2 := "334"
	t.Log(strings.Compare(s1, s2))
}
