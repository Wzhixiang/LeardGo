package main

import "fmt"

/*
	结构体（就是java中的实体类）
 	需要注意的是，定义结构体的形式是 type xxx struct {}
*/

type User struct {
	name string
	age  int8
}

func main() {
	user := User{"wzx", 18}

	fmt.Printf("%s is %d year's old", user.name, user.age)
}
