package main

import "fmt"

// MapKeys 返回Key的切片，comparable表示可比较大小，any表示任意类型
// any是interface{}的别名类型
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List表示一个具有任意类型值和头尾指针的单链表
type List[T any] struct {
	head, tail *element[T] // 这里的[T]相当于C++中的<typename T>
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		// 采用尾插法，这里添加一个元素后头尾指针都指向该元素
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		// 尾插法的例证
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// GetAll 将链表中的值存储在切片中返回
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// MapKeys在这里使用没有明确类型，因为编译器会帮我们自动推断
	fmt.Println("keys of m:", MapKeys(m))

	// 这里明确了类型
	_ = MapKeys[int, string](m) // 这里只是示意，所以用下划线避免编译器报错

	// 初始化时需要明确类型
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(42)
	fmt.Println("list", lst.GetAll())
}
