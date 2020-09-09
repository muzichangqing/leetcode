package leetcode

var combinations [][]int

func combinationSum(candidates []int, target int) [][]int {
	combinations = [][]int{}
	path := []int{}
	backtraceFroCombinationSum(0, candidates, path, 0, target)
	return combinations
}

func backtraceFroCombinationSum(begin int, candidates, path []int, sum, target int) {
	if sum > target {
		return
	} else if sum == target {
		combination := make([]int, len(path))
		for index, num := range path {
			combination[index] = num
		}
		combinations = append(combinations, combination)
		return
	}
	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtraceFroCombinationSum(i, candidates, path, sum+candidates[i], target)
		path = path[:len(path)-1]
	}
}
