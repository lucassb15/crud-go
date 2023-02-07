package controllers

import (
	"modulo/db"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create(name string, email string) error {

	s := User{Name: name, Email: email}

	return db.Insert("Teste", s)
}
