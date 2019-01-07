package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("当前循环到%d\n", i)
	}

	a := 0
	b := 10
	for a < b {
		fmt.Printf("当前循环到%d\n", a)
		a++
	}

	numbers := [6]int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("取numbers中的第%d，值为%d\n", i, x)
	}

}
