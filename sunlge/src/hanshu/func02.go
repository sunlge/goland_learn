package main

import (
	"fmt"
)

func add1(a, b int) int  {
	return a + b
}

func add2(a, b int, args ...int) int  {
	return 0
}

//callback其实是一个匿名函数
// 给print01定义了一个callback的形参，callback类型是函数类型的
// ...strings的参数是给callback形参， args ...string是给print01的形参
func print01(callback func(...string), args ...string)  {
	fmt.Println("print01函数输出")
	callback(args...)  //args... 是解包，打印传递的值
}

// 定义一个和callback签名一样的函数，然后将list01赋值给callback
func list01(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func main()  {
	fmt.Printf("%T\n", add1)
	fmt.Printf("%T\n", add2)

	var f func(int, int) int =add1
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 2))

	print01(list01, "a", "b", "c")


// 匿名函数
	SayHello := func(name01 string) {
		fmt.Println("hello:", name01)
	}
	SayHello("sunlge")

// 匿名函数直接调用
	func(name string){
		fmt.Println("Hi", name)
	}("Sophie")

	values := func(args ...string) {
		for _,  v := range args{
			fmt.Println(v)
		}
	}
	print01(values, "a", "b", "c")

	print01(func(args ...string) {
		for _, v := range args {
			fmt.Print(v, "\t")
		}
	}, "A", "C", "E")



}
