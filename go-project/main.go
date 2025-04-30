package main

import (
	"fmt"

	"go-project/algorithm/link_algo"
)

func main() {
	// demo
	listLink := CreateListNode([]int{1, 2, 2, 1})
	fmt.Println(link_algo.IsPalindrome(listLink))
}

func CreateListNode(list []int) *link_algo.ListNode {
	head := &link_algo.ListNode{
		Val:  list[0],
		Next: nil,
	}
	current := head
	for i := 1; i < len(list); i++ {
		temp := &link_algo.ListNode{Val: list[i]}
		current.Next = temp
		current = current.Next
	}
	return head
}
