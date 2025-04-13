package user

import (
	"context"

	"github.com/google/uuid"
)

const (
	ERR_USER_NOT_FOUND_OR_NOT_EXISTS             = "user not found or not exists"
	ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED = "only one of the parameters must be provided"
)

type Repository struct{}

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (uuid.UUID, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Authenticate(ctx context.Context, email, password string) (*User, error)
}

func NewRepository() *Repository {
	return &Repository{}
}
