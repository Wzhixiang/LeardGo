package main

import "fmt"

/*
	定义接口：type interface_name interface{}
*/
type Phone interface {
	call() string
}

/*
	定义结构体
*/
type HuaWeiPhone struct {
}

/*
	定义机构体
*/
type MiPhone struct {
}

/*
	结构体实现接口中的防范
*/
func (phone HuaWeiPhone) call() string {
	return "i am Huawei, who are you?"
}

/*
	结构体实现接口中的防范
*/
func (phone MiPhone) call() string {
	return "i am Mi, who are you?"
}

func main() {

	//声明Phone变量
	var phone Phone

	//将HuaWeiPhone赋值给Phone变量
	phone = new(HuaWeiPhone)
	fmt.Println(phone.call())

	//将MiPhone赋值给Phone变量
	phone = new(MiPhone)
	fmt.Println(phone.call())

}
