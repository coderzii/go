package main

import "fmt"

func main() {
	// 编译器会阻止类型不同的数组互相赋值

	var array1 [4]string

	array2 := [5]string{"Red","Blue","Green","Yellow","Pink"}

	// array1 = array2 // cannot use array2 (type [5]string) as type [4]string in assignment

	fmt.Println("Hello, World!", array1,array2)
}
