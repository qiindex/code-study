package int_info

import (
	"sort"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
*/

/*
思路：
代码解释​
​排序数组​：首先对输入数组 nums 进行排序，这样可以利用双指针技术高效地找到满足条件的三元组，并且方便跳过重复元素。
​外层循环​：遍历数组中的每一个元素 nums[i]，作为三元组的第一个元素。如果当前元素与前一个元素相同，则跳过以避免重复的三元组。
​双指针查找​：初始化两个指针 left 和 right，分别指向 i+1 和数组末尾。计算当前三元组的和：
如果和等于零，将三元组加入结果，并跳过所有重复的 left 和 right 元素。
如果和小于零，移动 left 指针向右以增加和。
如果和大于零，移动 right 指针向左以减少和。
​返回结果​：最终返回所有满足条件的不重复三元组。
这种方法的时间复杂度为 O(n²)，其中排序的时间复杂度为 O(n log n)，双指针遍历的时间复杂度为 O(n²)。空间复杂度为 O(1) 或 O(n)，取决于排序的实现。
*/
func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				//跳过重复的右边元素
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

/*
func main() {
	// 示例输入
	nums := []int{-1, 0, 1, 2, -1, -4}

	// 调用 threeSum 函数
	result := int_info.ThreeSum(nums)

*/
