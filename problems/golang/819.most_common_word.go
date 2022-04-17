package leetcode

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	counts := make(map[string]int)
	wordStart := -1
	for i, b := range []byte(paragraph) {
		if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
			if wordStart < 0 {
				wordStart = i
			}
		} else {
			if wordStart >= 0 {
				counts[strings.ToLower(string(paragraph[wordStart:i]))] += 1
				wordStart = -1
			}
		}
	}
	if wordStart >= 0 {
		counts[strings.ToLower(string(paragraph[wordStart:]))] += 1
	}
	for _, s := range banned {
		counts[s] = 0
	}
	var rsWord string
	var maxCount int
	for k, v := range counts {
		if v > maxCount {
			rsWord = k
			maxCount = v
		}
	}
	return rsWord
}
