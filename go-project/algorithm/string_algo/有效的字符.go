package string_algo

/*
思路：
// 思路：有一个单调战，
// 遇左边压入，
//遇右边括号比较，
//空， 不匹配 -f
// 匹配，弹出站顶
//极端：站需要为空，
*/
func IsValidString(s string) bool {
	var stack []rune
	mapping := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else if char == '}' || char == ']' || char == ')' {
			if len(stack) == 0 {
				return false
			}
			topChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if topChar != mapping[char] {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
