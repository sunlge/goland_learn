package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) PrintInfo() { //给Person添加方法
	fmt.Println(p.name, p.sex, p.age)
}

func main() {
	p := Person{"mike", 'm', 18} //初始化
	p.PrintInfo()                //调用func (p Person) PrintInfo()
}
