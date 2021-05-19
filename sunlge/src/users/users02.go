package main

import (
	"fmt"
	"strings"
)

// 添加用户
// string := strconv.Itoa(int) 等价于  string := strconv.FormatInt(int64(int), 10)

func AddUser(pk int, users map[string]map[string]string) {
	var (
		id   string = fmt.Sprintf("%d", pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入联系方式：")
	fmt.Scan(&tel)
	fmt.Print("请输入地址：")
	fmt.Scan(&addr)
	users[id] = map[string]string{
		"id":   id,
		"age":  age,
		"name": name,
		"tel":  tel,
		"addr": addr,
	}
}

// 查询用户
// q == "" 查找全部
// q != "" 必须有 名称,电话 或者住址任意一个属性 中包含q内容的显示

func Query(users map[string]map[string]string) {
	var q string
	fmt.Print("请输入查询信息: ")
	fmt.Scanln(&q)
	fmt.Print("您输入的为:", q)
	title := fmt.Sprintf("%10s|%10s|%10s|%10s|%10s", "id", "age", "name", "addr", "tel")
	fmt.Printf(strings.Repeat("-", len(title)))
	fmt.Println()
	fmt.Println(title)
	fmt.Printf(strings.Repeat("-", len(title)))
	fmt.Println()
	for _, user := range users {
		if q == " " || strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) || strings.Contains(user["id"], q) {
			fmt.Printf("%10s|%10s|%10s|%10s|%10s", user["id"], user["age"], user["name"], user["addr"], user["tel"])
			fmt.Println()
			fmt.Println(strings.Repeat("-", len(title)))
			//			fmt.Printf(strings.Repeat("-",len(title)))
			//			fmt.Println()
		} else if q == " " {
			fmt.Println(users)
			break
		}
	}
}

func main() {
	// 存储用户信息
	// users := map[string]map[string]string{"1": map[string]string{"name": "Sophie", "tel": "123", "addr": "ww", "age": "18"} }
	users := map[string]map[string]string{"1": {"name": "Sophie", "tel": "123", "addr": "ww", "age": "18", "id": "1"}}
	users["2"] = map[string]string{
		"id":   "2",
		"age":  "12",
		"name": "sunlge",
		"tel":  "123",
		"addr": "qq",
	}

	id := len(users)
	fmt.Println("欢迎来到小娄实验室")
	flags := true
	for flags {
		var op int
		fmt.Print(`
1. 新键用户
2. 修改用户
3. 查询用户
4. 删除用户
5. 退出程序
请输入指令：`)
		fmt.Scan(&op)
		switch op {
		case 1:
			id++
			AddUser(id, users)
		case 2:
			fmt.Println("修改用户")
		case 3:
			Query(users)
		case 4:
			fmt.Println("删除用户")
		case 5:
			flags = false
		default:
			fmt.Println("您输入错误")
		}
	}
}
