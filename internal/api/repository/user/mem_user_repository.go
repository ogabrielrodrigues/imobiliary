package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared/logger"
)

type MemUserRepository struct {
	users []*user.User
}

func NewMemUserRepository() *MemUserRepository {
	return &MemUserRepository{}
}

func (r *MemUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	logger.Info("Searching for user by ID", "id", id)

	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found or not exists")
}

func (r *MemUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	logger.Info("Searching for user by E-mail", "email", email)

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found or not exists")
}

func (r *MemUserRepository) Create(ctx context.Context, user *user.User) (uuid.UUID, error) {
	logger.Info("Creating user", "user", user)

	r.users = append(r.users, user)
	return user.ID, nil
}

func (r *MemUserRepository) Update(ctx context.Context, user *user.User) error {
	logger.Info("Updating user", "user", user)

	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}

	return errors.New("user not found or not exists")
}

func (r *MemUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	logger.Info("Deleting user", "id", id)

	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found or not exists")
}

func (r *MemUserRepository) Authenticate(ctx context.Context, email, password string) (*user.User, error) {
	logger.Info("Authenticating user", "email", email)

	for _, user := range r.users {
		if user.Email == email {
			if user.ComparePwd(password) {
				return user, nil
			}

			return nil, errors.New("invalid password")
		}
	}

	return nil, errors.New("user not found or not exists")
}
