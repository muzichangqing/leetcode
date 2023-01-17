package interviewtop

func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		h := height[i]
		if height[i] > height[j] {
			h = height[j]
		}
		area := h * (j - i)
		if area > max {
			max = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
