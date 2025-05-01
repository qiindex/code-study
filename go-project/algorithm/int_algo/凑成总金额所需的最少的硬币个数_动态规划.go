package int_algo

/*
思路：
2. 动态规划思路
(1) 定义状态
dp[i] 表示凑成金额 i 所需的最少硬币数。
(2) 初始化
dp[0] = 0（凑成金额 0 不需要硬币）。
其他 dp[i] = amount + 1（初始化为一个不可能的大值，方便后续比较）。
(3) 状态转移方程
对于每个金额 i（从 1 到 amount），遍历所有硬币 coin：

如果 coin <= i，则：
dp[i] = min(dp[i], dp[i - coin] + 1)
解释：
dp[i - coin] 表示凑成 i - coin 所需的最少硬币数。
+1 表示再拿一枚 coin 面额的硬币。
取 min 是因为可能有多种硬币组合方式。
(4) 最终结果
如果 dp[amount] 仍然是 amount + 1，说明无法凑成，返回 -1。
否则，返回 dp[amount]。
*/

import "sort"

/*
贪心算法：

贪心算法在这里失效，因为 [4, 3, 1] 的选择 4 + 1 + 1 不是最优解（3 + 3 更优）。
​根本原因​：贪心算法假设“优先使用大面额硬币”总是最优，但某些情况下并非如此。
*/
func coinChange1(coins []int, amount int) int {
	sort.Ints(coins)
	count := 0

	for i := 0; i < len(coins)/2; i++ {
		coins[i], coins[len(coins)-i-1] = coins[len(coins)-i-1], coins[i]
	}
	for _, value := range coins {

		for amount >= value {
			amount = amount - value
			count++
		}
		if amount == 0 {
			return count
		}

	}

	return -1
}

// 动态规划

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = minInt(dp[i], dp[i-coin]+1)
				//min dp[i]取最小是循环中的硬币组合不唯一，是和贪心算法最大的差异；不是局部最优解；是全局最优解
				//dp[i-coin]是初了当前循环中的coin之外，之前累积的记忆化中最小的数量，再+1就是当前硬币算上的数量
			}
		}
	}
	// 4 >3+1
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
