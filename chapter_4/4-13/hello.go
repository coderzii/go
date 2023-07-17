package main

import "fmt"

// 使用索引为多维数组赋值
func main() {
	// 声明二维数组
	

	array1 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}


	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array2 [2]int  = array1[1]


	// 将外层数组的索引为 1、内层数组的索引为 0 的整型值复制到新的整型变量里

	var value int = array1[1][0]

	fmt.Println("Hello, World!", array1,array2,value)
}
