package main

import (
	"fmt"
)

type Humaner  interface {
	SayHi()
}

type Student struct {
	name string
	grade int
	say string
}


func (s Student) SayHi()  {
	fmt.Printf("[%d]-[%s]: %s\n", s.grade, s.name, s.say)
}

// 多态的应用，同一个方法使用不同的数据,不使用这个值也是可以的
func Say(a Humaner) {
	a.SayHi()
}

func main()  {
	s := Student{"Sophie", 0602, "好好学习，天天向上"}
	s1 := Student{"Bob", 0702, "好好学习，天天向上"}
	s2 := Student{"Sunlge", 0603, "好好学习，天天向上"}
	s3 := Student{"Jack", 0604, "好好学习，天天向上"}
	s.SayHi()
	Say(s1)
	Say(s2)
	Say(s3)
}