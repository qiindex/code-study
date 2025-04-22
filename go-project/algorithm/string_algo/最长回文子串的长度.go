package string_algo

// 最长回文子串的内容 longestPalindrome
// 或者最长回文子串的长度 （maxLen就是需要的长度）longestPalindromeLengthInt
/*
思路：中心扩展法制
*/
func longestPalindrome(s string) string {
	maxstart := 0
	maxLen := 0
	if len(s) < 2 {
		return s
	}
	for i := 0; i < len(s); i++ {
		a := FindPalindromeString(s, i, i)
		b := FindPalindromeString(s, i, i+1)
		maxLenNew := MaxLength(a, b)
		if maxLenNew > maxLen {
			maxLen = maxLenNew //如果是最长回文子串的长度，返回maxLen就可以了；
			//如果是最长回文子串的内容，需要返回maxstart和maxLen，然后根据maxstart和maxLen来返回最长回文子串的内容；
			maxstart = i - (maxLenNew-1)/2 //兼容偶数的情况，如果是奇数，可以是i-len/2，但是偶数不行；所以可以利用int的向下取整思路；
			//如果是奇数，那么maxstart就是i-len/2，如果是偶数，那么maxstart就是i-len/2+1，所以可以统一写成i-(maxLen-1)/2;
		}
	}
	if maxLen == 0 {
		return ""
	}

	return s[maxstart : maxstart+maxLen]
}

func FindPalindromeString(s string, start int, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start + 1 - 2

}

func MaxLength(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最长回文子串的长度
func longestPalindromeLengthInt(s string) int {
	maxLen := 0
	if len(s) < 2 {
		return len(s)
	}
	for i := 0; i < len(s); i++ {
		a := FindPalindromeString(s, i, i)
		b := FindPalindromeString(s, i, i+1)
		maxLenNew := MaxLength(a, b)
		if maxLenNew > maxLen {
			maxLen = maxLenNew
		}
	}

	return maxLen
}
