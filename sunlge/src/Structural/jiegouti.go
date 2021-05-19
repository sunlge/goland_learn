package Structural

import "fmt"

type User struct {
	age    int
	name   string
	gender string
	height float64
	weight float64

}

func Jgt()  {

	// 匿名结构体
	var me10 struct {
		ID  int
		NAME string
	}
	fmt.Printf("%T\n%#v\n", me10, me10)
	me10.ID = 20
	me9 := &me10.NAME
	*me9 = "sunlge2"
	fmt.Printf("%v\n%#v\n", me10, *me9)

	//匿名结构体，通过字面量初始化
	_me := struct {
		ID int
		name string
	}{1, "kk"}
	fmt.Println(_me)

//-------------------------------------------------
	var me User
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)

	var me2 User
	me2 = User{20,"sunlge", "boy", 1.75, 140.52}
	fmt.Println(me2)
	}

