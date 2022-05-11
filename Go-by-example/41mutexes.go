package main

import (
	"fmt"
	"sync"
)

// sync.Mutex用于实现互斥访问
// 看来Go是比较推荐把互斥量封装在想要实现同步的数据结构里面
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// 照理来说这个函数应该属于公开的接口
// 但是这里首字母没有大写，可能Go没有那么严格的访问控制
func (c *Container) inc(name string) {
	c.mu.Lock()
	// 通过defer实现在函数结束时解锁
	// 这里每次加1，似乎有点原子操作的味道
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	// 注意不像map，互斥量不用初始化便可以直接使用
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		// 这里没加defer，影响不大
		wg.Done()
	}

	wg.Add(3)
	doIncrement("a", 1000)
	doIncrement("a", 1000)
	doIncrement("b", 1000)
	wg.Wait() // 注意别漏了这一句

	fmt.Println(c.counters) // map[a:2000 b:1000]
}
