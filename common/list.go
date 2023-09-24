package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) *ListNode {
	curNode := head
	var nextNode, prevNode *ListNode
	for curNode != nil {
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	return prevNode
}

func LenLinkedList(head *ListNode) int {
	var res int
	curNode := head
	for curNode != nil {
		res++
		curNode = curNode.Next
	}
	return res
}
