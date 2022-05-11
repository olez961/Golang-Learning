package main

import "os"

// panic通常表示不应该的错误，或者我们不准备优雅处理的错误
// 与error不同，触发panic会直接退出程序
func main() {
	// 触发panic后会输出一个错误信息和协程追踪信息
	// 并以非零的状态码退出程序
	panic("a problem")

	// 因为前面的panic，以下代码成为不可到达的代码
	// 在Go中，通常尽可能使用返回值来标识错误
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
