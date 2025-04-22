package link_algo

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 公共函数
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}
