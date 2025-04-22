package 滑动窗口

// 最长无重复字符的最长子串
/*
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

模版V1
//外层循环扩展右边界，内层循环扩展左边界
for (int l = 0, r = 0 ; r < n ; r++) {
	//当前考虑的元素
	while (l <= r && check()) {//区间[left,right]不符合题意
        //扩展左边界
    }
    //区间[left,right]符合题意，统计相关信息
}
百度二面

*/

// 元宝提供的，左边窗口直接跑到重复元素的下一个
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // 存储字符最后出现的位置
	maxLength := 0
	left := 0 // 窗口左边界

	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		// 如果字符已存在，且位置在窗口内，移动左边界
		if lastPos, exists := charIndex[currentChar]; exists && lastPos >= left {
			left = lastPos + 1
		}
		// 更新字符位置
		charIndex[currentChar] = right
		// 计算当前窗口长度
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

//自己手写的，用模版套写
/*
思路：滑动窗口计算长度，左右指针移动
关键步骤：用map记录是否存在，判断窗口的内容是否重复，key是字符，value是bool
 遍历string，先判断是否存在，如果存在，则删除map中的元素，移动左边指针，直到不存在
然后再添加元素;
计算长度，取最大值
*/
func LengthOfLongestSubStringByManually(s string) int {
	maxLength, left := 0, 0
	window := make(map[byte]bool)
	for right := 0; right < len(s); right++ {
		for window[s[right]] {
			//window[left]==false
			delete(window, s[left])
			left++
		}
		window[s[right]] = true
		if currentMaxLength := right - left + 1; currentMaxLength > maxLength {
			maxLength = currentMaxLength
		}
	}
	return maxLength

}

// abba 自己手写的V2，保存重复元素的下一个
func LengthOfLongestSubStringByManuallyV2(s string) int {
	maxLength, left := 0, 0
	hashMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if lastIndex, exist := hashMap[s[right]]; exist && lastIndex >= left {
			left = lastIndex + 1
		}
		hashMap[s[right]] = right
		if currentLength := right - left + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

// 自己手写官网的一种, 比较直观，直接添加
func LengthOfLongestSubStringByManuallyV3(s string) int {

	right, maxLength := -1, 0
	hashMap := make(map[byte]bool)

	for left := 0; left < len(s); left++ {
		if left != 0 {
			//hashMap[s[left-1]] = false
			delete(hashMap, s[left-1])
		}
		for right+1 < len(s) && !hashMap[s[right+1]] {
			hashMap[s[right+1]] = true
			right++
		}
		if currentLength := right - left + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

// 通一提供
func LengthOfLongestSubStringByTongYiV4(s string) int {
	if len(s) == 0 {
		return 0 // 提前处理空字符串
	}

	var hashMap [128]bool // 固定大小数组代替哈希表
	right, maxLength := 0, 0

	for left := 0; left < len(s); left++ {
		if left != 0 {
			hashMap[s[left-1]] = false // 移除左边界字符
		}
		for right < len(s) && !hashMap[s[right]] {
			hashMap[s[right]] = true // 标记右边界字符
			right++                  // 扩展右边界
		}
		currentLength := right - left
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

/*
	变形题：

这里计算的是最长K个字符可以重复的子串的长度，
所以map这里需要value变成int，统计字符重新的次数，每次判断value是否大于k，
这里主要的差异是，需要先添加 ，再判断；
*/
func LengthIfLongestSubStringOwnKReplace(s string, k int) int {
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口的模版，方便理解，但不要用
func Template(s string) int {
	left, right := 0, 0          // 窗口边界
	window := make(map[byte]int) // 窗口数据
	result := 0                  // 存储结果

	for right < len(s) {
		// 1. 右扩窗口
		c := s[right]
		window[c]++
		right++

		// 2. 判断是否需要收缩
		for window[c] > 1 { // 以无重复字符为例的条件
			// 3. 左缩窗口
			d := s[left]
			window[d]--
			left++
		}

		// 4. 更新结果
		result = Max(result, right-left)
	}
	return result
}
