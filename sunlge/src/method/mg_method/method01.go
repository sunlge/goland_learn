package main

import "fmt"

type Dog struct {
	name string
}

// 方法是添加了接收者的函数dog，接收者必须是自定义的类型
func (dog Dog) call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog Dog) SetName(name string) {
	dog.name = name
}

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {
	sunlge := Dog{name: "Bob"}
	sunlge.call()
	fmt.Println("使用值引用去修改名称的结果")
	sunlge.SetName("Sophie")
	sunlge.call()
	fmt.Println("使用指针引用去修改名称的结果")
	sunlge.PSetName("Sophie")
	sunlge.call()
	// 也可以写成(&sunlge).PSetName("Jinks")
	(&sunlge).PSetName("Jinks")
	sunlge.call()
}
