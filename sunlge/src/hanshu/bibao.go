package main

import "fmt"

func main()  {

	addBase := func(base int) func(int) int {
		return func(n int) int {
			return base + n
		}
	}

	addBase01 := func(base int) func(a int) int {
		return func(a int) int{
			return base + a
		}
	}

	fmt.Println(addBase(1)(10))
	add2 := addBase(2)
	fmt.Println(add2(2))

	fmt.Println(addBase01(2)(10))

}