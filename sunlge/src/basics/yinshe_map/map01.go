package main

import "fmt"

func main()  {
	var scores map[string]int
	fmt.Printf("%T,%#v\n", scores, scores)  // 没有初始化的map会返回一个nil
	fmt.Println(scores == nil)

	// 字面量
	scores = map[string]int{"乐乐": 1, "小白": 8}
	fmt.Println(scores)

	// 这种赋值方式不会给nil，会返回一个映射类型
	scores = make(map[string]int)
	fmt.Printf("%T,%#v\n", scores, scores)


	// add, delete, change, check
	// key
	scores = map[string]int{"乐乐": 10, "小白": 8}
	fmt.Println(scores["乐乐"])
	fmt.Println(scores["小黑"]) // 当这个key不存在的时候，会返回定义values的一个初始值，int的初始值就是0值

	value, ok := scores["小黑"]
	fmt.Println(value, ok) // go语言中如果key值存在会返回一个true,不存在返回一个false

	if v, back := scores["乐乐"]; back {
		fmt.Println(v)
	}

	scores["小黑"] = 6
	// 删除 乐乐 这个key
	delete(scores, "乐乐")
	fmt.Println(scores["乐乐"])
	scores["小黑"] =  7
	scores["小黑"] =  8 + 1
	fmt.Println(len(scores))
	for k , v01 := range scores {
		fmt.Println(k, v01)
	}

	fmt.Println(scores)

}