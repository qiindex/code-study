package int_algo

//package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//arr := []int{5, 3, 8, 4, 2, 7, 1, 10}
	arr := []int{110, 100, 0}
	QuickSort(arr, 0, len(arr)-1)
	//fmt.Println(arr) // 输出: [1 2 3 4 5 7 8 10]
	//value := int_algo.FindKthLargest(arr, 2)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	println(arr)

}

// 快速排序主函数
func QuickSort(arr []int, low, high int) {
	if low < high {
		// 找到分区点（pivot 的正确位置）
		pivotIdx := Partition(arr, low, high)
		// 递归排序左半部分
		QuickSort(arr, low, pivotIdx-1)
		// 递归排序右半部分
		QuickSort(arr, pivotIdx+1, high)
	}
}

// 分区函数：将数组分为 <=pivot 和 >pivot 的两部分
func Partition(arr []int, low, high int) int {
	//pivot := arr[high] // 选择最后一个元素作为基准
	i := low - 1 // i 指向 <=pivot 的最后一个元素

	// 随机选择 pivot
	randomIdx := low + rand.Intn(high-low+1)
	arr[randomIdx], arr[high] = arr[high], arr[randomIdx]
	pivot := arr[high]

	// 遍历 low 到 high-1
	for j := low; j < high; j++ {
		if arr[j] < pivot { // 在[110, 100, 0]例子中，需要改成<,不能等于，才能通过likou
			i++
			arr[i], arr[j] = arr[j], arr[i] // 交换
		}
	}
	/*
		为什么循环范围是 [left, right-1]？
		原因：nums[right] 是 pivot 的位置，循环中不需要处理它。
		目的：通过遍历 [left, right-1，将大于 pivot的元素移到左侧，最终将pivot` 放到正确位置。
	*/

	// 将 pivot 放到正确位置（i+1）
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // 返回 pivot 的索引
}
