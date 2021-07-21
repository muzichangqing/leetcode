package hot100

import "testing"

func Test_1_two_sum(t *testing.T) {
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
