package partitionarrayintotwoarraystominimizesumdifference

import "testing"

func TestPartitionArrayIntoTwoArraysToMinimizeSumDifference(t *testing.T) {
	nums := []int{7772197, 4460211, -7641449, -8856364, 546755, -3673029, 527497, -9392076, 3130315, -5309187, -4781283, 5919119, 3093450, 1132720, 6380128, -3954678, -1651499, -7944388, -3056827, 1610628, 7711173, 6595873, 302974, 7656726, -2572679, 0, 2121026, -5743797, -8897395, -9699694}
	if res := minimumDifference(nums); res != 1 {
		t.Errorf("Output: %d: Expected: %d\n", res, 1)
	}
}
