package main

import (
	"fmt"
	"sort"
)

// 自定义排序
type byLength []string

// 自定义排序需要实现接口的几个方法，分别是以下三个
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits) // [kiwi peach banana]
}
