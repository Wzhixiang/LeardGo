package main

import "fmt"

/*
 * 惯例学习一门语言，先从最基本的"Hello Word!"开始世界篇章
 *
 * 接下来，讲解下Go语言结构
 * package main：包声明。main表示这是一个可独立执行的程序，每个Go程序都包含一个main包
 * import "fmt"：引入包。fmt 包实现了格式化 IO（输入/输出）的函数
 * func：函数修饰。
 * main()：函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
 *
 * 语法：
 * ……
 */
func main() {
	fmt.Print("Hello word!")
}
