package cuowu

import (
	"errors"
	"fmt"
)

func Cuowu(a, b int) (int, error)  {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b , nil
}

func main()  {
	fmt.Println(Cuowu(1, 3))
	if v, err := Cuowu(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	e := fmt.Errorf("Errpr %s", "division by zero")
	fmt.Printf("%T, %v\n", e, e)
}