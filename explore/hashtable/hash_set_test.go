package hashtable

import (
	"math/rand"
	"testing"
)

func Test_MyHashSet(t *testing.T) {
	myHashSet := Constructor()
	for i := 0; i < 10000; i++ {
		randNumber := rand.Intn(1000000)
		myHashSet.Add(randNumber)
	}
	myHashSet.Add(270918)
	if !myHashSet.Contains(270918) {
		t.Fatalf("Error Add or Contains")
	}
	myHashSet.Remove(270918)
	if myHashSet.Contains(270918) {
		t.Fatalf("Error Remove or Contains")
	}
}
