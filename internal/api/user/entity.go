package user

import (
	"fmt"

	"github.com/google/uuid"
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

	fmt.Println("password", u.password)
	fmt.Println("hash", string(hash))

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

func New(creci_id, fullname, cellphone, email, password, avatar string) (*User, error) {
	u := &User{
		ID:        uuid.New(),
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

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) ChangePassword(password string) {
	u.password = password

	u.hashPwd()
}

func (u *User) ToDTO() *DTO {
	return &DTO{
		ID:        u.ID.String(),
		CreciID:   u.CreciID,
		Fullname:  u.Fullname,
		Cellphone: u.Cellphone,
		Email:     u.Email,
		Avatar:    u.Avatar,
	}
}

func UserFromDTO(dto *DTO) *User {
	return &User{
		ID:        uuid.MustParse(dto.ID),
		CreciID:   dto.CreciID,
		Fullname:  dto.Fullname,
		Cellphone: dto.Cellphone,
		Email:     dto.Email,
		Avatar:    dto.Avatar,
	}
}
