package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {

	message := "sunlge"

	// 编码消息
	EncodeDMessage := base64.StdEncoding.EncodeToString([]byte(message))

	// 输出编码完成的消息
	fmt.Println(EncodeDMessage)

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(EncodeDMessage)
	fmt.Printf("%T ----> %s --->>  %#v\n", data, data, err)
	fmt.Println(data)
	fmt.Println(string(data))

	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(nil)
	}
}