package main

import "fmt"

// 可以用range来遍历通道
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 可以看到关闭了的通道，当其内部缓存仍有值的时候
	// 依然是可以被遍历的
	for elem := range queue {
		fmt.Println(elem)
	}

	// 只能遍历一次，下面这个没有输出结果
	// range会消耗channel中的值
	// 或者说只要被读取channel中的值就会被消费
	for elem := range queue {
		fmt.Println(elem)
	}
}
