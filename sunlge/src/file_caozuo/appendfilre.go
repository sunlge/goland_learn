package main

import (
	"log"
	"os"
	"time"
)

func main()  {
	path := "src/file_caozuo/appenfile.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		//file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		//file.WriteString("\n")
		log.SetOutput(file)
		log.SetPrefix("user")
		log.SetFlags(log.Flags() | log.Llongfile)
		log.Print(time.Now().Unix(),"\ttest")
		file.Close()
	}
}
