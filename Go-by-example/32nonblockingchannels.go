package main

import "fmt"

// 非阻塞一般由select和default来实现
// 若当前没有发送或者接收者，直接跳转到default部分执行
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 一个非阻塞的接收，如果当前有合适的信息则会接收到
	// 否则会直接跳转到default
	// 这里由于messages中事先没有变量所以会直接default
	select {
	// no message received
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞的发送，这里不会发送
	// 因为channel无缓冲且当前没有接收者
	msg := "hi"
	select {
	// no message received
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 下面是一个多channels非阻塞的示例
	select {
	// no action
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signal", sig)
	default:
		fmt.Println("no action")
	}
}
