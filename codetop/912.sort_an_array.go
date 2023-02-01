package codetop

import (
	"math/rand"
	"time"
)

func sortArrayV2(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := quickSortPartition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func quickSortPartition(nums []int, left, right int) int {
	ri := rand.Intn(right-left) + left
	nums[ri], nums[right] = nums[right], nums[ri]
	key := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] < key {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
