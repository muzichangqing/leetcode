package top100

import "sort"

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	searchSpeed := sort.Search(max-1, func(speed int) bool {
		speed += 1
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}
		return time <= h
	})
	return searchSpeed + 1
}
