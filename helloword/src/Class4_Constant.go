package main

import "fmt"

/*
 * 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
 * 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
 *
 */

/*
 * 全局常量
 */
const (
	name = "全局常量"
)

func main() {

	fmt.Println(name)

	//和变量声明方式相似，只是使用const关键字，并且没有变量 := 声明方式 const identifier [type] = value
	const width = 3
	const height = 4

	fmt.Printf("面积：%d\n", width*height)

	/*
		 	iota
			iota，特殊常量，可以认为是一个可以被编译器修改的常量。
			iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
			iota 可以被用作枚举值：
	*/

	const (
		//第一个iota等于0
		a = iota
		b
		c
		//赋值后，iota就会变成该值
		d = "i am d"
		e
		//恢复计数
		f = iota
		g
		h = 10
		i = iota
		j
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
