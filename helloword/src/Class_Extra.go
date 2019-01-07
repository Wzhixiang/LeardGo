package main

/*
	todo import 导包时，在包名前面添加下划线表示只要初始化该类
*/
import (
	"fmt"
	_ "hello"
)

func main() {
	s := []int{0, 1, 1, 2, 3, 5, 8, 13}

	/*
		todo 在代码中使用下划线，代表忽略该变量
	*/
	for _, v := range s {
		fmt.Println("value = ", v)
	}
}
