package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	errors "github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type MemUserRepository struct {
	users []*user.User
}

func NewMemUserRepository() *MemUserRepository {
	return &MemUserRepository{}
}

func (r *MemUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, errors.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}

func (r *MemUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, errors.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}

func (r *MemUserRepository) Create(ctx context.Context, user *user.User) (uuid.UUID, *response.Err) {
	r.users = append(r.users, user)
	return user.ID, nil
}

func (r *MemUserRepository) Update(ctx context.Context, user *user.User) *response.Err {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}

	return response.NewErr(http.StatusNotFound, errors.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}

func (r *MemUserRepository) Delete(ctx context.Context, id uuid.UUID) *response.Err {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return response.NewErr(http.StatusNotFound, errors.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}

func (r *MemUserRepository) Authenticate(ctx context.Context, email, password string) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.Email == email {
			if user.ComparePwd(password) {
				return user, nil
			}

			return nil, response.NewErr(http.StatusUnauthorized, errors.ERR_PASSWORD_INVALID)
		}
	}

	return nil, response.NewErr(http.StatusNotFound, errors.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
