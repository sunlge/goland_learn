package main

import (
	"fmt"
	"os"
)

func main()  {
	path := "users2.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		_, err := file.Write([]byte("abctest"))
		if err != nil {
			fmt.Println(err)
		} else {
			file.Close()
		}
	}
}