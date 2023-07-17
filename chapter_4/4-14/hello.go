package main

import "fmt"

func foo(array [1e6]int){
	fmt.Println(array)
}

func main(){
	// 声明一个需要8mb的数组

	var array [1e6]int

	foo(array)
}

