package main

import "fmt"

// 声明一个结构类型

type user struct {
	name string

	email string

	ext int

	privileged bool
}

type admin struct {
	person user

	level string
}

/*
这个值要么用指定的值初始化，要么用零值（即变
量类型的默认值）做初始化。对数值类型来说，零值是 0；对字符串来说，零值是空字符串；对
布尔类型，零值是 false
*/

func main() {

	var bill user

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// 不使用字段名，创建结构类型的值
	tom := user{"Tom", "tom@email.com", 123, true}

	fred_admin := admin{
		person: lisa,

		level: "administrator",
	}

	fmt.Println(bill, lisa, tom, fred_admin)

}
