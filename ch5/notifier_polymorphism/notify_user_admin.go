package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Send notification to user %s<%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("Send notification to admin %s<%s>\n", a.name, a.email)
}

func main() {
	u := user{"bill", "bill@email.com"}
	a := admin{"lisa", "lisa@admin.email.com"}

	sendNotification(&u)
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
