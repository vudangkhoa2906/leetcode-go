package wherewilltheballfall

import "testing"

func TestWhereWillTheBallFall(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, -1, -1},
		{1, 1, 1, -1, -1},
		{-1, -1, -1, 1, 1},
		{1, 1, 1, 1, -1},
		{-1, -1, -1, -1, -1},
	}
	res := findBall(grid)
	t.Logf("Output: %v\n", res)
}
