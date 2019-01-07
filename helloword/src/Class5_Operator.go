package main

import "fmt"

/*
	运算符：

	1.算数运算符：+加 -减 *乘 /除 %余 ++自增 --自减
	2.关系运算符：==相等 !=不等 >大于 <小于 >=大于等于 <=小于等于
	3.逻辑运算符：&&与 ||或 ！非
	4.位预算符：&按位与 |按位或 ^按位异或 <<左移 >>右移
	5.赋值预算符：= += -= *= /= %= <<= >>= &= |= !=

	运算优先级：
	从上往下
	^ !
	* / % << >> & &^
	+ - | ^
	== != < <= >= >
	<-
	&&
	||
*/
func main() {

	//todo 算数运算
	a, b := 10, 5
	fmt.Printf("%d + %d = %d \n"+
		"%d - %d = %d \n"+
		"%d * %d = %d \n"+
		"%d / %d = %d \n"+
		"%d 余 %d = %d \n",
		a, b, a+b,
		a, b, a-b,
		a, b, a*b,
		a, b, a/b,
		a, b, a%b)

	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	//todo 关系运算
	if a == b {
		fmt.Println("a等于b")
	}

	if a != b {
		fmt.Println("a不等于b")
	}

	if a > b {
		fmt.Println("a大于b")
	}

	if a < b {
		fmt.Println("a小于b")
	}

	if a >= b {
		fmt.Println("a大于等于b")
	}

	if a <= b {
		fmt.Println("a小于等于b")
	}

	//逻辑运算
	if a > 0 && b > 0 {
		fmt.Println("a和b都大于0")
	}

	if a > 5 || b > 5 {
		fmt.Println("a大于5或者b大于5")
	}

	if !(a > 0 && b > 0) {
		fmt.Println("a和b都小于0")
	}

	/*
		 	位运算
			&按位与：二进制同位同为1时，结果为1，否则为0
			|按位或：二进制同位只要有1，结果为1，否则为0
			^按位或：二进制同位相同为0，不同为1
		    << n 向左移动n位
			>> n 向右移动n位
	*/
	a = 60 // 0011 1100
	b = 13 // 0000 1101

	//0000 1100 = 12
	fmt.Println(a & b)
	//0011 1101 = 61
	fmt.Println(a | b)
	//0011 0001 = 49
	fmt.Println(a ^ b)
	//1111 0000 = 240
	fmt.Println(a << 2)
	//0000 1111 = 15
	fmt.Println(a >> 2)

	/*
		赋值运算
		a = b 将b赋值给啊
		a += b a = a + b
		a -= b a = a - b
		a *= b a = a * b
		a /= b a = a / b
		a %= b a = a % b
		a <<= b a = a << b
		a >>= b a = a >> b
		a &= b a = a & b
		a |= b a = a | b
	*/
	a = 12
	fmt.Println(a)
	a += 1
	fmt.Println(a)
	a -= 1
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 5
	fmt.Println(a)
	a <<= 2
	fmt.Println(a)
	a >>= 2
	fmt.Println(a)
	a &= 0
	fmt.Println(a)
	a |= 1
	fmt.Println(a)

	//变量地址
	fmt.Println(&a)

	//指针变量
	var c *int
	c = &a
	fmt.Println(*c)
}
