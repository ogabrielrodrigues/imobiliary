package entity

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

type user struct {
	id          uuid.UUID
	creci_id    string
	fullname    string
	email       string
	access_code string
}

func NewUser(id uuid.UUID, creci_id, fullname, email string) (*user, error) {
	u := &user{
		id:       id,
		creci_id: creci_id,
		fullname: fullname,
		email:    email,
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *user) GenerateAccessCode() {
	var min = 48 // ASCII Char 0
	var max = 57 // ASCII Char 9

	var code []rune

	for range 8 {
		code = append(code, rune(rand.IntN(max-min)+min))
	}

	u.access_code = string(code)
}

// TODO: Getters e Setters
