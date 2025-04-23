package link_info

/*
19. 删除链表的倒数第 N 个结点
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

*/
// 思路：快慢指针；
// 但是要注意，为了好走，建议fast指向head，slow指向dummy，就不用fast.Next.Next走了；（fast先走n-0步，也就是n步）
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0 {
		return head
	}
	var dummy *ListNode = &ListNode{Next: head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
