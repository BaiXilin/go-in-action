package main

import (
	"fmt"

	"visibility/entities"
)

func main() {
	a := entities.Admin{
		Level: "admin",
	}
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("%+v\n", a)
}
