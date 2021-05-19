package main

import "fmt"

func tower(a, b ,c string, layer int)  {
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	// a layer-1 借助 c 移动 b
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	// b layer -1 借助a 移动到c
	tower(b, a, c, layer -1)
}

func main()  {
	fmt.Println("3曾：")
	tower("A", "B", "C", 3)
}
