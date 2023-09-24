package editdistance

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	if res := minDistance(word1, word2); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
