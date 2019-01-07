package main

import "fmt"

/*
	述责
*/
func main() {
	var numbers = [9]int{1, 2, 3, 4}

	var cities = [9]string{"北京", "上海", "广州"}

	for x, n := range numbers {
		fmt.Printf("%d, %d\n", x, n)
	}

	for x, n := range cities {
		fmt.Printf("%d, %s\n", x, n)
	}
}
