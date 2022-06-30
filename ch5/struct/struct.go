package main

import "fmt"

type user struct {
	name		string
	email		string
	ext		int
	privileged	bool
}

type admin struct {
	person user
	level string
}

func main() {
	// Different ways to declare a struct
	// Method 1: var
	var bill user
	fmt.Printf("%+v\n", bill)

	// Method 2: literal
	lisa := user{
		name: "Lisa",
		email: "l12ji@uwaterloo.ca",
		ext: 12345,
		privileged: true,
	}
	fmt.Printf("%+v\n", lisa)

	// Nested Structs
	fred := admin{
		person: lisa,
		level: "super",
	}
	fmt.Printf("%+v\n", fred)
}
