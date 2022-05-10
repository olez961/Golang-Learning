package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 当通道被关闭且其中所有值都被接收完毕时
			// more的值为false
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 这里是提前通知通道关闭
	close(jobs)
	fmt.Println("sent all jobs")

	// 阻塞，后面用WaitGroup处理起来应该更好一些
	<-done
}

/*
示例输出是一个发一个接，我这里似乎一次性发完了
输出：
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
*/
