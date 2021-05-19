package main

import (
	"bytes"
	"fmt"
)

func main()  {

	hammer := "sunlge"
	name := "Sophie"

	// 字符串写入缓冲
	var StringBuilder  bytes.Buffer
	StringBuilder.WriteString(hammer)
	StringBuilder.WriteString(name)

	// 将缓冲以字符串形式输出
	fmt.Println(StringBuilder.String())
	fmt.Printf("%T\n %v\n %#v\n", StringBuilder, StringBuilder, StringBuilder)

}