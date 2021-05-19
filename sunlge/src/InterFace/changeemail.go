package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) changeEmail0(NewEmail string) {
	u.email = NewEmail
	fmt.Println("changeEmail0: userName is ",u.name," userEmail is ",u.email)
}

func (u user) changeEmail1(NewEmail string) {
	u.email = NewEmail
	fmt.Println("changeEmail1: userName is ",u.name," userEmail is ",u.email)
}

func main() {
	usr0 := user{"pf", "xxxxxxxxx@qq.com"}
	usr0.changeEmail0("xxxxxxxxx@163.com")
	fmt.Println("main: userName is ",usr0.name," userEmail is ",usr0.email)

	usr1 := user{"pf", "xxxxxxxxx@qq.com"}
	usr1.changeEmail1("xxxxxxxxx@163.com")
	fmt.Println("main: userName is ",usr1.name," userEmail is ",usr1.email)
}
