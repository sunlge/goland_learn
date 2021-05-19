package Structural

import "fmt"



func NmzzJgt()  {
	type Address struct {
		Region string
		Street string
		No     string
	}

	type User struct {
		ID int
		Name string
		Addr Address
	}

	type Employee struct {
		*User
		Salary float64
		Name   string
	}

	var me Employee
	me =  Employee{&User{ID: 1, Name: "sdf", Addr: Address{"sxv", "sdf", "sxcv"}},10000, "sunlge"}
	fmt.Println(*(me.User))
	fmt.Printf("%#v\n", me)
	fmt.Println(me.Name)
	fmt.Println(me.Addr)
	fmt.Println(me.User.Name)
}
