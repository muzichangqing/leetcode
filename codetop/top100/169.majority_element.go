package top100

func majorityElement(nums []int) int {
	maj := nums[0]
	counter := 0
	for _, v := range nums {
		if counter == 0 {
			maj = v
		}
		if maj == v {
			counter++
		} else {
			counter--
		}
	}
	return maj
}
