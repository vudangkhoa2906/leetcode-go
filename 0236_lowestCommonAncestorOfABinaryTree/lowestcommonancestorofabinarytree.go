package lowestcommonancestorofabinarytree

import "vudangkhoa2906-leetcode-go/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	var lowestCommonAncestorRec func(node, p, q *common.TreeNode) *common.TreeNode
	lowestCommonAncestorRec = func(node, p, q *common.TreeNode) *common.TreeNode {
		if node == nil {
			return nil
		}
		if node == p || node == q {
			return node
		}
		leftNode, rightNode := lowestCommonAncestorRec(node.Left, p, q), lowestCommonAncestorRec(node.Right, p, q)
		if leftNode == nil && rightNode == nil {
			return nil
		}
		if leftNode == nil {
			return rightNode
		}
		if rightNode == nil {
			return leftNode
		}
		return node
	}
	return lowestCommonAncestorRec(root, p, q)
}
