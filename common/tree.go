package common

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type TreeElem struct {
	Idx  int
	Node *TreeNode
}
