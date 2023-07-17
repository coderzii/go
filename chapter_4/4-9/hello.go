package main

import "fmt"

func main() {
	// 把一个指针数组赋值给另一个

	var array1 [3]*string

	array2 := [3]*string{new(string),new(string),new(string)}

	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "White"

	array1 = array2

	fmt.Println("Hello, World!", array1,array2)
}
