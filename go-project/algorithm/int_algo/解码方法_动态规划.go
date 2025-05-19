package int_algo

import (
	"strconv"
)

func NumDecoding(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	if s[0] == '0' {
		dp[1] = 0 // 如果第一个字符是 '0'，则无法解码
	}

	for i := 2; i <= n; i++ {
		// 1-9
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		nums, _ := strconv.Atoi(s[i-2 : i])
		if nums >= 10 && nums <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
