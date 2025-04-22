package main

import (
	"go-study/algorithm/demo_design"
	"go-study/algorithm/link_algo"
	"go-study/business/link_info"
)

func main() {
	// demo
	demo_design.Demo()
	//
	linkList := link_algo.CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	link_algo.PrintList(linkList)
	reverseLink := link_info.ReverseLink(linkList)
	link_algo.PrintList(reverseLink)
}
