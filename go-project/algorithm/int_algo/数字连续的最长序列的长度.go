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
func FindLengthIndexV3(list []int) int {
	numSet := make(map[int]bool)
	for _, num := range list {
		numSet[num] = true
	}

	maxInt := 0
	for _, num := range list {
		if numSet[num-1] {
			continue
		}
		currentNum := num
		currentLength := 1
		for numSet[currentNum+1] {
			currentNum++
			currentLength++
		}

		maxInt = MaxInt(maxInt, currentLength)
	}
	return maxInt
}
