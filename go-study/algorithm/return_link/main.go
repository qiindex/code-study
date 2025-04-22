package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintList 打印链表
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// 创建链表
func CreateList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}

// ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next // 保存下一个节点
		current.Next = prev  // 反转当前节点的指针
		prev = current       // 移动 prev 和 current
		current = next
	}

	return prev // prev 将是新的头节点
}

func main() {
	// 创建链表 1 -> 2 -> 3 -> 4 -> 5
	values := []int{1, 2, 3, 4, 5}
	head := CreateList(values)

	fmt.Println("Original list:")
	PrintList(head)

	// 反转链表
	reversedHead := ReverseList(head)

	fmt.Println("Reversed list:")
	PrintList(reversedHead)
}
