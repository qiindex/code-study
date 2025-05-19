package main

import (
	"go-project/algorithm/link_algo"
)

var a int

func main() {
	//
	list2 := link_algo.CreateList([]int{1, 2, 3, 4, 5})
	reverse := link_algo.Reverse(list2, 2)
	link_algo.PrintList(reverse)
}
