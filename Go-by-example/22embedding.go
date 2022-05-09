package main

import "fmt"

// 菱形embedding会导致不明确ambiguous的错误
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%#v", b.num)
}

type back struct {
	idx int
}

func (b back) describe() string {
	return fmt.Sprintf("base with num=%#v", b.idx)
}

// embedding好像指的就是组合内部的类如果实现了某个接口
// 则组合后的类也能被当做接口来传递参数
// 那么问题来了，菱形组合的话会两个都跑一遍还是冲突报错呢
type container struct {
	base
	//back
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,	// 注意每个字段之后的逗号
		},
		//back: back{
		//	idx: 2,
		//},
		str : "some name",
	}
	fmt.Printf("co={num: %#v, str: %#v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.base.describe())
	//fmt.Println("describe:", co.back.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}