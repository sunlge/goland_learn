package rizhi

import "log"

func Rizhi()  {
	// SetPrefix 可以加一个前缀
	log.SetPrefix("info: ")
	//  log.SetFlags(log.Flags() | log.Llongfile) Llongfile 打印日志路径用
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.Print("我就是一条日志")
	log.Printf("%s,","谁说我是日志了，我是错误")
	// log.Panic("哈哈，我好痛")
}
