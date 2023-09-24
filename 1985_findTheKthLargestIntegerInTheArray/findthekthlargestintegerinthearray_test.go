package findthekthlargestintegerinthearray

import "testing"

func TestFindTheKthLargestIntegerInTheArray(t *testing.T) {
	nums := []string{"3", "6", "7", "10"}
	k := 4
	if res := kthLargestNumber(nums, k); res != "3" {
		t.Errorf("Output: %s: Expected: %s\n", res, "3")
	}
}
