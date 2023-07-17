package main

import "fmt"

// 同样类型的多维数组赋值

func main() {
	// 声明二维数组
	var array1 [4][2]int

	var array2 [4][2]int


	array1 = [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	// 访问二维数组的元素
	array1[0][1] = 100
	array1[0][0] = 10


	array2 = array1

	fmt.Println("Hello, World!", array1,array2)
}
