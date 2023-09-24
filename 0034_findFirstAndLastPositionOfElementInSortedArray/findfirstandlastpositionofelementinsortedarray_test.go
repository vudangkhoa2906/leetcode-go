package findfirstandlastpositionofelementinsortedarray

import (
	"testing"
)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	t.Logf("Output: %v\n", searchRange(nums, target))
}
