package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}

	realtime := time.Since(t)
	// 下面的程序执行耗时2.01秒，说明goruntine确实是并行执行的
	fmt.Println("real", realtime) // real 2.0131336s
}
