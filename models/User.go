package models

type User struct {
	Name string
}

func (u *User) CreateUser(name string) {
	u.Name = name
}
