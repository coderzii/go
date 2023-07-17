package main

import "fmt"

func main() {
	// 把同样类型的一个数组赋值给另外一个数组

	var array1 [5]string

	array2 := [5]string{"Red","Blue","Green","Yellow","Pink"}

	array1 = array2

	fmt.Println("Hello, World!", array1,array2)
}
