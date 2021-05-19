package main

import (
	"fmt"
)

func main()  {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)
	// 字面量设置
	nums = [10]int{1,2,3,4}
	// 下面如果在设置值那么上面设置的值就会被遗忘，因为它是 重新赋值的含义
	nums = [10]int{3: 10,4: 30, 9: 20}
	// 下面这种设置是通过索引更改单个的值
	nums[2], nums[5] = 15, 35
	fmt.Println(len(nums))
	fmt.Println(nums)
	a := []int{} // [] 这种写法是切片，不是数组
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)
	b := [...]int{1,2} // 写三个...如果不赋值会报错,...这也是数组的表现形式
	fmt.Println(b)
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	// 使用range方法遍历数组
	//for index, value := range nums{
	//	fmt.Println(index, ":", value)
	//}
	fmt.Printf("%T, %T, %T\n", nums, a, b)
	var sunlge [10][1]string
	sunlge[0][0] = "123"
	sunlge[1][0] = "sdf"
	fmt.Println(sunlge)

	// 多维类型不能用...
	var marrays [10][2]string
	marrays = [10][2]string{{"1sdf"}, {"3sdf", "4sdf"}, {"5xcv", "6sdfx"}}
	fmt.Println(marrays)
	fmt.Println("值为空： ",marrays[0][1])

}
