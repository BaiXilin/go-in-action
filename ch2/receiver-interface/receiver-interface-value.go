package main

import "fmt"

type Printer interface {
	print(name string)
}

type printName struct{}

// receiver function, 实现了Printer接口
// 1. 接收者是值
// func (p printName) print(name string) {
// 	fmt.Printf("My name is %s\n", name)
// }

// 2. 接收者是指针
func (p *printName) print(name string) {
	fmt.Printf("My name is %s\n", name)
}

// 使用接口的值来调用print函数

func main() {
	var name = "Mike"
	// 调用类型和接受类型相同
	// 值调用
	instance := printName{}
	// 指针调用
	//instance := &printName{}
	instance.print(name)

	// 使用接口类型调用
	pr := printName{}
	// 值调用
	var pn Printer = pr	// 当接收类型为指针时，不能使用接口值调用
	// 指针调用
	// var pn Printer = &pr
	pn.print(name)

}
