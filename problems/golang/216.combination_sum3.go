package leetcode

var combinations3 [][]int

func combinationSum3(k int, n int) [][]int {
	combinations3 = [][]int{}
	backtraceForCombinationSum3(k, 1, n, []int{})
	return combinations3
}

func backtraceForCombinationSum3(k, begine, rest int, path []int) {
	if len(path) == k {
		if rest == 0 {
			combination := make([]int, k)
			for index, element := range path {
				combination[index] = element
			}
			combinations3 = append(combinations3, combination)
		}
		return
	}
	for i := begine; i <= 9; i++ {
		path = append(path, i)
		backtraceForCombinationSum3(k, i+1, rest-i, path)
		path = path[:len(path)-1]
	}
}
