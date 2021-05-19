package test1

import (
	"fmt"
	"github.com/astaxie/beego"
)

func init()  {
	fmt.Println("Hello This is a init program")
}

func Test() {
	fmt.Println("My first go program")
}

func Test2() {
	beego.Run()
}