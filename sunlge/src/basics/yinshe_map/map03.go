package main

import (
	"fmt"
	"strings"
)

func main()  {

	article := `sunlge sunlge i me to your love like test name`
 	word := make(map[rune]int) // rune 就是一个int32的类型
	// 统计字符串
 	for _, ch := range article {
 		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
 			word[ch]++
		}
	}
	i := 0
	for k, v := range word {
		i++
		fmt.Printf("%c ===> %d\n", k, v)
	}
	fmt.Println(i)

	// 统计单词数量
	//定义并初始化map
	word01 := make(map[string]int)
	//切分字符串
	str2 := strings.Split(article," ") // strings.Split用来切分
	for _,v := range str2{
		word01[v] ++
	}
	fmt.Println(strings.Reader{})
	fmt.Println(word01)
}
