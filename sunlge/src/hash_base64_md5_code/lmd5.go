package lmd5

import (
	"crypto/md5"
	"fmt"
)

func Lmd5()  {
	fmt.Printf("%X\n", md5.Sum([]byte("sunlge")))
	fmt.Printf("%x", md5.Sum([]byte("sunlge")))
}
