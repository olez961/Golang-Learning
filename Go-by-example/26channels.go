package main

import (
	"fmt"
	"time"
)

// channels是连接多个协程的管道
// 可以从一个协程将值发送到通道然后在另一个协程中接收

func main() {
	// 创建一个无缓冲通道
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		// channel <- val 将值送入通道
		messages <- "ping"
		// 这样操作后下面输出会被直接跳过
		time.Sleep(time.Second * 3)
		fmt.Println("Finish")
	}()

	// <- channel 将通道内的值送入左侧，左侧可能是通道也可能是变量
	msg := <-messages // 这里会堵塞，上面没有执行结束这里不会开始执行
	fmt.Println(msg)
}
