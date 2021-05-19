package main

import "fmt"

type Dog struct {
	name string
}

// 方法是添加了接收者的函数dog，接收者必须是自定义的类型
func (dog Dog) call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog *Dog) PSetName(name string) {
	dog.name = name
}

func main() {
	m1 := Dog.call
	m2 := (*Dog).PSetName
	fmt.Printf("%T, %T\n", m1, m2)
	dog := Dog{"Sophie"}
	m1(dog)
	m2(&dog, "Bob")
	m1(dog)
}
