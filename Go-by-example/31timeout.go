package main

import (
	"fmt"
	"time"
)

// Golang可以很方便地处理超时信息
// 通过select和time.After方法可以使得超过一定时限收到的消息
// 通过不同的方式输出或处理
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	// timeout 1，说明此处等待c1结果超时
	case res := <-c1:
		fmt.Println(res)
	// 如果超时超过1秒，则按以下情况处理
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	// result 2，说明此处等待c2结果未超时
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
