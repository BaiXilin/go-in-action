package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Level string
}
