package CommandLine

import (
	"flag"
	"fmt"
)

func Mlh()  {
	var host string
	var port int
	var help bool
	// 绑定命令行参数与变量关系
	flag.IntVar(&port, "p", 22, "ssh port")
	flag.StringVar(&host, "H", "127.0.0.1", "ssh IP")
	// flag.String 这种写法定义的参数选项必须使用
	var species = flag.String("species", "gopher", "the species we are studying")

	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("Usage: flagargs -H -p -Species  Species must be specified")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	//fmt.Printf("%s:%d\n",host,port)
	fmt.Println(*species)
	//fmt.Println(help)

	if help {
		flag.Usage()
	} else {
		fmt.Println(host, port)
		fmt.Println(flag.Args())
	}


}
