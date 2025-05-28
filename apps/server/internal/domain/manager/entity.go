package manager

import (
	"net/http"

	"imobiliary/internal/domain/types"
	"imobiliary/internal/response"

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

func (u *Manager) hashPassword() *response.Err {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
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
