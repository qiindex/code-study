package link_algo

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
输入：head = [1,2,2,1]
输出：true

*/
// 思路，用数据先承接链表的数字，然后判断数组，估计不行，空间复杂度O(N)了；（时间复杂度O(n)）
func IsPalindromeV1(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	//fmt.Println(list)
	for i := 0; i < (len(list) / 2); i++ {
		if list[i] != list[len(list)-i-1] {
			return false
		}
	}
	return true

}

// 最优解：翻转链表，然后比较，双指针，O（1）的空间复杂度；
func IsPalindrome(head *ListNode) bool {
	//[1,2,3,2,1]
	//[1,2,2,1]
	//翻转n/2之后的链表，然后连接上
	count := 0
	dummy := head
	for dummy != nil {
		dummy = dummy.Next
		count++
	} // 1 2 3 1
	pre := &ListNode{Next: head}
	var reLink *ListNode
	for i := 0; i < count/2; i++ {
		pre = pre.Next
		reLink = pre.Next
	}

	// 翻转relink的首尾，返回新的首
	newHead := reverseListNode(reLink)
	pre.Next = newHead
	//  1. 2  1 3
	if pre != nil {
		pre = pre.Next
	}
	for i := 0; i < count/2; i++ {
		if pre != nil && head.Val != pre.Val {
			return false
		}
		if pre != nil {
			pre = pre.Next
		}
		head = head.Next
	}
	return true

}
func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

/*
func main() {
	// demo
	listLink := link_algo.CreateList([]int{1,2,2,1})
	fmt.Println(link_algo.IsPalindrome(listLink))
}

*/
