package main

// 文档更新了，这似乎是一种推荐的写法
import (
	"fmt"
	"time"
)

// switch ，方便的条件分支语句
func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	// 一个基本的 switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个 case 语句中，你可以使用逗号来分隔多个表达式。
	// 在这个例子中，我们很好地使用了可选的default 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekday")
	default:
		fmt.Println("it's a weekday")
	}

	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。
	// 这里展示了 case 表达式是如何使用非常量的
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	// 类型switch比较类型而不是值。可以使用它来发现接口值的类型。
	// 在此示例中，变量T将具有对应于其子句的类型。
	whatAmI := func(i interface{}) {
		switch  t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
