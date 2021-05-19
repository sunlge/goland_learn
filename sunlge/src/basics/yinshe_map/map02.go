package main

import (
	"fmt"
)

func main()  {
	users := []string{"Jack", "Tom", "Tom", "Sunlge", "Jack", "Sunlge", "Sophie","Tom", "Sophie", "Sophie", "Sophie"}
// 第一种方式
	vote := make(map[string]int)
	for index, user  := range users {
		fmt.Println(index, user)
		if _, ok := vote[user]; ok {
			vote[user] =1
		} else {
			vote[user] += 1
		}
//	fmt.Println(vote)
	}
	fmt.Println("---------------------------" )
// 第二种方式
	vote01 := make(map[string]int)
	for i := 0; i < len(users); i++ {
		fmt.Println(vote01[string(users[i])])
		if _, back := vote01[string(users[i])]; back {
			// +1
			vote01[string(users[i])]++

		} else {
			// 赋值
			vote01[string(users[i])] = 1
		}
	}
	fmt.Println(vote01)
}

