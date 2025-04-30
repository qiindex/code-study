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
