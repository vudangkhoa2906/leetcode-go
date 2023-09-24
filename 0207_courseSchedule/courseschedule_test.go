package courseschedule

import "testing"

func TestCourseSchedule(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	if res := canFinish(numCourses, prerequisites); !res {
		t.Errorf("Output: %t: Expected: true\n", res)
	}
}
