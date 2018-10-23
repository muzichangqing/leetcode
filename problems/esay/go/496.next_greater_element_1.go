package leetcode

func nextGreaterElement(findNums []int, nums []int) []int {
    res := make([]int, len(findNums))
    tmp_map := make(map[int]int)
    nums_size := len(nums)
    for fi := 0; fi < nums_size; fi++ {
        tmp_map[nums[fi]] = -1
        for si := fi + 1; si < nums_size; si++ {
            if nums[si] > nums[fi] {
                tmp_map[nums[fi]] = nums[si]
                break
            }
        }
    }
    for index, value := range(findNums) {
        res[index] = tmp_map[value]
    }
    return res
}
