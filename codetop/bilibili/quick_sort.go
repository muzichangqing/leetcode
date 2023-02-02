package bilibili

import (
	"math/rand"
	"time"
)

func quickSort(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	myQuickSort(nums, 0, len(nums)-1)
	return nums
}

func myQuickSort(nums []int, left, right int) {
	if left < right {
		mid := myQuickSortPartition(nums, left, right)
		myQuickSort(nums, left, mid-1)
		myQuickSort(nums, mid+1, right)
	}
}

func myQuickSortPartition(nums []int, left, right int) int {
	ri := rand.Intn(right-left+1) + left
	key := nums[ri]
	nums[ri], nums[right] = nums[right], nums[ri]
	i, j := left, left
	for ; j < right; j++ {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
