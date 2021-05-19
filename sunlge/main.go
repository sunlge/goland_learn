package main

import "fmt"

func main() {
	// 变量使用中，_相当于一个回收站，类似于/dev/null这样的东西，所有以_定义的变量都会被丢弃
	// fmt.Println("My first go program")
	// test1.Test()
	aaa()
	_, str := aaa() // 这里的_相当于把250丢弃了。
	fmt.Println(str)
	fmt.Printf("%#v\n", aaa)
}

func aaa() (int, string) {
	fmt.Println("Ha Ha Ha")
	return 250, "sb"
}
