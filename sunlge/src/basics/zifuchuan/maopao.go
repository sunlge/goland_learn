package main

import (
	"fmt"
)

func main()  {
	heights := []int{10, 6, 7, 5, 8, 9, 2, 4, 3, 1}
	// 先把最高的人移动到最后
	// 从第一个人开始，两两开始比较，如果前面的人高，这两个人就交换位置
	// n 个人比较 n -1 ,注意是从0开始
	for j := 0; j < len(heights) -1; j++ {
		fmt.Println("第", j, "轮交换")
		for i := 0; i < len(heights) -1; i++ {
			if heights[i] > heights[i + 1] {
				fmt.Println("交换：", heights[i], heights[i+1])
				// 如果要进行交换可以 a , b = b, a
				// 也可以借助中间数来交换，也就是需要多申明一个变量
				tmp := heights[i]
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println("第", i , "次", "交换", heights)
		}
		fmt.Println("第", j, "轮, 结果: ", heights)
	}
	//sort.Ints(heights)
	//fmt.Println(heights)
}