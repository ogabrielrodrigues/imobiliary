package entity

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/dto"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Fullname  string
	CreciID   string
	Cellphone string
	Email     string
	password  string
	Avatar    string
}

func (u *User) hashPwd() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), 14)
	if err != nil {
		return err
	}

	u.password = string(hash)
	return nil
}

func (u *User) ComparePwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func NewUser(id uuid.UUID, creci_id, fullname, cellphone, email, password, avatar string) (*User, error) {
	u := &User{
		ID:        id,
		CreciID:   creci_id,
		Fullname:  fullname,
		Cellphone: cellphone,
		Email:     email,
		Avatar:    avatar,
		password:  password,
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	if err := u.hashPwd(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:        u.ID.String(),
		CreciID:   u.CreciID,
		Fullname:  u.Fullname,
		Cellphone: u.Cellphone,
		Email:     u.Email,
		Avatar:    u.Avatar,
	}
}

func UserFromDTO(dto *dto.UserDTO) *User {
	return &User{
		ID:        uuid.MustParse(dto.ID),
		CreciID:   dto.CreciID,
		Fullname:  dto.Fullname,
		Cellphone: dto.Cellphone,
		Email:     dto.Email,
		Avatar:    dto.Avatar,
	}
}
