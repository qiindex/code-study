package link_algo

/*
要合并 K 个升序链表，可以使用分治法或者最小堆（优先队列）来实现。这里我将展示如何使用分治法来合并 K 个升序链表，并提供相应的 Go 语言代码示例。

分治法思路
​分治策略​：将 K 个链表两两合并，直到只剩下一个链表。这个过程类似于归并排序中的合并步骤。
​递归合并​：每次将链表数组分成两部分，分别递归地合并这两部分，然后再将两个结果合并。
​时间复杂度​：O(N log K)，其中 N 是所有链表中的节点总数，K 是链表的数量。
*/
func mergeTwoListsV1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}
