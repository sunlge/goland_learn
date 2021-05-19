// **值类型、引用类型**
// * 将变量赋值给新的一个变量，并修改新变量的值，如果对旧变量有影响则为引用类型，无影响则为值类型

// 引用类型：slice、map
// 值类型： int bool float array

package main

import "fmt"


func main()  {
	type vertex struct {
		x,y int
	}

	var s = map[string]int{
		"test": 1,
		"sss": 2,
	}
	var m = map[string]vertex{
		"home": vertex{0, 0},
		"school": vertex{1,1},
	}

	l := make(map[string]int)
	l ["apple"] = 2

	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(l)
}