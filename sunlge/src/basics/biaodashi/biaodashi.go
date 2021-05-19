package biaodashi

import (
	"fmt"
)

func Bdss1()  {
	var yes string
	fmt.Print("你买要一个西瓜吗?(Y/N): ")
	fmt.Scan(&yes)

	switch yes {
	case "y", "yes":
		fmt.Println("买了一个瓜")
	case "Y":
		fmt.Println("买了一个瓜")
	default:
		fmt.Println("不买瓜了")
	}
}

func Bdss2()  {
	var score int
	fmt.Print("请输入你的成绩：")
	fmt.Scan(&score)

	switch  {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("不及格")
	}
}

func Bdsf1()  {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	result := 0
	i := 1
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)
}

func Bdsf2() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 100 {
			break
		}
	}
}

func Bdsf3()  {
	desc := "我爱中国"
	for i, ch  := range desc {
		fmt.Printf("%d %T %q\n", i, ch, ch)
		//fmt.Println(i, ch)
	}
}

func Mygoto()  {
	var op string
	fmt.Println("输入你的操作：(c/u/r/d)")
	fmt.Scan(&op)
}