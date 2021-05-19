package main

import (
	"fmt"
	"strings"
)



func RemovePrefix(str string) string {
	// 返回移除字符串的前缀 go 的值
	return strings.TrimPrefix(str, "go")
}


func Sp(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
		fmt.Println(list[index])
	}
}

func t1(a string) string {
	return "sunlge"
}

func main() {
	fmt.Println(RemovePrefix("go test"))
	// 待处理的字符串列表
	list := []string{
		"go sunlge",
		"go parser",
		"go printer",
	}

	// 处理函数链
	chain := []func(string) string {
		RemovePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 处理字符串
	Sp(list, chain)

	//输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}

	Sp([]string{"s"}, []func(string) string {t1})

}