package main

import "fmt"

/*
	类型转换
*/
func main() {

	a, b := 15, 2

	var merchant float32

	//type_name(expression) type_name类型，expression表达式
	merchant = float32(a) / float32(b)

	fmt.Printf("int类型转换float类型：%f", merchant)
}
