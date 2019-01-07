package main

import (
	"fmt"
)

/*
	条件语句
	if
	if else
	switch
	select
*/
func main() {

	//if 条件语句
	if true {
		fmt.Println("条件为true")
	}

	//if else 条件语句
	if false {
		fmt.Println("条件为true")
	} else {
		fmt.Println("条件为false")
	}

	var position int8 = 3

	//switch 条件语句 方式一
	switch position {
	case 3:
		fmt.Println("当前为3")
	default:
		fmt.Println("默认")
	}

	//方式二
	switch {
	case position == 3:
		fmt.Println("当前为3")
	default:
		fmt.Println("默认")
	}

	//select 还没弄明白
	select {}
}
