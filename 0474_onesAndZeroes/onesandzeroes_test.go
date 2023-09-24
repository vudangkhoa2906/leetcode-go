package onesandzeroes

import "testing"

func TestOnesAndZeroes(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 3, 4
	if res := findMaxForm(strs, m, n); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
