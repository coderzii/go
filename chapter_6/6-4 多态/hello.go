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
	name string

	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {

	lisa := user{"Lisa", "lisa@user.com"}

	sendNotification(&lisa)

	bill := admin{"Bill", "bill@admin.com"}

	sendNotification(&bill)

}
