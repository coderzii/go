package main

import "fmt"

func main() {
	// 声明二维数组
	var array1 [4][2]int


	array1 = [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	// 访问二维数组的元素
	array1[0][1] = 100
	array1[0][0] = 10

	fmt.Println("Hello, World!", array1,)
}
