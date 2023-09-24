package shortestcommonsupersequence

import (
	"testing"
)

func TestShortestCommonSupersequence(t *testing.T) {
	str1, str2 := "abac", "cab"
	t.Log(shortestCommonSupersequence(str1, str2))
}
