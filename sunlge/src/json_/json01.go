// json 编码可以参考 https://learnku.com/go/t/23565/golang-json-coding-and-decoding-summary
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
	json.Marshal 序列化 内存 => 字符串/字节切片
	json.Unmarshal 反序列化 字符串/字节切片 => 内存
	json.MarshalIndent  可以格式化输出一个json格式
	json.Valid //判断json格式的字符串是否正确
	Marshal 是一次性读取
	NewEncoder 是流的读取，一点一点读
	*/

	names := []string{"Sunlge", "Sophie", "Bob", "Jack"}
	users := []map[string]string{{"name": "Jiessie", "Addr": "ShangHai"}, {"name": "Sophie", "Addr": "Beijing"}, {"name": "Bob", "Addr": "ChengDu"}}

	bytes, err := json.Marshal(users) // 返回两个值，一个是bytes切片，一个是错误内容，如果转换成功，那么错误内容为nil
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}



	bytes01, err01 := json.Marshal(names) // 返回两个值，一个是bytes切片，一个是错误内容，如果转换成功，那么错误内容为nil
	if err01 == nil {
		fmt.Println(string(bytes01))
	} else {
		fmt.Println(err01)
	}

	var names01 []string
	err02 := json.Unmarshal(bytes01, &names01)
	fmt.Println("********错误信息**********")
	fmt.Println(err02)
	fmt.Println(names01)
	names02, _ := json.MarshalIndent(names01, "", "\t") //格式化输出
	fmt.Println(string(names02))
	fmt.Println(json.Valid([]byte("[]x")))
	fmt.Println(json.Valid([]byte(names02)))
}
