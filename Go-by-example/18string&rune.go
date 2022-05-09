package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "你好"

	fmt.Println("Len:", len(s))	// 6

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])	// 十六进制输出所有字节
	}
	fmt.Println()

	// 用以下函数可以计算字符串包含多少rune，结果和len可能不一样
	fmt.Println("Rune count", utf8.RuneCountInString(s))
	// 输出 2

	// 注意到range是以rune为单位的
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// 输出结果
	// U+4F60 '你' starts at 0
	// U+597D '好' starts at 3


	fmt.Println("\nUsing DecodeRuneInString")
	// 以下提供了一种以rune为单位的遍历方法
	// 可能就是range的底层实现
	for i, w := 0, 0; i < len(s); i += w{
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '你' {	// 单括号表示rune常量，可以直接进行比较
		fmt.Println("found 你")
	}
}