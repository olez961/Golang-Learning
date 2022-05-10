package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// 以下等待一个定时器失效，失效前此处会一直阻塞
	<-timer1.C
	fmt.Println("Timer1 fired") // Timer1 fired

	timer2 := time.NewTimer(time.Second)
	go func() { // 输出空
		// 因为timer2被提前终止了，所以下面会一直阻塞
		// 直到程序退出
		// 这个协程很可能变成一个孤儿协程
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()

	// 以下证明定时器可以被提前终止
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	} // Timer2 stopped

	time.Sleep(3 * time.Second)
}
