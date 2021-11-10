package sort

import (
	"math"
	"math/rand"
	"time"
)

// 插入排序
func IntertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// 希尔排序
func ShellSort(nums []int) {
	length := len(nums)
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h - 1; i < length; i++ {
			key := nums[i]
			j := i - h
			for j >= 0 && nums[j] > key {
				nums[j+h] = nums[j]
				j -= h
			}
			nums[j+h] = key
		}
		h /= 3
	}
}

// 归并排序
func MergeSort(nums []int) {
	myMergeSort(nums, 0, len(nums)-1)
}

func myMergeSort(nums []int, p, r int) {
	if p < r {
		q := (r + p) / 2
		myMergeSort(nums, p, q)
		myMergeSort(nums, q+1, r)
		merge(nums, p, q, r)
	}
}

func merge(nums []int, p, q, r int) {
	ln := q - p + 2 // 左侧容量+哨兵
	rn := r - q + 1 // 右侧容量+哨兵
	lpart := make([]int, ln)
	rpart := make([]int, rn)
	for i := 0; i < ln-1; i++ {
		lpart[i] = nums[p+i]
	}
	lpart[ln-1] = math.MaxInt32
	rpart[rn-1] = math.MaxInt32
	for i := 0; i < rn-1; i++ {
		rpart[i] = nums[q+1+i]
	}

	var i, j int
	for k := 0; k < r-p+1; k++ {
		if lpart[i] < rpart[j] {
			nums[p+k] = lpart[i]
			i++
		} else {
			nums[p+k] = rpart[j]
			j++
		}
	}
}

// 冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// 选择排序
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minI := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minI] {
				minI = j
			}
		}
		nums[i], nums[minI] = nums[minI], nums[i]
	}
}

// 快速排序
func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	myQuickSort(nums, 0, len(nums)-1)
}

func myQuickSort(nums []int, lo, hi int) {
	if lo < hi {
		mid := myQuickSortPartition(nums, lo, hi)
		myQuickSort(nums, lo, mid-1)
		myQuickSort(nums, mid+1, hi)
	}
}

func myQuickSortPartition(nums []int, lo, hi int) int {
	randIndex := rand.Intn(hi-lo) + lo
	key := nums[randIndex]
	nums[randIndex], nums[hi] = nums[hi], nums[randIndex]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] < key {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[hi] = nums[hi], nums[i+1]
	return i + 1
}
