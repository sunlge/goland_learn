package main

import "fmt"

func main()  {
	list := [][2]int{{1,3}}
	list = append(list, [2]int{5,4})

	for i := 0 ;i < len(list); i++ {
		fmt.Println(list[i])
	}
}