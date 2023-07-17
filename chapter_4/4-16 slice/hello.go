package main

import "fmt"

// 使用指针在函数间传递大数组
func main(){
	
	// make关键字
	// slice1 := make([]string,4,2)  //容量小于长度的切片会在编译时报错
	slice1 := make([]string,2,4)

	// 字面量声明
	// 创造字符串切片 长度和容量都是5
	slice2 := []string{"Red","White"}

	// 创造整型切片 长度和容量都是2
	slice3 := []int{1,2}

	// 使用索引声明切片
	slice4 := []int{99:1}
	
	/* 不管是使用 nil 切片还是空切片，对其调用内置函数 append、len 和 cap 的效果都是一样的 */

	// 创造nil切片
	var slice5 []int

	// 声明空切片
	// 1.make
	slice6 := make([]int,0)
	// 2.字面量
	slice7 := []int{}


	fmt.Println(slice1,slice2,slice3,slice4,slice5,slice6,slice7)
}

