package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

// recover用于恢复panic，阻止其中止程序
// 保证程序的正常执行
// 服务器中当某个连接出现问题导致panic时
// 我们希望断开该连接而不是因此关闭服务器
// 这时便可以用recover来处理
func main() {
	// 必须通过defer来调用recover
	// 当引发panic时recover会捕获panic
	defer func() {
		if r := recover(); r != nil {
			// 这里用Println会导致第二行多一个空格
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// 以下代码不会被执行，因为main会在mayPanic()处终止
	// 并在继续处理完defer后结束
	fmt.Println("After mayPanic()")
}

// 输出：
/*
Recovered. Error:
 a problem
*/
