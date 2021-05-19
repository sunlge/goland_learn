package main

import "fmt"

type MyInt int //自定义类型，给int改名为MyInt

//在函数定义时，在其名字之前放上一个变量，即是一个 Add方法   函数a的方法Add
func (a MyInt) Add(b MyInt) MyInt { //面向对象
	return a + b
}

//传统方式的定义
func Add(a, b MyInt) MyInt { //面向过程
	return a + b
}

func main() {
	var a MyInt = 1
	var b MyInt = 1

	//调用func (a MyInt) Add(b MyInt)
	fmt.Println("a.Add(b) = ", a.Add(b)) //a.Add(b) =  2

	//调用func Add(a, b MyInt)
	fmt.Println("Add(a, b) = ", Add(a, b)) //Add(a, b) =  2
}
