package main

import "fmt"

func main() {
	// 声明二维数组
	var array1 [4][2]int


	array1 = [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	array2 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	//  声明并初始化外层数组中索引为 1 个和 3 的元素
	array3  := [4][2]int{1:{20,21},3:{40,41}}

	// 声明并初始化外层数组和内层数组的单个元素
	arry4 := [4][2]int{1:{0:20},3:{1:41}}

	fmt.Println("Hello, World!", array1,array2,array3,arry4)
}
