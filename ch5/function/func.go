package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
	u.name,
	u.email)
}

func  (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	lisa  := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

	bill.changeEmail("bill@outlook.com")
	bill.notify()

	lisa.changeEmail("lisa@outlook.com")
	lisa.notify()
}
