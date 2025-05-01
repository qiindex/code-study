package int_algo

import (
	"fmt"
)

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数
子数组: 是指数组中连续的元素序列。
*/

// 思路，暴力循环遍历，

// v2 前缀和记录+hash表
/*
关键点：
子数组和公式：sum(nums[i..j]) = prefixSum[j+1] - prefixSum[i]。
哈希表存储 prefixSum 的出现次数，避免重复计算

子数组和的推导公式：
子数组和公式
子数组 nums[i..j] 的和：
sum(nums[i..j]) = nums[i] + nums[i+1] + ... + nums[j]
如何用前缀和计算？
prefixSum[j+1] 表示 nums[0..j] 的和（即 nums[0] + nums[1] + ... + nums[j]）。
prefixSum[i] 表示 nums[0..i-1] 的和（即 nums[0] + nums[1] + ... + nums[i-1]）。
因此：
sum(nums[i..j])=prefixSum[j+1]−prefixSum[i]
因为 prefixSum[j+1] 包含了 nums[0..j]，而 prefixSum[i] 包含了 nums[0..i-1]，所以相减后剩下的就是 nums[i..j]。
*/

func main() {
	// demo

	fmt.Println(subarraySum([]int{1, 2, 3}, 3))

}

/*

pre - k 的意义：这个检查的意义在于，
如果 pre - k 存在于 Map 中，说明之前在某个点的累积和是 pre - k。
由于当前的累积和是 pre，这意味着从那个点到当前点的子数组之和恰好是 k
（因为 pre - (pre - k) = k）

*/

// 前缀和最优解
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	sumCount := make(map[int]int)
	sumCount[0] = 1 // 初始状态：前缀和为0出现1次

	for _, num := range nums {
		prefixSum += num
		if val, exists := sumCount[prefixSum-k]; exists {
			count += val
		}
		sumCount[prefixSum]++
	}
	return count
}

// 思路，暴力循环遍历，
func subarraySumV1(nums []int, k int) int {
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
