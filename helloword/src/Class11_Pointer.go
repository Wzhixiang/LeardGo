package main

import "fmt"

/*
	指针
*/
func main() {
	a := 10
	fmt.Printf("变量的地址：%d\n", &a)

	//指针变量
	var p *int
	p = &a

	fmt.Printf("指针变量地址：%d\n", p)
	fmt.Printf("指针变量值：%d", *p)
}
