package main

import (
	"fmt"

	"go-project/algorithm/int_algo"
)

func main() {
	// demo
	fmt.Println(int_algo.CountNumberNMV2(2, 3)) // 3 (1+1+1, 1+2, 2+1)
	fmt.Println(int_algo.CountNumberNMV2(3, 4)) // 7 (1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2, 1+3, 3+1)
}
