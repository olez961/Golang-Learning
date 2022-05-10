package main

import "fmt"

// channel是可以指明方向的
// 这里的pings相当于管道里面的fd[1]，只能进
// 如果强行让pings出的话会产生编译时错误
func ping(pings chan<- string, message string) {
	pings <- message
}

// 这里的pongs相当于管道里面的fd[0]，只能出
func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings // 注意一下这里的初始化方式
	pongs <- message
}

func main() {
	// 这里没有缓存的话可能会死锁
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed meaasge")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
