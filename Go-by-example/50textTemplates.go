package main

import (
	"os"
	"text/template"
)

func main() {

	t1 := template.New("t1") // 这里是t1还是t问题不大
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	}) // Value: [Go Rust C++ C#]

	// 这里相当于自动化了上面的几步操作，注意返回的是一个模板指针
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 如果数据是一个结构体，我们可以使用 {{.FieldName}} 动作来访问其字段。
	// 这些字段应该是导出（结构体的话就是首字母大写）的，以便在模板执行时可访问。
	t2 := Create("t2", "Name: {{.Name}}\n")

	// 以下函数第二个参数是any类型
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty") // yes
	t3.Execute(os.Stdout, "")          // no

	//  实现不定参数数目模板输出。参数数目可以为空
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{ // 切片为空的话初始化时不需要逗号
			"Go",
			"Rust",
			"C++",
			"C#",
		}) // Range: Go Rust C++ C#
}
