package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数
func workerInWGExample(id int, t time.Duration) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(t)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 用于等待所有该group中的协程结束
	// 该group中的协程指的是含有defer wg.Done()的协程
	// 通过wg.Add(n)来添加n个成员
	// 通过wg.Wait()等待所有该group中的协程运行结束
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// 去掉这一步后worker全部变成6了
		// 每个协程闭包中都用了同一个i
		i := i

		go func() {
			// defer用于延迟执行
			// 此处会在该匿名函数运行结束时自动执行
			// defer后的方法
			defer wg.Done()
			workerInWGExample(i, 1*time.Second)
		}()
	}

	go func() {
		defer wg.Done()
		// 这种情况下可以看到worker6开始运行
		// 但是由于未将其加入group，所以wg不会等待其结束
		workerInWGExample(6, 5*time.Second)
	}()

	wg.Wait()
}

// 输出：
// 每个协程的运行开始和结束不一定相同
/*
Worker 5 starting
Worker 2 starting
Worker 3 starting
Worker 4 starting
Worker 1 done
Worker 5 done
Worker 3 done
Worker 4 done
Worker 2 done
*/
