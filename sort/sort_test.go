package sort

import "testing"

func TestIntertionSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试插入排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		IntertionSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestShellSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试希尔排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		ShellSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试归并排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		MergeSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试冒泡排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		BubbleSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试选择排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		SelectionSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试快速排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		QuickSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
	numsList := [][]int{
		{},
		{1},
		{31, 41, 59, 26, 41, 58},
	}
	t.Log("开始测试快速排序")
	for _, nums := range numsList {
		t.Logf("排序前：%v", nums)
		heapSort(nums)
		t.Logf("排序后：%v", nums)
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				t.Errorf("排序错误：nums[%d] > nums[%d]", i, i+1)
			}
		}
	}
}
