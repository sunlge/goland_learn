package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Region string
	Street string
	Number string
}

type ReturnValue string

// ReturnValue 是设置的 string 自定义类型的一个返回值
func (addr Address) method() ReturnValue {
	return ReturnValue(fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.Number))
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func (user User) method() ReturnValue {
	return ReturnValue(fmt.Sprintf("[%d]%s: %s", user.ID, user.Name, user.Addr))
}

func main() {
	var u User = User{
		ID:   1,
		Name: "Sophie",
		Addr: Address{"成都市", "高新区", "001"},
	}
	fmt.Println(u)
	fmt.Println(u.Addr)

	var a ReturnValue
	var b string
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
