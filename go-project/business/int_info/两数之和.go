package int_info

import (
	"fmt"
)

/*
​两数之和（Two Sum）

​问题描述​
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 ​两个​ 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
输入：nums = [2, 7, 11, 15], target = 9
输出：[0, 1]
解释：因为 nums[0] + nums[1] = 2 + 7 = 9，所以返回 [0, 1]。
*/
// n2平方
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// n的复杂度,用hashmap

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := hashMap[target-nums[i]]; ok {
			return []int{value, i} // 必须要先value，
		}
		hashMap[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(nums, target)) // 输出: [0, 1]

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(TwoSum(nums, target)) // 输出: [1, 2]

	nums = []int{3, 3}
	target = 6
	fmt.Println(TwoSum(nums, target)) // 输出: [0, 1]
}
