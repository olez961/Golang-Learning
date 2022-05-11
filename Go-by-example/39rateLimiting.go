package main

import (
	"fmt"
	"time"
)

func main() {
	// 以下模拟五个请求，后续通过定时器限制处理速率
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 用于进行速率限制
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// 通过定时器实现阻塞的效果
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 短暂地允许并发请求
	burstyLimiter := make(chan time.Time, 3)
	// channel提前充入数据，允许并发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	
	// 设置速率限制
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 验证爆发请求的请求队列
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	// 以下前三个请求属于爆发请求，后两个请求间隔处理
	// 似乎是爆发请求结束后经过定时间隔才会处理新的请求
	for req := range burstRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// 可以看到前五个输出是按照200毫秒的间隔输出的
// 后五个输出中前三个是爆发bursts输出，后两个遵循了
// 200毫秒的间隔
// 输出：
/*
request 1 2022-05-11 14:28:46.1861651 +0800 CST m=+0.205435001
request 2 2022-05-11 14:28:46.3921015 +0800 CST m=+0.411371401
request 3 2022-05-11 14:28:46.5978106 +0800 CST m=+0.617080501
request 4 2022-05-11 14:28:46.8018027 +0800 CST m=+0.821072601
request 5 2022-05-11 14:28:46.9935264 +0800 CST m=+1.012796301
request 1 2022-05-11 14:28:46.9959107 +0800 CST m=+1.015180601
request 2 2022-05-11 14:28:46.9959107 +0800 CST m=+1.015180601
request 3 2022-05-11 14:28:46.9966429 +0800 CST m=+1.015912801
request 4 2022-05-11 14:28:47.2102202 +0800 CST m=+1.229490101
request 5 2022-05-11 14:28:47.399799 +0800 CST m=+1.419068901
*/
