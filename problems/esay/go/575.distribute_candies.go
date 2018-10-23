package leetcode

func distributeCandies(candies []int) int {
	max_length := len(candies) / 2
	diff_candies := make([]int, max_length)
	has_zero := false
	length := 0
	for _, candie := range candies {
		if length >= max_length {
			return length
		}
		if candie == 0 {
			has_zero = true
		} else {
			for index, diff_candie := range diff_candies {
				if diff_candie == 0 {
					diff_candies[index] = candie
					length++
					break
				}
				if candie == diff_candie {
					break
				}
			}
		}
	}
	if has_zero {
		length++
	}
	if length > max_length {
		return max_length
	}
	return length
}
