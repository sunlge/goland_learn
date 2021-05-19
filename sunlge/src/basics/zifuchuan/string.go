package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Compare("abc", "bac"))
	fmt.Println(strings.Contains("abc", "bc"))
	fmt.Println(strings.Count("abcabc", "a"))
	fmt.Println(strings.Fields("abc def\neeee\raaaaa\fxxxxx\vsdddd")) // 按空白字符分割
	fmt.Println(strings.HasPrefix("abcdef", "def"))
	fmt.Println(strings.HasSuffix("abcdef", "def"))
	fmt.Println(strings.Index("defabc", "dabc"))
	fmt.Println(strings.Index("defabcdef", "def"))

	fmt.Println(strings.LastIndex("defabcdef", "def"))
	fmt.Println(strings.Split("bac:def:def:abc", ":"))
	fmt.Println(strings.Join([]string{"abc", "def"}, ":"))

	fmt.Println(strings.Replace("abcabcabcabc", "ab", "sunlge", 2))
	fmt.Println(strings.Replace("abcabcabcabc", "ab", "sunlge", -1))
	fmt.Println(strings.ReplaceAll("abcabcabcabc", "ab", "sunlge"))

	fmt.Println(strings.ToLower("abcABC"))
	fmt.Println(strings.ToUpper("abcABC"))
	fmt.Println(strings.Title("abcABC"))

	fmt.Println(strings.Trim("xyazwxf", "xf"))
	fmt.Println(strings.TrimSpace(" abc xxx \n \r"))

}
