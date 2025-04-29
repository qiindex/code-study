package int_algo

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

// 跳台阶问题，如果一次跳N个台阶?，
// 一种是数学推导出来结果是2的（n-1)次方，一种是用dp，dp计算前面的累加和；建议第一种；
func CountNumberN(n int) int {
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
