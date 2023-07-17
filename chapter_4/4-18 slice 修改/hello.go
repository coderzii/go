package main

import "fmt"

// 修改切片内容可能导致的结果
func main(){

	slice1 := []int {10,20,30,40,50}

	// 其长度是 2 个元素，容量是 4 个元素 长度=3-1 容量= slice1容量-1
	new_slice1 := slice1[1:3]

	// slice1[2] =200
	// new_slice1[1] = 1000

	fmt.Println(slice1,new_slice1)

	// 表示索引越界的语言运行时错误

	slice2 := []int{1,2,3,4,5}

	new_slice2 := slice2[1:2]

	var value int = new_slice2[1]
	fmt.Println(slice2,new_slice2,value)
}

