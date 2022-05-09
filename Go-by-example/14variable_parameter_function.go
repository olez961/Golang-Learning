package main

// 可变参数函数。可以用任意数量的参数调用。
// 例如，fmt.Println 是一个常见的变参函数
import "fmt"

func sum(nums ...int) {	// 这个函数使用任意数目的 int 作为参数
	fmt.Print(nums, " ")	// 这里文档用的是Print而不是Println
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)	// 变参函数使用常规的调用方式，除了参数比较特殊
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)	// 如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)
}

// 接下来是go函数的另一个关键特性————闭包结构