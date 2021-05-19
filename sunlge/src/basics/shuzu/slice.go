// 数组是一个值类型
// 切片是一个*类型, 如果一个切片01的值赋值给另外一个切片02,如果修改02的内容，那么01也会改变。因为赋值的时候是直接传递的内存地址

package main

import (
	"fmt"
)

func main()  {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v\n", arr)
	s := arr[1:6:8]  // 因为是从0开始，所以容量是 8-1，如果从0开始，那么容量就是8 公式就是cap=cap - start  len=end-start
	fmt.Printf("%#v, %d\n", s, cap(s))

}


// 示例  array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// array[:6:8]	[0 1 2 3 4 5]	6	8	省略 low
// array[5:]	[5 6 7 8 9]	5	5	省略 high、 max
// array[:3]	[0 1 2]	3	10	省略 high、 max
// array[:]	[0 1 2 3 4 5 6 7 8 9]	10	10	全部省略

