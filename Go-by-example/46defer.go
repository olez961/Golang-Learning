package main

import (
	"fmt"
	"os"
)

// defer用于确保程序在执行完成后，会调用某个函数
// 常用于清理工作
// 其用途和其它语言的ensure或finally类似
func main() {
	// 若未找到当前目录会报错，网站上是"/tmp/defer.txt"
	f := createFile("./test/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	// 关闭文件时进行错误检查是非常重要的
	// 即使是在defer函数中
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %#v\n", err)
		os.Exit(1)
	}
}

// 输出：
/*
creating
writing
closing
*/
