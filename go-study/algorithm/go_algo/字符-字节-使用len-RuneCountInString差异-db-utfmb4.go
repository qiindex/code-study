package go_algo

import (
	"fmt"
	"unicode/utf8"
)

/*
	结论：

db的字段varchar定义的VARCHAR(2)，2表示的存储的字符数，不是字节数；
所以，例如,
'abc'：
字符数：3 个 ASCII 字符
每个字符的字节数：1 字节
总字节数：3 字节

**'齐天'**：

字符数：2 个中文字符
每个字符的字节数：3 字节（在 utf8mb4 中，大多数中文字符占用 3 字节）
总字节数：6 字节

一：
"abc", 尽管3个字符，3个字节占用 ASCII；
"齐天" 2个字符，6个字节，
但是Varchar（2）只能存下来齐天；

二：
同时：len（s）判断的是字节数；
utf8.RuneCountInString判断的是字符数，也就是rune；
*/
func CompareString() {
	a := "l. Sunter Jaya IV A No.0 11, RT.11/RW.2, Sunter Jaya, Kec. Tj. Priok, Jkt Utara, Daerah Khusus Ibukota Jakarta 14345"
	fmt.Println("len(a)", len(a))
	fmt.Println("rune len", utf8.RuneCountInString(a))
	fmt.Println()
	b := "abc"
	fmt.Println("len(b)", len(b))
	fmt.Println("rune len", utf8.RuneCountInString(b))
	fmt.Println()

	c := "abc abc 齐"
	fmt.Println("len(c)", len(c)) // 11个：3 + 1+3 +3
	/*for i, s := range c {
		fmt.Print(i, c[i], "->", s, "  ")
	}*/
	fmt.Println("rune len", utf8.RuneCountInString(c)) // 9 个
	fmt.Println()
	//fmt.Println(strings.Spl)

	d := "单词就需要统计字符数。此外有些翻译人员按不计空格的字符数来确"
	fmt.Println("len(d)", len(d))
	fmt.Println("rune len", utf8.RuneCountInString(d))
	fmt.Println()

	e := "齐天"
	fmt.Println("len(d)", len(e))                      //6
	fmt.Println("rune len", utf8.RuneCountInString(e)) //2
	fmt.Println()

}
