package main

import "fmt"

/*
	切片（动态数组）
	和数组区别：数组长度固定，切片可动态增长


	通过make([]type, len)|make([]type, len, capacity)声明切片
	len：初始长度
	capacity：可选参数
*/
func main() {

	//声明切片 make(type, len)
	var sliceInt []int = make([]int, 5)
	sliceInt = []int{1, 2, 4, 8, 16, 32}
	for x, v := range sliceInt {
		fmt.Printf("当前为%d值为%d\n", x, v)
	}

	//声明切片
	slice2 := make([]string, 5)
	slice2 = []string{"aaa", "bbb", "ccc"}
	for x, v := range slice2 {
		fmt.Printf("当前为%d值为%s\n", x, v)
	}

	//声明并初始化slice3 []代表切片类型
	slice3 := []int{1, 2, 4}
	for x, v := range slice3 {
		fmt.Printf("当前为%d值为%d\n", x, v)
	}

	//len(type)切片长度 cap(type)切片最大长度
	fmt.Printf("len = %d\ncap = %d\n value = %v\n", len(slice3), cap(slice3), slice3)

	//切片为空
	slice3 = nil
	fmt.Printf("%v\n", slice3)

	//切片截取
	slice3 = []int{1, 2, 4, 8, 16, 32, 64}
	//需要注意不包含5
	fmt.Printf("截取[1:5]%v\n", slice3[1:5])
	//从开始位到第5位
	fmt.Printf("截取[:5]%v\n", slice3[:5])
	//从第1位到最后一位
	fmt.Printf("截取[1:]%v\n", slice3[1:])
}
