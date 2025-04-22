package link_algo

/*

206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表

*/

/*
	type ListNode struct {
	    Val int
	    Next *ListNode
	}
*/

/*
思路：循环遍历，标记当前节点及其next；
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

func reverseListBetweenNode(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	current := head
	for prev != tail {
		nextTemp := prev.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return tail, head

}
