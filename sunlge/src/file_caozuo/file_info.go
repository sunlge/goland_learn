package main

import (
	"fmt"
	"os"
)

func main()  {
	// 获取文件信息
	for _, path := range []string{"xxx", "appenfile.log"} {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println(err)
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.IsDir()) // 是不是文件夹，如果是文件的话那么输出结果为false
			fmt.Println(fileInfo.Size()) // 文件总共有多少个字节
			fmt.Println(fileInfo.ModTime())

			if fileInfo.IsDir() {
				dirfile, err := os.Open(path)
				if err == nil {
					defer dirfile.Close()

					childerns, err := dirfile.Readdir(-1)
					for _, children := range childerns {
						fmt.Println(children.Name(),children.Size())
						fmt.Println(err)
					}
				}
			}
		}
	}

}
