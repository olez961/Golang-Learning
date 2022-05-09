package main

// Go 的结构体 是各个字段字段的类型的集合。这在组织数据时非常有用
import "fmt"

// 这里的 person 结构体包含了 name 和 age 两个字段
type person struct {
	name string
	age int
}

func main() {
	// 使用这个语法创建了一个新的结构体元素
	fmt.Println(person{"Bob", 20})

	// 初始化的时候一个参数用了name：另一个就不能漏掉age：
	// 可以在初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "Fred"})

	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 50})	// &{Ann 50}

	s := person{"Sean", 50}
	fmt.Printf("%#v\n", s)	// main.person{name:"Sean", age:50}会输出详细类型名
	fmt.Printf("%+v\n", s)	// {name:Sean age:50}，输出字段名和值
	// 使用点来访问结构体字段
	fmt.Println(s.name)

	sp := &s
	// 也可以对结构体指针使用. - 指针会被自动解引用
	fmt.Println(sp.age)

	// 结构体是可变的
	sp.age = 51
	fmt.Println(sp)	// &{Sean 51}
	fmt.Println(*sp)	//{Sean 51}
}
