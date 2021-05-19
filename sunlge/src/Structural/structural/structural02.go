package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthady time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {

	var me User = User{1,
		"sunlge",
		time.Now(),
		"Chengdu",
		"151xxxxxxxx",
		"test",
	}
	fmt.Printf("%#v\n", me)

	me2 := &User{
		ID:   2,
		Name: "test",
	}
	fmt.Println(me2)

	(*me2).Name = "test2"
	// 也可以写成me2.Name 它会自动的解引用
	me2.Addr = "北京"
	fmt.Println(me2)

}
