package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{1,2,6,5,4,3,7}
	sort.Ints(nums)
	fmt.Println(nums)
}