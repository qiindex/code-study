package string_algo

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 支持三个数比较的 min 函数
func Min3(a, b, c int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
