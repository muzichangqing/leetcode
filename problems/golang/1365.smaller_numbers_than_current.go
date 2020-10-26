package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	cnt := [101]int{}

	for _, v := range nums {
		cnt[v]++
	}

	ans := make([]int, len(nums))
	for i := 1; i < 101; i++ {
		cnt[i] += cnt[i-1]
	}
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}
