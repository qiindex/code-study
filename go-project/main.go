package main

import (
	"fmt"
)

func main() {
	// demo

	list := make([]int, 0, 3)
	for i := 0; i < len(list); i++ {
		list = append(list, i)
	}

	fmt.Println(list)

}
