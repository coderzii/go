package main

import (
	"fmt"
)

var ()

func main() {

	// 无缓冲的整型通道
	// unbuffered := make(chan int)

	// 有缓冲的字符串通道
	buffered := make(chan string, 10)

	// 向通道发送值或者指针需要用到<-操作符

	buffered <- "Gopher"

	// 当从通道里接收一个值或者指针时，<-运算符在要操作的通道变量的左侧
	value := <-buffered

	fmt.Println("buffered", value)

	//通道是否带有缓冲，其行为会有一些不同

}
