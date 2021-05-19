package main

import "fmt"

type Element interface{}

type Person struct {
	name string
	age  int
}

func main() {
	list := make([]Element, 3)
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person{"mike", 18}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
	fmt.Printf("%T ----> %T\n", list[1], list[2])
}
