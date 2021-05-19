package main

import (
	"fmt"
)

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

type Employee struct {
	*User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	me = Employee{&User{ID: 1, Name: "Sophie", Addr: Address{"北京", "昌平", "01"}}, 10000, "sunlge"}
	fmt.Printf("%#v\n", me)

	var me01 Employee
	me01 = Employee{&User{ID: 1, Name: "sdf", Addr: Address{"sxv", "sdf", "sxcv"}}, 10000, "sunlge"}
	fmt.Println(*(me.User))
	fmt.Printf("%#v\n", me)
	fmt.Println(me.Name)
	fmt.Println(me.Addr)
	fmt.Println(me01.User.Name)
}
