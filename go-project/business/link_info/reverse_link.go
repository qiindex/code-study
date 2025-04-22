package link_info

import (
	"go-study/algorithm/link_algo"
)

func ReverseLink(linkList *link_algo.ListNode) *link_algo.ListNode {
	var prev *link_algo.ListNode
	current := linkList
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
