package main

import (
	"fmt"
	"math"
)

// 方法签名的集合叫做接口

// 这是一个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

// 我们这里为rect和circle实现该接口
type rect struct {
	width, height float64	// 可以放两个变量定义
}
type circle struct {
	radius float64
}

// 要在Go中实现一个接口，我们只需要实现接口中的所有方法
// 首先为rect实现接口
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * r.height * r.width
}
// 接着为circle实现该接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法
// 并且接口可以被当做参数
func measure(g geometry) {
	fmt.Printf("%#v\n", g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 只有实现了相关接口的结构体的实例，才能作为measure的参数
	measure(r)	// 这里调用接口的参数实际上是我们的结构体
	measure(c)
}