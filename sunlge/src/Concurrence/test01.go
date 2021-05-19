package main

import (
	"flag"
	"fmt"
	"time"
)

func newTask() {
	for i :=0 ;  i<= 3; i++ {
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
}

func main() {
	//创建一个 goroutine，启动另外一个任务
	go newTask()

	//main goroutine 循环打印
	for i :=0 ;  i<= 3; i++ {
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}

	var mode = flag.String("mode", "", "process mode")
	// 每一个flag参数都必须设置Parse()才能够正确解析到这个参数内容
	flag.Parse()
	fmt.Println(*mode)
	if *mode == "sunlge" {
		fmt.Println(" This is test go program")

	} else {
		fmt.Println("测试失败")
	}

	var a int = 10              //声明一个变量，同时初始化
	fmt.Printf("&a = %p\n", &a) //操作符 "&" 取变量地址

	var p *int = nil //声明一个变量p, 类型为 *int, 指针类型
	p = &a
	fmt.Printf("p = %p\n", p)
	fmt.Printf("a = %d, *p = %v\n", a, *p)
	fmt.Printf("%T\n", *p)

	//*p = 111 //*p操作指针所指向的内存，即为a
	//fmt.Printf("a = %d, *p = %d\n", a, *p)

}
