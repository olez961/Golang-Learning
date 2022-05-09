package main

// Go依靠其多值返回的特性一般采用一个返回值来传递错误信息
// 似乎这里没有直接用error的接口输出错误信息

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if 42 == arg {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// 似乎实现了下面这个函数就实现了Error相关的接口
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if 42 == arg {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg + 3, nil // nil表示没错
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("fWithSync failed:", e)
		} else {
			fmt.Println("fWithSync worked", r)
		}
	}

	_, e := f2(42)
	// 下面这个用法感觉argError已经embedding到error里面了
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
