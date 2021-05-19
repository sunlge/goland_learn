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
	Addr Address
}

func change(u User) {
	u.Name = "Sophie"
	u.Addr.Region = "北京"
}

func Pointerchange(u *User) {
	u.Name = "Sophie"
	u.Addr.Region = "北京"
}

func main() {
	var me User = User{Name: "Bob"}
	me.Addr.Region = "北极"
	fmt.Println(me.Name)
	change(me)
	fmt.Println("****使用值类型****")
	fmt.Println(me.Name)
	fmt.Println(me.Addr.Region)
	fmt.Println("****使用引用类型****")
	Pointerchange(&me)
	fmt.Println(me.Name)
	fmt.Println(me.Addr.Region)

}
