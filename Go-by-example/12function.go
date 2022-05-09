package main

import "fmt"

// 函数 是 Go 的中心
// 一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	// Go 需要明确的返回值，它不会自动返回最后一个表达式的值
	return a + b
}

// 有多个相同类型的连续参数时，可以省略键入参数的类型名称，直到最终参数。
func plusplus(a, b, c int) int {
	return a + b + c;
}

func main() {

	// 通过 name(args) 来调用一个函数
	res := plus(1 , 2)
	fmt.Println("1 + 2 =", res)

	// 注意:=是定义符号，不能重复定义新变量，否则会报错
	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}

// 接下来是多值返回
