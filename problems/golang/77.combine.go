package leetcode

var combines [][]int

func combine(n int, k int) [][]int {
	combines = make([][]int, 0)
	if k <= 0 || k > n {
		return combines
	}
	path := make([]int, k)
	backtraceForCombine(n, k, 1, 0, path)
	return combines
}

func backtraceForCombine(n, k, begine, pathIndex int, path []int) {
	if pathIndex == k {
		copyToCombine(pathIndex, path)
		return
	}
	for i := begine; i <= n; i++ {
		path[pathIndex] = i
		backtraceForCombine(n, k, i+1, pathIndex+1, path)
	}
}

func copyToCombine(size int, path []int) {
	element := make([]int, size)
	for i := 0; i < size; i++ {
		element[i] = path[i]
	}
	combines = append(combines, element)
}
