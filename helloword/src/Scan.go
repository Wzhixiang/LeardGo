package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Printf("你输入的信息为：%s", input.Text())
}
