package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	// 如果数组为空，直接返回 0
	if n == 0 {
		return 0
	}
	// 初始化最小长度为无穷大
	minLength := math.MaxInt32
	// 初始化左指针和当前窗口的和
	left, sum := 0, 0
	// 右指针从 0 开始遍历数组
	for right := 0; right < n; right++ {
		// 将当前元素加入窗口
		sum += nums[right]
		// 当窗口内元素的和大于等于 s 时，尝试缩小窗口
		for sum >= s {
			// 更新最小长度
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			// 缩小窗口，将左指针右移，并更新窗口内元素的和
			sum -= nums[left]
			left++
		}
	}
	// 如果最小长度仍然是无穷大，说明没有找到满足条件的子数组，返回 0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	result := minSubArrayLen(s, nums)
	fmt.Println(result)
}
