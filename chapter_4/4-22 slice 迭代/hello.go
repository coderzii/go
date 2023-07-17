package main

import "fmt"

// 迭代
func main() {

	slice := []int{10, 20, 30, 40}

	for index, value := range slice {

		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	//  range 提供了每个元素的副本
	for index, value := range slice {
		/* 因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以 value 的地址总
		是相同的。要想获取每个元素的地址，可以使用切片变量和索引值 */
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}

	// 使用空白标识符（下划线）来忽略索引值

	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}


	// 使用传统的 for 循环对切片进行迭代(函数 len 返回切片的长度，函数 cap 返回切片的容量)

	for index :=2; index < len(slice);index++{
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}
