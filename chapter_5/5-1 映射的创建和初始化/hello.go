package main

import "fmt"

func main() {

	dict_1 := make(map[string]int)

	dict_2 := map[string]string{"Red": "#1", "Blue": "#2"}

	/* 映射的键可以是任何值。这个值的类型可以是内置的类型，也可以是结构类型，只要这个值
	可以使用==运算符做比较。切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
	不能作为映射的键，使用这些类型会造成编译错误 */

	// dict_3 := map[[]string]string{} invalid map key type

	dict_4 := map[string][]int{"Red": []int{1, 2, 3}}

	fmt.Println(dict_1, dict_2, dict_4)
}
