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
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return tail, head

}

// 翻转节点和后面k个
// 翻转后，节点后面是断开的
func reverseKNode(head *ListNode, k int) (*ListNode, *ListNode) {
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

//反转后，pre指向这一段的末尾，cur指向这一段的后续的第一个节点

/*
92. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
给你单链表的头指针 head 和两个整数 left 和 right
，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{Next: head}
	p0 := dummy
	//var pre *ListNode =nil
	for i := 1; i < left; i++ { // left-1
		p0 = p0.Next
	}
	var pre *ListNode = nil
	current := p0.Next
	for i := left; i <= right; i++ {
		nxt := current.Next
		current.Next = pre
		pre = current
		current = nxt

	}
	p0.Next.Next = current
	p0.Next = pre

	return dummy.Next
}

/* 反转后：
核心思路（茶同学的方法）
pre 指向这一段的末尾位置
current指向下一段的开始位置
*/
