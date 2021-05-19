package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main()  {
	bytes01 := []byte{}
	strs  := []string{}
	ints  := []int{}
	fmt.Printf(" %T %#v \n %T %#v \n %T  %#v \n", bytes01, bytes01, strs, strs, ints, ints)

	// 字节转字符串
	fmt.Println("字节转字符串")
	s := string(bytes01)
	fmt.Printf("%T %#v \n", s, s)

	// 字符串转byte
	bs := []byte(s)
	fmt.Printf("%T %#v\n", bs, bs)

	fmt.Println(bytes.Compare([]byte("abc"), []byte("def")))
	fmt.Println(bytes.Index([]byte("abcdefabc"), []byte("def")))
	fmt.Println(bytes.Contains([]byte("abcdefabc"), []byte("defd")))

	s01 := "我爱小美"
	fmt.Println(len(s01))
	fmt.Println("转为utf8格式的bytes，utf8中一个中文等于一个bytes")
	fmt.Println(utf8.RuneCountInString(s01))

	// 字符串转变未bool
	v := "true"
	if s02, err := strconv.ParseBool(v); err == nil {
		fmt.Println(s02)
	}

	// int
	if v, err := strconv.Atoi("1024"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}


}