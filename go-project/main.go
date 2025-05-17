package main

import (
	"fmt"

	"go-project/algorithm/int_algo"
)

func main() {
	testCases := []struct {
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{{1}},
			want:   []int{1},
		},
	}

	for _, tc := range testCases {
		got := int_algo.SpiralOrderV2(tc.matrix)
		fmt.Printf("Input: %v\nOutput: %v\nExpected: %v\n\n", tc.matrix, got, tc.want)
	}
}
