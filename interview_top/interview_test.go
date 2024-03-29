package interviewtop

import (
	"math"
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

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2, 3, 6}
	nums2 := []int{1, 3, 4, 5, 9, 10}
	correct := 3.5

	rs := findMedianSortedArrays(nums1, nums2)
	if math.Abs(correct-rs) > 0.00001 {
		t.Fatalf("expect %v, got %v", correct, rs)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	correct = 2.5

	rs = findMedianSortedArrays(nums1, nums2)
	if math.Abs(correct-rs) > 0.00001 {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
}

func TestLongestPalindrome(t *testing.T) {
	s, correct := "babad", "bab"
	if rs := longestPalindrome(s); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
}

func TestReverseInt32(t *testing.T) {
	var x, correct int
	x, correct = 123, 321
	if rs := reverse(x); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
	x, correct = 0, 0
	if rs := reverse(x); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
	x, correct = -123, -321
	if rs := reverse(x); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
	x, correct = math.MaxInt32, 0
	if rs := reverse(x); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
	x, correct = math.MinInt32, 0
	if rs := reverse(x); rs != correct {
		t.Fatalf("expect %v, got %v", correct, rs)
	}
}

func TestMyAtoi(t *testing.T) {
	var (
		s       string
		correct int
	)
	s, correct = "42", 42
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "  +00042xyz", 42
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "  000-42xyz", 0
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "  -00042xyz", -42
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "21474836460", 2147483647
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "-2147483649", -2147483648
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
	s, correct = "", 0
	if rs := myAtoi(s); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
}

func TestIsMatch(t *testing.T) {
	var s, p string
	var correct bool
	s, p = "aa", "a"
	correct = false
	if rs := isMatch(s, p); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}

	s, p = "aa", "a*"
	correct = true
	if rs := isMatch(s, p); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}

	s, p = "ab", ".*"
	correct = true
	if rs := isMatch(s, p); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}

	s, p = "aab", "c*a*b"
	correct = true
	if rs := isMatch(s, p); rs != correct {
		t.Errorf("expect %v, got %v", correct, rs)
	}
}

func TestMaxArea(t *testing.T) {
	var (
		height  []int
		correct int
	)
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	correct = 49
	if rs := maxArea(height); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestIntRaise(t *testing.T) {
	var i8 int8 = 127
	t.Log(i8+1 == -128)
	t.Log(math.MaxInt32, math.MinInt32)
	t.Log(-11 % 10)
}

func TestRomanToInt(t *testing.T) {
	ss := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	ii := []int{3, 4, 9, 58, 1994}
	for i, v := range ss {
		if rs := romanToInt(v); rs != ii[i] {
			tErrorf(ii[i], rs, t)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	var (
		strs    []string
		correct string
	)

	strs = []string{"flower", "flow", "flight"}
	correct = "fl"
	if rs := longestCommonPrefix(strs); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestIsValid(t *testing.T) {
	var (
		s       string
		correct bool
	)

	s = "()[]{}"
	correct = true
	if rs := isValid(s); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestLetterCombinations(t *testing.T) {
	var (
		digits string
	)
	digits = "23"
	t.Log(letterCombinations(digits))
	digits = ""
	t.Log(letterCombinations(digits))
	digits = "2"
	t.Log(letterCombinations(digits))
}

func TestThreeSum(t *testing.T) {
	var (
		nums []int
	)
	nums = []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))
	nums = []int{0, 1, 1}
	t.Log(threeSum(nums))
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
	t.Log(generateParenthesis(1))
}

func TestStrStr(t *testing.T) {
	var haystack, needle string
	var correct int

	haystack, needle, correct = "mississippi", "issip", 4
	if rs := strStr(haystack, needle); rs != correct {
		tErrorf(correct, rs, t)
	}
	haystack, needle, correct = "leetcode", "leeto", -1
	if rs := strStr(haystack, needle); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestDivide(t *testing.T) {
	var dividend, divisor, correct int

	dividend, divisor, correct = 10, 3, 3
	if rs := divide(dividend, divisor); rs != correct {
		tErrorf(correct, rs, t)
	}

	dividend, divisor, correct = 7, -3, -2
	if rs := divide(dividend, divisor); rs != correct {
		tErrorf(correct, rs, t)
	}

	dividend, divisor, correct = -7, -3, 2
	if rs := divide(dividend, divisor); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestSearch(t *testing.T) {
	var (
		nums    []int
		target  int
		correct int
	)

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	correct = 4

	if result := search(nums, target); correct != result {
		tErrorf(correct, result, t)
	}
}

func TestSearchRange(t *testing.T) {
	var (
		nums    []int
		target  int
		correct []int
	)

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	correct = []int{3, 4}
	rs := searchRange(nums, target)
	if len(rs) != 2 || rs[0] != correct[0] || rs[1] != correct[1] {
		tErrorf(correct, rs, t)
	}
}

func TestCountAndSay(t *testing.T) {
	var (
		n       int
		correct string
	)

	n, correct = 1, "1"
	if rs := countAndSay(n); rs != correct {
		tErrorf(correct, rs, t)
	}
	n, correct = 4, "1211"
	if rs := countAndSay(n); rs != correct {
		tErrorf(correct, rs, t)
	}
}

func TestTest(t *testing.T) {
	removeNthFromEnd(nil, 0)
	mergeKLists(nil)
	isValidSudoku(nil)
}

func tErrorf(correct, rs interface{}, t *testing.T) {
	t.Errorf("expect %v, got %v", correct, rs)
}
