package main

import (
	"fmt"
	"time"
)

// 可以通过通道来实现goruntine的同步
// 下面是接收一个传送bool值的通道
// 这个参数没有传指针，似乎通道传的是本题而不是拷贝
func worker(done chan bool) {
	// 这里是为了实现这样的效果：working...done
	// Print和Printf类似，只不过不用添加格式控制
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 向该通道传送bool值，这里true和false都是一样的
	// 因为该值仅充当一个消费品
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	fmt.Println("sending done signal")

	// 此处接收通道的值，若尚未收到值，会阻塞
	<-done
}
