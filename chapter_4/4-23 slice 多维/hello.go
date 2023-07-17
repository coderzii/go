package main

import "fmt"

// 迭代
func main() {

	slice := [][]int{{10}, {100, 200}}

	fmt.Println(slice)

	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)

	fmt.Println(slice)

}
