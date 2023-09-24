package maximalrectangle

import "testing"

func TestMaximalRectangle(t *testing.T) {
	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}
	matrix := [][]byte{
		{'1', '0'},
		{'1', '0'},
	}

	if res := maximalRectangle(matrix); res != 6 {
		t.Errorf("Output: %d: Expected: %d\n", res, 6)
	}
}
