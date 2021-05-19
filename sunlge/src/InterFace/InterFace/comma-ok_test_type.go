package main

import "fmt"

type Element01 interface{}

type Person01 struct {
	name string
	age  int
}

func main() {
	list := make([]Element01, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person01{"mike", 18}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	/*  打印结果：
	    list[0] is an int and its value is 1
	    list[1] is a string and its value is Hello
	    list[2] is a Person and its value is [mike, 18]
	*/
}
