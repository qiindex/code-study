package int_info

import (
	"fmt"
)

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/
func main位置() {
	nums := []int{5, 7, 7, 8, 8, 10}
	list := searchRange(nums, 8)
	fmt.Println(list)
}

func searchRange(nums []int, target int) []int {
	first, last := -1, -1
	left, right := 0, len(nums)-1
	ans := make([]int, 0)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			first = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	ans = append(ans, first)
	ans = append(ans, last)
	return ans
}
