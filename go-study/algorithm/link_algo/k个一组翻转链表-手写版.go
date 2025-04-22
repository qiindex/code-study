package link_algo

/*func main() {
	list1 := link_algo.CreateList([]int{1, 2, 3, 4, 5})
	k := 2
	result := link_algo.ReverseKGroup(list1, k)
	link_algo.PrintList(result)
}
*/
// 思路：每K个元素一组；用一个小函数翻转，标记头，尾，用来穿起来
// 最后不满足k个的就不用翻转，直接返回

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next //如果是2，2的下一个就是3
	current := head
	//  1 2 3 4。 -----  2 1   3 4
	for pre != tail { //  ！2
		// 1 指向->3。  2指向 ->1
		next := current.Next //next 2，      3 ，
		current.Next = pre   // 1的下一个4    2的下一个是1，
		pre = current        // pre 是 1 ，   pre是2 ，
		current = next       // current是 2， current是3，
	}

	return tail, head

}
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		headNew, tailNew := reverse(head, tail)
		// 连接
		prev.Next = headNew
		// 下次循环
		prev = tailNew
		head = tailNew.Next
	}
	return dummy.Next
}

/*func reverseKNumber(head * ListNode,k int)	(*ListNode,*ListNode)  {
	current := head
	for

}
*/
