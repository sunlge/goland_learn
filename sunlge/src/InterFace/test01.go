package main

import (
	"fmt"
)

type Test struct {}

func (t Test) test01() {
	fmt.Println("这是一个test")
}

func main()  {
	t1 := Test{}
	t1.test01()

	// 方法表达式，必须显示的传一个Test类型的数据
	t2 := Test.test01
	t2(t1)
}