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

// 翻转两个节点之间
// 翻转后，节点后面不是断开的，不要追加在原来的链表上，就可以直接用了
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

// 翻转节点和后面k个
// 翻转后，节点后面是断开的
func reverseKNode(head, tail *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode = nil
	current := head

	for i := 0; i < k; i++ {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre, head

}
