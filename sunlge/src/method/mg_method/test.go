package main

import "fmt"

type Address struct {
	Region string
	Street string
	Number string
}

// String() 是go 语言中的一个默认方法
func (addr Address) String() string {
	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.Number)
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func (user User) String() string {
	return fmt.Sprintf("[%d]%s: %s", user.ID, user.Name, user.Addr)
}

func main() {
	var u User = User{
		ID:   1,
		Name: "Sophie",
		Addr: Address{"成都市", "高新区", "001"},
	}
	fmt.Println(u)
	fmt.Println(u.Addr)
}
