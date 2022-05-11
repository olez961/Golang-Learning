package main

import (
	"fmt"
	"sort"
)

func main() {
	// Go中的排序是针对内置数据类型的，采用的是原地排序
	// 所以不用传指针
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs) // Strings: [a b c]

	// 似乎无法对数组进行排序，只支持对切片进行排序
	// 暂时未找到对数组进行排序的方法
	// ints1 := [3]int{7, 2, 4}
	// sort.Sort(ints1)
	// 原地排序，会改变原切片
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints   :", ints) // Ints   : [2 4 7]

	// 检查数组是否已经有序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted :", s) // Sorted : true
}
