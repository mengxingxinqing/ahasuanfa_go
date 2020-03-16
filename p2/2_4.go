package p2

//递归 链表翻转
type Node struct {
	val  int
	next *Node
}

func revert(head *Node) *Node {
	if head.next == nil {
		return head
	}
	last := revert(head.next)
	head.next.next = head
	head.next = nil
	return last
}
