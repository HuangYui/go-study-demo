package main

import (
	"errors"
	"fmt"
)

// error 类型的异常 不会导致程序中断
func main() {
	res1, err1 := div(1, 1)
	fmt.Println(res1, err1)

	res2, err2 := div(1, 0)
	fmt.Println(res2, err2)

	//返回一个error
	e := fmt.Errorf("自定义error")
	fmt.Println(e)
}

func div(n, m int) (int, error) {
	if m == 0 {
		return 0, errors.New("0不能作为分母")
	}
	return m / n, nil
}
