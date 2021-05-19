package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//Person定义了方法
func (p *Person) PrintInfo() {
	fmt.Printf("Person: %s,%c,%d\n", p.name, p.sex, p.age)
}

func (p *Person) PrintInfo01() {
	fmt.Printf("Person: %s,%c,%d\n", p.name, p.sex, p.age)
}

type Student struct {
	Person // 匿名字段，那么Student包含了Person的所有字段
	id     int
	addr   string
}

//Student定义了方法
func (s *Student) PrintInfo() {
	fmt.Printf("Student：%s,%c,%d\n", s.name, s.sex, s.age)
}

type Nationality struct {
	ID   int
	Name string
}

func main() {
	p := Person{"mike", 'm', 18}
	p.PrintInfo() //Person: mike,m,18

	s := Student{Person{"yoyo", 'f', 20}, 2, "sz"}
	s.PrintInfo()        //Student：yoyo,f,20
	s.Person.PrintInfo() //Person: yoyo,f,20

	p.PrintInfo01()

}
