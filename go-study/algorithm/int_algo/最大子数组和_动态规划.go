package int_algo

/*
	最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

/*
思路：暴力枚举法 n2
*/
func MaxSubArray(nums []int) int {
	maxInt := 0
	for i := 0; i < len(nums); i++ {
		current := 0
		for j := i; j < len(nums); j++ {
			current += nums[j]
			maxInt = Max(current, maxInt)

		}
	}
	return maxInt
}

// 动态规划 O（n）
/*
理解 currentMax = max(nums[i], currentMax + nums[i]) 这一行代码是掌握 ​Kadane's Algorithm（卡达内算法）​​ 的关键。让我们逐步解析这行代码的含义，并通过示例帮助你更好地理解。

背景知识：Kadane's Algorithm
Kadane's Algorithm 是一种用于在数组中寻找最大子数组和的高效算法。其基本思想是通过遍历数组，动态地决定每个位置的最大子数组和，从而在 O(n) 的时间复杂度内解决问题。
这行代码的核心在于决定当前元素是否应该加入之前的子数组，或者从当前元素开始一个新的子数组。具体来说：

​**currentMax + nums[i]**​：
表示将当前元素 nums[i] 加入到之前的连续子数组中，形成一个新的子数组和。
如果之前的子数组和为正数，加上当前元素可能会使总和更大，因此保留这个操作是有意义的。
​**nums[i]**​：
表示从当前元素开始一个新的子数组，忽略之前的所有元素。
如果之前的子数组和为负数，加上当前元素反而会减小总和，因此选择从当前元素重新开始更有利。
​**max(nums[i], currentMax + nums[i])**​：
比较上述两种情况，选择使 currentMax 更大的那个值作为新的 currentMax。
这样可以确保 currentMax 始终记录的是截至当前位置的最大子数组和。
*/

func MaxSubArray2(nums []int) int {
	globalMax := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		currentMax = Max(nums[i], currentMax+nums[i])
		globalMax = Max(globalMax, currentMax)
	}

	return globalMax
}
