package matrix01

import (
	"fmt"
	"testing"
)

func TestMatrix01(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	res := updateMatrix(mat)
	fmt.Println(res)
}
