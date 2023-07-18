package main

import "fmt"

// 基于 int64 声明一个新类型
type Duration int

func main() {
	// 给不同类型的变量赋值会产生编译错误

	var dur Duration

	/*
			虽然 int64 类型是基础类型，Duration 类型依然是一个独立的类型。两种不同类型的
		值即便互相兼容，也不能互相赋值。编译器不会对不同类型的值做隐式转换。
	*/

	// dur = int64(1000) /

	fmt.Println(dur)

}
