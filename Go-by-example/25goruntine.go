package main

// 协程是轻量级的执行线程

// 考虑到Goland会自动import包，感觉这一步可以不写
import (
	"fmt"
	"sync"
	"time"
)

func fWithoutSync(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func fWithSync(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	fWithoutSync("direct")
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 协程执行该函数，Go协程会并发地执行该函数
	go fWithSync("goruntine", &wg)

	// 可以为匿名函数启动一个协程
	go func(msg string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(msg)
	}("going", &wg)

	// 等待两个协程异步执行结束，更好的方式是用WaitGroup
	time.Sleep(time.Second)
	// 由于是并发执行，所以两个协程输出可能没有按照先后顺序
	wg.Wait()
	fmt.Println("done")
}
