package main

import (
	"fmt"
)

type Callback func(...string)

func main() {

	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}
	list("a", "B", "c")
}
