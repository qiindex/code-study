package link_algo

// 定义链表节点结构

// 翻转链表的函数
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// 创建一个虚拟头节点，方便处理头节点的翻转
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		// 检查是否有k个节点
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		// 记录k个节点的起始和结束位置
		next := tail.Next
		headNew, tailNew := reverse(head, tail)

		// 将翻转后的子链表连接到主链表中
		prev.Next = headNew
		tailNew.Next = next

		// 移动prev到下一次翻转的起始位置
		prev = tailNew
		head = tailNew.Next
	}

	return dummy.Next
}

// 翻转从head到tail的链表，并返回翻转后的头尾节点
func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head

	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return tail, head
}

// 辅助函数：打印链表
/*func main() {
	//
	list2 := link_algo.CreateList([]int{1, 2, 3, 4, 5})
	reverse := link_algo.Reverse(list2, 2)
	link_algo.PrintList(reverse)
}

*/
