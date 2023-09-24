package findpeakelement

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2}
	expected := []int{1}
	if res := findPeakElement(nums); common.IndexOf(expected, res) == -1 {
		t.Errorf("Output: %d: Expected: %v\n", res, expected)
	}
}
