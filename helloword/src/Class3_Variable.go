package main

import "fmt"

/*
 * Go 语言变量名由字母、数字、下划线组成，其中首个字母不能为数字。
 */

/*
 * 因式分解关键字的写法一般用于声明全局变量
 */
var (
	describe = "这是一家充满活力，拥有无限可能性的公司"
)

func main() {
	//声明变量的一般形式是使用 var 关键字
	//可以在变量名后声明类型，也可以不声明（）

	//声明name变量不指定类型，赋值为 wzx
	var name = "wzx"
	fmt.Println(name)
	//声明 age int8类型变量，不对变量赋值，使用默认值
	var age int8
	fmt.Println(age)
	//对 age 变量赋值
	age = 18
	fmt.Printf("今年%d岁,", age)

	//省略var声明变量，注意 := ，如果该变量已经声明，会导致编译错误
	city := "广州"
	fmt.Printf("目前在%s", city)

	//多变量声明，可以使用关键字var或者使用:=
	var company, job = "金融会所", "CEO"
	fmt.Printf("分公司--%s，任职%s。%s\n", company, job, describe)

	//
	a, b, c := 1, "小于", 2
	fmt.Printf("%d %s %d", a, b, c)
}
