package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(name string) {
	for i := 'a'; i <= 'z' ; i++ {
		// fmt.Printf("%c->%T->%v  ", i, i, i)
		fmt.Println(name, ":", string(i))
		runtime.Gosched()
	}
}

func main()  {
	go PrintChars("1")
	go PrintChars("2")
	PrintChars("3")
	time.Sleep(time.Second * 3)
}