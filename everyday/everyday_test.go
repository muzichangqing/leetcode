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
