package main

import "fmt"

// 感觉slice有点像C++里面的vector
// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口
func main() {
	// 不像数组，slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。
	// 要创建一个长度非零的空slice，需要使用内建的方法 make。
	// 这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）。

	// var s []string	// 也可以这样初始化，初始化为零值
	s := make([]string, 3)
	fmt.Println("emp", s)	// 输出emp [  ]，[]括号中是两个空格，里面的string元素为空

	// 还能这样写，第二个参数是len，第三个参数是cap
	// 若cap小于len会报错：len larger than cap in make([]string)
	s = make([]string, 3, 5)
	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	// len 返回 slice 的长度
	fmt.Println("len", len(s))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 append，它返回一个包含了一个或者多个新值的 slice。
	// 注意我们接受返回由 append返回的新的 slice 值。
	// 直接append无法完成添加元素的工作，而且goland会报错
	s = append(s, "d")
	s = append(s, "e", "f", "gh")
	fmt.Println("apd", s)

	// Slice 也可以被 copy。
	// 这里我们创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c
	// 注意c长度可以不等于s，若len(c) < len(s)，则只会复制len(c)长度到c中
	// 若len(c) > len(s)，则会将c的前len(s)个元素全部替换成s中的元素
	c := make([]string, len(s) + 1)
	c[len(s)] = "g"
	copy(c, s)
	fmt.Println("cpy", c)

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作。
	// 例如，这里得到一个包含元素 s[2], s[3],s[4] 的 slice
	l := s[2:5]
	fmt.Println("sl1", l)	// sl1 [c d e]

	//l := s[2:cap(s) + 1]	// 写成这样会报错，因为超出slices的容量了

	// 这个 slice 从 s[0] 到（但是包含）s[5]
	l = s[:5]
	fmt.Println("sl2", l)	// sl2 [a b c d e]

	// 这个 slice 从（包含）s[2] 到 slice 的后一个值
	l = s[2:]
	fmt.Println("sl3", l)	// sl3 [c d e f gh]

	// 我们可以在一行代码中声明并初始化一个 slice 变量，列表初始化
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// Slice 可以组成多维数据结构。
	// 内部的 slice 长度可以不同，这和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 通过以下操作我们可以发现切片后切片前的数据访问不到了
	// 但是可以通过slice扩容的方式得到原来切片之后的数据
	s = s[2:4]
	fmt.Println("s: ", s, "and the capacity of s is: ", cap(s), len(s))
	// s:  [c d] and the capacity of s is:  10 2
	s = s[0:cap(s)]
	fmt.Println("s: ", s, "and the capacity of s after slicing to cap is: ", cap(s), len(s))
	// s:  [c d e f gh     ] and the capacity of s after slicing to cap is:  10 10

	// 以下l将成为s的一个引用
	l = s[:]
	fmt.Println(l, s)	// [c d e f gh     ] [c d e f gh     ]
	l[0] = "b"
	fmt.Println("after changing l, l and s: ", l, s)
	// after changing l, l and s:  [b d e f gh     ] [b d e f gh     ]

	// s = append(s, l)	// 无法将 'l' (类型 []string) 用作类型 string
	// 将一个slice贴到另一个slice后面，应该这样写：
	s = append(s, l...)
	fmt.Println("s after appending :", s)
	// s after appending : [b d e f gh    b d e f gh   ]
}

// 注意，slice 和数组不同，虽然它们通过 fmt.Println 输出的格式差不多
