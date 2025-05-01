package link_algo

/*
前缀和
*/
func subarraySum(nums []int, k int) int {

	count := 0
	prefixSum := 0
	sumCount := make(map[int]int)
	sumCount[prefixSum] = 1
	for _, num := range nums {
		prefixSum += num
		if value, exists := sumCount[prefixSum-k]; exists {
			count += value
		}
		sumCount[prefixSum]++
	}
	return count

}

/*

pre - k 的意义：这个检查的意义在于，
如果 pre - k 存在于 Map 中，说明之前在某个点的累积和是 pre - k。
由于当前的累积和是 pre，这意味着从那个点到当前点的子数组之和恰好是 k
（因为 pre - (pre - k) = k）

*/
