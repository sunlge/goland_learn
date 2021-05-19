package main

import (
	"fmt"
	"strconv"
)

func auth(v1 string) bool {
	return true
}

func main() {

	if !auth("a") {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbbb")
	}

	list01 := []string{"a", "b", "c", "false", "true"}
	for i := 0; i < len(list01); i++ {
		fmt.Printf("%T\n", list01[i])
	}
	a := "true"
	b, err := strconv.ParseBool(a) // string 转bool
	fmt.Printf("%T\n %T, %#v", a, b, err)
}
