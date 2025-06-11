package link_algo

import "fmt"

// ListNode 定义单链表节点
/*type ListNode struct {
	Val  int
	Next *ListNode
}*/

// createList 根据整数数组创建单链表
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

// printList 打印链表中的所有节点值
func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func main3() {
	// 示例 1: 链表 1->2->3->4->5，翻转第 2 到第 4 个节点
	nums := []int{1, 2, 3, 4, 5}
	head := createList(nums)
	fmt.Print("原链表: ")
	printList(head)

	m, n := 2, 4
	newHead := reverseBetween1(head, m, n)
	fmt.Printf("翻转 [%d,%d] 后的链表: ", m, n)
	printList(newHead)

	// 示例 2: 链表 1->2->3->4->5，翻转第 1 到第 5 个节点（整个链表）
	nums = []int{1, 2, 3, 4, 5}
	head = createList(nums)
	fmt.Print("\n原链表: ")
	printList(head)

	m, n = 1, 5
	newHead = reverseBetween1(head, m, n)
	fmt.Printf("翻转 [%d,%d] 后的链表: ", m, n)
	printList(newHead)

	// 示例 3: 链表 1->2->3->4->5，翻转第 3 到第 3 个节点（无需翻转）
	nums = []int{1, 2, 3, 4, 5}
	head = createList(nums)
	fmt.Print("\n原链表: ")
	printList(head)

	m, n = 3, 3
	newHead = reverseBetween1(head, m, n)
	fmt.Printf("翻转 [%d,%d] 后的链表: ", m, n)
	printList(newHead)
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head, Val: 0}
	P0 := dummy
	for i := 1; i < left; i++ {
		//left-1
		P0 = P0.Next
	}
	current := P0.Next
	var pre *ListNode = nil
	for i := 0; i < right-left+1; i++ {
		// right-left+1 4-2+1=3
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	P0.Next.Next = current
	P0.Next = pre

	return dummy.Next
}

/*
原链表: 1 2 3 4 5
翻转 2 后的链表: 2 1 3 4 5
*/
func reverseBetweenK(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head, Val: 0}
	p0 := dummy
	n := 0
	current1 := head
	for current1 != nil {
		n += 1
		current1 = current1.Next
	}
	for n-k >= 0 {
		n = n - k
		var pre *ListNode = nil
		current := p0.Next
		fmt.Println(current.Val)
		for i := 0; i < k; i++ {
			next := current.Next
			current.Next = pre
			pre = current
			current = next
		}
		nxt := p0.Next //next是1
		p0.Next.Next = current
		p0.Next = pre
		p0 = nxt //p0 这里变成1
	}
	return dummy.Next
}
func main() {
	//
	list2 := createList([]int{1, 2, 3, 4, 5})
	reverse := reverseBetweenK(list2, 2)
	printList(reverse)
}
