package differentwaystoaddparentheses

import "testing"

func TestDifferentWaysToAddParentheses(t *testing.T) {
	expression := "2*3-4*5"
	res := diffWaysToCompute(expression)
	t.Logf("Output: %v\n", res)
}
