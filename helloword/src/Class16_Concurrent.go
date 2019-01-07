package main

import "fmt"

/*
	TODO 并发

	goroutine
	通过go关键字，开启goroutine
	goroutine是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的
	需要注意在同一个程序中的所有 goroutine 共享同一个地址空间

	channel
	用来传递数据的一个数据结构。
	通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
*/

func say(s string) {

	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	//遍历数组，_代表忽略改变量，v是s对应位值
	for _, v := range s {
		// 把 s[i] 发送到通道
		c <- v
	}

	//关闭通道
	close(c)
}

func main() {
	//启动一个goroutine
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2)

	go sum(s, c)

	//需要关闭通道，才能遍历输出
	for i := range c {
		fmt.Println(i)
	}

}
