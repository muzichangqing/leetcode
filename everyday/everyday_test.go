package everyday

import "testing"

func TestDecodeTheMessage(t *testing.T) {
	var key, message, correct string

	key = "the quick brown fox jumps over the lazy dog"
	message = "vkbs bs t suepuv"
	correct = "this is a secret"
	if result := decodeMessage(key, message); result != correct {
		t.Fatalf("For key: %s, message: %s: expect \"%s\", got \"%s\"",
			key, message, correct, result)
	}
}

func TestShortestAlternatingPaths(t *testing.T) {
	t.Log(shortestAlternatingPaths(3, [][]int{{0, 1}, {0, 2}}, [][]int{}))
}

func TestMaximumConsecutive(t *testing.T) {
	var (
		coins   []int
		correct int
	)

	coins, correct = []int{1, 3}, 2
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
	coins, correct = []int{1, 4, 1, 1}, 8
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
	coins, correct = []int{1, 4, 10, 3, 1}, 20
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
}
