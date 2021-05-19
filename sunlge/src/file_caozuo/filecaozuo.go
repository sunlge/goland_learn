package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	path := "src/file_caozuo/users.txt"
	if file, err := os.Open(path); err != nil {
		//a := errors.New("The system cannot find the file specified")
		//fmt.Println(a)
		fmt.Println(err)
	} else {
		fmt.Println(file)
		var bytes []byte = make([]byte, 20)

		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(n, err)
				fmt.Println(bytes)
				fmt.Println(bytes[:n])
				fmt.Print(string(bytes[:n]))
			}
		}
		file.Close()
	}
}
