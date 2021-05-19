package main

import "fmt"

type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct{}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给:", to, ", 消息内容是:", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type SmsSender struct{}

func (s SmsSender) Send(to, msg string) error {
	fmt.Println("发送短信给:", to, ", 消息内容是:", msg)
	return nil
}

func (s SmsSender) SendAll(tos []string, msg string) error {
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

func main() {
	var sender Sender = EmailSender{}
	// var sender Sender = SmsSender{}
	//fmt.Printf("%T  %v\n",sender,  sender)
	//
	//sender.Send("小美", "我爱你")
	//sender.SendAll([]string{"Sophie", "Bob"}, "帮我求婚")
	do(sender)
	sender = SmsSender{}
	do(sender)
}
