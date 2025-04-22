package main

import (
	"fmt"
)

func main() {
	ss := removeDuplicateLetters("cbacdcbc")
	fmt.Println(ss)

}

// 找出来有可能最小的字典序
func removeDuplicateLetters(s string) string {

	//贪心算法，遍历循环，当前字段小于栈顶并且栈顶还存在，则一直删除栈顶的元素；
	// 单调栈一直增加
	//frequency := make([]int,26)
	//frequency := make([]rune,26)
	//frequency := make([]byte, 26)
	frequency := [26]int{}
	//frequency := [26]int{}
	for _, char := range s {
		frequency[char-'a']++
	}
	stack := make([]int32, 0)
	inStack := make(map[rune]bool) // inter32其实就是rune

	for _, char := range s {
		frequency[char-'a']--
		if inStack[char] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > char && frequency[char-'a'] > 0 {
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, char)
		inStack[char] = true
	}

	return string(stack)

}
