package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	u.name = "bill-notified"
}

func main() {
	u := user{"bill", "bill@gmail.com"}
	sendNotification(&u)

	fmt.Printf("%+v\n", u)

}

func sendNotification(u notifier) {
	u.notify()
}
