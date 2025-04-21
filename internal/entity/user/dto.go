package user

import (
	"github.com/google/uuid"
)

type DTO struct {
	ID        string `json:"id"`
	CreciID   string `json:"creci_id"`
	Fullname  string `json:"fullname"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
}

func (d *DTO) ToUser() *User {
	return &User{
		ID:        uuid.MustParse(d.ID),
		CreciID:   d.CreciID,
		Fullname:  d.Fullname,
		Cellphone: d.Cellphone,
		Email:     d.Email,
		Avatar:    d.Avatar,
	}
}

type CreateDTO struct {
	CreciID   string `json:"creci_id"`
	Fullname  string `json:"fullname"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (d *CreateDTO) ToUser() *User {
	u := &User{
		CreciID:   d.CreciID,
		Fullname:  d.Fullname,
		Cellphone: d.Cellphone,
		Email:     d.Email,
	}

	u.SetPassword(d.Password)

	return u
}

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
