package main

import (
	"fmt"
)


func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("program errpr")
	return err
}

func main()  {
	err := test()
	fmt.Println(err)
}
