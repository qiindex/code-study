package int_window_algo

import (
	"container/list"
	"fmt"
)

/*
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

*/
/*

方法一：优先队列
思路与算法
对于「最大值」，我们可以想到一种非常合适的数据结构，那就是优先队列（堆），其中的大根堆可以帮助我们实时维护一系列元素中的最大值。
对于本题而言，初始时，我们将数组 nums 的前 k 个元素放入优先队列中。每当我们向右移动窗口时，我们就可以把一个新的元素放入优先队列中，此时堆顶的元素就是堆中所有元素的最大值。然而这个最大值可能并不在滑动窗口中，在这种情况下，这个值在数组 nums 中的位置出现在滑动窗口左边界的左侧。因此，当我们后续继续向右移动窗口时，这个值就永远不可能出现在滑动窗口中了，我们可以将其永久地从优先队列中移除。
我们不断地移除堆顶的元素，直到其确实出现在滑动窗口中。此时，堆顶元素就是滑动窗口中的最大值。为了方便判断堆顶元素与滑动窗口的位置关系，我们可以在优先队列中存储二元组 (num,index)，表示元素 num 在数组中的下标为 index。

*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	result := make([]int, 0, len(nums)-k+1)
	deque := list.New() // 存储的是索引

	for i := 0; i < len(nums); i++ {
		// 移除不在窗口范围内的元素
		for deque.Len() > 0 && deque.Front().Value.(int) <= i-k {
			deque.Remove(deque.Front())
		}

		// 移除所有小于当前元素的元素
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] <= nums[i] {
			deque.Remove(deque.Back())
		}

		// 添加当前元素索引
		deque.PushBack(i)

		// 当窗口大小达到k时，记录结果
		if i >= k-1 {
			result = append(result, nums[deque.Front().Value.(int)])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) // 输出: [3 3 5 5 6 7]
}

// 自己写的遍历，会超时
// 方法二：暴力法
func MaxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i+k > len(nums) {
			break
		}
		currentWindow := nums[i : i+k]
		currentMax := nums[i]
		for j := 1; j < k; j++ {
			if currentWindow[j] > currentMax {
				currentMax = currentWindow[j]
			}
		}
		result = append(result, currentMax)

	}
	return result
}

// 双端队列，自己实现双端队列
func maxSlidingWindow3(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	deque := make([]int,0)
	res :=make([]int,0)
	for
}
