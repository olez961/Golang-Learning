package main

import "fmt"

// Go 支持在结构体类型中定义方法
type rectangle struct {
	width, height int
}

// 这里的 area 方法有一个*rect类型接收器
// 这里类似于类的成员函数
func (r *rectangle) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。
// 这里是一个值类型接收器的例子
func (r *rectangle) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	// 默认顺序是width在前，这样可以改变默认赋值顺序
	r := rectangle{height: 10, width: 5}

	// 这里我们调用上面为结构体定义的两个方法
	fmt.Println(r.area())
	fmt.Println(r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}

// 接下来我们将介绍 Go 语言中组织和命名相关的方法集合的机制：接口