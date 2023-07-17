package main

import "fmt"

func foo(array *[1e6]int){
	fmt.Println(*array)
}

// 使用指针在函数间传递大数组
func main(){
	var array [1e6]int

	foo(&array)
}

