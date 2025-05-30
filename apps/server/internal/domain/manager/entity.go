package manager

import (
	"context"

	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	ID       uuid.UUID
	Fullname string
	Phone    types.Phone
	Email    types.Email
	Password string
}

func NewManager(fullname string, phone types.Phone, email types.Email, password string) (*Manager, *httperr.HttpError) {
	newManager := &Manager{
		ID:       uuid.New(),
		Fullname: fullname,
		Phone:    phone,
		Email:    email,
		Password: password,
	}

	if err := newManager.hashPassword(); err != nil {
		return nil, err
	}

	return newManager, nil
}

func (u *Manager) hashPassword() *httperr.HttpError {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return httperr.NewInternalServerError(context.Background(), "error creating user")
	}

	u.Password = string(hash)
	return nil
}

func (u *Manager) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *Manager) ChangePassword(password string) {
	u.Password = password

	u.hashPassword()
}
