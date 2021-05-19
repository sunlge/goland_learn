package suijishu

import (
	"fmt"
	"math/rand"
	"time"
)

func Sjs(a int)  {
	rand.Seed(time.Now().Unix())
	// 0 - 100
	fmt.Println(rand.Int() % 100)
	// 或者用下面写法
	fmt.Println(rand.Intn(a))
}
