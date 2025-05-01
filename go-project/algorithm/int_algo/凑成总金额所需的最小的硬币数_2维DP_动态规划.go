package int_algo

import (
	"fmt"
	"math"
)

/*
2dp：
二维 DP 思路
定义状态
● dp[i][j] 表示使用前 i 种硬币凑出金额 j 所需的最少硬币数。
状态转移方程
对于第 i 种硬币（面额为 coin = coins[i-1]），我们有两种选择：
1. 不选第 i 种硬币：
  ○ 此时凑出金额 j 的最少硬币数与使用前 i-1 种硬币凑出 j 相同：dp[i][j]=dp[i−1][j]
2. 选第 i 种硬币（前提是 coin <= j）：
  ○ 此时凑出金额 j 的最少硬币数为：dp[i][j]=dp[i][j−coin]+1
  ○ 即：用前 i 种硬币凑出 j - coin 的最少硬币数，再加上当前硬币 1 枚。
最终，dp[i][j] 取上述两种情况的最小值：
dp[i][j]=min(dp[i−1][j],dp[i][j−coin]+1)if coin≤j
*/

func coinChangeTwoDP(coins []int, amount int) int {
	n := len(coins)
	if n == 0 {
		if amount == 0 {
			return 0
		}
		return -1
	}

	// dp[i][j] 表示用前 i 种硬币凑出金额 j 的最少硬币数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32 // 初始化为一个较大的值（表示不可达）
		}
	}

	// 凑出金额 0 不需要硬币
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	// 动态规划填充 dp 表
	for i := 1; i <= n; i++ {
		coin := coins[i-1]
		for j := 1; j <= amount; j++ {
			if coin > j {
				// 当前硬币面额大于目标金额，不能选
				dp[i][j] = dp[i-1][j]
			} else {
				// 可以选择当前硬币或不选
				dp[i][j] = MinInt(dp[i-1][j], dp[i][j-coin]+1)
			}
		}
	}

	if dp[n][amount] == math.MaxInt32 {
		return -1 // 无法凑出
	}
	return dp[n][amount]
}

func main() {
	// 示例 1
	coins1 := []int{1, 2, 5}
	amount1 := 6
	fmt.Printf("示例 1: coins = %v, amount = %d -> 最少硬币数: %d\n", coins1, amount1, coinChange(coins1, amount1))

	// 示例 2
	coins2 := []int{2}
	amount2 := 3
	fmt.Printf("示例 2: coins = %v, amount = %d -> 最少硬币数: %d\n", coins2, amount2, coinChange(coins2, amount2))

	// 示例 3
	coins3 := []int{1}
	amount3 := 0
	fmt.Printf("示例 3: coins = %v, amount = %d -> 最少硬币数: %d\n", coins3, amount3, coinChange(coins3, amount3))

	// 示例 4
	coins4 := []int{1}
	amount4 := 2
	fmt.Printf("示例 4: coins = %v, amount = %d -> 最少硬币数: %d\n", coins4, amount4, coinChange(coins4, amount4))

	// 示例 5
	coins5 := []int{1, 3, 4}
	amount5 := 6
	fmt.Printf("示例 5: coins = %v, amount = %d -> 最少硬币数: %d\n", coins5, amount5, coinChange(coins5, amount5))
}
