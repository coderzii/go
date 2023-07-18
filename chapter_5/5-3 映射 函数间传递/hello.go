package main

import "fmt"

/*
在函数间传递映射并不会制造出该映射的一个副本。实际上，当传递映射给一个函数，并对
这个映射做了修改时，所有对这个映射的引用都会察觉到这个修改
*/
func main() {

	colors := map[string]string{"Red": "#00cece"}

	removeColor(colors, "Red")

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	fmt.Println(colors)
}

func removeColor(color_dict map[string]string, key string) {
	delete(color_dict, key)
}
