package main

import (
	"fmt"
	"time"
)

// 定时器，固定时间间隔发送一次信号到其对应的channel中
// channel是其内部的C，C内部的格式如下：
// 2022-05-10 22:33:18.0040026 +0800 CST m=+0.519556001
// 最后一个字段似乎是当前时间相对程序开始运行的时刻经过的秒数
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // 依旧是之前的同步手段

	go func() {
		for {
			select {
			case <-done:
				return
			// 从通道中读取当前时间
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 保证定时器的触发
	time.Sleep(1600 * time.Millisecond)

	// 可以中断定时器
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}

// 输出
/*
Tick at 2022-05-10 22:33:18.0040026 +0800 CST m=+0.519556001
Tick at 2022-05-10 22:33:18.5032321 +0800 CST m=+1.018785501
Tick at 2022-05-10 22:33:18.9899696 +0800 CST m=+1.505523001
Ticker stopped
*/
