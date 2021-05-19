package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID    int
	Name  string
	Score int
}

func main() {
	list := [][2]int{{1, 3}, {5, 9}, {8, 7}, {4, 5}}
	list = append(list, [2]int{3, 4})

	fmt.Printf("%T ---> %d\n", list, cap(list))
	fmt.Println(len(list))
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	fmt.Println("****进行排序****")

	// 排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
	fmt.Println(list)

	users := []User{{1, "Bob", 4}, {2, "Sophie", 2}, {3, "sunlge", 7}}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Score < users[j].Score
	})

	fmt.Println(users)
}
