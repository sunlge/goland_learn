package fanfa

import (
	"fmt"
)

//结构体
type Cat struct {
	name string
}

func (cat Cat) Call() {
	fmt.Printf("%s: 喵喵\n", cat.name)
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func Fanfa_expre()  {
	m1 := Cat.Call
	m2 := (*Cat).SetName
	fmt.Printf("%T,%T\n", m1, m2)

	ll := Cat{"小美"}
	m1(ll)
	ll.SetName("可可")
	m1(ll)
	m2(&ll, "sdf")
}