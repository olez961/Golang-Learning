package main

import "fmt"

// range 迭代各种各样的数据结构
func main() {

	nums := []int{2, 3, 4}
	sum := 0	// 第一个参数是索引号，若不适用则以下划线代替
	for _, num := range nums {	// 使用 range 来统计一个 slice 的元素个数。数组也可以采用这种方法。
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中都同样提供每个项的索引和值。
	// 上面我们不需要索引，所以我们使用 空值定义符_ 来忽略它。
	// 有时候我们实际上是需要这个索引的
	for i, num := range nums {
		if num == 3{
			fmt.Println("index:", i)
		}
	}

	// range 在 map 中迭代键值对
	kvs := map[string]string{"a" : "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)	// 这个格式控制和c很像，感觉一脉相承
	}

	// range也能只遍历map中的keys
	for k := range kvs {
		fmt.Println("key: ", k);
	}

	// 即使这边声明了c，下面range语句还是用的新的c，应该是覆盖了
	var c byte = 'a'
	fmt.Printf("%c\n", c)
	// range 在字符串中迭代 unicode 编码。
	// 第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己
	// rune 等价于int32
	for i, c := range "go"{
		fmt.Println(i, c)
	}
}