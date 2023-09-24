package kokoeatingbananas

import "testing"

func TestKokoEatingBananas(t *testing.T) {
	piles := []int{2, 2}
	h := 2
	if res := minEatingSpeed(piles, h); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
