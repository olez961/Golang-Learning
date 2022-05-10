package main

import "fmt"

func main() {
	// 带缓存的channels，后面表示通道缓存大小
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// 若没有消费，这里还是会阻塞
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
