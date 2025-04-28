package main

import (
	"go-project/algorithm/demo_design"
	"go-project/business/link_info"
)

func main() {
	// demo
	demo_design.Demo()
	//
	linkList := link_info.CreateList([]int{1, 2, 3, 4, 5})
	link_info.PrintList(linkList)
	/*reverseLink := link_info.RemoveNthFromEnd(linkList, 2)
	link_info.PrintList(reverseLink)*/
}
