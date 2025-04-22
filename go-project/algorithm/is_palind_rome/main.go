package is_palind_rome

import (
	"strconv"
)

/*
回文数：
9. 回文数 https://leetcode.cn/problems/palindrome-number/description/
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
*/
func IsPalindromeFund(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	//1221
	//12321
	reverseNumber := 0
	for reverseNumber < x {
		reverseNumber = reverseNumber*10 + x%10
		x = x / 10
	}

	return reverseNumber == x || reverseNumber/10 == x

}

func IsPalindromeV2(x int) bool {
	s := strconv.Itoa(x)
	println(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
