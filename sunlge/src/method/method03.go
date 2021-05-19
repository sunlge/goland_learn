package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer() {
	//给成员赋值
	(*p).name = "yoyo"
	p.sex = 'f'
	p.age = 22
}

//值作为接收者，值语义
func (p Person) SetInfoValue() {
	//给成员赋值
	p.name = "yoyo"
	p.sex = 'f'
	p.age = 22
}

func main() {
	//指针作为接收者，引用语义
	p1 := Person{"mike", 'm', 18} //初始化
	fmt.Println("函数调用前 = ", p1)   //函数调用前 =  {mike 109 18}
	(&p1).SetInfoPointer()
	fmt.Println("函数调用后 = ", p1) //函数调用后 =  {yoyo 102 22}

	fmt.Println("==========================")

	p2 := Person{"mike", 'm', 18} //初始化
	//值作为接收者，值语义
	fmt.Println("函数调用前 = ", p2) //函数调用前 =  {mike 109 18}
	p2.SetInfoValue()
	fmt.Println("函数调用后 = ", p2) //函数调用后 =  {mike 109 18}

}
