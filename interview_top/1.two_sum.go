package interviewtop

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		if j, ok := dict[target-v]; ok {
			return []int{j, i}
		}
		dict[v] = i
	}
	return nil
}
