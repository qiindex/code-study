package main

import (
	"fmt"

	"go-study/algorithm/demo_design"
	"go-study/business/int_info"
	"go-study/business/link_info"
)

func main1() {
	// demo
	demo_design.Demo()
	//
	linkList := link_info.CreateList([]int{1, 2, 3, 4, 5})
	link_info.PrintList(linkList)
	/*reverseLink := link_info.RemoveNthFromEnd(linkList, 2)
	link_info.PrintList(reverseLink)*/
}
func main() {
	// 示例输入
	nums := []int{-1, 0, 1, 2, -1, -4}

	// 调用 threeSum 函数
	result := int_info.ThreeSum(nums)

	// 打印结果
	fmt.Println("输入数组:", nums)
	fmt.Println("所有和为 0 的不重复三元组:")
	for _, triplet := range result {
		fmt.Println(triplet)
	}
}
