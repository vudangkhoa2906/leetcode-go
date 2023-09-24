package matrix01

import (
	"testing"
)

func TestMatrix01(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	t.Logf("Output: %v\n", updateMatrix(mat))
}
