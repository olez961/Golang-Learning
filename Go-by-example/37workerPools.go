package main

import (
	"fmt"
	"time"
)

// 工作池
func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	t := time.Now()
	// 注意一下常量的初始化方式
	const numJobs = 5

	jobs := make(chan int, numJobs) // 第二个参数也能是变量
	results := make(chan int, numJobs)

	// 启动了三个worker，因为尚无任务，所以这里会堵塞
	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	elapse := time.Since(t)
	// 只消耗了两秒多，因为3个worker是并行的
	fmt.Println("time cost:", elapse) // time cost: 2.0214179s
}

// 输出
/*
worker 1 started job 1
worker 2 started job 2
worker 3 started job 3
worker 3 finished job 3
worker 3 started job 4
worker 2 finished job 2
worker 2 started job 5
worker 1 finished job 1
worker 3 finished job 4
worker 2 finished job 5
time cost: 2.0214179s
*/
