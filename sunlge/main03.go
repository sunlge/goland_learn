package main

import "fmt"

func main()  {
	for i := 'a'; i <= 'z' ; i++ {
		fmt.Printf("%c->%T->%v  ", i, i, i)
	}
	var sunlge = 'a'
	fmt.Printf("\n%v  %T", sunlge, sunlge)
}