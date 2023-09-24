package maximumwidthofbinarytree

import (
	"vudangkhoa2906-leetcode-go/common"
)

func widthOfBinaryTree(root *common.TreeNode) int {
	res := 1
	queue := []common.TreeElem{common.TreeElem{0, root}}
	for count := len(queue); count > 0; count = len(queue) {
		if count > 1 {
			if tempRes := queue[count-1].Idx - queue[0].Idx + 1; tempRes > res {
				res = tempRes
			}
		}
		for c := 0; c < count; c++ {
			if queue[c].Node.Left != nil {
				queue = append(queue, common.TreeElem{2*queue[c].Idx + 1, queue[c].Node.Left})
			}
			if queue[c].Node.Right != nil {
				queue = append(queue, common.TreeElem{2*queue[c].Idx + 2, queue[c].Node.Right})
			}
		}
		queue = queue[count:]
	}
	return res
}
