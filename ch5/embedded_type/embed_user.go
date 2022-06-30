package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Send notification to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	a := admin{
		user: user{
			name:  "lisa",
			email: "lisa@email.com",
		},
		level: "admin",
	}

	a.user.notify()
	a.notify()
}
