package main

import "fmt"

// for 是 Go 中唯一的循环结构。以下是 for 循环的三个基本使用方式
func main(){
	// 最常用的方式，带单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// 似乎不能写++j，经典的初始化/条件/后续形式 for 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。
	for{
		fmt.Println("loop")
		break
	}

	for n:= 0; n < 8; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// 在教程后面，当我们学到 range 语句，channels，以及其他数据结构时，将会看到一些 for 的其它使用形式
}
