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
			为 admin 类型增加了 notifier 接口的实现。当 admin 类型的实现被调用
		时，会显示"Sending admin email"。作为对比，user 类型的实现被调用时，会显示"Sending
		user email"。
	*/

	ad.user.notify() //Sending user email to john smith<john@yahoo.com>

	ad.notify() //Sending admin email to john smith<john@yahoo.com>
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(u notifier) {
	u.notify()
}
