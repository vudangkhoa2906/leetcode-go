package maxconsecutiveones

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	if res := findMaxConsecutiveOnes(nums); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
