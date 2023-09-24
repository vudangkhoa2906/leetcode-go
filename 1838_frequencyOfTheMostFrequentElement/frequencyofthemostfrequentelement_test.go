package frequencyofthemostfrequentelement

import "testing"

func TestFrequencyOfTheMostFrequentElement(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1}
	k := 2
	if res := maxFrequency(nums, k); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
