package string_algo

/*
用go写算法，简述解释；// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符


// 示例 1：

// 输入：word1 = "horse", word2 = "ros"
*/

/*
	思路：

这个问题是经典的 ​编辑距离（Edit Distance）​​ 问题，
也称为 ​Levenshtein 距离。我们需要计算将 word1 转换成 word2 所需的最少操作次数，操作包括插入、删除和替换字符。

状态转移方程​
​**如果 word1[i-1] == word2[j-1]**​
当前字符相同，不需要操作，直接继承 dp[i-1][j-1] 的值：
dp[i][j]=dp[i−1][j−1]
​**如果 word1[i-1] != word2[j-1]**​
可以进行 ​插入、删除或替换​ 操作，取三者中的最小值 + 1：
dp[i][j]=min(dp[i−1][j],dp[i][j−1],dp[i−1][j−1])+1
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 动态规划填充 dp 表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i][j]=min( dp[i−1][j],dp[i][j−1],dp[i−1][j−1]     )+1
				dp[i][j] = Min(dp[i-1][j],
					Min(dp[i][j-1],
						dp[i-1][j-1]),
				) + 1
			}
		}
	}

	return dp[m][n]
}

/*func main() {
	word1 := "horse"
	word2 := "ros"
	result := minTimes(word1, word2)
	fmt.Println(result) // 输出 3
}
*/

// min函数是三个不同的比较
func minTimes(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// 初始化 dp 表
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 边界条件：空字符串转换
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删除 i 次
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插入 j 次
	}

	// 填充 dp 表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 字符相同，无需操作
			} else {
				// 取 左（插入）、上（删除）、左上（替换）的最小值 + 1
				dp[i][j] = Min3(
					dp[i-1][j],   // 删除
					dp[i][j-1],   // 插入
					dp[i-1][j-1], // 替换
				) + 1
			}
		}
	}

	return dp[m][n]
}
