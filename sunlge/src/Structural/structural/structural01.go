package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthady time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {

	var me User
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)
	fmt.Println(me)

	// 赋值, 这种方式必须按照顺序并且每一个属性都要赋值, "," 号千万不能省略
	var me2 User = User{1,
		"sunlge",
		time.Now(),
		"Chengdu",
		"151xxxxxxxx",
		"test",
	}
	fmt.Printf("%#v\n", me2)

	// var me3 User 类似这种方式去赋值，如果想要设置一个全是0值的指针需要这种方式
	var me3 User = User{} // 0值默认就是{}
	fmt.Printf("%#v\n", me3)

	// 指定属性名进行初始化, 不用指定顺序，以及可以忽略一些值的赋值
	var me4 User = User{ID: 1, Name: "sunlge", Addr: "chengdu", Tel: "151xxxxxxx"}
	fmt.Printf("%#v\n", me4)

	// 结构体的指针类型
	var pointer *User
	fmt.Printf("pointer : %%T  ---->  %T\n", pointer)

	// 指针类型的默认值都为 nil
	fmt.Printf("pointer : %%#v  ---->  %#v\n", pointer)

	// 指针类型赋值必须为&
	var pointer2 *User = &User{ID: 2}
	fmt.Printf("pointer2: %%#v  ---->  %#v\n", pointer2)
	fmt.Printf("pointer2: %%T  ---->  %T\n", pointer2)
	fmt.Printf("&pointer2: %%T  ---->  %T\n", &pointer2)
	fmt.Printf("&pointer2: %%#v  ---->  %#v\n", &pointer2)
	fmt.Printf("*pointer2: %%#v  ---->  %#v\n", *pointer2)
	fmt.Printf("*pointer2: %%T  ---->  %T\n", *pointer2)

	fmt.Println("##########################")
	var pointer3 *User = &me4 // 先声明一个变量然后在进行赋值，这种方式定义可以修改赋值变量的值。
	fmt.Printf("pointer3: %%#v  ---->  %#v\n", pointer3)
	fmt.Printf("pointer3: %%T  ---->  %T\n", pointer3)
	pointer3.Name = "test"
	fmt.Printf("%#v\n", me4)

	//new 出来的一个对象，也是一个指针对象。
	var pointer4 *User = new(User)
	fmt.Printf("&pointer4: %%#v  ---->  %#v\n", &pointer4)
	fmt.Printf("pointer4: %%T  ---->  %T\n", pointer4)

	// 定义一个匿名的结构体 匿名结构体赋值只有两种方式。
	//    一种通过方法，方法只能使用var 初始化完成后再去调用赋值。
	//    一种通过字面量,字面量只能在简短申明的时候去赋值
	// 应用场景有两种
	//    一种是项目的配置，这个配置只有一份。
	//	  一种是外部开发的时候当作模板使用。
	var me5 struct {
		ID   int
		Name string
	}
	// 通过方法去赋值
	me5.Name = "test"
	fmt.Printf("me5: %%#v  ---->  %#v\n", me5)
	fmt.Printf("me5: %%T   ---->  %T\n", me5)
	// 通过字面量去赋值
	me6 := struct {
		ID   int
		Name string
	}{1, "sunlge"}
	fmt.Printf("me6: %%#v  ---->  %#v\n", me6)
	fmt.Printf("me6: %%T   ---->  %T\n", me6)

	// 结构体的嵌套
	type Address struct {
		Region string
		Street string
		Number string
	}

	type User struct {
		ID   int
		Name string
		Addr Address
	}

	var user01 User // 也可以写成 var user01 = User{}
	fmt.Printf("%#v\n", user01)
	user02 := User{
		ID:   1,
		Name: "Sophie",
		Addr: Address{"成都", "高新区", "01"},
	}
	fmt.Printf("%#v\n", user02)

	addr := Address{"北京", "海淀区", "02"}
	user03 := User{
		ID:   1,
		Name: "Sophie",
		Addr: addr,
	}
	fmt.Printf("%#v\n", user03)
	fmt.Printf("%#v\n", user03.Addr)
	user03.Addr.Street = "昌平区"
	fmt.Printf("%#v\n", user03.Addr)

	type Employee struct {
		User   // 这个User相当于 User User, 也是一种匿名的方式
		Salary float32
	}

	var user04 = Employee{}
	fmt.Printf("%#v\n", user04)

	// 嵌套中使用引用类型
	type User05 struct {
		ID   int
		Name string
		Addr *Address
	}
	// Address{"北京", "海淀区", "02"}
	me07 := User05{
		ID:   1,
		Name: "Bob",
		Addr: &Address{},
	}
	me07.Addr.Number = "02"
	fmt.Printf("%#v\n", me07.Addr)
	fmt.Printf("%#v\n", me07) // 如果写成*me07那么会语法错误

	type User06 struct {
		ID   int
		Name string
		*Address
	}
	// Address{"北京", "海淀区", "02"}
	var me08 User06
	me08 = User06{Address: &Address{"北京", "海淀区", "02"}}
	fmt.Println(me08.Address)
	fmt.Printf("%#v\n", me08.Address)
	fmt.Printf("%T\n", me08)
	fmt.Printf("%#v\n", me08)
	fmt.Println(*me08.Address)
	fmt.Println(me08) // 如果只打印一个me08, 那么它就会输出一个内存地址

	me09 := User06{
		ID:      1,
		Name:    "Sophie",
		Address: &Address{"北京", "海淀区", "03"},
	}
	fmt.Println(*me09.Address)
}
