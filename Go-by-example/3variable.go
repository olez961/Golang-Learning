package main

import "fmt"

func main(){
	// var 声明 1 个或者多个变量
	var a string = "initial"
	fmt.Println(a)

	// 可以申明一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 将自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如
	var e int
	fmt.Println(e)

	// := 语句是声明并初始化变量的简写，例如这个例子
	f := "short"
	fmt.Println(f)
}
