package main

import "fmt"

func foo(slice []int)[]int {
	return slice
}

// 迭代
func main() {

	slice := make([]int, 10)

	res := foo(slice)

	fmt.Println(res,&res,&slice)

}
