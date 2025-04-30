package int_algo

import (
	"fmt"
)

/*import (
	"fmt"
)

func main() {
	// demo

	fmt.Println(CountNumberNV2(3))
}
*/
// 跳台阶问题，一次跳一个，一次跳2个，跳N个台阶需要多少步？
//思路：（假设n>0,如果n=0，还需要增加逻辑直接返回0）
func CountNumber(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1 // 如果这里不写dp[0] =1,则需要判断n=0时，直接return 1
	dp[1] = 1
	dp[2] = 2
	if n <= 2 {
		return dp[n]
	}

	//dp[3] = 3  推导
	//dp[4] = 5
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}

// 跳台阶问题扩展2，如果一次可以跳N个台阶,那么跳到N个台阶，有多少种可能性?，
// 思路：
// 一种是数学推导出来结果是2的（n-1)次方，一种是用dp，dp计算前面的累加和；建议第一种；
func CountNumberNV1(n int) int {
	if n == 0 {
		// 跳到0个台阶，可以选择不跳
		return 1
	}
	return 1 << (n - 1)
}

// 使用dp计算一次可以跳N个台阶
func CountNumberNV2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	//dp[3] = 4 推导
	//dp[4] = 8
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j]
		}
	}
	return dp[n]
}

// 跳台阶问题扩展3，如果一次可以跳N个台阶,那么跳到M个台阶，有多少种可能性?，
/*
每次可以跳 1 到 n 个台阶，问跳到第 m 级台阶有多少种跳法？
*/
/*
思路：
递推关系​
对于 m ≥ n，每次可以跳 1 到 n 步，因此：

f(m)=f(m−1)+f(m−2)+⋯+f(m−n)
其中：

f(0) = 1（不跳）
f(1) = 1（跳 1 步）
f(2) = f(1) + f(0) = 2
f(3) = f(2) + f(1) + f(0) = 4
f(4) = f(3) + f(2) + f(1) + f(0) = 8

观察规律​：

当 m ≤ n 时，f(m) = 2^{m-1}（因为每次可以选择跳或不跳某些步）。
当 m > n 时，f(m) = 2 * f(m-1) - f(m-n-1)（但更简单的是直接递推）
*/
func CountNumberNMV1(n, m int) int {
	if m == 0 {
		return 1
	}
	dp := make([]int, m+1)
	dp[0] = 1 // 不跳也是一种方法
	for i := 1; i <= m; i++ {
		for j := 1; j <= n && j <= i; j++ {
			dp[i] += dp[i-j]
		}
	}
	return dp[m]
}

func main1() {
	fmt.Println(CountNumberNMV1(2, 3)) // 3 (1+1+1, 1+2, 2+1)
	fmt.Println(CountNumberNMV1(3, 4)) // 7 (1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2, 1+3, 3+1)
}

// 小于等于M时，可以用2（n-1）次方，大于m时，用dp,难理解
func CountNumberNMV2(n, m int) int {
	if m == 0 {
		return 1
	}
	if m <= n {
		return 1 << (m - 1) // 2^(m-1)
	}
	// 当 m > n 时，仍然需要递推
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n && j <= i; j++ {
			dp[i] += dp[i-j]
		}
	}
	return dp[m]
}
