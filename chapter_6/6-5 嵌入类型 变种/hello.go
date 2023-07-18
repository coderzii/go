package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string

	email string
}

type admin struct {
	user

	level string
}

func main() {

	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	/*
	   这里才是事情变得有趣的地方。我们创建了一个名为 ad 的变
	   量，其类型是外部类型 admin。这个类型内部嵌入了 user 类型。之后第 48 行，我们将这个外
	   部类型变量的地址传给 sendNotification 函数。编译器认为这个指针实现了 notifier 接
	   口，并接受了这个值的传递。不过如果看一下整个示例程序，就会发现 admin 类型并没有实现
	   这个接口。
	*/
	sendNotification(&ad)
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(u notifier) {
	u.notify()
}
