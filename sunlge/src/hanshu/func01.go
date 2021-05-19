package main

import "fmt"

func AddN(a, b int, args ...int) int {
	c := a + b
	for _, s_  := range args {
		c = c + s_
	}
	return c
}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return AddN(a, b, args...) //解包，将AddN的args参数传递进来
	}
	return -1   // 必须定义一个return -1 || 0
}

func main() {
	fmt.Println(calc("add", 1, 2, 3 ))

	nums := []int{1,3,5,7}
	fmt.Println(nums[2:])
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
}


