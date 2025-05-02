package int_algo

/*
给定一个未排序的整数数组，找出数字连续的最长序列的长度。例如：
● 输入: [100, 4, 200, 1, 3, 2]
● 输出: 4（因为最长连续序列是 [1, 2, 3, 4]）
 //百度一面
*/

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
思路，用hashSet判断，从当前节点，开始的，i+1，i+2是否存在，递增
*/

func longestConsecutive(nums []int) int {
	sumLength := 0
	hashMap := make(map[int]bool)

	for _, num := range nums {
		hashMap[num] = true
	}

	// 这里循环list，会超时的，需要循环map，因为map已经去重了
	for num, _ := range hashMap {
		if hashMap[num-1] {
			// 简单来说就是每个数都判断一次这个数是不是连续序列的开头那个数。
			continue
		}
		//if !hashMap[num-1] {
		currentNum := num
		currentLength := 1
		for hashMap[currentNum+1] {
			currentLength++
			currentNum++
		}
		if currentLength > sumLength {
			sumLength = currentLength
		}
		//	}

	}
	return sumLength

}
