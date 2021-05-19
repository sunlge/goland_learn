// 关键字 defer ⽤于延迟一个函数或者方法（或者当前所创建的匿名函数）的执行。注意，defer语句只能出现在函数或方法的内部。
package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("main over")
}