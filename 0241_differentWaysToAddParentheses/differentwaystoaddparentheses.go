package differentwaystoaddparentheses

import (
	"strconv"
	"strings"
	"vudangkhoa2906-leetcode-go/common"
)

func diffWaysToCompute(expression string) []int {
	lenExpression := len(expression)
	var operands []int
	var operators []byte
	var operandBuilder strings.Builder
	var operand int
	for idx := 0; idx < lenExpression; idx++ {
		if expression[idx] == '+' || expression[idx] == '-' || expression[idx] == '*' {
			operators = append(operators, expression[idx])
			operand, _ = strconv.Atoi(operandBuilder.String())
			operands = append(operands, operand)
			operandBuilder.Reset()
		} else {
			operandBuilder.WriteByte(expression[idx])
		}
	}
	operand, _ = strconv.Atoi(operandBuilder.String())
	operands = append(operands, operand)
	lenOperands := len(operands)
	dp := make([][][]int, lenOperands)
	for idx1 := range dp {
		dp[idx1] = make([][]int, lenOperands)
	}

	var diffWaysToComputeRec func(lowIdx, highIdx int) []int
	diffWaysToComputeRec = func(lowIdx, highIdx int) []int {
		if dp[lowIdx][highIdx] != nil {
			return dp[lowIdx][highIdx]
		}
		var res []int
		if lowIdx == highIdx {
			res = make([]int, 1)
			res[0] = operands[lowIdx]
			dp[lowIdx][highIdx] = res
			return res
		}

		res = make([]int, 0)
		for midIdx := lowIdx; midIdx < highIdx; midIdx++ {
			for _, lowOperand := range diffWaysToComputeRec(lowIdx, midIdx) {
				for _, highOperand := range diffWaysToComputeRec(midIdx+1, highIdx) {
					res = append(res, common.Compute(lowOperand, highOperand, operators[midIdx]))
				}
			}
		}
		return res
	}

	return diffWaysToComputeRec(0, lenOperands-1)
}
