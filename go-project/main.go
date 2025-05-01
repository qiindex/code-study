package main

import (
	"fmt"
)

func main() {
	// demo

	list := make([]int, 3)
	/*for i := 0; i < len(list); i++ {
		list = append(list, i)
	}*/

	for i := range list {
		fmt.Println(i)

	}
	fmt.Println(list)

	maps := make(map[string]int)
	maps["a"] = 1
	maps["b"] = 2

	for k, v := range maps {
		fmt.Println(k, v)
	}

}
