package main

import (
	"fmt"
	// 感觉类似与python的import http as h
	strLib "strings" // 一种方便的写法
)

// 以下写法挺方便的，类似于一个1宏
var p = fmt.Println

func main() {
	p("Contains:  ", strLib.Contains("test", "es"))
	p("Count:     ", strLib.Count("test", "t"))
	p("HasPrefix: ", strLib.HasPrefix("test", "te"))
	p("HasSuffix: ", strLib.HasSuffix("test", "st"))
	p("Index:     ", strLib.Index("test", "e")) // 1
	p("Join:      ", strLib.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strLib.Repeat("a", 5))
	p("Replace:   ", strLib.Replace("foo", "o", "0", -1)) // -1表示不限制替换次数
	p("Replace:   ", strLib.Replace("foo", "o", "0", 1))  // 1表示替换一次
	p("Split:     ", strLib.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strLib.ToLower("TEST"))
	p("ToUpper:   ", strLib.ToUpper("test"))
	p()

	p("Len:       ", len("hello"))
	// 下面两个结果一致
	p("Char:      ", "hello"[1])       // 101，获取的似乎是一个编码而不是一个char
	p("Char:      ", rune("hello"[1])) // 101，获取的似乎是一个编码而不是一个char
}

// 输出：
/*
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:        5
Char:       101
*/
