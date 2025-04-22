package int_algo

import "sort"
import "math/rand"

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/

/*
思路：
1、如果直接用排序+找K，--> nlogN，会超过
所以里面快排的基准原则，升级为「快速选择」算法。q 为倒数第 k 个下标的时候
2、「快速选择」算法
于快速排序的 分区（Partition） 思想：
随机选择一个 pivot，将数组分为两部分：
左侧元素 ≥ pivot，右侧元素 < pivot。
如果 pivot 的位置正好是 len(nums)-k，则返回 pivot。
否则递归处理左侧或右侧子数组。
*/

// 升级的「快速选择」算法。，最优解
// 同时pivot是随机数，不是制定的，不会有n-1多长度空间

func FindKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// 豆包提供（看懂了，但是提交也超时，很多1）
// findKthLargest +quickSelect+partition
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, target int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := partition(nums, left, right)

	if pivotIdx == target {
		return nums[pivotIdx]
	} else if pivotIdx < target {
		return quickSelect(nums, pivotIdx+1, right, target)
	} else {
		return quickSelect(nums, left, pivotIdx-1, target)
	}
}

func partition(nums []int, left, right int) int {
	pivotIdx := rand.Intn(right-left+1) + left
	pivot := nums[pivotIdx]
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	storeIdx := left - 1

	for i := left; i < right; i++ {
		if nums[i] < pivot { // 不需要，先按照从小到大的顺序排列
			storeIdx++
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]

		}
	}
	nums[storeIdx+1], nums[right] = nums[right], nums[storeIdx+1]
	return storeIdx + 1
}

// likou官方 (没看懂)
// findKthLargestv1   +quickselect
func findKthLargestv1(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// 利用代码库，不推荐；
func FindKthLargestV2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 借助快排实现，但是会超时；不推荐
func FindKthLargestV1(nums []int, k int) int {
	QuickSort(nums, 0, len(nums)-1)

	return nums[len(nums)-k]
}

// 手写partition
func quickSelectV3(nums []int, left, right, target int) int {
	pivot := partition(nums, 0, len(nums)-1)
	if pivot == target {
		return nums[pivot]
	} else if pivot < target {
		return quickSelectV3(nums, pivot+1, right, target)
	} else {
		return quickSelectV3(nums, left, pivot-1, pivot)
	}

}
func partitionV3(nums []int, left, right int) int {
	pivotIndex := left + rand.Intn(right-left+1)
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex] // 让right的内容为pivot的值；（为了后面方便循环）
	startIndex := left - 1

	for i := startIndex; i < right; i++ {
		//从 left 到 right-1 遍历元素，将 ​**大于 pivot**​ 的元素依次交换到 storeIdx 位置，并递增 storeIdx。
		if nums[i] < pivot {
			startIndex++
			nums[startIndex], nums[i] = nums[i], nums[startIndex] //交换
		}
	}
	nums[startIndex+1], nums[right] = nums[right], nums[startIndex+1] // right的值，和pivot的值交换，此时pivot的值回到他应该去的位置；
	return startIndex + 1
}
