package main

import "fmt"

// SelectionSort 实现选择排序
func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// PrintArray 打印数组
func PrintArray(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	//arr := []int{3, 2, 1, 5, 6, 4}
	arr := []int{5, 2, 4, 1, 3, 6, 0}

	/*fmt.Println("Original array:")
	PrintArray(arr)

	SelectionSort(arr)
	fmt.Println("Sorted array:")
	PrintArray(arr)
	arr = []int{5, 2, 4, 1, 3, 6, 0}*/
	arr = []int{3, 1, 2, 5, 6, 4}
	fmt.Println(" myself Sorted array:")
	//arrList := s(arr)
	arrList := SelectedSort(arr)
	fmt.Println(arrList)

	//println("Original:", arr[2-1])

}

func SelectedSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr

}
