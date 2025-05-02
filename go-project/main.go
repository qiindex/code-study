package main

import (
	"fmt"
)

func main() {
	// demo
	/*b := make([]int, 1024)
	c := make(map[int]int)
	for v := range c {
		c[v]++
	}*/
	s := "中国 你好"
	for i, v := range s {
		fmt.Println("s[i]:", s[i])
		fmt.Println(i, string(v))
	}
	fmt.Println(s)
	for v := range s {
		fmt.Println(string(v))
	}

}
