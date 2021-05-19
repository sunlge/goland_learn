package MyVar

import (
	"fmt"
	"reflect"
)

func MyVar()  {
	/*
		定义变量
	*/
	var me string
	me = "My is a first variable"
	var v1 int = 10
	var v2 string = "v2"
	v3 := "v3 ,编译器自动推导出v3的类型，简短声明，只能在函数内部使用"
	var v4, v5 string = "v4", "v5"
	var (
		age int = 22
		height float64 = 1.70
	)
	var (
		v6 = 10
		v7 = "abc"
	)
	fmt.Println(me)
	fmt.Println("v1 type is ", reflect.TypeOf(v1))
	fmt.Println("v2 type is ", reflect.TypeOf(v2))
	fmt.Println("v3 type is ", reflect.TypeOf(v3))
	fmt.Println("v4 type is ", reflect.TypeOf(v4))
	fmt.Println("v5 type is ", reflect.TypeOf(v5))
	fmt.Println("我今年", age )
	fmt.Println("我身高", height)
	fmt.Println("v6 type is ", reflect.TypeOf(v6))
	fmt.Println("v7 type is ", reflect.TypeOf(v7))

}

func MyConst() {
	const ( // 常量组,常量定义好之后不可被修改。const必须要有一个默认值
		size int64 = 1024
		eof        =  -1 // 整型常量, 自动推导类型
	)

	const (
		B1 = (iota + 2) * 100 // 每一个枚举(iota)都是从0开始
		B2
		B3
	)


	fmt.Println("这是一个常量", size)
	fmt.Println("这是一个常量", eof)
	fmt.Println(B1, B2, B3 )

}