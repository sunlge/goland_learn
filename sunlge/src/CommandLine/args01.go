package CommandLine

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args //获取用户输入的所有参数

	//如果用户没有输入,或参数个数不够,则调用该函数提示用户
	if args == nil || len(args) < 2 {
		fmt.Println("err: 没有输入 ip port")
		return
	}
	ip := args[1]   //获取输入的第一个参数
	port := args[2] //获取输入的第二个参数
	fmt.Printf("ip = %s, port = %s\n", ip, port)
}

