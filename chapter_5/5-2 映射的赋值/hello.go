package main

import "fmt"

func main() {

	// 声明一个nil映射
	/* 可以通过声明一个未初始化的映射来创建一个值为 nil 的映射（称为 nil 映射 ）。nil 映射
	不能用于存储键值对，否则，会产生一个语言运行时错误 */

	// var colors map[string]string

	// colors["Red"] = "#00cece" //panic: runtime error: assignment to entry in nil map

	// fmt.Println(colors)

	// 判断键是否存在

	// 1

	colors := map[string]string{"Red": "#00cece"}

	_, exists := colors["Red"]

	// 2

	value := colors["Blue"]

	// 这个键存在吗？
	if value != "" {
		fmt.Println(value)
	}

	fmt.Println(colors, exists)

	//使用 range 迭代映射
	colors_dict := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 从映射中 删除一项

	delete(colors_dict, "Coral")

	for key, value := range colors_dict {
		fmt.Printf("Key : %s Value: %s\n", key, value)

	}

	fmt.Println(colors_dict)
}
