package main

import "fmt"

func AddN(a, b int, args ...int) int {
	c := a + b
	for _, s_  := range args {
		c = c + s_
	}
	return c
}

func main() {
	a := AddN(1, 2,3, 4)
	fmt.Println(a)
}


