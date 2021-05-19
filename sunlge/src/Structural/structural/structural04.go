package main

import "fmt"

type Address struct {
	Region string
	Street string
	Number string
}

type User struct {
	ID   int
	Name string
	Addr *Address
}

func change(u User) {
	u.Name = "Sophie"
	u.Addr.Region = "北京"
}

func Pointerchange(u *User) {
	u.Name = "Sophie"
	u.Addr.Region = "北京"
}

// 使用函数来创建strct
func NewUser(id int, name string, region, street, number string) User {
	return User{
		ID:   id,
		Name: name,
		Addr: &Address{region, street, number},
	}
}

// 也可以写指针类型的
func NewUser91(id int, name string, region, street, number string) *User {
	return &User{
		ID:   id,
		Name: name,
		Addr: &Address{region, street, number},
	}
}

func main() {

	me3 := NewUser(3, "Bob", "上海", "陆家嘴", "3")
	me4 := NewUser91(3, "Bob", "上海", "陆家嘴", "3")
	fmt.Println(me3)
	fmt.Println(me4)
	var me User = User{}
	me2 := me
	me2.Name = "Sophie"
	fmt.Println(me2)
	fmt.Println(me2.Addr) // 此时访问Addr它是一个空的值， nil
	// go 语言中nil值是无法进行处理的，否则会报错
	/*
		在 Go 中 * 代表取指针地址中存的值，& 代表取一个值的地址
		对于指针，我们一定要明白指针储存的是一个值的地址，但本身这个指针也需要地址来储存
		如上 Addr 是一个指针，他的值为内存地址 0xc04204a080
		而 Addr 的内存地址为 0xc042068018
		内存地址 0xc04204a080 储存的值为 1
		地址 0xc042068018 0xc04204a080
		值 0xc04204a080 1

		报错：
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x1 addr=0x0 pc=0x498025]

		报这个错的原因是 go 初始化指针的时候会为指针 Addr 的值赋为 nil ，但 Addr 的值代表的是 *Addr 的地址，
		nil 的话系统还并没有给 *Addr 分配地址，所以这时给 *Addr 赋值肯定会出错。
		解决这个问题非常简单，在给指针赋值前可以先创建一块内存分配给赋值对象即可
	*/
	me2.Addr.Region = "北极"

}
