package main

import (
	"errors"
	"fmt"
)

/*
	TODO 错误

	通过errors。New(string)抛异常
*/

func divide(a int32, b int32) (int32, error) {
	if b <= 0 {
		return 0, errors.New("被除数不能小于等于0")
	}

	return a / b, nil
}

func main() {

	fmt.Println(divide(3, 2))
	fmt.Println(divide(4, 0))
}
