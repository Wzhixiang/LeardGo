package main

import (
	"fmt"
	"math"
)

/*
	函数
*/
func main() {

	log("没有返回值")

	fmt.Println(add(1, 1))

	a, b := judge("第一个参数", "第二个参数")

	fmt.Println(a, b)

	fmt.Println(math.Max(10, 9))
}

func log(msg string) {
	fmt.Println(msg)
}

/*
	单一返回值
*/
func add(num1, num2 int) int {
	return num1 + num2
}

/**
多返回值
*/
func judge(key, value string) (string, string) {
	return key, value
}
