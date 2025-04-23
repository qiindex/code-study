package link_info

func ReverseLink(linkList *ListNode) *ListNode {
	var prev *ListNode
	current := linkList
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
