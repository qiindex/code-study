package int_algo

/*
*
300. 最长递增子序列
已解答
中等
相关标签
相关企业
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：

输入：nums = [10,9,2,5,3,7,101,18]

输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
*/

/*
思路
定义 dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度。
初始化时，每个元素自身可以看作是一个长度为 1 的子序列，因此 dp 数组初始化为全 1。
对于每个 i，遍历其之前的所有元素 j（0 <= j < i），如果 nums[j] < nums[i]，则更新 dp[i] = max(dp[i], dp[j] + 1)。
最后，取 dp 数组中的最大值即为最长递增子序列的长度。

*/
// todo
// lengthOfLIS 使用动态规划求解最长递增子序列的长度
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // 每个元素自身可以作为一个子序列
	}

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

/*func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("最长递增子序列的长度是:", lengthOfLIS(nums)) // 输出: 4
}*/

/*
 思路：
 解题思路：
 状态定义：

 dp[i] 的值代表 nums 以 nums[i] 结尾的最长子序列长度。
 转移方程： 设 j∈[0,i)，考虑每轮计算新 dp[i] 时，遍历 [0,i) 列表区间，做以下判断：

 当 nums[i]>nums[j] 时： nums[i] 可以接在 nums[j] 之后（此题要求严格递增），此情况下最长上升子序列长度为 dp[j]+1 ；
 当 nums[i]<=nums[j] 时： nums[i] 无法接在 nums[j] 之后，此情况上升子序列不成立，跳过。
 上述所有 1. 情况 下计算出的 dp[j]+1 的最大值，为直到 i 的最长上升子序列长度（即 dp[i] ）。实现方式为遍历 j 时，每轮执行 dp[i]=max(dp[i],dp[j]+1)。
 转移方程： dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)。
 初始状态：

 dp[i] 所有元素置 1，含义是每个元素都至少可以单独成为子序列，此时长度都为 1。
 返回值：

 返回 dp 列表最大值，即可得到全局最长上升子序列长度。

*/
