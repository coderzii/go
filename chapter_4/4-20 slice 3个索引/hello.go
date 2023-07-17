package main

import "fmt"

// 修改切片内容可能导致的结果
func main(){

	// 其长度为 5 个元素，容量为 5 个元素
	slice1 := []string {"Apple","Orange","Plum","Banana","Grape"}


	// 将第三个元素切片，并限制容量
	// 其长度为 1 个元素，容量为 2 个元素
	new_slice1 := slice1[2:3:4]  //[i:j:k]

	// 长度: j – i 或 3 - 2 = 1
	// 容量: k – i 或 4 - 2 = 2
	fmt.Println(slice1,new_slice1)

	// 设置容量大于已有容量的语言运行时错误 可用容量 5 - 2 = 3

	// new_slice2 := slice1[2:3:6] //slice bounds out of range [::6] with capacity 5
	// fmt.Println(new_slice2)

	// 设置长度和容量一样的好处

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	/* 如果不加第三个索引，由于剩余的所有容量都属于 slice，向 slice 追加 Kiwi 会改变
原有底层数组索引为 3 的元素的值 Banana。不过在代码清单 4-36 中我们限制了 slice 的容
量为 1。当我们第一次对 slice 调用 append 的时候，会创建一个新的底层数组，这个数组包
括 2 个元素，并将水果 Plum 复制进来，再追加新水果 Kiwi，并返回一个引用了这个底层数组
的新切片 */
	source_slice := source[2:3:3] //长度和容量一样 : Plum

	source_slice = append(source_slice,"Kiwi")

	fmt.Println(source,source_slice)
}

