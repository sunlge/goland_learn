package zhizhen

import "fmt"

func Zhizhen()  {
	a := 2
	b := 3
	fmt.Println(a,b)
	var d *int
	d = &a
	fmt.Println(&a)
	fmt.Println(d)
	*d = 4
	fmt.Println(*d)
	fmt.Println(a)

//------------------------------------
	A := 2
	B := A
	B = 3
	fmt.Println(A, B)
	/*
	指针， 指针的类型就是 如果值为int，那么指针类型就是*int，如果为bool，那么指针类型就是*bool *相当于调用内存地址，输出为内存地址所存储的值
	如果写&，那么这个值就是内存地址
	 */
	var  cc *int = &A
	//cc = &A
	c := &A
	fmt.Printf("%T %T\n", cc, c)
	fmt.Println(*cc, *c)

	var name string
	fmt.Print("请输入名字: ")
	fmt.Scan(&name)
	fmt.Println("你输入的内容为: ", name)

	var age int
	fmt.Print("输入年龄：")
	fmt.Scan(&age)
	fmt.Println("你输入的内容为: ", age)

	qq := "hehehe"
	fmt.Println(qq, &qq)
}
