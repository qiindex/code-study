package link_algo

/*
	K 个一组翻转链表

输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
*/

func Reverse(head *ListNode, k int) *ListNode {
	// 异常处理 ，k，head
	if head == nil || k < 1 {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				//没有就提前返回，例如刚才的7-8，不用翻转
				return dummy.Next
			}
		}
		// 保留现场
		//nextGroup := tail.Next

		// 用小函数翻转制定范围的链表//  2 1 3 4 5       tail:  1 3 4 5
		headNew, tailNew := reverseK(head, tail)
		// 子链表接上
		pre.Next = headNew
		//tail.Next = nextGroup

		//移动pre ，head都下一组K个元素的起始位置
		pre = tailNew
		head = tailNew.Next
	}
	return dummy.Next
}

/*
思路:

	记录k个翻转的一组节点的翻转位置，（方便翻转函数翻转后，把链表穿起来）
	有一个翻转函数，每次传入翻转起始点，翻转后返回node

移动到下一组k个元素继续处理，到头  1 2 3 4    2 1 3 4
*/
func reverseK(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail { // abc
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}

func main1() {
	//
	list2 := CreateList([]int{1, 2, 3, 4, 5})
	reverse := Reverse(list2, 2)
	PrintList(reverse)
}
