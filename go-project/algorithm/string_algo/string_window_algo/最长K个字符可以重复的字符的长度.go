package string_window_algo

/*
	变形题：

这里计算的是最长K个字符可以重复的子串的长度，
所以map这里需要value变成int，统计字符重新的次数，每次判断value是否大于k，
这里主要的差异是，需要先添加 ，再判断；
*/
func LengthIfLongestSubStringOwnKReplaceString(s string, k int) int {
	hashMap := make(map[byte]int)
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		hashMap[s[right]]++
		for hashMap[s[right]] > k {
			hashMap[s[left]]--
			left++
		}
		if currentMaxLength := right - left + 1; currentMaxLength > maxLength {
			maxLength = currentMaxLength
		}
	}
	return maxLength

}
