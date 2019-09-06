package leetcode

func distributeCandies(candies []int) int {
	maxLength := len(candies) / 2
	diffCandies := make([]int, maxLength)
	hasZero := false
	length := 0
	for _, candie := range candies {
		if length >= maxLength {
			return length
		}
		if candie == 0 {
			hasZero = true
		} else {
			for index, diffCandie := range diffCandies {
				if diffCandie == 0 {
					diffCandies[index] = candie
					length++
					break
				}
				if candie == diffCandie {
					break
				}
			}
		}
	}
	if hasZero {
		length++
	}
	if length > maxLength {
		return maxLength
	}
	return length
}
