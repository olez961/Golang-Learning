package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子计数器，若不采用原子操作，可能会发生竞争导致结果混乱
func main() {
	var ops uint64
	var anotherOps uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// 通过原子操作避免竞争
				atomic.AddUint64(&ops, 1)
				anotherOps += 1
			}
		}()
	}

	wg.Wait()
	// 如下所示，anotherOps无法实现我们的目标
	fmt.Println("ops:", ops)               // ops: 50000
	fmt.Println("anotherOps:", anotherOps) // anotherOps: 25409
}

// 通过运行时加上-race标志获取数据竞争失败的详情
/*
PS *\Golang-learning\Go-by-example> go run -race .\40atomicCounters.go
==================
WARNING: DATA RACE
Read at 0x00c0000100e8 by goroutine 8:
  main.main.func1()
      D:/Users/LZ/GolandProjects/Golang-learning/Go-by-example/40atomicCoun
ters.go:24 +0xce

Previous write at 0x00c0000100e8 by goroutine 7:
  main.main.func1()
      D:/Users/LZ/GolandProjects/Golang-learning/Go-by-example/40atomicCoun
ters.go:24 +0xe4

Goroutine 8 (running) created at:
  main.main()
      D:/Users/LZ/GolandProjects/Golang-learning/Go-by-example/40atomicCoun
ters.go:19 +0xb1

Goroutine 7 (finished) created at:
  main.main()
      D:/Users/LZ/GolandProjects/Golang-learning/Go-by-example/40atomicCoun
ters.go:19 +0xb1
==================
ops: 50000
anotherOps: 49954
Found 1 data race(s)
exit status 66
*/
