package main

import "fmt"

// 使用指针在函数间传递大数组
func main(){

	slice1 := []int {1,2}

	slice1[0] =10

	// 使用切片创建切片


	slice2 := []int{10,20,30,40,50}

	/* 第一个值表示新切片开始的元素的索引位置，这个例子中是 1。第二个值表示开始的索引位置（1），加上希望包含的元素的个数（2），1+2 的结果是 3，所以第二个值就是 3 */

	slice2_new := slice2[1:3] //[20 30]

	fmt.Println(slice1,slice2,slice2_new)


}

