package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
我们要学习的处理 JSON 的第二个方面是，使用 json 包的 MarshalIndent 函数进行编码。
这个函数可以很方便地将 Go 语言的 map 类型的值或者结构类型的值转换为易读格式的 JSON 文
档。序列化（marshal）是指将数据转换为 JSON 字符串的过程
*/

func main() {

	c := make(map[string]interface{})

	c["name"] = "Gopher"

	c["title"] = "programmer"

	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 映射 序列化 JSON

	data, err := json.MarshalIndent(c, "", "	")

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))

}
