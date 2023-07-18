package main

import (
	"fmt"
)

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

	ad.user.notify()

	ad.notify()

}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}
