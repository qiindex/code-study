package string_algo

/*
定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
*/

//
/*
思路；
​最长公共子序列（Longest Common Subsequence, LCS）​​
​问题描述​
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列（LCS）的长度。如果不存在公共子序列，则返回 0。

​子序列定义​：

子序列是指在不改变字符相对顺序的情况下，删除某些字符（或不删除）后形成的新字符串。
例如：
"ace" 是 "abcde" 的子序列。
"aec" 不是 "abcde" 的子序列。
​示例​：

markdown
复制
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，其长度为 3。
​动态规划解法​
我们可以使用 ​动态规划（Dynamic Programming, DP）​​ 来解决这个问题。

​状态定义​
定义 dp[i][j] 表示 text1[0..i-1] 和 text2[0..j-1] 的最长公共子序列的长度。
i 表示 text1 的前 i 个字符（text1[0..i-1]）。
j 表示 text2 的前 j 个字符（text2[0..j-1]）。
​状态转移方程​
​**如果 text1[i-1] == text2[j-1]**​：
当前字符匹配，可以加入 LCS，因此：
dp[i][j]=dp[i−1][j−1]+1
​**如果 text1[i-1] != text2[j-1]**​：
当前字符不匹配，不能同时加入 LCS，因此：
dp[i][j]=max(dp[i−1][j],dp[i][j−1])
即取 text1 去掉当前字符 或 text2 去掉当前字符 的最大值。
​初始化​
dp[0][j] = 0：text1 为空时，LCS 长度为 0。
dp[i][0] = 0：text2 为空时，LCS 长度为 0。
​最终结果​
dp[m][n] 即为 text1 和 text2 的最长公共子序列的长度，其中：
m = len(text1)

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp { //i是index
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//dp[i-1][j]上边，      dp[i][j-1] 左边
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
