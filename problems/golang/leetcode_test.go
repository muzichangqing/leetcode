package leetcode

import "testing"

func TestMostCommonWord(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	correctResult := "ball"

	if rs := mostCommonWord(paragraph, banned); rs != correctResult {
		t.Fatalf("expected \"%s\", got \"%s\"", correctResult, rs)
	}
}
