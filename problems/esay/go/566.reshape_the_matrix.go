package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rows := len(nums)
	cols := len(nums[0])
	if r*c != rows*cols {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	count := 0
	for _, row := range nums {
		for _, num := range row {
			res[count/c][count%c] = num
			count++
		}
	}
	return res
}
