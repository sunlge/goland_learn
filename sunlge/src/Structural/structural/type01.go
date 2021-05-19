package main

import "fmt"

type Counter int

type Users map[string]string

func main() {
	var counter Counter = 10
	counter += 10
	fmt.Println(counter)
	fmt.Printf("%T\n", counter)

	me := make(Users)
	me["name"] = "test"
	fmt.Printf("%T, %T\n", me, counter)
}
