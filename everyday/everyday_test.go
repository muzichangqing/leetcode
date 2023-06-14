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

func TestAlertNames(t *testing.T) {
	keyNames := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTimes := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}

	t.Log(alertNames(keyNames, keyTimes))
}

func TestNumTimesAllBlue(t *testing.T) {
	flips := []int{3, 2, 4, 1, 5}
	correct := 2
	if ans := numTimesAllBlue(flips); ans != correct {
		t.Errorf("for %v, expect %d, got %d", flips, correct, ans)
	}
}
