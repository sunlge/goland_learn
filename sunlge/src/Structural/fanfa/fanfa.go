package fanfa


import "fmt"

//结构体
type Dog struct {
	name string
}

// 为结构体Dog 定义方法Call (dog *Dog) 这一块叫做接收者 也叫 类型名 类型，由接收者去调用方法Call()
func (ss *Dog) Call()  {
	fmt.Printf("%s: 汪汪\n", ss.name)
}

// 为结构体Dog定义方法SetName
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {

	dog := Dog{"乐乐"}
	m1 := dog.Call
	fmt.Printf("%T\n m1", m1)
	m1()

	dog.SetName("豆豆")
	dog.Call()
	m1()

	dog2 := &Dog{"呵呵"}
	m2 := dog2.Call
	fmt.Printf("%T\n", m2)
	m2()

	dog2.SetName("拉拉")
	dog2.Call()
	m2()

	// //方法表达式， 须显式传参... 方法表达式中，go会自动的去为类型构建一个同名的方法
	q := Dog{
		name: "Bob",
	}
	q1 := (*Dog).Call
	q1(&q)

}
