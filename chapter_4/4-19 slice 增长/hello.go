package main

import "fmt"

// 修改切片内容可能导致的结果
func main(){

	// 其长度为 5 个元素，容量为 5 个元素
	slice1 := []int {10,20,30,40,50}

	// 其长度为 2 个元素，容量为 4 个元素
	new_slice1 := slice1[1:3]

	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60\
	
	new_slice1 = append(new_slice1,60)//由于和原始的 slice 共享同一个底层数组，slice 中索引为 3 的元素的值也被改动了

	fmt.Println(slice1,new_slice1)

// 如果切片的底层数组没有足够的可用容量，append 函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新的值，如代码清单 4-31 所示。
	slice2 := []int{10,20,30}

	// new_slice2 := slice2[1:2]
	
	/* 函数 append 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是
会成倍地增加容量 */
	new_slice2 := append(slice2,999) //当这个 append 操作完成后，newSlice 拥有一个全新的底层数组，这个数组的容量是原来的两倍


	new_slice2=append(new_slice2,111)
	new_slice2=append(new_slice2,222)


	fmt.Println(slice2,new_slice2)

}

