package main

import "fmt"

type SignalSender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
	EmaAddr string
}
func (s EmailSender) Send(to, msg string) error  {
	fmt.Println("发送邮件给:", to, ", 消息内容是:", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

func (s EmailSender) A() {
	fmt.Println("调用方法A")
}

type SmsSender struct {}
func (s SmsSender) Send(to, msg string) error  {
	fmt.Println("发送短信给:", to, ", 消息内容是:", msg)
	return nil
}

func (s *SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

//func do(sender EmailSender) {
//func do(sender SmsSender) {
func do(sender Sender) {
	sender.Send("亲朋", "结婚消息")
}

func main()  {
	var sender Sender = EmailSender{}
	// var sender Sender = SmsSender{}
	//fmt.Printf("%T  %v\n",sender,  sender)
	//
	//sender.Send("小美", "我爱你")
	//sender.SendAll([]string{"Sophie", "Bob"}, "帮我求婚")
	do(sender)
	sender = &SmsSender{}
	do(sender)

	// 接口传递给接口，重点是接口里面定义的方法类型值要一致才可以进行传递
	var sender01 SignalSender = sender
	sender01.Send("Jack", "周六一起吃饭吗？")
	sender01.SendAll([]string{"Jack", "Jack", "Sophie"}, "明天一起去野炊吗？")

	// 接口断言，将sender01的赋值SignalSender, 也就是&SmsSender
	// go语言里接口断言和类型断言一样，是用于判断某个对象是否有接口的所有方法
	sender02, ok1 := sender01.(SignalSender)
	fmt.Printf("返回是&SmsSender接口的类型和值：%T --> %v, %v\n",sender02, sender02, ok1)
	sender02.SendAll([]string{"Sophie", "Bob"}, "周日一起复习功课吗？")

	fmt.Println("#####")
	fmt.Println(sender01.(SignalSender))
	fmt.Printf("%T", sender01.(SignalSender))
	// 赋值为结构体*SmsSender 的值
	sender03, ok2 := sender01.(*SmsSender)
	fmt.Printf("%T, %v\n", sender03, ok2)
	// 下面这个会转换失败，因为sender01的本身没有EmailSender的类型
	sender04, ok3 := sender01.(EmailSender)
	sender04.A()
	fmt.Printf("%T, %v\n", sender04, ok3)
	if !ok3 {
		sender04.A()
		fmt.Println("判断失败，但是调用成功。。这是因为sender04 现在等是EmailSender实例出来的对象")
	}

	sender01  = EmailSender{"EmaAddr"}
	// 接口断言，接口查询，查找sender01的类型，看看它现在属于哪个类型
	switch v := sender01.(type) { //.(type)只能用在switch语句中
	case EmailSender:
		fmt.Println("EmailSender", v.EmaAddr)
	case *SmsSender:
		fmt.Println("SmsSender")
	default:
		fmt.Println("不匹配所有类型")
	}
}